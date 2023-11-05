LIBRARY()

PEERDIR(
    contrib/ydb/library/yql/ast
    contrib/ydb/library/yql/core
    contrib/ydb/library/yql/dq/common
    contrib/ydb/library/yql/dq/expr_nodes
    contrib/ydb/library/yql/dq/integration
    contrib/ydb/library/yql/dq/proto
    contrib/ydb/library/yql/dq/type_ann
    contrib/ydb/library/yql/providers/dq/expr_nodes
)

SRCS(
    dq_opt.cpp
    dq_opt_build.cpp
    dq_opt_join.cpp
    dq_opt_log.cpp
    dq_opt_peephole.cpp
    dq_opt_phy_finalizing.cpp
    dq_opt_phy.cpp
    dq_opt_stat.cpp
    dq_opt_stat_transformer_base.cpp
    dq_opt_join_cost_based.cpp
    dq_opt_join_cost_based_generic.cpp
    dq_opt_predicate_selectivity.cpp
)

YQL_LAST_ABI_VERSION()

END()

RECURSE_FOR_TESTS(ut)
