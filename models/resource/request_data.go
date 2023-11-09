package resource

import (
	"time"

	"github.com/swanwish/swhelper/swerror"
)

// Save method for table sw_resource.t_request_data
func (model *RequestData) Save() error {
	currentTime := time.Now().Unix()
	if model.RequestId == 0 {
		// Insert code snippet
		model.CreateTime = currentTime
		insertSql := `INSERT INTO t_request_data (request_url, query_param, status, url_content, create_time) VALUES (?, ?, ?, ?, ?)`
		result, err := resourceDB.Exec(insertSql, model.RequestUrl, model.QueryParam, model.Status, model.UrlContent, model.CreateTime)
		if err != nil {
			return err
		}
		lastInsertId, err := result.LastInsertId()
		model.RequestId = lastInsertId
	} else {
		// Update code snippet
		updateSql := `UPDATE t_request_data SET request_url=?, query_param=?, status=?, url_content=?, create_time=? WHERE request_id=?`
		_, err := resourceDB.Exec(updateSql, model.RequestUrl, model.QueryParam, model.Status, model.UrlContent, model.CreateTime, model.RequestId)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUrlContent(requestUrl, queryParam string) (RequestData, error) {
	querySql := `SELECT request_id, request_url, query_param, status, url_content, create_time FROM t_request_data
				WHERE request_url=? AND query_param=?`
	params := []interface{}{requestUrl, queryParam}
	list := []RequestData{}
	err := resourceDB.Select(&list, querySql, params...)
	if err != nil {
		return RequestData{}, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return RequestData{}, swerror.ErrNotExist
}
