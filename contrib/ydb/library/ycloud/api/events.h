#pragma once
#include <util/generic/ptr.h>
#include <util/generic/string.h>
#include <contrib/ydb/library/actors/core/actor.h>
#include <contrib/ydb/library/actors/core/event_local.h>
#include <contrib/ydb/library/grpc/client/grpc_client_low.h>

namespace NCloud {

template <typename TEv, ui32 TEventType, typename TProtoMessage>
struct TEvGrpcProtoRequest : NActors::TEventLocal<TEv, TEventType> {
    TProtoMessage Request;
    TString Token;
    TString RequestId;
};

template <typename TEv, ui32 TEventType, typename TProtoMessage>
struct TEvGrpcProtoResponse : NActors::TEventLocal<TEv, TEventType> {
    THolder<NActors::IEventHandle> Request;
    TProtoMessage Response;
    NYdbGrpc::TGrpcStatus Status;
};

}
