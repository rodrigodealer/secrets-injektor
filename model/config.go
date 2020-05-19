package model

import (
	"github.com/kpango/glg"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	data, err := ioutil.ReadFile("config.yaml")
	check("Error loading file: %s", err)
	return data
}

func check(message string, e error) {
	if e != nil {
		glg.Errorf(message, e.Error())
	}
}

func (m *Config) Load(data []byte) {
	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		glg.Errorf("error: %v", err)
	}

	if len(data) > 0 && err == nil {
		glg.Info("Config loaded")
	}
}
