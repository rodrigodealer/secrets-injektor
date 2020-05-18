package providers

import (
	"testing"

	"github.com/rodrigodealer/secrets-injektor/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVaultValue(t *testing.T) {
	var data = `
provider:
  name: vault
  token: s.kJeKLfTtFfOM2yxXzzYQAxtF
  address: "http://localhost:8200"
environment:
- ONE_ENV=hello_set
- TWO_ENV=hello_set`
	var c = model.Config{}

	c.Load([]byte(data))

	mock := new(ProviderChooserMock)
	var result = DecideProvider(c, mock)
	assert.Equal(t, "vault", result, "they should be equal")
}

type ProviderChooserMock struct {
	mock.Mock
}

func (m *ProviderChooserMock) callVault(c model.Config) string {
	return "vault"
}
