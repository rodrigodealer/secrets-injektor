package providers

import "strings"

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
