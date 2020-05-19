package providers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetEnvs(t *testing.T) {
	var theArray = []string{"VAR_A=bla", "VAR_B=bla"}

	var values = GetEnvs(theArray)

	assert.Equal(t, 2, len(values), "they should be equal")
	assert.Equal(t, "bla", values[0].Value, "they should be equal")
}

func TestSetEnvs(t *testing.T) {
	var theArray = []EnvironmentVariable{EnvironmentVariable{Name: "bla", Value: "blu"}}

	SetEnvs(theArray)

	assert.Equal(t, "blu", os.Getenv("bla"), "they should be equal")
}

func TestGetEnvParts(t *testing.T) {
	var theString = "VAR_A=bla"

	var env = GetEnvParts(theString)

	assert.Equal(t, "bla", env.Value, "they should be equal")
}

type VaultMock struct {
	mock.Mock
}
