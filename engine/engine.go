package engine

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var (
	logger = log.WithField("pkg", "engine")
)

type EngineFunc func(engine *Engine) error
type Engine struct {
}

func New(ctx *cli.Context) (*Engine, error) {
	en := new(Engine)
	return en, nil
}

func (p *Engine) Serve() error {
	logger.Infof("Engine::Shutdown...")

	return nil
}

func (p *Engine) Shutdown() error {
	logger.Infof("Engine::Shutdown...")

	return nil
}

func (e *Engine) Execute(fn EngineFunc) error {
	return fn(e)
}
