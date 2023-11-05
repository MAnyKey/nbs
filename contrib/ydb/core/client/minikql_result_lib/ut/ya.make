UNITTEST_FOR(contrib/ydb/core/client/minikql_result_lib)

FORK_SUBTESTS()

TIMEOUT(300)

SIZE(MEDIUM)

SRCS(
    converter_ut.cpp
    objects_ut.cpp
)

PEERDIR(
    library/cpp/testing/unittest
    contrib/ydb/core/testlib/default
)

YQL_LAST_ABI_VERSION()

REQUIREMENTS(network:full ram:13)

END()
