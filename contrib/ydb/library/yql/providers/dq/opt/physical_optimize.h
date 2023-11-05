#pragma once

#include <contrib/ydb/library/yql/providers/dq/common/yql_dq_settings.h>
#include <contrib/ydb/library/yql/core/yql_graph_transformer.h>
#include <contrib/ydb/library/yql/core/yql_type_annotation.h>

#include <util/generic/ptr.h>

namespace NYql::NDqs {

THolder<IGraphTransformer> CreateDqsPhyOptTransformer(TTypeAnnotationContext* typeCtx, const TDqConfiguration::TPtr& config);

} // namespace NYql::NDqs
