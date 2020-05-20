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
	  token: s.kJeKLfTtFfOM2yxXzzYQAxtF
	  address: "http://localhost:8200"
	environment:
	- ONE_ENV=hello_set
	- TWO_ENV=hello_set`
	var c = model.Config{}
	c.Load([]byte(data))

	mock := new(VaultMock)

	var result = GetOnVault(c, mock)

	var expect = []EnvironmentVariable{}

	assert.Equal(t, expect, result, "they should be equal")
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
