package config

import "github.com/Sirupsen/logrus"

var (
	logger = logrus.WithField("pkg", "config")
)

type config struct {
}

func TemplatePath() string {
	return ""
}
