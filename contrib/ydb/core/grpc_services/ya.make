LIBRARY()

SRCS(
    audit_log.cpp
    audit_dml_operations.cpp
    db_metadata_cache.h
    grpc_endpoint_publish_actor.cpp
    grpc_helper.cpp
    grpc_mon.cpp
    grpc_publisher_service_actor.cpp
    grpc_request_proxy.cpp
    grpc_request_proxy_simple.cpp
    local_rate_limiter.cpp
    operation_helpers.cpp
    resolve_local_db_table.cpp
    rpc_alter_coordination_node.cpp
    rpc_alter_table.cpp
    rpc_begin_transaction.cpp
    rpc_calls.cpp
    rpc_cancel_operation.cpp
    rpc_cms.cpp
    rpc_commit_transaction.cpp
    rpc_dynamic_config.cpp
    rpc_copy_table.cpp
    rpc_copy_tables.cpp
    rpc_export.cpp
    rpc_create_coordination_node.cpp
    rpc_create_table.cpp
    rpc_describe_coordination_node.cpp
    rpc_describe_path.cpp
    rpc_describe_table.cpp
    rpc_describe_table_options.cpp
    rpc_drop_coordination_node.cpp
    rpc_drop_table.cpp
    rpc_discovery.cpp
    rpc_execute_data_query.cpp
    rpc_execute_scheme_query.cpp
    rpc_execute_yql_script.cpp
    rpc_explain_yql_script.cpp
    rpc_explain_data_query.cpp
    rpc_forget_operation.cpp
    rpc_fq_internal.cpp
    rpc_fq.cpp
    rpc_get_operation.cpp
    rpc_get_shard_locations.cpp
    rpc_import.cpp
    rpc_import_data.cpp
    rpc_keep_alive.cpp
    rpc_keyvalue.cpp
    rpc_kh_describe.cpp
    rpc_kh_snapshots.cpp
    rpc_kqp_base.cpp
    rpc_list_operations.cpp
    rpc_login.cpp
    rpc_load_rows.cpp
    rpc_log_store.cpp
    rpc_long_tx.cpp
    rpc_node_registration.cpp
    rpc_maintenance.cpp
    rpc_make_directory.cpp
    rpc_modify_permissions.cpp
    rpc_monitoring.cpp
    rpc_prepare_data_query.cpp
    rpc_rate_limiter_api.cpp
    rpc_read_columns.cpp
    rpc_read_table.cpp
    rpc_read_rows.cpp
    rpc_remove_directory.cpp
    rpc_rename_tables.cpp
    rpc_rollback_transaction.cpp
    rpc_scheme_base.cpp
    rpc_stream_execute_scan_query.cpp
    rpc_stream_execute_yql_script.cpp
    rpc_whoami.cpp
    table_settings.cpp

    rpc_common/rpc_common_kqp_session.cpp

    query/rpc_execute_query.cpp
    query/rpc_execute_script.cpp
    query/rpc_fetch_script_results.cpp
    query/rpc_attach_session.cpp
    query/service_query.h
)

PEERDIR(
    contrib/libs/xxhash
    library/cpp/cgiparam
    library/cpp/digest/old_crc
    contrib/ydb/core/actorlib_impl
    contrib/ydb/core/audit
    contrib/ydb/core/base
    contrib/ydb/core/control
    contrib/ydb/core/discovery
    contrib/ydb/core/engine
    contrib/ydb/core/formats
    contrib/ydb/core/fq/libs/actors
    contrib/ydb/core/fq/libs/control_plane_proxy
    contrib/ydb/core/fq/libs/control_plane_proxy/events
    contrib/ydb/core/grpc_services/base
    contrib/ydb/core/grpc_services/counters
    contrib/ydb/core/grpc_services/local_rpc
    contrib/ydb/core/grpc_services/cancelation
    contrib/ydb/core/grpc_services/auth_processor
    contrib/ydb/core/health_check
    contrib/ydb/core/io_formats
    contrib/ydb/core/kesus/tablet
    contrib/ydb/core/kqp/common
    contrib/ydb/core/protos
    contrib/ydb/core/scheme
    contrib/ydb/core/sys_view
    contrib/ydb/core/tx
    contrib/ydb/core/tx/datashard
    contrib/ydb/core/tx/sharding
    contrib/ydb/core/tx/long_tx_service/public
    contrib/ydb/core/tx/ev_write
    contrib/ydb/core/ydb_convert
    contrib/ydb/core/security
    contrib/ydb/library/aclib
    contrib/ydb/library/binary_json
    contrib/ydb/library/dynumber
    contrib/ydb/library/mkql_proto
    contrib/ydb/library/persqueue/topic_parser
    contrib/ydb/library/yql/parser/pg_wrapper/interface
    contrib/ydb/library/yql/public/types
    contrib/ydb/library/services
    contrib/ydb/public/api/grpc/draft
    contrib/ydb/public/api/protos
    contrib/ydb/public/lib/fq
    contrib/ydb/public/lib/operation_id
    contrib/ydb/public/sdk/cpp/client/resources
    contrib/ydb/services/ext_index/common
)

YQL_LAST_ABI_VERSION()

END()

RECURSE(
    base
    counters
    local_rpc
)

RECURSE_FOR_TESTS(
    ut
)
