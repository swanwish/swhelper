package cmd

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/go-common/utils"
	"github.com/urfave/cli"
	"github.com/xuri/excelize/v2"
)

var (
	UpdateServerCommandCmd = cli.Command{
		Name:        "usc",
		Usage:       "swhelper usc",
		Description: "This command is used to update server commands",
		Action:      updateServerCommandAction,
		Flags: []cli.Flag{
			stringFlag("p", "/Users/Stephen/OneDrive/Documents/servers.xlsx", "The excel file path for the server list"),
			stringFlag("d", "/Users/Stephen/tools/wd/libexec", "The destination for the generated files"),
		},
	}
)

const tmpl = `
#!/usr/bin/env bash
# Usage: Connect to {{.Name}}
# Summary: Connect to the server {{.Name}}
# Help: This command will connect to {{.Name}}

set -e

export user=themartian
export host={{.Ip}}
export port=2022

exec {{.ConnectCommand}} "$@"`

type ServerInfo struct {
	Type string
	Name string
	Ip   string
}

func (s ServerInfo) ConnectCommand() string {
	switch s.Type {
	case "wd":
		return "connect-wd"
	default:
		return "connect"
	}
}

func updateServerCommandAction(c *cli.Context) error {
	excelFilePath := c.String("p")
	if !utils.FileExists(excelFilePath) {
		return fmt.Errorf("The file path %s does not exists", excelFilePath)
	}
	dest := c.String("d")
	if !utils.FileExists(dest) {
		return fmt.Errorf("The dest path %s does not exists", dest)
	}
	// Open the Excel file.
	f, err := excelize.OpenFile(excelFilePath)
	if err != nil {
		logs.Errorf("Cannot open file: %v", err)
		return err
	}
	defer func() {
		// Close the Excel file.
		if err := f.Close(); err != nil {
			logs.Errorf("Cannot close file: %v", err)
		}
	}()

	sheetList := f.GetSheetList()
	for _, sheet := range sheetList {
		logs.Debugf("The sheet is %s", sheet)
	}
	// Get all the rows in the first sheet.
	rows, err := f.GetRows(sheetList[0])
	if err != nil {
		logs.Errorf("Cannot get rows: %v", err)
		return err
	}

	// Create a new template with a name.
	t, err := template.New("cmd").Parse(tmpl)
	if err != nil {
		logs.Errorf("parsing: %s", err)
		return err
	}

	// Loop through the rows.
	for index, row := range rows {
		if index == 0 {
			continue
		}
		var buf bytes.Buffer
		// Loop through the columns in each row.
		if len(row) == 3 {
			serverInfo := ServerInfo{
				Type: row[0],
				Name: getServerName(row[1]),
				Ip:   row[2],
			}
			err = t.Execute(&buf, serverInfo)
			if err != nil {
				logs.Errorf("execution: %s", err)
				return err
			}

			// The string is now in the buffer. Convert it to a string.
			err = saveFile(dest, serverInfo.Name, buf.Bytes())
			if err != nil {
				logs.Errorf("Failed to save command file for server %s, the error is %#v", serverInfo.Name, err)
				return err
			}
		}

	}
	return nil
}

func getServerName(serverName string) string {
	return strings.ReplaceAll(strings.ToLower(serverName), "-", "_")
}

func saveFile(dest, serverName string, command []byte) error {
	destFile := filepath.Join(dest, fmt.Sprintf("wd-%s", serverName))
	if utils.FileExists(destFile) {
		if err := utils.DeleteFile(destFile); err != nil {
			return err
		}
	}
	return utils.SaveFile(destFile, command, 0755)
}
