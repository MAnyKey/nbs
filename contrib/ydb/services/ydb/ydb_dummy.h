#pragma once

#include <contrib/ydb/library/grpc/server/grpc_server.h>
#include <library/cpp/actors/core/actorsystem.h>
#include <library/cpp/monlib/counters/counters.h>
#include <contrib/ydb/public/api/grpc/draft/dummy.grpc.pb.h>

namespace NKikimr {
namespace NGRpcService {

class TGRpcYdbDummyService
    : public NYdbGrpc::TGrpcServiceBase<Draft::Dummy::DummyService>
{
public:
    TGRpcYdbDummyService(NActors::TActorSystem* system, TIntrusivePtr<::NMonitoring::TDynamicCounters> counters, NActors::TActorId proxyActorId);

    void InitService(grpc::ServerCompletionQueue* cq, NYdbGrpc::TLoggerPtr logger) override;
    void SetGlobalLimiterHandle(NYdbGrpc::TGlobalLimiter* limiter) override;

    bool IncRequest();
    void DecRequest();

private:
    void SetupIncomingRequests(NYdbGrpc::TLoggerPtr logger);

    NActors::TActorSystem* ActorSystem_;
    grpc::ServerCompletionQueue* CQ_ = nullptr;

    TIntrusivePtr<::NMonitoring::TDynamicCounters> Counters_;
    NActors::TActorId GRpcRequestProxyId_;
    NYdbGrpc::TGlobalLimiter* Limiter_ = nullptr;
};

}
}
