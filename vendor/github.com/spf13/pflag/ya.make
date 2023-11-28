GO_LIBRARY()

LICENSE(BSD-3-Clause)

SRCS(
    bool.go
    bool_slice.go
    bytes.go
    count.go
    duration.go
    duration_slice.go
    flag.go
    float32.go
    float32_slice.go
    float64.go
    float64_slice.go
    golangflag.go
    int.go
    int16.go
    int32.go
    int32_slice.go
    int64.go
    int64_slice.go
    int8.go
    int_slice.go
    ip.go
    ip_slice.go
    ipmask.go
    ipnet.go
    ipnet_slice.go
    string.go
    string_array.go
    string_slice.go
    string_to_int.go
    string_to_int64.go
    string_to_string.go
    uint.go
    uint16.go
    uint32.go
    uint64.go
    uint8.go
    uint_slice.go
)

GO_TEST_SRCS(
    bool_slice_test.go
    bool_test.go
    bytes_test.go
    count_test.go
    duration_slice_test.go
    export_test.go
    flag_test.go
    float32_slice_test.go
    float64_slice_test.go
    golangflag_test.go
    int32_slice_test.go
    int64_slice_test.go
    int_slice_test.go
    ip_slice_test.go
    ip_test.go
    ipnet_slice_test.go
    ipnet_test.go
    printusage_test.go
    string_array_test.go
    string_slice_test.go
    string_to_int64_test.go
    string_to_int_test.go
    string_to_string_test.go
    uint_slice_test.go
)

GO_XTEST_SRCS(example_test.go)

END()

RECURSE(gotest)
