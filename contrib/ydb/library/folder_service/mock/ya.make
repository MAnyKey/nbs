LIBRARY()

SRCS(
    mock_folder_service_adapter.cpp
)

PEERDIR(
    library/cpp/actors/core
    contrib/ydb/library/folder_service
    contrib/ydb/library/folder_service/proto
)

YQL_LAST_ABI_VERSION()

END()
