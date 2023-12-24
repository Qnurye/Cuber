package tests

import (
	"fmt"
	"log"
	"qnurye/Cuber/pkg/config"
	"qnurye/Cuber/pkg/rubiksCube"
	"strconv"
	"testing"
	"time"
)

func TestCubeParser_ParseFormula_Performance(t *testing.T) {
	// 加载配置文件
	cmdConfig, err := config.LoadCmd("../config/cmd.json")
	if err != nil {
		t.Fatalf("Failed to load command config: %v", err)
	}
	delayConfig, err := config.LoadCmdDelay("../config/delay.json")
	if err != nil {
		t.Fatalf("Failed to load delay config: %v", err)
	}

	// 创建 CubeParser
	parser := rubiksCube.NewCubeParser(cmdConfig, delayConfig)

	// 运行测试
	formula := "U2 F L' F2 B L2 U' D' L F U R U' F2 D' R2 U F2 R2 U2"
	//formula, err := os.ReadFile("../config/mofang.txt")
	commands, err := parser.ParseFormula(string(formula))
	if err != nil {
		t.Fatalf("ParseFormula failed: %v", err)
	}

	timeCost := time.Duration(0)
	bytes := 0

	current := commands.Head
	for current != nil {
		d, err := time.ParseDuration(strconv.Itoa(current.Command.Delay) + "ms")
		if err != nil {
			log.Fatal(err)
		}

		bytes += 1
		timeCost += d

		current = current.Next
	}

	fmt.Printf("Bytes: %v | Duration: %v (%v)\n", bytes, timeCost, timeCost.Milliseconds())
}
