package sources

import (
	"context"

	"github.com/taskfactory/http_access/client/admin"
	"github.com/taskfactory/http_access/common/errs"
	"github.com/taskfactory/http_access/entity"

	proto "github.com/taskfactory/http_access/tars-protocol/admin"
)

// GetSources 查询数据源列表
func GetSources(ctx context.Context, req *entity.GetSourcesReq) (*proto.SourcePagination, error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	return admin.GetSources(ctx, req.SName, req.Page, req.PageSize)
}

// UpsertSource 插入或更新数据源
func UpsertSource(ctx context.Context, req *entity.UpsertSourceReq) (*proto.Source, error) {
	if req.SName == "" {
		return nil, errs.New(errs.CodeParam, "数据源名称不得为空")
	}
	return admin.Upsert(ctx, req.ID, req.SName, req.Desc)
}
