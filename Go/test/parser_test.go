package test

import (
	"os"
	"qnurye/Cuber/pkg/config"
	"qnurye/Cuber/pkg/rubiksCube"
	"testing"
)

func TestCubeParser_ParseFormula(t *testing.T) {
	// 创建一个临时的 JSON 文件，用于测试
	tmpDir := "./config_test"
	cmdFile := tmpDir + "/cmd_test.json"
	delayFile := tmpDir + "/delay_test.json"

	// 创建测试用例的 JSON 内容
	testCmdJSON := `{
		"CMD_R_GRIP_CLOSE":    "test_CMD_R_GRIP_CLOSE",
		"CMD_R_GRIP_OPEN":     "test_CMD_R_GRIP_OPEN",
		"CMD_L_GRIP_CLOSE":    "test_CMD_L_GRIP_CLOSE",
		"CMD_L_GRIP_OPEN":     "test_CMD_L_GRIP_OPEN",
		"CMD_R_ROTATE_CW_90":  "test_CMD_R_ROTATE_CW_90",
		"CMD_R_ROTATE_CCW_90": "test_CMD_R_ROTATE_CCW_90",
		"CMD_R_ROTATE_CW_180": "test_CMD_R_ROTATE_CW_180",
		"CMD_L_ROTATE_CW_90":  "test_CMD_L_ROTATE_CW_90",
		"CMD_L_ROTATE_CCW_90": "test_CMD_L_ROTATE_CCW_90",
		"CMD_L_ROTATE_CW_180": "test_CMD_L_ROTATE_CW_180",
		"CMD_L_ROTATE_CCW_180": "test_CMD_L_ROTATE_CCW_180",
		"CMD_R_ROTATE_CCW_180": "test_CMD_R_ROTATE_CCW_180"
	}`
	testDelayJSON := `{
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
	err = os.WriteFile(cmdFile, []byte(testCmdJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create test command file: %v", err)
	}
	err = os.WriteFile(delayFile, []byte(testDelayJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create test delay file: %v", err)
	}

	// 加载配置文件
	cmdConfig, err := config.LoadCmd(cmdFile)
	if err != nil {
		t.Fatalf("Failed to load command config: %v", err)
	}
	delayConfig, err := config.LoadCmdDelay(delayFile)
	if err != nil {
		t.Fatalf("Failed to load delay config: %v", err)
	}

	// 创建 CubeParser
	parser := rubiksCube.NewCubeParser(cmdConfig, delayConfig)

	// 运行测试
	formula := "R D (f20)"
	commands, err := parser.ParseFormula(formula)
	if err != nil {
		t.Fatalf("ParseFormula failed: %v", err)
	}

	// 验证解析的命令序列是否符合预期
	expectedList := &rubiksCube.LinkedList{}

	// Right commands
	expectedList.Append(rubiksCube.Command{Operation: "test_CMD_R_ROTATE_CCW_90", Delay: 400})
	expectedList.Append(rubiksCube.Command{Operation: "test_CMD_R_GRIP_OPEN", Delay: 200})
	expectedList.Append(rubiksCube.Command{Operation: "test_CMD_R_ROTATE_CW_90", Delay: 300})
	expectedList.Append(rubiksCube.Command{Operation: "test_CMD_R_GRIP_CLOSE", Delay: 100})

	// Left commands
	expectedList.Append(rubiksCube.Command{Operation: "test_CMD_L_ROTATE_CCW_90", Delay: 700})
	expectedList.Append(rubiksCube.Command{Operation: "test_CMD_L_GRIP_OPEN", Delay: 250})
	expectedList.Append(rubiksCube.Command{Operation: "test_CMD_L_ROTATE_CW_90", Delay: 600})
	expectedList.Append(rubiksCube.Command{Operation: "test_CMD_L_GRIP_CLOSE", Delay: 150})

	if !commandsEqual(commands, expectedList) {
		t.Errorf("Parsed commands do not match expected values.\nGot: %+v\nExpected: %+v", commands, expectedList)
	}
}

// helper function to compare two slices of Command instances
func commandsEqual(list1, list2 *rubiksCube.LinkedList) bool {
	current1 := list1.Head
	current2 := list2.Head

	for current1 != nil && current2 != nil {
		if current1.Operation != current2.Operation || current1.Delay != current2.Delay {
			return false
		}

		current1 = current1.Next
		current2 = current2.Next
	}

	// If lengths are not equal, or if one list is longer than the other
	return current1 == nil && current2 == nil
}
