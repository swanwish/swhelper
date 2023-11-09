package oxford

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/swhelper/common"
	"github.com/swanwish/swhelper/utils"
)

func GetWordEntries(wordId string) (RetrieveEntry, error) {
	wordId = strings.ToLower(wordId)
	getUrl := fmt.Sprintf("%s/entries/%s/%s/regions=us", common.OD_API_BASE_URL, common.OD_API_SOURCE_LANG, wordId)
	data := url.Values{}
	content, err := utils.GetOxfordRequest(getUrl, data, nil)
	if err != nil {
		logs.Errorf("The error is %#v", err)
		return RetrieveEntry{}, err
	}
	retrieveEntry := RetrieveEntry{}
	err = json.Unmarshal([]byte(content), &retrieveEntry)
	if err != nil {
		logs.Errorf("Failed to get retrieved entry for wordId %s, the error is %#v", wordId, err)
	}
	return retrieveEntry, err
}

func (retrieveEntry RetrieveEntry) ToString() string {
	results := []string{}
	for _, headwordEntry := range retrieveEntry.Results {
		subString := headwordEntry.ToString()
		if subString != "" {
			results = append(results, subString)
		}
	}
	return strings.Join(results, "\n")
}

func (headwordEntry HeadwordEntry) ToString() string {
	results := []string{}
	for _, item := range headwordEntry.LexicalEntries {
		subString := item.ToString()
		if subString != "" {
			results = append(results, subString)
		}
	}
	return strings.Join(results, "\n")
}

func (lexicalEntry LexicalEntry) ToString() string {
	results := []string{}
	wordText := lexicalEntry.Text
	if wordText != "" {
		spellings := ""
		for _, pronunciation := range lexicalEntry.Pronunciations {
			phoneticSpelling := pronunciation.PhoneticSpelling
			if phoneticSpelling != "" {
				if strings.Index(spellings, phoneticSpelling) != -1 {
					continue
				}
				spellings = fmt.Sprintf("%s /%s/", spellings, phoneticSpelling)
			}
		}
		results = append(results, fmt.Sprintf("%s%s", wordText, spellings))
	}
	for _, item := range lexicalEntry.Entries {
		subString := item.ToString()
		if subString != "" {
			results = append(results, subString)
		}
	}
	return strings.Join(results, "\n")
}

func (entry Entry) ToString() string {
	results := []string{}
	for index, item := range entry.Senses {
		if index > 0 {
			results = append(results, "")
		}
		subString := item.ToString(fmt.Sprintf("%d.", index+1), common.DEFAULT_INDENT)
		if subString != "" {
			results = append(results, subString)
		}
	}
	return strings.Join(results, "\n")
}

func (sense Sense) ToString(pos, prefix string) string {
	results := []string{}
	for _, subString := range sense.Definitions {
		if subString != "" {
			results = append(results, fmt.Sprintf("%s%s %s", prefix, pos, subString))
		}
	}
	subPrefix := fmt.Sprintf("%s%s", prefix, common.DEFAULT_INDENT)
	for _, example := range sense.Examples {
		results = append(results, example.ToString(subPrefix))
	}
	for index, subSense := range sense.Subsenses {
		results = append(results, "")
		results = append(results, subSense.ToString(fmt.Sprintf("%s%d.", pos, index+1), subPrefix))
	}
	return strings.Join(results, "\n")
}

func (example Example) ToString(prefix string) string {
	return fmt.Sprintf("%s> %s", prefix, example.Text)
}
