YQL_UDF_CONTRIB(ip_udf)

YQL_ABI_VERSION(
    2
    28
    0
)

SRCS(
    ip_base.cpp
)

PEERDIR(
    contrib/ydb/library/yql/udfs/common/ip_base/lib
)

END()

RECURSE_FOR_TESTS(
    test
)