package providers

import (
	"github.com/kpango/glg"
	"github.com/rodrigodealer/secrets-injektor/model"
)

type ProviderChooserI interface {
	callVault(c model.Config) string
}

type ProviderChooser struct {
}

func (p *ProviderChooser) callVault(c model.Config) string {
	GetOnVault(c)
	return "vault"
}

func DecideProvider(c model.Config, chooser ProviderChooserI) string {
	switch c.Provider.Name {
	case "vault":
		return chooser.callVault(c)
	case "ssm":
		println("TBD")
	default:
		glg.Infof("Please choose a provider")
	}
	return ""
}
