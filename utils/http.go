package utils

import (
	"net/http"
	"net/url"

	"wayshare.cn/common/logs"
	"wayshare.cn/common/utils"
	"wayshare.cn/common/wserror"
	"wdbible.com/wdbibleisimporter/common"
	"wdbible.com/wdbibleisimporter/models/bibleis"
)

func GetRequest(requestUrl string, data url.Values, headers http.Header) (string, error) {
	queryParam := ""
	if data == nil {
		data = url.Values{}
	}
	data.Set("v", common.BIBLE_IS_API_VERSION)
	data.Set("key", common.BIBLE_IS_KEY_APP)
	queryParam = data.Encode()
	requestData, err := bibleis.GetUrlContent(requestUrl, queryParam)
	if err == nil {
		logs.Debugf("Get request data from database")
		return requestData.UrlContent, nil
	}
	if err != wserror.ErrNotExist {
		logs.Errorf("Failed to get request data from database, the error is %#v", err)
	}
	status, content, err := utils.GetRequest(requestUrl, data, headers)
	if err != nil {
		logs.Errorf("Failed to get request from requestUrl %s, the error is %#v", requestUrl, err)
		return "", err
	}
	if status != http.StatusOK {
		logs.Errorf("Invalid status %d", status)
		return string(content), wserror.ErrInvalidStatus
	}
	requestData = bibleis.RequestData{RequestUrl: requestUrl, QueryParam: queryParam, Status: int64(status), UrlContent: string(content)}
	if err = requestData.Save(); err != nil {
		logs.Errorf("Failed to save data to database, the error is %#v", err)
	}
	return string(content), nil
}
