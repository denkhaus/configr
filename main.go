package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/denkhaus/configr/command"
)

var (
	logger = logrus.WithField("pkg", "main")
)

func main() {

	app := cli.NewApp()
	app.Name = "configr"
	app.Version = fmt.Sprintf("%s-%s", AppVersion, Revision)
	app.Usage = "A settings manager with web backend."
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "config, c", Value: "mirsvc.yaml", Usage: "Configuration File path"},
		cli.BoolFlag{Name: "revision, r", Usage: "Print revision"},
		//		cli.IntFlag{"port, P", 993, "Port number to connect to.", ""},
		//		cli.StringFlag{"user, u", "", "Your username at host", ""},
		//		cli.StringFlag{"pass, p", "", "Your IMAP password. For security reasons prefer IMAP_PASSWORD='yourpassword'", "IMAP_PASSWORD"},

	}

	app.Action = func(ctx *cli.Context) {
		if ctx.GlobalBool("revision") {
			fmt.Println(Revision)
			return
		}

		cli.ShowAppHelp(ctx)
	}

	command.New(app).Run(os.Args)
}
