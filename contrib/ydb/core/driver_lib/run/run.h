#pragma once
#include "config.h"
#include "factories.h"
#include "service_initializer.h"

#include <library/cpp/actors/core/actorsystem.h>
#include <library/cpp/actors/core/log_settings.h>
#include <library/cpp/actors/interconnect/poller_tcp.h>
#include <library/cpp/actors/util/should_continue.h>
#include <contrib/ydb/library/grpc/server/grpc_server.h>
#include <contrib/ydb/core/base/appdata.h>
#include <contrib/ydb/core/base/statestorage.h>
#include <contrib/ydb/core/scheme_types/scheme_types.h>
#include <contrib/ydb/core/formats/factory.h>
#include <contrib/ydb/public/lib/base/msgbus.h>
#include <contrib/ydb/core/client/server/grpc_server.h>
#include <contrib/ydb/core/fq/libs/shared_resources/interface/shared_resources.h>
#include <contrib/ydb/core/kqp/common/kqp.h>
#include <contrib/ydb/core/base/memobserver.h>
#include <contrib/ydb/core/tablet/node_tablet_monitor.h>
#include <contrib/ydb/core/tablet/tablet_setup.h>
#include <contrib/ydb/core/ymq/http/http.h>

#include <library/cpp/monlib/dynamic_counters/counters.h>

namespace NKikimr {

class TKikimrRunner : public virtual TThrRefBase, private IGlobalObjectStorage {
protected:
    static TProgramShouldContinue KikimrShouldContinue;
    static void OnTerminate(int);

    std::shared_ptr<TModuleFactories> ModuleFactories;
    TIntrusivePtr<NScheme::TTypeRegistry> TypeRegistry;
    TIntrusivePtr<NMiniKQL::IMutableFunctionRegistry> FunctionRegistry;
    TIntrusivePtr<TFormatFactory> FormatFactory;
    NFq::IYqSharedResources::TPtr YqSharedResources;

    TAutoPtr<TMon> Monitoring;
    ::NMonitoring::TDynamicCounterPtr Counters;
    NMonitoring::TIndexMonPage *ActorsMonPage = nullptr;

    bool EnabledGrpcService = false;
    bool GracefulShutdownSupported = false;
    THolder<NSQS::TAsyncHttpServer> SqsHttp;

    THolder<NYdb::TDriver> YdbDriver;

    std::vector<std::shared_ptr<void>> GlobalObjects;

    THolder<NKqp::TKqpShutdownController> KqpShutdownController;

    TIntrusivePtr<NInterconnect::TPollerThreads> PollerThreads;
    TAutoPtr<TAppData> AppData;

    TVector<std::pair<TString, TAutoPtr<NYdbGrpc::TGRpcServer>>> GRpcServers;

    TIntrusivePtr<NActors::NLog::TSettings> LogSettings;
    std::shared_ptr<TLogBackend> LogBackend;
    TAutoPtr<TActorSystem> ActorSystem;

    TIntrusivePtr<TMemObserver> MemObserver;

    TKikimrRunner(std::shared_ptr<TModuleFactories> factories = {});

    virtual ~TKikimrRunner();

    virtual void InitializeRegistries(const TKikimrRunConfig& runConfig);

    void InitializeAllocator(const TKikimrRunConfig& runConfig);

    void InitializeLogSettings(const TKikimrRunConfig& runConfig);

    void ApplyLogSettings(const TKikimrRunConfig& runConfig);

    void InitializeMonitoring(const TKikimrRunConfig& runConfig, bool includeHostName = true);

    void InitializeControlBoard(const TKikimrRunConfig& runConfig);

    void InitializeMonitoringLogin(const TKikimrRunConfig& runConfig);

    void InitializeGRpc(const TKikimrRunConfig& runConfig);

    void InitializeKqpController(const TKikimrRunConfig& runConfig);

    void InitializeGracefulShutdown(const TKikimrRunConfig& runConfig);

    void InitializeAppData(const TKikimrRunConfig& runConfig);

    void InitializeActorSystem(
        const TKikimrRunConfig& runConfig,
        TIntrusivePtr<TServiceInitializersList> serviceInitializers,
        const TBasicKikimrServicesMask& serviceMask = {});

    TIntrusivePtr<TServiceInitializersList> CreateServiceInitializersList(
        const TKikimrRunConfig& runConfig,
        const TBasicKikimrServicesMask& serviceMask = {});

private:
    void AddGlobalObject(std::shared_ptr<void> object) override;

public:
    static void SetSignalHandlers();

    virtual void KikimrStart();
    virtual void BusyLoop();
    virtual void KikimrStop(bool graceful);

    static TIntrusivePtr<TKikimrRunner> CreateKikimrRunner(
            const TKikimrRunConfig& runConfig,
            std::shared_ptr<TModuleFactories> factories);
};

int MainRun(const TKikimrRunConfig &runConfig, std::shared_ptr<TModuleFactories> factories);

}
