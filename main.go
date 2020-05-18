package main

import (
	"io/ioutil"

	"github.com/kpango/glg"
	"github.com/rodrigodealer/secrets-injektor/providers"

	"github.com/rodrigodealer/secrets-injektor/model"
)

func main() {
	data, err := ioutil.ReadFile("config.yaml")
	config := model.Config{}
	config.Load(data)
	check("Error loading file: %s", err)
	providers.DecideProvider(config, &providers.ProviderChooser{})
}

func check(message string, e error) {
	if e != nil {
		glg.Errorf(message, e.Error())
	}
}
