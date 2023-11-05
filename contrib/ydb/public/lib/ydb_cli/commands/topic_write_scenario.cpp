#include "topic_write_scenario.h"

#include <contrib/ydb/public/lib/ydb_cli/commands/topic_workload/topic_workload_defines.h>
#include <contrib/ydb/public/lib/ydb_cli/commands/topic_workload/topic_workload_describe.h>
#include <contrib/ydb/public/lib/ydb_cli/commands/topic_workload/topic_workload_stats_collector.h>
#include <contrib/ydb/public/lib/ydb_cli/commands/topic_workload/topic_workload_writer.h>

#include <util/random/random.h>

namespace NYdb::NConsoleClient {

int TTopicWriteScenario::DoRun(const TClientCommand::TConfig& config)
{
    auto describeTopicResult = TCommandWorkloadTopicDescribe::DescribeTopic(config.Database, TopicName, *Driver);
    ui32 partitionCount = describeTopicResult.GetTotalPartitionsCount();
    ui32 partitionSeed = RandomNumber<ui32>(partitionCount);

    std::vector<TString> generatedMessages =
        TTopicWorkloadWriterWorker::GenerateMessages(MessageSize);

    std::vector<std::future<void>> threads;

    StartConsumerThreads(threads, config.Database);
    StartProducerThreads(threads, partitionCount, partitionSeed, generatedMessages);

    StatsCollector->PrintWindowStatsLoop();

    JoinThreads(threads);

    StatsCollector->PrintTotalStats();

    if (AnyErrors() || !AnyOutgoingMessages()) {
        return EXIT_FAILURE;
    }

    return EXIT_SUCCESS;
}

}
