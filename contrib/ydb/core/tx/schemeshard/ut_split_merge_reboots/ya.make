IF (NOT WITH_VALGRIND)
    UNITTEST_FOR(contrib/ydb/core/tx/schemeshard)

    FORK_SUBTESTS()

    SPLIT_FACTOR(20)

    IF (SANITIZER_TYPE OR WITH_VALGRIND)
        TIMEOUT(3600)
        SIZE(LARGE)
        TAG(ya:fat)
    ELSE()
        TIMEOUT(1800)
        SIZE(LARGE)
        TAG(ya:fat)
    ENDIF()

    PEERDIR(
        library/cpp/getopt
        library/cpp/regex/pcre
        library/cpp/svnversion
        contrib/ydb/core/testlib/default
        contrib/ydb/core/tx
        contrib/ydb/core/tx/schemeshard/ut_helpers
        contrib/ydb/library/yql/public/udf/service/exception_policy
    )

    YQL_LAST_ABI_VERSION()

    SRCS(
        ut_split_merge_reboots.cpp
    )

    END()
ENDIF()
