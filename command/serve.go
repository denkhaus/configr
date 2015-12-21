package command

import (
	"github.com/codegangsta/cli"
	"github.com/denkhaus/configr/engine"
)

func (c *Commander) NewServeCommand() {
	c.RegisterCmd(cli.Command{
		Name:  "serve",
		Usage: "Serve setting backend",
		//Flags: []cli.Flag{
		//cli.BoolFlag{"force, f", "rebuild all images"},
		//cli.BoolFlag{"kill, k", "kill containers"},
		//},

		Action: func(ctx *cli.Context) {
			c.Execute(func(engine *engine.Engine) error {
				return engine.Serve()
			}, ctx)
		},
	})
}
