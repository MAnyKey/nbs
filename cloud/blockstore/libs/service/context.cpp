#include "context.h"

#include <util/datetime/cputimer.h>

namespace NCloud::NBlockStore {

////////////////////////////////////////////////////////////////////////////////

TCallContext::TCallContext(ui64 requestId)
    : TCallContextBase(requestId)
{}

TRequestTime TCallContext::CalcRequestTime(ui64 nowCycles) const
{
    const ui64 startCycles = GetRequestStartedCycles();
    if (!startCycles || startCycles >= nowCycles) {
        return TRequestTime {
            .TotalTime = TDuration::Zero(),
            .ExecutionTime = TDuration::Zero(),
        };
    }

    TRequestTime requestTime;
    requestTime.TotalTime = CyclesToDurationSafe(nowCycles - startCycles);

    const ui64 postponeStart = AtomicGet(PostponeTsCycles);
    if (postponeStart && startCycles < postponeStart && postponeStart < nowCycles) {
        nowCycles = postponeStart;
    }

    const auto postponeDuration = Time(EProcessingStage::Postponed);
    const auto backoffTime = Time(EProcessingStage::Backoff);

    auto responseSentCycles = GetResponseSentCycles();
    auto responseDuration = CyclesToDurationSafe(
        (responseSentCycles ? responseSentCycles : nowCycles) - startCycles);

    requestTime.ExecutionTime = responseDuration - postponeDuration -
        backoffTime - GetPossiblePostponeDuration();

    return requestTime;
}

TDuration TCallContext::GetPossiblePostponeDuration() const
{
    return TDuration::MicroSeconds(AtomicGet(PossiblePostponeMicroSeconds));
}

void TCallContext::SetPossiblePostponeDuration(TDuration d)
{
    AtomicSet(PossiblePostponeMicroSeconds, d.MicroSeconds());
}

bool TCallContext::GetSilenceRetriableErrors() const
{
    return AtomicGet(SilenceRetriableErrors);
}

void TCallContext::SetSilenceRetriableErrors(bool silence)
{
    AtomicSet(SilenceRetriableErrors, silence);
}

bool TCallContext::GetHasUncountableRejects() const
{
    return AtomicGet(HasUncountableRejects);
}

void TCallContext::SetHasUncountableRejects()
{
    AtomicSet(HasUncountableRejects, true);
}

}   // namespace NCloud::NBlockStore
