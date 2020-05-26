package model

import (
	"io/ioutil"

	"github.com/kpango/glg"
	"gopkg.in/yaml.v2"

	"github.com/rodrigodealer/secrets-injektor/util"
)

type Provider struct {
	Name      string `yaml:"name"`
	Token     string `yaml:"token"`
	Address   string `yaml:"address"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
}

type Config struct {
	Provider    Provider `yaml:"provider"`
	Environment []string `yaml:"environment"`
}

func OpenConfig(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	check("Error loading file: %s", err)
	return data
}

func check(message string, e error) {
	if e != nil {
		if util.IsLoggingEnabled() {
			glg.Errorf(message, e.Error())
		}
	}
}

func (m *Config) Load(data []byte) {
	err := yaml.Unmarshal([]byte(data), &m)
	check("error: %v", err)

	if len(data) > 0 && err == nil {
		if util.IsLoggingEnabled() {
			glg.Info("Config loaded")
		}
	}
}
