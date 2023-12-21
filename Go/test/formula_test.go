package test

import (
	"os"
	config2 "qnurye/Cuber/pkg/config"
	"testing"
)

func TestLoadFormula(t *testing.T) {
	// 创建一个临时的 JSON 文件，用于测试
	tmpDir := "./config_test"
	tmpFile := tmpDir + "/formula_test.json"

	// 创建测试用例的 JSON 内容
	testJSON := `{
		"formula": "test_formula"
	}`

	// 创建临时目录
	err := os.MkdirAll(tmpDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Fatalf("Failed to remove test directory: %v", err)
		}
	}(tmpDir) // 确保测试完成后删除目录

	// 将测试用例写入临时文件
	err = os.WriteFile(tmpFile, []byte(testJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// 运行测试
	config, err := config2.LoadFormula(tmpFile)
	if err != nil {
		t.Fatalf("LoadFormula failed: %v", err)
	}

	// 验证加载的配置是否与测试用例一致
	expected := &config2.FormulaConfig{
		Formula: "test_formula",
	}

	if !formulaConfigEqual(config, expected) {
		t.Errorf("Loaded configuration does not match expected values.\nGot: %+v\nExpected: %+v", config, expected)
	}
}

// helper function to compare two FormulaConfig instances
func formulaConfigEqual(c1, c2 *config2.FormulaConfig) bool {
	return c1.Formula == c2.Formula
}
