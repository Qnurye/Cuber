package main

import (
	"qnurye/Cuber/pkg/config"
	"qnurye/Cuber/pkg/rubiksCube"
)

func main() {
	cmd, err := config.LoadCmd()

	parser := rubiksCube.NewCubeParser(cmd)
	formula := "D R B R L' D L F L' F L2 F2 D F2 U D L2 D' F2 R2"
	commands, err := parser.ParseFormula(formula)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Commands:", commands)
	}

	//
	//c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	//s, err := serial.OpenPort(c)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//n, err := s.Write([]byte(cmd.CMD_R_ROTATE_CCW_90 + "\n"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//buf := make([]byte, 128)
	//n, err = s.Read(buf)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("%q", buf[:n])
}
