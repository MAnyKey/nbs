#pragma once

#include <contrib/ydb/core/testlib/actors/test_runtime.h>
#include <contrib/ydb/core/fq/libs/shared_resources/interface/shared_resources.h>

#include <library/cpp/actors/core/actorsystem.h>

namespace NFq {

NActors::IActor* CreateYqlMockActor(int grpcPort);
void InitTest(NActors::TTestActorRuntime* runtime, int httpPort, int grpcPort, const IYqSharedResources::TPtr& yqSharedResources);

} // namespace NFq
