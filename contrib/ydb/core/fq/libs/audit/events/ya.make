LIBRARY()

SRCS(
    events.cpp
)

PEERDIR(
    library/cpp/actors/core
    library/cpp/actors/interconnect
    contrib/ydb/core/fq/libs/checkpointing_common
    contrib/ydb/core/fq/libs/control_plane_storage/events
    contrib/ydb/library/yql/public/issue
    contrib/ydb/public/api/protos
)

END()
