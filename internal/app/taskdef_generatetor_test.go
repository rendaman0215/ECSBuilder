package app_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/rendaman0215/ECSbuilder/internal/app"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateTaskDefinition(t *testing.T) {
	success := map[string]struct {
		dir     string
		tplFile string
		envFile string
		outFile string
	}{
		"全パラメータ指定": {
			dir:     "testdata",
			tplFile: "taskdef.template.json",
			envFile: ".env",
			outFile: "taskdef.json",
		},
	}

	t.Run("成功", func(t *testing.T) {
		for tt, tc := range success {
			t.Run(tt, func(t *testing.T) {
				// テスト用の一時ディレクトリを作成
				tempDir := t.TempDir()

				// テスト用の.envファイルを作成
				envFilePath := filepath.Join(tempDir, tc.envFile)
				envContent := "ENVIRONMENT=test\nCONTAINER_IMAGE=test-image"
				err := os.WriteFile(envFilePath, []byte(envContent), 0644)
				require.NoError(t, err)

				// テスト用のテンプレートファイルを作成
				templateFilePath := filepath.Join(tempDir, tc.tplFile)
				templateContent := `{
					"family": "sample-task-definition-${ENVIRONMENT}",
					"containerDefinitions": [
						{
							"name": "sample-container",
							"image": "${CONTAINER_IMAGE}"
						}
					]
				}`
				err = os.WriteFile(templateFilePath, []byte(templateContent), 0644)
				require.NoError(t, err)

				// 出力先ファイルパス
				outputFilePath := filepath.Join(tempDir, tc.outFile)

				// GenerateTaskDefinition関数を実行
				err = GenerateTaskDefinition(templateFilePath, envFilePath, outputFilePath, tempDir)
				require.NoError(t, err)

				// 生成されたファイルの内容を検証
				outputContent, err := os.ReadFile(outputFilePath)
				require.NoError(t, err)

				// 期待する結果
				expectedContent := `{
					"family": "sample-task-definition-test",
					"containerDefinitions": [
						{
							"name": "sample-container",
							"image": "test-image"
						}
					]
				}`
				assert.JSONEq(t, expectedContent, string(outputContent), "The generated task definition content should match expected output")

				// 生成されたファイルを削除
				err = os.Remove(outputFilePath)
				require.NoError(t, err)

			})
		}
	})
}
