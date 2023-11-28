GO_LIBRARY()

SRCS(
    reader.go
    register.go
    struct.go
    writer.go
)

GO_TEST_SRCS(
    fuzz_test.go
    reader_test.go
    writer_test.go
    zip_test.go
)

GO_XTEST_SRCS(example_test.go)

END()

RECURSE(
)
