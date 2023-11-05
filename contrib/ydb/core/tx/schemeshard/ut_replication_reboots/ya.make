UNITTEST_FOR(contrib/ydb/core/tx/schemeshard)

FORK_SUBTESTS()

SPLIT_FACTOR(20)

IF (SANITIZER_TYPE OR WITH_VALGRIND)
    TIMEOUT(3600)
    SIZE(LARGE)
    TAG(ya:fat)
ELSE()
    TIMEOUT(600)
    SIZE(MEDIUM)
ENDIF()

PEERDIR(
    contrib/ydb/core/tx/schemeshard/ut_helpers
    contrib/ydb/library/yql/sql/pg_dummy
)

SRCS(
    ut_replication_reboots.cpp
)

YQL_LAST_ABI_VERSION()

END()
