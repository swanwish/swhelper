package resource

// RequestData for table sw_resource.t_request_data
type RequestData struct {
	RequestId  int64  `db:"request_id" json:"requestId" form:"requestId"`
	RequestUrl string `db:"request_url" json:"requestUrl" form:"requestUrl"`
	QueryParam string `db:"query_param" json:"queryParam" form:"queryParam"`
	Status     int64  `db:"status" json:"status" form:"status"`
	UrlContent string `db:"url_content" json:"urlContent" form:"urlContent"`
	CreateTime int64  `db:"create_time" json:"createTime" form:"createTime"`
}
