LIBRARY()

SRCS(
    actor.cpp
    counters.cpp
    task.cpp
    events.cpp
)

PEERDIR(
    contrib/ydb/core/protos
    contrib/ydb/library/actors/core
    contrib/ydb/core/tablet_flat
)

END()
