LIBRARY()

SRCS(
    merge_context.cpp
    column_cursor.cpp
    column_portion_chunk.cpp
    merged_column.cpp
)

PEERDIR(
    contrib/ydb/core/tx/tiering
)

END()
