#include "part2_actor.h"

#include "part2_counters.h"

#include <cloud/blockstore/libs/storage/core/probes.h>

namespace NCloud::NBlockStore::NStorage::NPartition2 {

using namespace NActors;

LWTRACE_USING(BLOCKSTORE_STORAGE_PROVIDER);

namespace {

////////////////////////////////////////////////////////////////////////////////

#define COPY_FIELD(l, r, name)      l.Set##name(r.Get##name());

template <typename T1, typename T2>
void CopyPartitionConfig(T1& l, const T2& r)
{
    COPY_FIELD(l, r, InstanceId);
    COPY_FIELD(l, r, ChannelsCount);
}

template <typename T1, typename T2>
void CopyPartitionStats(T1& l, const T2& r)
{
    COPY_FIELD(l, r, MixedBlobsCount);
    COPY_FIELD(l, r, MixedBlocksCount);
    COPY_FIELD(l, r, MergedBlobsCount);
    COPY_FIELD(l, r, MergedBlocksCount);
    COPY_FIELD(l, r, GarbageBlocksCount);
    COPY_FIELD(l, r, UsedBlocksCount);
    COPY_FIELD(l, r, LogicalUsedBlocksCount);
    COPY_FIELD(l, r, CheckpointBlocksCount);
}

#undef COPY_FIELD

}   // namespace

////////////////////////////////////////////////////////////////////////////////

void TPartitionActor::HandleStatPartition(
    const TEvPartition::TEvStatPartitionRequest::TPtr& ev,
    const TActorContext& ctx)
{
    auto* msg = ev->Get();

    auto requestInfo = CreateRequestInfo(
        ev->Sender,
        ev->Cookie,
        msg->CallContext);

    LWTRACK(
        RequestReceived_Partition,
        requestInfo->CallContext->LWOrbit,
        "StatPartition",
        requestInfo->CallContext->RequestId);

    auto response = std::make_unique<TEvPartition::TEvStatPartitionResponse>();

    CopyPartitionConfig(
        *response->Record.MutableVolume(),
        State->GetConfig());

    CopyPartitionStats(
        *response->Record.MutableStats(),
        State->GetStats());

    UpdatePartitionCounters(
        *response->Record.MutableStats(),
        State->GetStats());

    response->Record.MutableStats()->SetFreshBlocksCount(
        State->GetFreshBlockCount());

    response->Record.MutableStats()->SetNonEmptyRangeCount(
        State->GetCompactionMap().GetNonEmptyRangeCount());

    LWTRACK(
        ResponseSent_Partition,
        requestInfo->CallContext->LWOrbit,
        "StatPartition",
        requestInfo->CallContext->RequestId);

    NCloud::Reply(ctx, *requestInfo, std::move(response));
}

}   // namespace NCloud::NBlockStore::NStorage::NPartition2
