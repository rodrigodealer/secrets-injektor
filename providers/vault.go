package providers

import (
	"github.com/hashicorp/vault/api"
	"github.com/kpango/glg"
	"strings"
)

type EnvironmentVariable struct {
	Name  string
	Value string
}

func GetEnvs(envs []string) []EnvironmentVariable {
	m := []EnvironmentVariable{}
	for _, element := range envs {
		m = append(m, GetEnvParts(element))
	}
	return m
}

func GetEnvParts(env string) EnvironmentVariable {
	parts := strings.Split(env, "=")
	return EnvironmentVariable{Name: parts[0], Value: parts[1]}
}

func getOnVault() {

}

type Vault struct {
	Client *api.Logical
}

func (m *Vault) New(addr string) *api.Client {
	config := &api.Config{
		Address: addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		glg.Errorf("Error connecting to Vault: %s", err.Error())
	}
	return client
}

func (m *Vault) Connect(token string, addr string) {
	client := m.New(addr)
	client.SetToken(token)
	m.Client = client.Logical()
}

func (m *Vault) GetSecret(secretName string) string {
	secret, err := m.Client.Read(secretName)
	if err != nil {
		glg.Errorf("Error getting secret in Vault: %s", err.Error())
	}
	return secret.Data["hello"].(string)
}
