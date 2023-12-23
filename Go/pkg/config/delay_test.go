package config

import (
	"os"
	"testing"
)

func TestLoadCmdDelay(t *testing.T) {
	// 创建一个临时的 JSON 文件，用于测试
	tmpDir := "./config_test"
	tmpFile := tmpDir + "/delay_test.json"

	// 创建测试用例的 JSON 内容
	testJSON := `{
		"CMD_R_GRIP_CLOSE":    100,
		"CMD_R_GRIP_OPEN":     200,
		"CMD_L_GRIP_CLOSE":    150,
		"CMD_L_GRIP_OPEN":     250,
		"CMD_R_ROTATE_CW_90":  300,
		"CMD_R_ROTATE_CCW_90": 400,
		"CMD_R_ROTATE_CW_180": 500,
		"CMD_L_ROTATE_CW_90":  600,
		"CMD_L_ROTATE_CCW_90": 700,
		"CMD_L_ROTATE_CW_180": 800,
		"CMD_L_ROTATE_CCW_180": 900,
		"CMD_R_ROTATE_CCW_180": 1000
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
	config, err := LoadCmdDelay(tmpFile)
	if err != nil {
		t.Fatalf("LoadCmdDelay failed: %v", err)
	}

	// 验证加载的配置是否与测试用例一致
	expected := &CommandDelayConfig{
		CmdRGripClose:    100,
		CmdRGripOpen:     200,
		CmdLGripClose:    150,
		CmdLGripOpen:     250,
		CmdRRotateCw90:   300,
		CmdRRotateCcw90:  400,
		CmdRRotateCw180:  500,
		CmdLRotateCw90:   600,
		CmdLRotateCcw90:  700,
		CmdLRotateCw180:  800,
		CmdLRotateCcw180: 900,
		CmdRRotateCcw180: 1000,
	}

	if !delayConfigEqual(config, expected) {
		t.Errorf("Loaded configuration does not match expected values.\nGot: %+v\nExpected: %+v", config, expected)
	}
}

// helper function to compare two CommandDelayConfig instances
func delayConfigEqual(c1, c2 *CommandDelayConfig) bool {
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
