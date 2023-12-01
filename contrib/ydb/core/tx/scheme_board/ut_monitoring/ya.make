UNITTEST_FOR(contrib/ydb/core/tx/scheme_board)

FORK_SUBTESTS()

SIZE(MEDIUM)

TIMEOUT(600)

PEERDIR(
    contrib/ydb/library/actors/core
    library/cpp/testing/unittest
    contrib/ydb/core/testlib/basics/default
)

YQL_LAST_ABI_VERSION()

SRCS(
    monitoring_ut.cpp
    ut_helpers.cpp
)

END()
