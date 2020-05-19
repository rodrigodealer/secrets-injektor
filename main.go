package main

import (
	"github.com/rodrigodealer/secrets-injektor/providers"

	"github.com/rodrigodealer/secrets-injektor/model"
)

func main() {
	var data = model.OpenConfig("config.yaml")
	config := model.Config{}
	config.Load(data)

	providers.DecideProvider(config, &providers.ProviderChooser{})
}
