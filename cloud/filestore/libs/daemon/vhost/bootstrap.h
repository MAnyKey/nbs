#pragma once

#include "public.h"

#include <cloud/filestore/libs/client/public.h>
#include <cloud/filestore/libs/daemon/common/bootstrap.h>
#include <cloud/filestore/libs/endpoint/public.h>
#include <cloud/filestore/libs/endpoint_vhost/public.h>
#include <cloud/filestore/libs/server/public.h>
#include <cloud/filestore/libs/service/public.h>
#include <cloud/filestore/libs/vfs/public.h>

#include <cloud/storage/core/libs/common/public.h>
#include <cloud/storage/core/libs/keyring/public.h>

namespace NCloud::NFileStore::NDaemon {

////////////////////////////////////////////////////////////////////////////////

struct TVhostModuleFactories
{
    std::function<NVFS::IFileSystemLoopFactoryPtr(
        ILoggingServicePtr logging,
        ITimerPtr timer,
        ISchedulerPtr scheduler,
        IRequestStatsRegistryPtr requestStats,
        IProfileLogPtr profileLog)> LoopFactory;
};

////////////////////////////////////////////////////////////////////////////////

class TBootstrapVhost final
    : public TBootstrapCommon
{
private:
    const TVhostModuleFactoriesPtr VhostModuleFactories;

    TConfigInitializerVhostPtr Configs;
    IFileStoreEndpointsPtr FileStoreEndpoints;
    IEndpointListenerPtr EndpointListener;
    IEndpointStoragePtr EndpointStorage;
    IEndpointManagerPtr EndpointManager;

    NServer::IServerPtr Server;

public:
    TBootstrapVhost(
        std::shared_ptr<NKikimr::TModuleFactories> kikimrFactories,
        TVhostModuleFactoriesPtr vhostFactories);
    ~TBootstrapVhost();

private:
    TConfigInitializerCommonPtr InitConfigs(int argc, char** argv) override;

    void InitComponents() override;
    void StartComponents() override;
    void StopComponents() override;

private:
    void InitLWTrace();
    void InitConfig();
    void InitEndpoints();
    void InitNullEndpoints();
    void RestoreKeyringEndpoints();
};

}   // namespace NCloud::NFileStore::NServer
