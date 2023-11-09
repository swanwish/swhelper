package utils

import (
	"net/http"
	"net/url"

	"github.com/swanwish/swhelper/settings"
)

func GetOxfordRequest(requestUrl string, data url.Values, headers http.Header) (string, error) {
	if headers == nil {
		headers = http.Header{}
	}
	headers.Set("app_id", settings.AppId)
	headers.Set("app_key", settings.AppKey)
	headers.Set("Accept", "application/json")
	return GetRequest(requestUrl, data, headers)
}
