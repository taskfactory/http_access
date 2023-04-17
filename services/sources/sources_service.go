package sources

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taskfactory/http_access/common/errs"
	"github.com/taskfactory/http_access/common/tool/query"
	"github.com/taskfactory/http_access/entity"
	"github.com/taskfactory/http_access/logic/sources"
)

// GetSources 获取数据来源
func GetSources(c *gin.Context) {
	req := &entity.GetSourcesReq{
		SName:    query.GetString(c, "sname"),
		Page:     int32(query.GetInt(c, "page")),
		PageSize: int32(query.GetInt(c, "pageSize")),
	}
	rsp, err := sources.GetSources(c.Request.Context(), req)
	if err != nil {
		c.JSON(200, gin.H{
			"code": errs.Code(err),
			"msg":  errs.Msg(err),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": errs.CodeOK,
		"msg":  "succeed",
		"data": rsp,
	})
}

// Upsert 更新或创建数据
func Upsert(c *gin.Context) {
	req := new(entity.UpsertSourceReq)
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(200, gin.H{
			"code": errs.CodeInternal,
			"msg":  fmt.Sprintf("解析参数失败,err:%v", err),
		})
		return
	}

	rsp, err := sources.UpsertSource(c.Request.Context(), req)
	if err != nil {
		c.JSON(200, gin.H{
			"code": errs.Code(err),
			"msg":  errs.Msg(err),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": errs.CodeOK,
		"msg":  "succeed",
		"data": rsp,
	})
}
