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
	//formula := "U F' D' L' F2 R2 F R2 U' B2 L' F2 D2 R2 L F2 D2 R2 U2 L' F2"
	formula := formulaCfg.Formula
	commands, err := parser.ParseFormula(formula)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		//fmt.Println("Commands:", commands)
		slept := time.Duration(0)
		sent := 0
		for _, cmd := range commands {
			d, err := time.ParseDuration(strconv.Itoa(cmd.Delay) + "ms")
			if err != nil {
				return
			}
			n, err := s.Write([]byte(cmd.Operation + "\n"))
			if err != nil {
				log.Fatal(err)
			}
			sent += n
			slept += d
			//sent += 2
			log.Printf("Sent %v byte: %v, sleeping for %v ms", 2, cmd.Operation, d.Milliseconds())
			time.Sleep(d)
		}
		log.Printf("Finished !\n Sent %v bytes, slept %v ms (which is %v) in total", sent, slept.Milliseconds(), slept)
	}
}
