package providers

import (
	"testing"

	"github.com/hashicorp/vault/api"
	"github.com/rodrigodealer/secrets-injektor/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetOnVault(t *testing.T) {
	var data = `
provider:
  name: vault
  token: s.7LXjNtEutLhDwmW8Ff0VuTRt
  address: "http://localhost:8200"
environment:
  - ONE_ENV=hello_set
  - TWO_ENV=hello_set
`
	var c = model.Config{}
	c.Load([]byte(data))

	vaultMock := new(VaultMock)

	vaultMock.On("GetSecret", mock.Anything).Return("bla")

	var result = GetOnVault(c, vaultMock)

	assert.Equal(t, 2, len(result), "they should be equal")
}

type VaultMock struct {
	mock.Mock
}

func (m *VaultMock) Connect(token string, addr string) {
}

func (m *VaultMock) New(addr string) *api.Client {
	return nil
}

func (m *VaultMock) GetSecret(secretName string) string {
	return ""
}
