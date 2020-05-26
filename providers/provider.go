package providers

import (
	"github.com/kpango/glg"
	"github.com/rodrigodealer/secrets-injektor/model"
	"github.com/rodrigodealer/secrets-injektor/util"
)

type ProviderChooserI interface {
	callVault(c model.Config) []EnvironmentVariable
}

type ProviderChooser struct {
}

func (p *ProviderChooser) callVault(c model.Config) []EnvironmentVariable {
	return GetOnVault(c, &Vault{})
}

func DecideProvider(c model.Config, chooser ProviderChooserI) interface{} {
	switch c.Provider.Name {
	case "vault":
		return chooser.callVault(c)
	case "ssm":
		println("TBD")
		return "ssm"
	default:
		if util.IsLoggingEnabled() {
			glg.Infof("Please choose a provider")
		}
		return ""
	}
	return ""
}
