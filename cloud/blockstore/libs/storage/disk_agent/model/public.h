#pragma once

#include <util/generic/strbuf.h>

#include <memory>

namespace NCloud::NBlockStore::NStorage {

////////////////////////////////////////////////////////////////////////////////

class TDeviceClient;
using TDeviceClientPtr = std::shared_ptr<TDeviceClient>;

class TDiskAgentConfig;
using TDiskAgentConfigPtr = std::shared_ptr<TDiskAgentConfig>;

constexpr TStringBuf BackgroundOpsClientId = "migration";
constexpr TStringBuf AnyWriterClientId = "any-writer";

}   // namespace NCloud::NBlockStore::NStorage
