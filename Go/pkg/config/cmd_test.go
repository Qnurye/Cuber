package config

import (
	"os"
	"testing"
)

func TestLoadCmd(t *testing.T) {
	// 创建一个临时的 JSON 文件，用于测试
	tmpDir := "./config_test"
	tmpFile := tmpDir + "/cmd_test.json"

	// 创建测试用例的 JSON 内容
	testJSON := `{
		"CMD_R_GRIP_CLOSE": "test_CMD_R_GRIP_CLOSE",
		"CMD_R_GRIP_OPEN": "test_CMD_R_GRIP_OPEN",
		"CMD_L_GRIP_CLOSE": "test_CMD_L_GRIP_CLOSE",
		"CMD_L_GRIP_OPEN": "test_CMD_L_GRIP_OPEN",
		"CMD_R_ROTATE_CW_90": "test_CMD_R_ROTATE_CW_90",
		"CMD_R_ROTATE_CCW_90": "test_CMD_R_ROTATE_CCW_90",
		"CMD_R_ROTATE_CW_180": "test_CMD_R_ROTATE_CW_180",
		"CMD_L_ROTATE_CW_90": "test_CMD_L_ROTATE_CW_90",
		"CMD_L_ROTATE_CCW_90": "test_CMD_L_ROTATE_CCW_90",
		"CMD_L_ROTATE_CW_180": "test_CMD_L_ROTATE_CW_180",
		"CMD_L_ROTATE_CCW_180": "test_CMD_L_ROTATE_CCW_180",
		"CMD_R_ROTATE_CCW_180": "test_CMD_R_ROTATE_CCW_180"
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
	config, err := LoadCmd(tmpFile)
	if err != nil {
		t.Fatalf("LoadCmd failed: %v", err)
	}

	// 验证加载的配置是否与测试用例一致
	expected := &CommandConfig{
		CmdRGripClose:    "test_CMD_R_GRIP_CLOSE",
		CmdRGripOpen:     "test_CMD_R_GRIP_OPEN",
		CmdLGripClose:    "test_CMD_L_GRIP_CLOSE",
		CmdLGripOpen:     "test_CMD_L_GRIP_OPEN",
		CmdRRotateCw90:   "test_CMD_R_ROTATE_CW_90",
		CmdRRotateCcw90:  "test_CMD_R_ROTATE_CCW_90",
		CmdRRotateCw180:  "test_CMD_R_ROTATE_CW_180",
		CmdLRotateCw90:   "test_CMD_L_ROTATE_CW_90",
		CmdLRotateCcw90:  "test_CMD_L_ROTATE_CCW_90",
		CmdLRotateCw180:  "test_CMD_L_ROTATE_CW_180",
		CmdLRotateCcw180: "test_CMD_L_ROTATE_CCW_180",
		CmdRRotateCcw180: "test_CMD_R_ROTATE_CCW_180",
	}

	if !configEqual(config, expected) {
		t.Errorf("Loaded configuration does not match expected values.\nGot: %+v\nExpected: %+v", config, expected)
	}
}

// helper function to compare two CommandConfig instances
func configEqual(c1, c2 *CommandConfig) bool {
	return c1.CmdRGripClose == c2.CmdRGripClose &&
		c1.CmdRGripOpen == c2.CmdRGripOpen &&
		c1.CmdLGripClose == c2.CmdLGripClose &&
		c1.CmdLGripOpen == c2.CmdLGripOpen &&
		c1.CmdRRotateCw90 == c2.CmdRRotateCw90 &&
		c1.CmdRRotateCcw90 == c2.CmdRRotateCcw90 &&
		c1.CmdRRotateCw180 == c2.CmdRRotateCw180 &&
		c1.CmdLRotateCw90 == c2.CmdLRotateCw90 &&
		c1.CmdLRotateCcw90 == c2.CmdLRotateCcw90 &&
		c1.CmdLRotateCw180 == c2.CmdLRotateCw180 &&
		c1.CmdLRotateCcw180 == c2.CmdLRotateCcw180 &&
		c1.CmdRRotateCcw180 == c2.CmdRRotateCcw180
}
