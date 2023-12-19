package rubiksCube

import (
	"errors"
	"qnurye/Cuber/pkg/config"
	"strings"
)

// TODO: Finish and clean this AI generated shit

// CubeParser represents a Rubik's Cube formula parser.
type CubeParser struct {
	Config config.CommandConfig
}

// NewCubeParser creates a new CubeParser with the specified configuration.
func NewCubeParser(config config.CommandConfig) *CubeParser {
	return &CubeParser{Config: config}
}

// ParseFormula parses a Rubik's Cube formula and returns the sequence of commands.
func (parser *CubeParser) ParseFormula(formula string) ([]string, error) {
	ts := strings.Fields(formula) // tokens
	cmds := make([]string, 0)     // commands

	for _, t := range ts {
		command, err := parser.parseToken(t)
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, command)
	}

	return cmds, nil
}

type Direction int

const (
	CCW Direction = 1 << iota // Counter Clock Wise
	CW2                       // Clock Wise Twice
	CW                        // Clock Wise
)

func (parser *CubeParser) parseToken(token string) (string, error) {
	if len(token) < 2 {
		return "", errors.New("invalid token length")
	}

	d := token[len(token)-1:] // direction
	switch d {
	case "'":
		return parser.getCmd(token[:len(token)-1], CCW)
	case "2":
		return parser.getCmd(token[:len(token)-1], CW2)
	default:
		return parser.getCmd(token, CW)
	}
}

func (parser *CubeParser) getCmd(move string, direction Direction) (string, error) {
	switch move {
	case "R":
		return parser.getRotateCmd("R", direction)
	case "L":
		return parser.getRotateCmd("L", direction)
	case "D":
		return parser.getGripCmd("R", direction)
	case "U":
		return parser.getGripCmd("L", direction)
	case "F":
		return parser.getGripCmd("R", direction)
	case "B":
		return parser.getGripCmd("L", direction)
	default:
		return "", errors.New("unknown move: " + move)
	}
}

func (parser *CubeParser) getRotateCmd(side string, rotation Direction) (string, error) {
	switch side {
	case "R":
		return parser.getCmdByRotation(cmd.CMD_R_ROTATE_CW_90, cmd.CMD_R_ROTATE_CCW_90, cmd.CMD_R_ROTATE_CW_180, rotation)
	case "L":
		return parser.getCmdByRotation(cmd.CMD_L_ROTATE_CW_90, cmd.CMD_L_ROTATE_CCW_90, cmd.CMD_L_ROTATE_CW_180, rotation)
	default:
		return "", errors.New("unknown side: " + side)
	}
}

func (parser *CubeParser) getGripCmd(side string, rotation Direction) (string, error) {
	switch side {
	case "R":
		return parser.getCmdByRotation(cmd.CMD_R_GRIP_CLOSE, cmd.CMD_R_GRIP_OPEN, "", rotation)
	case "L":
		return parser.getCmdByRotation(cmd.CMD_L_GRIP_CLOSE, cmd.CMD_L_GRIP_OPEN, "", rotation)
	default:
		return "", errors.New("unknown side: " + side)
	}
}

func (parser *CubeParser) getCmdByRotation(cwCommand string, ccwCommand string, cw180Command string, rotation Direction) (string, error) {
	switch rotation {
	case CW:
		return cwCommand, nil
	case CCW:
		return ccwCommand, nil
	case CW2:
		if cw180Command == "" {
			return "", errors.New("unsupported 180 rotation for this command")
		}
		return cw180Command, nil
	default:
		return "", errors.New("unknown rotation: " + rotation)
	}
}

// Example usage:
// config := CommandConfig{
// 	CMD_R_GRIP_CLOSE:     "CMD_R_GRIP_CLOSE",
// 	CMD_R_GRIP_OPEN:      "CMD_R_GRIP_OPEN",
// 	CMD_L_GRIP_CLOSE:     "CMD_L_GRIP_CLOSE",
// 	CMD_L_GRIP_OPEN:      "CMD_L_GRIP_OPEN",
// 	CMD_R_ROTATE_CW_90:   "CMD_R_ROTATE_CW_90",
// 	CMD_R_ROTATE_CCW_90:  "CMD_R_ROTATE_CCW_90",
// 	CMD_R_ROTATE_CW_180:  "CMD_R_ROTATE_CW_180",
// 	CMD_L_ROTATE_CW_90:   "CMD_L_ROTATE_CW_90",
// 	CMD_L_ROTATE_CCW_90:  "CMD_L_ROTATE_CCW_90",
// 	CMD_L_ROTATE_CW_180:  "CMD_L_ROTATE_CW_180",
// 	CMD_L_ROTATE_CCW_180: "CMD_L_ROTATE_CCW_180",
// 	CMD_R_ROTATE_CCW_180: "CMD_R_ROTATE_CCW_180",
// }
//
// parser := NewCubeParser(config)
// formula := "D R B R L' D L F L' F L2 F2 D F2 U D L2 D' F2 R2"
// commands, err := parser.ParseFormula(formula)
// if err != nil {
// 	fmt.Println("Error:", err)
// } else {
// 	fmt.Println("Commands:", commands)
// }
