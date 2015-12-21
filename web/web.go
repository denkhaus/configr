package web

import (
	"bytes"
	"io"
	"io/ioutil"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/denkhaus/configr/config"
)

var (
	logger = logrus.WithField("pkg", "web")
)

func LoadTemplate(name string) (io.Reader, error) {
	filePath := path.Join(config.TemplatePath(), name)
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(b), nil
}
