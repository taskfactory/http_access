package entity

// GetSourcesReq 获取数据源列表请求
type GetSourcesReq struct {
	SName    string `json:"sname"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
}

// UpsertSourceReq 创建或更新数据源请求
type UpsertSourceReq struct {
	ID    int64  `json:"id"`
	SName string `json:"sname"`
	Desc  string `json:"desc"`
}
