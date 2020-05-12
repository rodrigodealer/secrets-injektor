package main

import (
	"github.com/kpango/glg"
	"gopkg.in/yaml.v2"
)

type Provider struct {
	Name      string `yaml:"name"`
	Token     string `yaml:"token"`
	Address   string `yaml:"address"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
}

type Environment struct {
	Envs map[string]string `yaml:"envs"`
}

type Config struct {
	Provider Provider `yaml:"provider"`
}

func (m *Config) Load(data []byte) {
	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		glg.Errorf("error: %v", err)
	}
}
