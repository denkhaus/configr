package command

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/denkhaus/configr/engine"
	"github.com/juju/errors"
)

var (
	logger = logrus.WithField("pkg", "command")
)

type Commander struct {
	app    *cli.App
	engine *engine.Engine
}

////////////////////////////////////////////////////////////////////////////////
func (c *Commander) Execute(fn engine.EngineFunc, ctx *cli.Context) {
	if err := c.InitEngine(ctx); err != nil {
		logger.Errorf("engine init::%q", err)
	}

	if err := c.engine.Execute(fn); err != nil {
		logger.Errorf(errors.ErrorStack(err))
	}
}

////////////////////////////////////////////////////////////////////////////////
func (c *Commander) InitEngine(ctx *cli.Context) (err error) {
	if c.engine == nil {
		c.engine, err = engine.New(ctx)
		if err != nil {
			return errors.Annotate(err, "new engine")
		}
	}
	return
}

////////////////////////////////////////////////////////////////////////////////
func New(app *cli.App) *Commander {
	cmd := &Commander{app: app}
	cmd.NewServeCommand()
	return cmd
}

////////////////////////////////////////////////////////////////////////////////
func (c *Commander) RegisterCmd(cmd cli.Command) {
	c.app.Commands = append(c.app.Commands, cmd)
}

////////////////////////////////////////////////////////////////////////////////
func (c *Commander) Run(args []string) {
	c.app.Run(args)
	if c.engine != nil {
		c.engine.Shutdown()
	}
}
