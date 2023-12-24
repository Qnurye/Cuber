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
	formula := "U L F D U U U U U F' D' L' F2 R2 F R2 U' B2 L' F2 D2 R2 L F2 D2 R2 U2 L' F2"
	commands, err := parser.ParseFormula(formula)
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