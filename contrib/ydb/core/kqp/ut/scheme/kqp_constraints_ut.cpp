#include <contrib/ydb/core/kqp/ut/common/kqp_ut_common.h>
#include <contrib/ydb/core/formats/arrow/arrow_helpers.h>
#include <contrib/ydb/core/tx/tx_proxy/proxy.h>
#include <contrib/ydb/core/tx/columnshard/columnshard_ut_common.h>
#include <contrib/ydb/public/sdk/cpp/client/ydb_proto/accessor.h>
#include <contrib/ydb/public/sdk/cpp/client/ydb_scheme/scheme.h>
#include <contrib/ydb/public/sdk/cpp/client/ydb_topic/topic.h>
#include <contrib/ydb/public/sdk/cpp/client/draft/ydb_long_tx.h>
#include <contrib/ydb/core/testlib/cs_helper.h>
#include <contrib/ydb/core/testlib/common_helper.h>
#include <contrib/ydb/core/formats/arrow/serializer/full.h>

#include <library/cpp/threading/local_executor/local_executor.h>

#include <util/generic/serialized_enum.h>
#include <util/string/printf.h>

namespace NKikimr::NKqp {

using namespace NYdb;
using namespace NYdb::NTable;

Y_UNIT_TEST_SUITE(KqpConstraints) {

    static NKikimrPQ::TPQConfig DefaultPQConfig() {
        NKikimrPQ::TPQConfig pqConfig;
        pqConfig.SetEnabled(true);
        pqConfig.SetEnableProtoSourceIdInfo(true);
        pqConfig.SetTopicsAreFirstClassCitizen(true);
        pqConfig.SetRequireCredentialsInNewProtocol(false);
        pqConfig.AddClientServiceType()->SetName("data-streams");
        return pqConfig;
    }

    Y_UNIT_TEST(CreateTableSerialTypeForbidden) {
        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(false);
        auto serverSettings = TKikimrSettings().SetAppConfig(appConfig);
        TKikimrRunner kikimr(serverSettings);
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/SerialTableDisabled` (
                    Key Serial,
                    Value String,
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::GENERIC_ERROR,
                                       result.GetIssues().ToString());
        }
    }

    Y_UNIT_TEST(AddSerialColumnForbidden) {
        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(false);
        auto serverSettings = TKikimrSettings().SetAppConfig(appConfig);
        TKikimrRunner kikimr(serverSettings);
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/SerialTableCreateAndAlter` (
                    Key Int32,
                    Value String,
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }

        {
            auto query = R"(
                --!syntax_v1
                ALTER TABLE `/Root/SerialTableCreateAndAlter` ADD COLUMN SeqNo Serial;
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::GENERIC_ERROR,
                                       result.GetIssues().ToString());
        }
    }

    Y_UNIT_TEST(CreateTableWithDefaultForbidden) {
        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(false);
        appConfig.MutableTableServiceConfig()->SetEnableColumnsWithDefault(false);
        auto serverSettings = TKikimrSettings().SetAppConfig(appConfig);
        TKikimrRunner kikimr(serverSettings);
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/CreateAndAlterDefaultsDisabled` (
                    Key Int32,
                    Value String Default "empty",
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::GENERIC_ERROR,
                                       result.GetIssues().ToString());
        }
    }

    Y_UNIT_TEST(AddColumnWithDefaultForbidden) {
        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(false);
        appConfig.MutableTableServiceConfig()->SetEnableColumnsWithDefault(false);
        auto serverSettings = TKikimrSettings().SetAppConfig(appConfig);
        TKikimrRunner kikimr(serverSettings);
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/SerialTableCreateAndAlter` (
                    Key Int32,
                    Value String,
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }

        {
            auto query = R"(
                --!syntax_v1
                ALTER TABLE `/Root/SerialTableCreateAndAlter` ADD COLUMN SeqNo DEFAULT 5;
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::GENERIC_ERROR,
                                       result.GetIssues().ToString());
        }
    }

    Y_UNIT_TEST(SerialTypeNegative1) {
        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(true);
        TKikimrRunner kikimr(TKikimrSettings().SetPQConfig(DefaultPQConfig()).SetAppConfig(appConfig));
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/SerialTableNeg1` (
                    Key SerialUnknown,
                    Value String,
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::GENERIC_ERROR,
                                       result.GetIssues().ToString());
        }
    }

    Y_UNIT_TEST(SerialTypeNegative2) {
        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(true);
        TKikimrRunner kikimr(TKikimrSettings().SetPQConfig(DefaultPQConfig()).SetAppConfig(appConfig));
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/SerialTableNeg2` (
                    Key Uint32,
                    Value Serial,
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::BAD_REQUEST,
                                       result.GetIssues().ToString());
        }
    }


    void TestSerialType(TString serialType) {
        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(true);
        TKikimrRunner kikimr(TKikimrSettings().SetPQConfig(DefaultPQConfig()).SetAppConfig(appConfig));
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = Sprintf(R"(
                --!syntax_v1
                CREATE TABLE `/Root/SerialTable%s` (
                    Key %s,
                    Value String,
                    PRIMARY KEY (Key)
                );
            )", serialType.c_str(), serialType.c_str());

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }

        {
            TString query = Sprintf(R"(
                UPSERT INTO `/Root/SerialTable%s` (Value) VALUES ("New");
            )", serialType.c_str());

            NYdb::NTable::TExecDataQuerySettings execSettings;
            execSettings.KeepInQueryCache(true);
            execSettings.CollectQueryStats(ECollectQueryStatsMode::Basic);

            auto result =
                session
                    .ExecuteDataQuery(query, TTxControl::BeginTx().CommitTx(),
                                      execSettings)
                    .ExtractValueSync();

            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }
        {
            TString query = Sprintf(R"(
                SELECT * FROM `/Root/SerialTable%s`;
            )", serialType.c_str());

            NYdb::NTable::TExecDataQuerySettings execSettings;
            execSettings.KeepInQueryCache(true);
            execSettings.CollectQueryStats(ECollectQueryStatsMode::Basic);

            auto result =
                session
                    .ExecuteDataQuery(query, TTxControl::BeginTx().CommitTx(),
                                      execSettings)
                    .ExtractValueSync();

            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
            CompareYson(R"(
                    [
                        [[1];["New"]]
                    ]
                )",
                NYdb::FormatResultSetYson(result.GetResultSet(0)));
        }
    }

    Y_UNIT_TEST(SerialTypeSmallSerial) {
        TestSerialType("SmallSerial");
    }

    Y_UNIT_TEST(SerialTypeSerial2) {
        TestSerialType("serial2");
    }

    Y_UNIT_TEST(SerialTypeSerial) {
        TestSerialType("Serial");
    }

    Y_UNIT_TEST(SerialTypeSerial4) {
        TestSerialType("Serial4");
    }

    Y_UNIT_TEST(SerialTypeBigSerial) {
        TestSerialType("BigSerial");
    }

    Y_UNIT_TEST(SerialTypeSerial8) {
        TestSerialType("serial8");
    }

    Y_UNIT_TEST(AlterTableAddColumnWithDefaultValue) {
        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(false);
        auto serverSettings = TKikimrSettings().SetAppConfig(appConfig);
        TKikimrRunner kikimr(serverSettings);
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/SerialTableCreateAndAlter` (
                    Key Int32,
                    Value String,
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }

        {
            auto query = R"(
                --!syntax_v1
                ALTER TABLE `/Root/SerialTableCreateAndAlter` ADD COLUMN Exists bool DEFAULT false;
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }
    }

    Y_UNIT_TEST(DefaultValuesForTable) {
        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(false);
        appConfig.MutableTableServiceConfig()->SetEnableColumnsWithDefault(true);
        TKikimrRunner kikimr(TKikimrSettings().SetPQConfig(DefaultPQConfig()).SetAppConfig(appConfig));
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/TableWithDefaults` (
                    Key Int32 Default 7,
                    Value String,
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }        

        {
            TString query = R"(
                UPSERT INTO `/Root/TableWithDefaults` (Value) VALUES ("New");
            )";

            NYdb::NTable::TExecDataQuerySettings execSettings;
            execSettings.KeepInQueryCache(true);
            execSettings.CollectQueryStats(ECollectQueryStatsMode::Basic);

            auto result =
                session
                    .ExecuteDataQuery(query, TTxControl::BeginTx().CommitTx(),
                                      execSettings)
                    .ExtractValueSync();

            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }
        {
            TString query = R"(
                SELECT * FROM `/Root/TableWithDefaults`;
            )";

            NYdb::NTable::TExecDataQuerySettings execSettings;
            execSettings.KeepInQueryCache(true);
            execSettings.CollectQueryStats(ECollectQueryStatsMode::Basic);

            auto result =
                session
                    .ExecuteDataQuery(query, TTxControl::BeginTx().CommitTx(),
                                      execSettings)
                    .ExtractValueSync();

            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
            CompareYson(R"(
                    [
                        [[7];["New"]]
                    ]
                )",
                NYdb::FormatResultSetYson(result.GetResultSet(0)));
        }
    }

    Y_UNIT_TEST(DefaultValuesForTableNegative2) {
        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(false);
        appConfig.MutableTableServiceConfig()->SetEnableColumnsWithDefault(true);
        TKikimrRunner kikimr(TKikimrSettings().SetPQConfig(DefaultPQConfig()).SetAppConfig(appConfig));
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/TableWithDefaults2` (
                    Key Int32 Default 1,
                    Value String,
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }        
    }

    Y_UNIT_TEST(DefaultValuesForTableNegative3) {
        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(false);
        appConfig.MutableTableServiceConfig()->SetEnableColumnsWithDefault(true);
        TKikimrRunner kikimr(TKikimrSettings().SetPQConfig(DefaultPQConfig()).SetAppConfig(appConfig));
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/TableWithDefaults` (
                    Key Int32 NOT NULL Default null,
                    Value String,
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::GENERIC_ERROR, result.GetIssues().ToString());
            UNIT_ASSERT_STRING_CONTAINS(result.GetIssues().ToString(),
                                        "Default expr Key is nullable or optional, but column has not null constraint");
        }        
    }

    Y_UNIT_TEST(DefaultValuesForTableNegative4) {

        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(false);
        appConfig.MutableTableServiceConfig()->SetEnableColumnsWithDefault(true);

        TKikimrRunner kikimr(TKikimrSettings().SetPQConfig(DefaultPQConfig()).SetAppConfig(appConfig));
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/TableWithDefaults` (
                    Key Uint32 Default "someText",
                    Value String,
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::GENERIC_ERROR,
                                       result.GetIssues().ToString());
        }        
    }

    Y_UNIT_TEST(AlterTableAddNotNullWithDefault) {

        NKikimrConfig::TAppConfig appConfig;
        appConfig.MutableTableServiceConfig()->SetEnableSequences(false);
        appConfig.MutableTableServiceConfig()->SetEnableColumnsWithDefault(true);

        TKikimrRunner kikimr(TKikimrSettings().SetPQConfig(DefaultPQConfig()).SetAppConfig(appConfig));
        auto db = kikimr.GetTableClient();
        auto session = db.CreateSession().GetValueSync().GetSession();

        {
            auto query = R"(
                --!syntax_v1
                CREATE TABLE `/Root/AlterTableAddNotNullColumn` (
                    Key Uint32,
                    Value String,
                    PRIMARY KEY (Key)
                );
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }

        {
            TString query = R"(
                UPSERT INTO `/Root/AlterTableAddNotNullColumn` (Key, Value) VALUES (1, "Old");
            )";

            NYdb::NTable::TExecDataQuerySettings execSettings;
            execSettings.KeepInQueryCache(true);
            execSettings.CollectQueryStats(ECollectQueryStatsMode::Basic);

            auto result =
                session
                    .ExecuteDataQuery(query, TTxControl::BeginTx().CommitTx(),
                                      execSettings)
                    .ExtractValueSync();

            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }

        {
            auto query = R"(
                --!syntax_v1
                ALTER TABLE `/Root/AlterTableAddNotNullColumn` ADD COLUMN Exists Int32 NOT NULL DEFAULT 1;
            )";

            auto result = session.ExecuteSchemeQuery(query).GetValueSync();
            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }

        Sleep(TDuration::Seconds(3));

        {
            TString query = R"(
                INSERT INTO `/Root/AlterTableAddNotNullColumn` (Key, Value) VALUES (2, "New");
            )";

            NYdb::NTable::TExecDataQuerySettings execSettings;
            execSettings.KeepInQueryCache(true);
            execSettings.CollectQueryStats(ECollectQueryStatsMode::Basic);

            auto result =
                session
                    .ExecuteDataQuery(query, TTxControl::BeginTx().CommitTx(),
                                      execSettings)
                    .ExtractValueSync();

            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());
        }

        {
            TString query = R"(
                SELECT * FROM `/Root/AlterTableAddNotNullColumn` ORDER BY Key;
            )";

            NYdb::NTable::TExecDataQuerySettings execSettings;
            execSettings.KeepInQueryCache(true);
            execSettings.CollectQueryStats(ECollectQueryStatsMode::Basic);

            auto result =
                session
                    .ExecuteDataQuery(query, TTxControl::BeginTx().CommitTx(),
                                      execSettings)
                    .ExtractValueSync();

            UNIT_ASSERT_VALUES_EQUAL_C(result.GetStatus(), EStatus::SUCCESS,
                                       result.GetIssues().ToString());

            Cerr << NYdb::FormatResultSetYson(result.GetResultSet(0)) << Endl;
            CompareYson(R"(
                    [
                        [[1u];["Old"];[1]];[[2u];["New"];[1]]
                    ]
                )",
                NYdb::FormatResultSetYson(result.GetResultSet(0)));
        }

    }

}
} // namespace NKikimr::NKqp