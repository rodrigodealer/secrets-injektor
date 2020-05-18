package providers

import (
	"github.com/kpango/glg"
	"github.com/rodrigodealer/secrets-injektor/model"
)

func DecideProvider(c model.Config) string {
	switch c.Provider.Name {
	case "vault":
		GetOnVault(c)
		return "vault"
	case "ssm":
		println("TBD")
	default:
		glg.Infof("Please choose a provider")
	}
	return ""
}
