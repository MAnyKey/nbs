#pragma once

#include <contrib/ydb/core/mon/mon.h>

#include <contrib/ydb/core/fq/libs/compute/common/run_actor_params.h>
#include <contrib/ydb/core/fq/libs/config/protos/pinger.pb.h>
#include <contrib/ydb/core/fq/libs/events/events.h>
#include <contrib/ydb/core/fq/libs/private_client/private_client.h>
#include <contrib/ydb/core/fq/libs/shared_resources/shared_resources.h>
#include <contrib/ydb/core/fq/libs/signer/signer.h>

#include <contrib/ydb/library/yql/minikql/computation/mkql_computation_node.h>
#include <contrib/ydb/library/yql/providers/dq/provider/yql_dq_gateway.h>
#include <contrib/ydb/library/yql/providers/dq/worker_manager/interface/counters.h>
#include <contrib/ydb/library/yql/providers/dq/actors/proto_builder.h>
#include <contrib/ydb/library/yql/providers/common/http_gateway/yql_http_gateway.h>
#include <contrib/ydb/library/yql/providers/common/metrics/service_counters.h>
#include <contrib/ydb/library/yql/providers/pq/cm_client/client.h>

#include <contrib/ydb/public/lib/fq/scope.h>

#include <contrib/ydb/library/actors/core/actorsystem.h>
#include <library/cpp/random_provider/random_provider.h>
#include <library/cpp/time_provider/time_provider.h>

#include <util/datetime/base.h>

namespace NKikimr  {
    namespace NMiniKQL {
        class IFunctionRegistry;
    }
}

namespace NFq {

NActors::TActorId MakeYqlAnalyticsHttpProxyId();
NActors::TActorId MakePendingFetcherId(ui32 nodeId);

NActors::IActor* CreatePendingFetcher(
    const NFq::TYqSharedResources::TPtr& yqSharedResources,
    const NKikimr::TYdbCredentialsProviderFactory& credentialsProviderFactory,
    const ::NFq::NConfig::TConfig& config,
    const NKikimr::NMiniKQL::IFunctionRegistry* functionRegistry,
    TIntrusivePtr<ITimeProvider> timeProvider,
    TIntrusivePtr<IRandomProvider> randomProvider,
    NKikimr::NMiniKQL::TComputationNodeFactory dqCompFactory,
    const ::NYql::NCommon::TServiceCounters& serviceCounters,
    NYql::ISecuredServiceAccountCredentialsFactory::TPtr credentialsFactory,
    NYql::IHTTPGateway::TPtr s3Gateway,
    NYql::NConnector::IClient::TPtr clientConnector,
    ::NPq::NConfigurationManager::IConnections::TPtr pqCmConnections,
    const ::NMonitoring::TDynamicCounterPtr& clientCounters,
    const TString& tenantName,
    NActors::TMon* monitoring
    );

NActors::IActor* CreateRunActor(
    const NActors::TActorId& fetcherId,
    const ::NYql::NCommon::TServiceCounters& serviceCounters,
    TRunActorParams&& params);

struct TResultId {
    TString Id;
    int SetId;
    TString HistoryId;
    TString Owner;
    TString CloudId;
};

NActors::IActor* CreateResultWriter(
    const NActors::TActorId& executerId,
    const TString& resultType,
    const TResultId& resultId,
    const TVector<TString>& columns,
    const TString& traceId,
    const TInstant& deadline,
    ui64 resultBytesLimit);

NActors::IActor* CreateRateLimiterResourceCreator(
    const NActors::TActorId& parent,
    const TString& ownerId,
    const TString& queryId,
    const NYdb::NFq::TScope& scope,
    const TString& tenant);

NActors::IActor* CreateRateLimiterResourceDeleter(
    const NActors::TActorId& parent,
    const TString& ownerId,
    const TString& queryId,
    const NYdb::NFq::TScope& scope,
    const TString& tenant);

TString MakeInternalError(const TString& text);

} /* NFq */
