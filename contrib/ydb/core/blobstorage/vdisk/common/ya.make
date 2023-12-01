LIBRARY()

PEERDIR(
    contrib/ydb/library/actors/protos
    library/cpp/monlib/service/pages
    contrib/ydb/core/base
    contrib/ydb/core/blobstorage/vdisk/hulldb/base
    contrib/ydb/core/blobstorage/vdisk/protos
    contrib/ydb/core/protos
)

SRCS(
    align.h
    blobstorage_dblogcutter.cpp
    blobstorage_dblogcutter.h
    blobstorage_event_filter.cpp
    blobstorage_event_filter.h
    blobstorage_status.cpp
    blobstorage_status.h
    blobstorage_vdisk_guids.cpp
    blobstorage_vdisk_guids.h
    defs.h
    disk_part.h
    sublog.h
    vdisk_config.cpp
    vdisk_config.h
    vdisk_context.cpp
    vdisk_context.h
    vdisk_costmodel.cpp
    vdisk_costmodel.h
    vdisk_dbtype.h
    vdisk_defrag.h
    vdisk_events.cpp
    vdisk_events.h
    vdisk_handle_class.cpp
    vdisk_handle_class.h
    vdisk_histogram_latency.cpp
    vdisk_histogram_latency.h
    vdisk_histograms.cpp
    vdisk_histograms.h
    vdisk_hugeblobctx.cpp
    vdisk_hugeblobctx.h
    vdisk_hulllogctx.h
    vdisk_log.cpp
    vdisk_log.h
    vdisk_lsnmngr.h
    vdisk_mon.h
    vdisk_mongroups.h
    vdisk_outofspace.cpp
    vdisk_outofspace.h
    vdisk_pdisk_error.h
    vdisk_pdiskctx.h
    vdisk_private_events.h
    vdisk_queues.h
    vdisk_recoverylogwriter.cpp
    vdisk_recoverylogwriter.h
    vdisk_response.cpp
    vdisk_response.h
    vdisk_syncneighbors.h
)

END()

RECURSE_FOR_TESTS(
    ut
)
