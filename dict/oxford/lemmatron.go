package oxford

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/swhelper/common"
	"github.com/swanwish/swhelper/swerror"
	"github.com/swanwish/swhelper/utils"
)

func GetWordLemmatron(wordId string) (Lemmatron, error) {
	wordId = strings.ToLower(wordId)
	getUrl := fmt.Sprintf("%s/inflections/%s/%s", common.OD_API_BASE_URL, common.OD_API_SOURCE_LANG, wordId)
	data := url.Values{}
	content, err := utils.GetOxfordRequest(getUrl, data, nil)
	if err != nil {
		logs.Errorf("The error is %#v", err)
		return Lemmatron{}, err
	}
	lemmatron := Lemmatron{}
	err = json.Unmarshal([]byte(content), &lemmatron)
	if err != nil {
		logs.Errorf("Failed to unmarshal json %s", content)
	}
	return lemmatron, err
}

func GetFirstInflection(wordId string) (string, error) {
	lemmatron, err := GetWordLemmatron(wordId)
	if err != nil {
		logs.Errorf("Failed to get lemmatron, the error is %#v", err)
		return "", err
	}
	results := lemmatron.Results
	if len(results) > 0 {
		lexicalEntries := results[0].LexicalEntries
		if len(lexicalEntries) > 0 {
			inflectionList := lexicalEntries[0].InflectionOf
			if len(inflectionList) > 0 {
				return inflectionList[0].Id, nil
			}
		}
	}
	return "", swerror.ErrNotExist
}
