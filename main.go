package main

import (
	"os"
	"runtime"

	"github.com/swanwish/swhelper/cmd"
	"github.com/urfave/cli"
)

const APP_VERSION = "1.0.0"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "swhelper"
	app.Usage = "Helper Tool for Swanwish"
	app.Version = APP_VERSION
	app.Commands = []cli.Command{
		cmd.ShowCurrentTimeCmd,
		cmd.ShowCurrentWeekNumCmd,
		cmd.ShowDoubleSideBookletPagesCmd,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
