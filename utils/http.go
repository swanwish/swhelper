package utils

import (
	"net/http"
	"net/url"

	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/go-common/utils"
	"github.com/swanwish/swhelper/models/resource"
	"github.com/swanwish/swhelper/swerror"
)

func GetRequest(requestUrl string, data url.Values, headers http.Header) (string, error) {
	queryParam := ""
	queryParam = data.Encode()
	requestData, err := resource.GetUrlContent(requestUrl, queryParam)
	if err == nil {
		logs.Debugf("Get request data from database")
		return requestData.UrlContent, nil
	}
	if err != swerror.ErrNotExist {
		logs.Errorf("Failed to get request data from database, the error is %#v", err)
	}
	status, content, err := utils.GetRequest(requestUrl, data, headers)
	if err != nil {
		logs.Errorf("Failed to get request from requestUrl %s, the error is %#v", requestUrl, err)
		return "", err
	}
	if status != http.StatusOK {
		logs.Errorf("Invalid status %d", status)
		return string(content), swerror.ErrInvalidStatus
	}
	requestData = resource.RequestData{RequestUrl: requestUrl, QueryParam: queryParam, Status: int64(status), UrlContent: string(content)}
	if err = requestData.Save(); err != nil {
		logs.Errorf("Failed to save data to database, the error is %#v", err)
	}
	return string(content), nil
}
