package main

import (
	"fmt"
	"log"
	"qnurye/Cuber/pkg/config"
	"qnurye/Cuber/pkg/rubiksCube"
	"strconv"
	"testing"
	"time"
)

func Test(t *testing.T) {
	cmd, err := config.LoadCmd("./config/cmd.json")
	delay, err := config.LoadCmdDelay("./config/delay.json")
	if err != nil {
		log.Fatal(err)
	}

	parser := rubiksCube.NewCubeParser(cmd, delay)
	//formula := "U F' D' L' F2 R2 F R2 U' B2 L' F2 D2 R2 L F2 D2 R2 U2 L' F2"
	//formula := "F B U D L R F B U D L F B U D L F B U D L" // between B and U
	formula := "F B U D L R F B U D L F B"
	commands, err := parser.ParseFormula(formula)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		slept := time.Duration(0)
		step := 0

		current := commands.Head
		for current != nil {
			d, err := time.ParseDuration(strconv.Itoa(current.Command.Delay) + "ms")
			if err != nil {
				log.Fatal(err)
			}

			n := 1
			step += n
			slept += d

			log.Printf("Sent %v step: %v, sleeping for %v ms", n, current.Command.Operation, d.Milliseconds())

			current = current.Next
		}
		log.Printf("Finished !\n Sent %v steps, slept %v ms (which is %v) in total", step, slept.Milliseconds(), slept)
	}
}
