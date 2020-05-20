package providers

import (
	"fmt"

	"github.com/hashicorp/vault/api"
	"github.com/kpango/glg"

	"github.com/rodrigodealer/secrets-injektor/model"
)

type VaultI interface {
	Connect(token string, addr string)

	GetSecret(secretName string) string
}

type Vault struct {
	Client *api.Logical
}

func GetOnVault(c model.Config, v VaultI) []EnvironmentVariable {
	v.Connect(c.Provider.Token, c.Provider.Address)
	var envs = GetEnvs(c.Environment)
	for _, element := range envs {
		var secret = v.GetSecret(fmt.Sprintf("secret/data/data/%s", element.Value))
		element.Value = secret
		glg.Infof("Secret found: %s", element.Name)
	}
	return envs
}

func (m *Vault) Connect(token string, addr string) {
	config := &api.Config{
		Address: addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		glg.Errorf("Error connecting to Vault: %s", err.Error())
	}
	client.SetToken(token)
	m.Client = client.Logical()
}

func (m *Vault) GetSecret(secretName string) string {
	secret, err := m.Client.Read(secretName)
	if err != nil {
		glg.Errorf("Error getting secret in Vault: %s", err.Error())
	}
	var data = secret.Data["data"].(map[string]interface{})
	return data["value"].(string)
}
