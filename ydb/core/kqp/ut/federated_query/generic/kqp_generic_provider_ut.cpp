#include <ydb/core/kqp/ut/common/kqp_ut_common.h>
#include <ydb/core/kqp/ut/federated_query/common/common.h>
#include <ydb/library/yql/providers/common/structured_token/yql_token_builder.h>
#include <ydb/library/yql/providers/generic/connector/libcpp/client.h>
#include <ydb/library/yql/providers/generic/connector/libcpp/ut_helpers/connector_client_mock.h>
#include <ydb/library/yql/providers/generic/connector/libcpp/ut_helpers/database_resolver_mock.h>
#include <ydb/public/sdk/cpp/client/ydb_operation/operation.h>
#include <ydb/public/sdk/cpp/client/ydb_query/query.h>
#include <ydb/public/sdk/cpp/client/ydb_types/status_codes.h>

#include <library/cpp/testing/unittest/registar.h>

#include <util/system/defaults.h>
#include <util/system/env.h>

#include <arrow/api.h>

#include <google/protobuf/util/message_differencer.h>

#include <fmt/format.h>

namespace NKikimr::NKqp {
    using namespace NYdb;
    using namespace NYdb::NQuery;
    using namespace NYql::NConnector;
    using namespace NYql::NConnector::NTest;
    using namespace NKikimr::NKqp::NFederatedQueryTest;
    using namespace testing;
    using namespace fmt::literals;

    enum class EProviderType {
        PostgreSQL,
        ClickHouse,
    };

    NApi::TDataSourceInstance MakeDataSourceInstance(EProviderType providerType) {
        switch (providerType) {
            case EProviderType::PostgreSQL:
                return TConnectorClientMock::TPostgreSQLDataSourceInstanceBuilder<>().GetResult();
            case EProviderType::ClickHouse:
                return TConnectorClientMock::TClickHouseDataSourceInstanceBuilder<>().GetResult();
        }
    }

    void CreateExternalDataSource(EProviderType providerType, const std::shared_ptr<NKikimr::NKqp::TKikimrRunner>& kikimr) {
        switch (providerType) {
            case EProviderType::PostgreSQL:
                return CreatePostgreSQLExternalDataSource(kikimr);
            case EProviderType::ClickHouse:
                return CreateClickHouseExternalDataSource(kikimr);
        }
    }

    Y_UNIT_TEST_SUITE(GenericFederatedQuery) {
        void TestSelectAllFields(EProviderType providerType) {
            // prepare mock
            auto clientMock = std::make_shared<TConnectorClientMock>();

            const NApi::TDataSourceInstance dataSourceInstance = MakeDataSourceInstance(providerType);

            // step 1: DescribeTable
            clientMock->ExpectDescribeTable()
                .DataSourceInstance(dataSourceInstance)
                .Response()
                    .Column("col1", Ydb::Type::UINT16);

            // step 2: ListSplits
            clientMock->ExpectListSplits()
                .Select()
                    .DataSourceInstance(dataSourceInstance)
                    .What()
                        .Column("col1", Ydb::Type::UINT16)
                        .Done()
                    .Done()
                .Response()
                    .Split()
                        .Description("some binary description")
                        .Select()
                            .DataSourceInstance(dataSourceInstance)
                            .What()
                                .Column("col1", Ydb::Type::UINT16);

            // step 3: ReadSplits
            std::vector<ui16> colData = {10, 20, 30, 40, 50};
            clientMock->ExpectReadSplits()
                .DataSourceInstance(dataSourceInstance)
                .Split()
                    .Description("some binary description")
                    .Select()
                        .DataSourceInstance(dataSourceInstance)
                        .What()
                            .Column("col1", Ydb::Type::UINT16)
                            .Done()
                        .Done()
                    .Done()
                .Response()
                    .RecordBatch(MakeRecordBatch<arrow::UInt16Builder>("col1", colData, arrow::uint16()));

            // prepare database resolver mock
            std::shared_ptr<TDatabaseAsyncResolverMock> databaseAsyncResolverMock;
            if (providerType == EProviderType::ClickHouse) {
                databaseAsyncResolverMock = std::make_shared<TDatabaseAsyncResolverMock>();
                databaseAsyncResolverMock->AddClickHouseCluster();
            }

            // run test
            auto kikimr = MakeKikimrRunner(nullptr, clientMock, databaseAsyncResolverMock);

            CreateExternalDataSource(providerType, kikimr);

            const TString query = fmt::format(
                R"(
                SELECT * FROM {data_source_name}.`{database_name}.{table_name}`;
            )",
                "data_source_name"_a = DEFAULT_DATA_SOURCE_NAME,
                "database_name"_a = DEFAULT_DATABASE,
                "table_name"_a = DEFAULT_TABLE);

            auto db = kikimr->GetQueryClient();
            auto scriptExecutionOperation = db.ExecuteScript(query).ExtractValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(scriptExecutionOperation.Status().GetStatus(), EStatus::SUCCESS, scriptExecutionOperation.Status().GetIssues().ToString());
            UNIT_ASSERT(scriptExecutionOperation.Metadata().ExecutionId);

            NYdb::NQuery::TScriptExecutionOperation readyOp = WaitScriptExecutionOperation(scriptExecutionOperation.Id(), kikimr->GetDriver());
            UNIT_ASSERT_C(readyOp.Metadata().ExecStatus == EExecStatus::Completed, readyOp.Status().GetIssues().ToString());
            TFetchScriptResultsResult results = db.FetchScriptResults(scriptExecutionOperation.Id(), 0).ExtractValueSync();
            UNIT_ASSERT_C(results.IsSuccess(), results.GetIssues().ToString());

            TResultSetParser resultSet(results.ExtractResultSet());
            UNIT_ASSERT_VALUES_EQUAL(resultSet.ColumnsCount(), 1);
            UNIT_ASSERT_VALUES_EQUAL(resultSet.RowsCount(), colData.size());

            // check every row
            MATCH_RESULT_WITH_INPUT(colData, resultSet, GetUint16);
        }

        Y_UNIT_TEST(PostgreSQLLocal) {
            TestSelectAllFields(EProviderType::PostgreSQL);
        }

        Y_UNIT_TEST(ClickHouseManaged) {
            TestSelectAllFields(EProviderType::ClickHouse);
        }

        void TestSelectConstant(EProviderType providerType) {
            // prepare mock
            auto clientMock = std::make_shared<TConnectorClientMock>();

            const NApi::TDataSourceInstance dataSourceInstance = MakeDataSourceInstance(providerType);

            constexpr size_t ROWS_COUNT = 5;

            // step 1: DescribeTable
            clientMock->ExpectDescribeTable()
                .DataSourceInstance(dataSourceInstance)
                .Response()
                    .Column("col1", Ydb::Type::UINT16)
                    .Column("col2", Ydb::Type::DOUBLE);

            // step 2: ListSplits
            clientMock->ExpectListSplits()
                .Select()
                    .DataSourceInstance(dataSourceInstance)
                    .What()
                        // Empty
                        .Done()
                    .Done()
                .Response()
                    .Split()
                        .Description("some binary description")
                        .Select()
                            .DataSourceInstance(dataSourceInstance)
                            .What();

            // step 3: ReadSplits
            clientMock->ExpectReadSplits()
                .DataSourceInstance(dataSourceInstance)
                .Split()
                    .Description("some binary description")
                    .Select()
                        .DataSourceInstance(dataSourceInstance)
                        .What()
                            .Done()
                        .Done()
                    .Done()
                .Response()
                    .RecordBatch(MakeEmptyRecordBatch(ROWS_COUNT));

            // prepare database resolver mock
            std::shared_ptr<TDatabaseAsyncResolverMock> databaseAsyncResolverMock;
            if (providerType == EProviderType::ClickHouse) {
                databaseAsyncResolverMock = std::make_shared<TDatabaseAsyncResolverMock>();
                databaseAsyncResolverMock->AddClickHouseCluster();
            }

            // run test
            auto kikimr = MakeKikimrRunner(nullptr, clientMock, databaseAsyncResolverMock);

            CreateExternalDataSource(providerType, kikimr);

            const TString query = fmt::format(
                R"(
                SELECT 42 FROM {data_source_name}.`{database_name}.{table_name}`;
            )",
                "data_source_name"_a = DEFAULT_DATA_SOURCE_NAME,
                "database_name"_a = DEFAULT_DATABASE,
                "table_name"_a = DEFAULT_TABLE);

            auto db = kikimr->GetQueryClient();
            auto queryResult = db.ExecuteQuery(query, TTxControl::BeginTx().CommitTx(), TExecuteQuerySettings()).ExtractValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(queryResult.GetStatus(), EStatus::SUCCESS, queryResult.GetIssues().ToString());

            TResultSetParser resultSet(queryResult.GetResultSetParser(0));
            UNIT_ASSERT_VALUES_EQUAL(resultSet.ColumnsCount(), 1);
            UNIT_ASSERT_VALUES_EQUAL(resultSet.RowsCount(), ROWS_COUNT);

            // check every row
            std::vector<i32> constants(ROWS_COUNT, 42);
            MATCH_RESULT_WITH_INPUT(constants, resultSet, GetInt32);
        }

        Y_UNIT_TEST(PostgreSQLSelectConstant) {
            TestSelectConstant(EProviderType::PostgreSQL);
        }

        Y_UNIT_TEST(ClickHouseManagedSelectConstant) {
            TestSelectConstant(EProviderType::ClickHouse);
        }

        void TestSelectCount(EProviderType providerType) {
            // prepare mock
            auto clientMock = std::make_shared<TConnectorClientMock>();

            const NApi::TDataSourceInstance dataSourceInstance = MakeDataSourceInstance(providerType);

            constexpr size_t ROWS_COUNT = 5;

            // step 1: DescribeTable
            clientMock->ExpectDescribeTable()
                .DataSourceInstance(dataSourceInstance)
                .Response()
                    .Column("col1", Ydb::Type::UINT16)
                    .Column("col2", Ydb::Type::DOUBLE);

            // step 2: ListSplits
            clientMock->ExpectListSplits()
                .Select()
                    .DataSourceInstance(dataSourceInstance)
                    .What()
                        // Empty
                        .Done()
                    .Done()
                .Response()
                    .Split()
                        .Description("some binary description")
                        .Select()
                            .DataSourceInstance(dataSourceInstance)
                            .What();

            // step 3: ReadSplits
            clientMock->ExpectReadSplits()
                .DataSourceInstance(dataSourceInstance)
                .Split()
                    .Description("some binary description")
                    .Select()
                        .DataSourceInstance(dataSourceInstance)
                        .What()
                            .Done()
                        .Done()
                    .Done()
                .Response()
                    .RecordBatch(MakeEmptyRecordBatch(ROWS_COUNT));

            // prepare database resolver mock
            std::shared_ptr<TDatabaseAsyncResolverMock> databaseAsyncResolverMock;
            if (providerType == EProviderType::ClickHouse) {
                databaseAsyncResolverMock = std::make_shared<TDatabaseAsyncResolverMock>();
                databaseAsyncResolverMock->AddClickHouseCluster();
            }

            // run test
            auto kikimr = MakeKikimrRunner(nullptr, clientMock, databaseAsyncResolverMock);

            CreateExternalDataSource(providerType, kikimr);

            const TString query = fmt::format(
                R"(
                SELECT COUNT(*) FROM {data_source_name}.`{database_name}.{table_name}`;
            )",
                "data_source_name"_a = DEFAULT_DATA_SOURCE_NAME,
                "database_name"_a = DEFAULT_DATABASE,
                "table_name"_a = DEFAULT_TABLE);

            auto db = kikimr->GetQueryClient();
            auto queryResult = db.ExecuteQuery(query, TTxControl::BeginTx().CommitTx(), TExecuteQuerySettings()).ExtractValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(queryResult.GetStatus(), EStatus::SUCCESS, queryResult.GetIssues().ToString());

            TResultSetParser resultSet(queryResult.GetResultSetParser(0));
            UNIT_ASSERT_VALUES_EQUAL(resultSet.ColumnsCount(), 1);
            UNIT_ASSERT_VALUES_EQUAL(resultSet.RowsCount(), 1);

            // check every row
            std::vector<ui64> result = { ROWS_COUNT };
            MATCH_RESULT_WITH_INPUT(result, resultSet, GetUint64);
        }

        Y_UNIT_TEST(PostgreSQLSelectCount) {
            TestSelectCount(EProviderType::PostgreSQL);
        }

        Y_UNIT_TEST(ClickHouseSelectCount) {
            TestSelectCount(EProviderType::ClickHouse);
        }
    }
}
