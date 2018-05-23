package cmd

import (
	"fmt"
	"time"

	"github.com/urfave/cli"
)

var (
	ShowCurrentTimeCmd = cli.Command{
		Name:        "ct",
		Usage:       "Show current time",
		Description: "This command is used to show current time",
		Action:      showCurrentTime,
		Flags:       []cli.Flag{},
	}

	ShowCurrentWeekNumCmd = cli.Command{
		Name:        "wn",
		Usage:       "Show current week number",
		Description: "This command is used to show current week number",
		Action:      showCurrentWeekNum,
		Flags:       []cli.Flag{},
	}
)

func showCurrentTime(c *cli.Context) error {
	now := time.Now()
	currentTime := now.Format("2006-01-02 15:04:05")
	fmt.Printf("The current time is: %s\nThe unix timestamp is: %d, Nanoseconds: %d\n", currentTime, now.Unix(), now.UnixNano()/1000000)
	return nil
}

func showCurrentWeekNum(c *cli.Context) error {
	now := time.Now()
	year, week := now.ISOWeek()
	fmt.Printf("The current week number is: year:%d week:%d\n", year, week)
	return nil
}
