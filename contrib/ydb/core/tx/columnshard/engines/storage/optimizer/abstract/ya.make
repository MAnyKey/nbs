LIBRARY()

SRCS(
    optimizer.cpp
)

PEERDIR(
    contrib/libs/apache/arrow
    contrib/ydb/core/protos
    contrib/ydb/core/formats/arrow
    contrib/ydb/core/tx/columnshard/engines/changes/abstract
)

END()
