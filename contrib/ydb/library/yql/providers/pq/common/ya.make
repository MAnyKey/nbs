LIBRARY()

SRCS(
    pq_meta_fields.cpp
    yql_names.cpp
)

PEERDIR(
    contrib/ydb/library/yql/public/types
)

YQL_LAST_ABI_VERSION()

END()
