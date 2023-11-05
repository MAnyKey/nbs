UNITTEST_FOR(contrib/ydb/core/tx/tiering)

FORK_SUBTESTS()

SPLIT_FACTOR(60)

IF (SANITIZER_TYPE == "thread" OR WITH_VALGRIND)
    TIMEOUT(3600)
    SIZE(LARGE)
    TAG(ya:fat)
    REQUIREMENTS(ram:16)
ELSE()
    TIMEOUT(600)
    SIZE(MEDIUM)
ENDIF()

PEERDIR(
    library/cpp/getopt
    library/cpp/regex/pcre
    library/cpp/svnversion
    contrib/ydb/core/testlib/default
    contrib/ydb/services/metadata
    contrib/ydb/core/tx
    contrib/ydb/core/tx/columnshard
    contrib/ydb/public/lib/yson_value
)

YQL_LAST_ABI_VERSION()

SRCS(
    ut_tiers.cpp
)

END()
