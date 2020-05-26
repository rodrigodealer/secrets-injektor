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
  token: s.7LXjNtEutLhDwmW8Ff0VuTRt
  address: "http://localhost:8200"
environment:
  - ONE_ENV=hello_set
  - TWO_ENV=hello_set
`
	var c = model.Config{}

	c.Load([]byte(data))

	mock := new(ProviderChooserMock)
	var result = DecideProvider(c, mock)
	assert.Equal(t, []EnvironmentVariable{}, result, "they should be equal")
}

func TestSSMValue(t *testing.T) {
	var data = `
provider:
  name: ssm
  token: s.kJeKLfTtFfOM2yxXzzYQAxtF
  address: "http://localhost:8200"
environment:
- ONE_ENV=hello_set
- TWO_ENV=hello_set`
	var c = model.Config{}

	c.Load([]byte(data))

	mock := new(ProviderChooserMock)
	var result = DecideProvider(c, mock)
	assert.Equal(t, "ssm", result, "they should be equal")
}

func TestSMValue(t *testing.T) {
	var data = `
provider:
  name: sm
  token: s.kJeKLfTtFfOM2yxXzzYQAxtF
  address: "http://localhost:8200"
environment:
- ONE_ENV=hello_set
- TWO_ENV=hello_set`
	var c = model.Config{}

	c.Load([]byte(data))

	mock := new(ProviderChooserMock)
	var result = DecideProvider(c, mock)
	assert.Equal(t, "", result, "they should be equal")
}

func TestEmptyValue(t *testing.T) {
	var data = `
environment:
- ONE_ENV=hello_set
- TWO_ENV=hello_set`
	var c = model.Config{}

	c.Load([]byte(data))

	mock := new(ProviderChooserMock)
	var result = DecideProvider(c, mock)
	assert.Equal(t, "", result, "they should be equal")
}

type ProviderChooserMock struct {
	mock.Mock
}

func (m *ProviderChooserMock) callVault(c model.Config) []EnvironmentVariable {
	return []EnvironmentVariable{}
}
