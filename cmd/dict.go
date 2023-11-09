package cmd

import (
	"errors"

	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/swhelper/dict/oxford"
	"github.com/swanwish/swhelper/settings"
	"github.com/urfave/cli"
)

var (
	ShowWordDefinitionCmd = cli.Command{
		Name:        "def",
		Usage:       "Show the definition of a given word",
		Description: "This command will get the word definition from oxford website, and print it",
		Action:      showWordDefinition,
		Flags: []cli.Flag{
			stringFlag("word", "", "The word to get the definition"),
		},
	}
)

func showWordDefinition(c *cli.Context) error {
	word := c.String("word")
	if word == "" {
		logs.Errorf("The word is not specified")
		return errors.New("word not specified")
	}
	settings.LoadAppSetting()
	logs.Debugf("The word is %s", word)
	inflectionId, err := oxford.GetFirstInflection(word)
	if err != nil {
		logs.Errorf("Failed to get first inflection for word %s, the error is %#v", word, err)
		return err
	}
	logs.Debugf("The inflection id is %#v", inflectionId)
	retrievedEntry, err := oxford.GetWordEntries(inflectionId)
	if err != nil {
		logs.Errorf("Failed to get word entry, the error is %#v", err)
	}
	definition := retrievedEntry.ToString()
	logs.Debugf("The definition is \n%s", definition)
	return nil
}
