OWNER(g:cloud-nbs)

GO_LIBRARY()

SRCS(
    s3.go
    s3_metrics.go
    ydb.go
    ydb_logger.go
    ydb_metrics.go
)

GO_TEST_SRCS(
    ydb_test.go
)

END()

RECURSE(
    config
)

RECURSE_FOR_TESTS(
    tests
)
