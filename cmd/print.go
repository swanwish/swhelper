package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/swanwish/go-common/logs"
	"github.com/urfave/cli"
)

var (
	ShowDoubleSideBookletPagesCmd = cli.Command{
		Name:        "dsbp",
		Usage:       "Show double side booklet pages",
		Description: "This command is used to show double side booklet pages",
		Action:      showDoubleSideBookletPages,
		Flags: []cli.Flag{
			intFlag("p, page", 0, "The total pages"),
		},
	}
)

func showDoubleSideBookletPages(c *cli.Context) error {
	page := c.Int("p")
	if page <= 0 {
		logs.Errorf("The page does not specified")
		return cli.NewExitError("page parameter does not specified", 1)
	}
	fixedPage := page
	mod := page % 4
	if mod > 0 {
		blankPages := 4 - mod
		fmt.Printf("Add %d blank pages\n", blankPages)
		fixedPage += blankPages
	}
	round := fixedPage / 4
	pageIndexes := make([]string, 0)
	for i := 0; i < round; i++ {
		offset := i * 2
		pageIndexes = append(pageIndexes, strconv.Itoa(fixedPage-offset), strconv.Itoa(offset+1), strconv.Itoa(offset+2), strconv.Itoa(fixedPage-offset-1))
	}

	fmt.Printf("The page indexes are: [%s]\n", strings.Join(pageIndexes, ", "))
	return nil
}
