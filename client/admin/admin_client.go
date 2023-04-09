package admin

import (
	"context"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/taskfactory/http_access/common/errs"

	proto "github.com/taskfactory/http_access/tars-protocol/admin"
)

const (
	adminServantObj = "TaskFactory.admin.admin"
)

// GetSources 获取数据源列表
func GetSources(ctx context.Context, sname string, page, pageSize int32) ([]proto.Source, error) {
	app := new(proto.AdminService)
	tars.NewCommunicator().StringToProxy(adminServantObj, app)
	req := &proto.GetSourcesReq{
		Sname:    sname,
		Page:     page,
		PageSize: pageSize,
	}
	rsp, err := app.GetSourcesWithContext(ctx, req)
	if err != nil {
		err = errs.Newf(rsp.Code, "get source list failed with message:%s, err:%v", rsp.Msg, err)
		return nil, err
	}

	return rsp.Sources, nil
}

// Upsert 创建或更新数据源
func Upsert(ctx context.Context, id int64, sname, desc string) (*proto.Source, error) {
	app := new(proto.AdminService)
	tars.NewCommunicator().StringToProxy(adminServantObj, app)
	req := &proto.UpsertSourceReq{
		Id:    id,
		Sname: sname,
		Desc:  desc,
	}
	rsp, err := app.UpsertSourceWithContext(ctx, req)

	if err != nil {
		err = errs.Newf(rsp.Code, "upsert source failed with message:%s, err:%v", rsp.Msg, err)
		return nil, err
	}

	return &rsp.Source, nil
}
