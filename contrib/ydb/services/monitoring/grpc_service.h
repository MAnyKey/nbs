#pragma once

#include <contrib/ydb/library/actors/core/actorsystem.h>

#include <contrib/ydb/public/api/grpc/ydb_monitoring_v1.grpc.pb.h>

#include <contrib/ydb/library/grpc/server/grpc_server.h>
#include <contrib/ydb/core/grpc_services/base/base_service.h>

namespace NKikimr {
namespace NGRpcService {

 class TGRpcMonitoringService : public TGrpcServiceBase<Ydb::Monitoring::V1::MonitoringService>
 {
 public:
     using TGrpcServiceBase<Ydb::Monitoring::V1::MonitoringService>::TGrpcServiceBase;
 private:
     void SetupIncomingRequests(NYdbGrpc::TLoggerPtr logger);
 };

} // namespace NGRpcService
} // namespace NKikimr
