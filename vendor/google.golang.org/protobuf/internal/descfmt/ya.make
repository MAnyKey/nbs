GO_LIBRARY()

LICENSE(BSD-3-Clause)

SRCS(stringer.go)

GO_TEST_SRCS(desc_test.go)

END()

RECURSE(gotest)
