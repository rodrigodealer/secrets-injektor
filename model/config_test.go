package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLoad(t *testing.T) {
	var data = `
provider:
  name: vault
  token: s.kJeKLfTtFfOM2yxXzzYQAxtF
  address: "http://localhost:8200"
environment:
- ONE_ENV=hello_set
- TWO_ENV=hello_set`
	var c = Config{}

	c.Load([]byte(data))

	assert.Equal(t, "vault", c.Provider.Name, "they should be equal")
	assert.Equal(t, 2, len(c.Environment), "they should be equal")
}

func TestConfigLoadWithError(t *testing.T) {
	var data = `
	{error:error}`
	var c = Config{}

	c.Load([]byte(data))
}

func TestOpenConfig(t *testing.T) {
	var c = OpenConfig("../config.yaml")

	assert.True(t, len(c) > 0, "they should be equal")
}

func TestOpenConfigWithFailure(t *testing.T) {
	var c = OpenConfig("config.yaml")

	assert.True(t, len(c) == 0, "they should be equal")
}
