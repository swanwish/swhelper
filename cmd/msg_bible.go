package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/swanwish/go-common/logs"
	cu "github.com/swanwish/go-common/utils"
	"github.com/swanwish/swhelper/settings"
	"github.com/swanwish/swhelper/swerror"
	"github.com/swanwish/swhelper/utils"
	"github.com/urfave/cli"
)

const URL_QUOTE = "&#8212;"

var (
	FetchMsgBibleCmd = cli.Command{
		Name:        "fmsg",
		Usage:       "Fetch msg bible content and audio",
		Description: "This command is used to fetch the bible content and audio for msg bible",
		Action:      fetchMsgBible,
		Flags: []cli.Flag{
			intFlag("id", 0, "The id for the book to fetch"),
			stringFlag("dest", "", "The destination for the audio files"),
			//intFlag("p", 0, "Page number for fetch"),
		},
	}
)

func fetchMsgBible(c *cli.Context) error {
	settings.LoadAppSetting()
	id := c.Int64("id")
	if id <= 0 {
		msg := fmt.Sprintf("Invalid id %d", id)
		logs.Errorf(msg)
		return cli.NewExitError(msg, 1)
	}
	dest := c.String("dest")

	var page int64 = 1
	responseMsg, err := fetchMsg(id, page, dest)
	if err != nil {
		return err
	}
	pages := responseMsg.Pages
	for page = 2; page <= pages; page++ {
		responseMsg, err = fetchMsg(id, page, dest)
		if err != nil {
			return err
		}
		//logs.Debugf("The response message is %#v", responseMsg)
	}
	return nil
}

func fetchMsg(id, page int64, dest string) (ResponseMsg, error) {
	logs.Debugf("Fetch msg bible for id %d and page %d", id, page)
	requestUrl := fmt.Sprintf("http://enmsg.bibleinfo.info/api/get_category_posts/?id=%d&page=%d", id, page)

	//data := url.Values{}
	//data.Set("id", fmt.Sprintf("%d", id))
	//data.Set("page", fmt.Sprintf("%d", page))
	//data.Set("count", "8")

	headers := http.Header{}
	headers.Set("Accept", "application/json")
	headers.Set("Content-Type", "application/json; charset=UTF-8")

	content, err := utils.GetRequest(requestUrl, nil, headers)
	if err != nil {
		msg := fmt.Sprintf("Failed to get bible content, the error is %#v", err)
		logs.Errorf(msg)
		return ResponseMsg{}, cli.NewExitError(msg, 1)
	}
	logs.Debugf("The content is %s", content)

	responseMsg := ResponseMsg{}
	err = json.Unmarshal([]byte(content), &responseMsg)
	if err != nil {
		logs.Errorf("Failed to unmarshal response message, the error is %#v", err)
		return responseMsg, cli.NewExitError("Unmarshal response failed", 1)
	}
	for _, post := range responseMsg.Posts {
		audioUrl, err := getAudioUrl(post.Content)
		if err != nil {
			logs.Errorf("Failed to get audio url, the error is %#v", err)
			return responseMsg, cli.NewExitError("Failed to get audio url", 1)
		}
		logs.Debugf("The audio url is %s", audioUrl)
		err = cu.DownloadFromUrl(audioUrl, dest, "")
		if err != nil {
			if err == cu.ErrAlreadyExists {
				continue
			}
			logs.Errorf("Failed to download audio %s, the error is %#v", audioUrl, err)
			return responseMsg, cli.NewExitError("Download audio failed", 1)
		}
	}
	return responseMsg, nil
}

type ResponseMsg struct {
	Status   string           `json:"status"`
	Count    int64            `json:"count"`
	Pages    int64            `json:"pages"`
	Category ResponseCategory `json:"category"`
	Posts    []ResponsePost   `json:"posts"`
}

type ResponseCategory struct {
	Id          int64  `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Parent      int64  `json:"parent"`
	PostCount   int64  `json:"post_count"`
}

type ResponsePost struct {
	Id         int64              `json:"id"`
	Type       string             `json:"type"`
	Slug       string             `json:"slug"`
	Url        string             `json:"url"`
	Status     string             `json:"status"`
	Title      string             `json:"title"`
	TitlePlain string             `json:"title_plain"`
	Content    string             `json:"content"`
	Excerpt    string             `json:"excerpt"`
	Date       string             `json:"date"`
	Modified   string             `json:"modified"`
	Categories []ResponseCategory `json:"categories"`
}

func getAudioUrl(content string) (string, error) {
	startIndex := strings.Index(content, URL_QUOTE)
	endIndex := strings.LastIndex(content, URL_QUOTE)
	if startIndex != endIndex {
		url := content[startIndex+len(URL_QUOTE) : endIndex]
		return url, nil
	}
	return "", swerror.ErrNotExist
}
