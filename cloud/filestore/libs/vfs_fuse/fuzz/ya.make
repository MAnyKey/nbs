FUZZ()

TAG(
    ya:not_autocheck
    ya:manual
)

SRCS(
    main.cpp
    starter.cpp
)

ADDINCL(
    cloud/contrib/vhost
)

PEERDIR(
    cloud/filestore/libs/diagnostics
    cloud/filestore/libs/endpoint_vhost
    cloud/filestore/libs/service
    cloud/filestore/libs/service_null
    cloud/filestore/libs/vfs_fuse/vhost
    
    cloud/storage/core/libs/common
    cloud/storage/core/libs/diagnostics
    cloud/storage/core/libs/vhost-client

    contrib/libs/virtiofsd
)

END()
