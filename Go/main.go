package main

import (
	"fmt"
	"log"
	"qnurye/Cuber/pkg/config"
	"qnurye/Cuber/pkg/rubiksCube"
	"strconv"
	"time"
)

func main() {
	cmd, err := config.LoadCmd("./config/cmd.json")
	delay, err := config.LoadCmdDelay("./config/delay.json")
	//formulaCfg, err := config.LoadFormula("./config/formula.json")
	if err != nil {
		log.Fatal(err)
	}
	//c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	//s, err := serial.OpenPort(c)
	//if err != nil {
	//	log.Fatal(err)
	//}

	parser := rubiksCube.NewCubeParser(cmd, delay)
	//formula := "U F' D' L' F2 R2 F R2 U' B2 L' F2 D2 R2 L F2 D2 R2 U2 L' F2"
	formula := "U F D L F R F R U B L2 F D"
	//formula := formulaCfg.Formula
	commands, err := parser.ParseFormula(formula)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		slept := time.Duration(0)
		sent := 0

		current := commands.Head
		for current != nil {
			d, err := time.ParseDuration(strconv.Itoa(current.Command.Delay) + "ms")
			if err != nil {
				log.Fatal(err)
			}

			//n, err := s.Write([]byte(current.Command.Operation + "\n"))
			//if err != nil {
			//	log.Fatal(err)
			//}

			n := 2
			sent += n
			slept += d

			log.Printf("Sent %v byte: %v, sleeping for %v ms", n, current.Command.Operation, d.Milliseconds())
			//time.Sleep(d)

			current = current.Next
		}
		log.Printf("Finished !\n Sent %v bytes, slept %v ms (which is %v) in total", sent, slept.Milliseconds(), slept)
	}
}
