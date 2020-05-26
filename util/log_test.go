package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLoggingEnabled(t *testing.T) {
	os.Setenv("DEBUG_LEVEL", "true")

	result := IsLoggingEnabled()

	assert.True(t, result, "they should be equal")

	os.Setenv("DEBUG_LEVEL", "")
}

func TestIsLoggingDisabled(t *testing.T) {
	os.Setenv("DEBUG_LEVEL", "")

	result := IsLoggingEnabled()

	assert.False(t, result, "they should be equal")

	os.Setenv("DEBUG_LEVEL", "")

}
