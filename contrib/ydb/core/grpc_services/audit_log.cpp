#include "defs.h"

#include <contrib/ydb/core/util/address_classifier.h>
#include <contrib/ydb/core/audit/audit_log.h>

#include "base/base.h"
#include "audit_log.h"

namespace NKikimr {
namespace NGRpcService {

void AuditLogConn(const IRequestProxyCtx* ctx, const TString& database, const TString& userSID)
{
    static const TString GrpcConnComponentName = "grpc-conn";

    AUDIT_LOG(
        AUDIT_PART("component", GrpcConnComponentName)

        AUDIT_PART("remote_address", NKikimr::NAddressClassifier::ExtractAddress(ctx->GetPeerName()))
        AUDIT_PART("subject", userSID)
        AUDIT_PART("database", database)
        AUDIT_PART("operation", ctx->GetRequestName())
    );

    // and transitional, to be removed, output to the common log
    LOG_NOTICE_S(TlsActivationContext->AsActorContext(), NKikimrServices::GRPC_SERVER, "AUDIT: "
        << "request name: " << ctx->GetRequestName()
        << ", database: " << database
        << ", peer: " << ctx->GetPeerName()
        << ", subject: " << (userSID ? userSID : "no subject")
    );
}

void AuditLog(ui32 status, const TAuditLogParts& parts)
{
    static const TString GrpcProxyComponentName = "grpc-proxy";

    //NOTE: EmptyValue couldn't be an empty string as AUDIT_PART() skips parts with an empty values
    static const TString EmptyValue = "{none}";

    AUDIT_LOG(
        AUDIT_PART("component", GrpcProxyComponentName)

        // all parts are considered as required, so all empty values gets replaced with a special stub
        for (const auto& [name, value] : parts) {
            AUDIT_PART(name, (!value.empty() ? value : EmptyValue))
        }

        AUDIT_PART("status", (status == Ydb::StatusIds::SUCCESS ? TString("SUCCESS") : TString("ERROR")))
        AUDIT_PART("detailed_status", (Ydb::StatusIds::StatusCode_IsValid(status)
            ? TString(Ydb::StatusIds::StatusCode_Name(status))
            : ToString(status)
        ))
    );
}

}
}

