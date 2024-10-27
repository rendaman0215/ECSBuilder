package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ECSbuilder",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Prompt for the input directory
		dirPrompt := promptui.Prompt{
			Label:   "Enter input directory for template and env files (default: current directory)",
			Default: ".",
		}
		inputDir, err := dirPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		// Prompt for template and env file names, with defaults
		templatePrompt := promptui.Prompt{
			Label:   "Enter task definition template file (default: taskdef.template.json)",
			Default: filepath.Join(inputDir, "taskdef.template.json"),
		}
		templateFile, err := templatePrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		envPrompt := promptui.Prompt{
			Label:   "Enter environment file (default: .env)",
			Default: filepath.Join(inputDir, ".env"),
		}
		envFile, err := envPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		// Load environment variables from specified .env file
		err = godotenv.Load(envFile)
		if err != nil {
			fmt.Printf("Error loading %s: %v\n", envFile, err)
			return
		}

		// Read the template file
		templateData, err := os.ReadFile(templateFile)
		if err != nil {
			fmt.Printf("Error reading template file %s: %v\n", templateFile, err)
			return
		}

		// Convert template content to string for replacements
		output := string(templateData)

		// Replace placeholders in the form ${VAR_NAME} with environment values
		for _, env := range os.Environ() {
			pair := strings.SplitN(env, "=", 2)
			placeholder := fmt.Sprintf("${%s}", pair[0])
			output = strings.ReplaceAll(output, placeholder, pair[1])
		}

		// Write the result to taskdef.json in the current directory
		err = os.WriteFile("taskdef.json", []byte(output), 0644)
		if err != nil {
			fmt.Printf("Error writing to taskdef.json: %v\n", err)
			return
		}

		fmt.Println("Generated taskdef.json successfully.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
