#pragma once

#include "feature_flags.h"

#include <contrib/ydb/core/base/appdata.h>
#include <contrib/ydb/core/base/channel_profiles.h>
#include <contrib/ydb/core/base/domain.h>
#include <contrib/ydb/core/formats/factory.h>
#include <contrib/ydb/core/scheme/scheme_type_registry.h>
#include <contrib/ydb/core/testlib/actors/test_runtime.h>
#include <contrib/ydb/core/tx/datashard/export_iface.h>
#include <contrib/ydb/core/tx/datashard/export_s3.h>

namespace NKikimr {

    // FIXME
    // Split this factory
    class TDataShardExportFactory : public NDataShard::IExportFactory {
        using IExport = NDataShard::IExport;

    public:
        IExport* CreateExportToYt(
                const IExport::TTask& task, const IExport::TTableColumns& columns) const override
        {
            Y_UNUSED(task);
            Y_UNUSED(columns);
            return nullptr;
        }

        IExport* CreateExportToS3(
                const IExport::TTask& task, const IExport::TTableColumns& columns) const override
        {
        #ifndef KIKIMR_DISABLE_S3_OPS
            return new NDataShard::TS3Export(task, columns);
        #else
            Y_UNUSED(task);
            Y_UNUSED(columns);
            return nullptr;
        #endif
        }

        void Shutdown() override {
        }
    };

    struct TAppPrepare : public TTestFeatureFlagsHolder<TAppPrepare> {
        struct TMine : public NActors::IDestructable {
            TIntrusivePtr<NScheme::TTypeRegistry> Types;
            TIntrusivePtr<NMiniKQL::IFunctionRegistry> Funcs;
            TIntrusivePtr<TFormatFactory> Formats;
            std::shared_ptr<NDataShard::IExportFactory> DataShardExportFactory;
            std::shared_ptr<NPDisk::IIoContextFactory> IoContext;

            ~TMine();
        };

        using TFnReg = std::function<NMiniKQL::IFunctionRegistry*(const NScheme::TTypeRegistry&)>;

        TAppPrepare(std::shared_ptr<NKikimr::NDataShard::IExportFactory> ef = {});

        NActors::TTestActorRuntime::TEgg Unwrap() noexcept;

        void AddDomain(TDomainsInfo::TDomain* domain);
        void AddHive(ui32 hiveUid, ui64 hive);
        void ClearDomainsAndHive();
        void SetChannels(TIntrusivePtr<TChannelProfiles> channels);
        void SetBSConf(NKikimrBlobStorage::TNodeWardenServiceSet config);
        void SetFnRegistry(TFnReg func);
        void SetFormatsFactory(TIntrusivePtr<TFormatFactory> formats);
        void SetKeyForNode(const TString& path, ui32 node);
        void SetEnableKqpSpilling(bool value);
        void SetNetDataSourceUrl(const TString& value);
        void SetKeepSnapshotTimeout(TDuration value);
        void SetChangesQueueItemsLimit(ui64 value);
        void SetChangesQueueBytesLimit(ui64 value);
        void SetMinRequestSequenceSize(ui64 value);
        void SetRequestSequenceSize(ui64 value);
        void SetHiveStoragePoolFreshPeriod(ui64 value);
        void AddSystemBackupSID(const TString& sid);
        void SetEnableProtoSourceIdInfo(std::optional<bool> value);
        void SetEnablePqBilling(std::optional<bool> value);
        void SetEnableDbCounters(bool value);
        void SetAwsRegion(const TString& value);

        TIntrusivePtr<TChannelProfiles> Channels;
        NKikimrBlobStorage::TNodeWardenServiceSet BSConf;
        TIntrusivePtr<TDomainsInfo> Domains;
        TMap<ui32, NKikimrProto::TKeyConfig> Keys;
        bool EnableKqpSpilling = false;
        NKikimrConfig::TCompactionConfig CompactionConfig;
        TString NetDataSourceUrl;
        NKikimrConfig::THiveConfig HiveConfig;
        NKikimrConfig::TDataShardConfig DataShardConfig;
        NKikimrConfig::TColumnShardConfig ColumnShardConfig;
        NKikimrConfig::TSchemeShardConfig SchemeShardConfig;
        NKikimrConfig::TMeteringConfig MeteringConfig;
        NKikimrPQ::TPQConfig PQConfig;
        NKikimrConfig::TAwsCompatibilityConfig AwsCompatibilityConfig;

    private:
        TAutoPtr<TMine> Mine;
    };

}
