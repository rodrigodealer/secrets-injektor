package providers

import (
	"os"
	"strings"

	"github.com/kpango/glg"
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

func SetEnvs(envs []EnvironmentVariable) {
	for _, element := range envs {
		os.Setenv(element.Name, element.Value)
	}
	glg.Infof("%v environment variables set.", len(envs))
}

func GetEnvParts(env string) EnvironmentVariable {
	parts := strings.Split(env, "=")
	return EnvironmentVariable{Name: parts[0], Value: parts[1]}
}
