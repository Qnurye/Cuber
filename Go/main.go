package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"qnurye/Cuber/pkg/config"
	"qnurye/Cuber/pkg/rubiksCube"
	"strconv"
	"time"
)

func main() {
	cmd, err := config.LoadCmd("./config/cmd.json")
	delay, err := config.LoadCmdDelay("./config/delay.json")
	formulaCfg, err := config.LoadFormula("./config/formula.json")
	if err != nil {
		log.Fatal(err)
	}
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	parser := rubiksCube.NewCubeParser(cmd, delay)
	formula := formulaCfg.Formula
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

			n, err := s.Write([]byte(current.Command.Operation + "\n"))
			if err != nil {
				log.Fatal(err)
			}

			step += 1
			slept += d

			log.Printf("Sent %v step: %v, sleeping for %v ms", n, current.Command.Operation, d.Milliseconds())
			time.Sleep(d)

			current = current.Next
		}
		log.Printf("Finished !\n Sent %v steps, slept %v ms (which is %v) in total", step, slept.Milliseconds(), slept)
	}
}
