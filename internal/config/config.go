package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(filePath string) (map[string]string, error) {
	err := godotenv.Load(filePath)
	if err != nil {
		return nil, err
	}

	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		envVars[pair[0]] = pair[1]
	}

	return envVars, nil
}
