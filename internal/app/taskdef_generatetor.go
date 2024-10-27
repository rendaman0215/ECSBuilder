package app

import (
	"fmt"
	"os"

	"github.com/rendaman0215/ECSbuilder/internal/config"
	"github.com/rendaman0215/ECSbuilder/pkg/template_parser"

	"github.com/manifoldco/promptui"
)

func GenerateTaskDefinition(templatePath, envFilePath, outputPath, inputDir string) error {
	// パスが指定されていない場合はpromptuiで入力を求める
	if inputDir == "" {
		prompt := promptui.Prompt{
			Label:   "Path to input files directory",
			Default: ".",
		}
		result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("failed to get input: %v", err)
		}
		inputDir = result
	}

	if templatePath == "" {
		prompt := promptui.Prompt{
			Label:   "template file name",
			Default: "taskdef.template.json",
		}
		result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("failed to get input: %v", err)
		}
		templatePath = fmt.Sprintf("%s/%s", inputDir, result)
	} else {
		templatePath = fmt.Sprintf("%s/%s", inputDir, templatePath)
	}

	if envFilePath == "" {
		prompt := promptui.Prompt{
			Label:   "env file name",
			Default: ".env",
		}
		result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("failed to get input: %v", err)
		}
		envFilePath = fmt.Sprintf("%s/%s", inputDir, result)
	} else {
		envFilePath = fmt.Sprintf("%s/%s", inputDir, envFilePath)
	}

	if outputPath == "" {
		prompt := promptui.Prompt{
			Label:   "Path to output task definition file",
			Default: "taskdef.json",
		}
		result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("failed to get input: %v", err)
		}
		outputPath = result
	}

	// 環境変数の読み込み
	envVars, err := config.LoadEnvVariables(envFilePath)
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %v", err)
	}

	// テンプレートファイルの読み込みと変数埋め込み
	content, err := template_parser.ParseTemplate(templatePath, envVars)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// ファイルの書き込み
	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	fmt.Println("Task definition file generated successfully")

	return nil
}
