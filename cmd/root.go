package cmd

import (
	"github.com/rendaman0215/ECSbuilder/internal/app"
	"github.com/spf13/cobra"
)

var (
	templateFile string
	envFile      string
	outputFile   string
	inputDir     string
)

var rootCmd = &cobra.Command{
	Short: "Generate ECS task definition files with embedded environment variables",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.GenerateTaskDefinition(templateFile, envFile, outputFile, inputDir)
	},
}

func Execute() error {
	// コマンドライン引数を追加
	rootCmd.Flags().StringVarP(&inputDir, "input", "i", "", "Path to input files directory")
	rootCmd.Flags().StringVarP(&templateFile, "template", "t", "", "Path to template file")
	rootCmd.Flags().StringVarP(&envFile, "env", "e", "", "Path to environment file")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Path to output task definition file")

	return rootCmd.Execute()
}
