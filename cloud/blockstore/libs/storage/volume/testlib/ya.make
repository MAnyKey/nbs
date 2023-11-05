LIBRARY()

SRCS(
    test_env.cpp
)

PEERDIR(
    cloud/blockstore/libs/storage/api
    cloud/blockstore/libs/storage/core
    cloud/blockstore/libs/storage/model
    cloud/blockstore/libs/storage/stats_service
    cloud/blockstore/libs/storage/testlib
    cloud/blockstore/libs/storage/volume

    cloud/storage/core/libs/api
    cloud/storage/core/libs/diagnostics

    library/cpp/lwtrace/mon
    library/cpp/testing/unittest

    cloud/blockstore/libs/storage/partition
    cloud/blockstore/libs/storage/partition_nonrepl
    cloud/blockstore/libs/storage/service

    contrib/ydb/core/blockstore
    contrib/ydb/core/mind
    contrib/ydb/core/testlib
    contrib/ydb/core/testlib/basics
)

END()

