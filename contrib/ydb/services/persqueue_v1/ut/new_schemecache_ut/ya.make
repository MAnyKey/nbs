UNITTEST_FOR(contrib/ydb/services/persqueue_v1)

FORK_SUBTESTS()

IF (WITH_VALGRIND)
    TIMEOUT(1800)
    SIZE(LARGE)
    TAG(ya:fat)
    REQUIREMENTS(ram:32)
ELSE()
    TIMEOUT(600)
    SIZE(MEDIUM)
ENDIF()

SRCS(
    persqueue_new_schemecache_ut.cpp
    persqueue_common_new_schemecache_ut.cpp
    ut/api_test_setup.h
    ut/test_utils.h
    ut/pq_data_writer.h
    ut/rate_limiter_test_setup.h
    ut/rate_limiter_test_setup.cpp
)

PEERDIR(
    library/cpp/getopt
    library/cpp/svnversion
    contrib/ydb/library/persqueue/tests
    contrib/ydb/core/testlib/default
    contrib/ydb/public/api/grpc
    contrib/ydb/public/sdk/cpp/client/resources
    contrib/ydb/public/sdk/cpp/client/ydb_persqueue_core/ut/ut_utils
    contrib/ydb/public/sdk/cpp/client/ydb_table
    contrib/ydb/services/persqueue_v1
)

YQL_LAST_ABI_VERSION()

END()
