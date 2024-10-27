package template_parser

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ParseTemplate(templatePath string, envVars map[string]string) (string, error) {
	data, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %v", err)
	}

	content := string(data)

	// プレースホルダー ${VAR_NAME} の置換
	for key, value := range envVars {
		placeholder := fmt.Sprintf("${%s}", key)
		content = strings.ReplaceAll(content, placeholder, value)
	}

	return content, nil
}
