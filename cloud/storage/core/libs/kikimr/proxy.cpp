#include "proxy.h"

#include "actorsystem.h"

#include <library/cpp/monlib/dynamic_counters/counters.h>
#include <library/cpp/monlib/service/pages/mon_page.h>

namespace NCloud::NStorage {

using namespace NMonitoring;

////////////////////////////////////////////////////////////////////////////////

TLog TLoggingProxy::CreateLog(const TString& component)
{
    Y_ABORT_UNLESS(ActorSystem);
    return ActorSystem->CreateLog(component);
}

////////////////////////////////////////////////////////////////////////////////

NMonitoring::IMonPagePtr TMonitoringProxy::RegisterIndexPage(
    const TString& path,
    const TString& title)
{
    Y_ABORT_UNLESS(ActorSystem);
    return ActorSystem->RegisterIndexPage(path, title);
}

void TMonitoringProxy::RegisterMonPage(IMonPagePtr page)
{
    Y_ABORT_UNLESS(ActorSystem);
    ActorSystem->RegisterMonPage(std::move(page));
}

IMonPagePtr TMonitoringProxy::GetMonPage(const TString& path)
{
    Y_ABORT_UNLESS(ActorSystem);
    return ActorSystem->GetMonPage(path);
}

TDynamicCountersPtr TMonitoringProxy::GetCounters()
{
    Y_ABORT_UNLESS(ActorSystem);
    return ActorSystem->GetCounters();
}

}   // namespace NCloud::NStorage
