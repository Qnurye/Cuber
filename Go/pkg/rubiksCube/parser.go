package rubiksCube

import (
	"errors"
	"qnurye/Cuber/pkg/config"
	"strings"
)

// CubeParser represents a Rubik's Cube formula parser.
type CubeParser struct {
	Command *config.CommandConfig
	Delay   *config.CommandDelayConfig
}

// NewCubeParser creates a new CubeParser with the specified configuration.
func NewCubeParser(cmd *config.CommandConfig, delay *config.CommandDelayConfig) *CubeParser {
	return &CubeParser{Command: cmd, Delay: delay}
}

// ParseFormula parses a Rubik's Cube formula and returns the sequence of commands.
func (p *CubeParser) ParseFormula(formula string) ([]Command, error) {
	ts := strings.Fields(formula) // tokens
	cmd := make([]Command, 0)     // commands

	for _, t := range ts {
		c, err := p.parseToken(t)
		if err != nil {
			return nil, err
		}
		cmd = append(cmd, c[:]...)
	}

	return cmd, nil
}

type Direction int

const (
	CCW Direction = 1 << iota // Counter-Clock Wise
	CW                        // Clock Wise
)

func (p *CubeParser) parseToken(token string) ([]Command, error) {
	if len(token) < 1 {
		return make([]Command, 0), errors.New("invalid token length")
	}

	var d string
	if len(token) == 1 {
		d = ""
	} else {
		d = token[len(token)-1:] // direction
	}
	switch d {
	case "'":
		return p.getCommand(token[:len(token)-1], CCW)
	case "2":
		c, nil := p.getCommand(token[:len(token)-1], CW)
		return append(c, c[:]...), nil
	default:
		return p.getCommand(token, CW)
	}
}

type Command struct {
	Operation string
	Delay     int
}

func (p *CubeParser) getCommand(token string, dir Direction) ([]Command, error) {
	var (
		LClose  = Command{Operation: p.Command.CmdLGripClose, Delay: p.Delay.CmdLGripClose}
		RClose  = Command{Operation: p.Command.CmdRGripClose, Delay: p.Delay.CmdRGripClose}
		LOpen   = Command{Operation: p.Command.CmdLGripOpen, Delay: p.Delay.CmdLGripOpen}
		ROpen   = Command{Operation: p.Command.CmdRGripOpen, Delay: p.Delay.CmdRGripOpen}
		LCw90   = Command{Operation: p.Command.CmdLRotateCw90, Delay: p.Delay.CmdLRotateCw90}
		RCw90   = Command{Operation: p.Command.CmdRRotateCw90, Delay: p.Delay.CmdRRotateCw90}
		LCcw90  = Command{Operation: p.Command.CmdLRotateCcw90, Delay: p.Delay.CmdLRotateCcw90}
		RCcw90  = Command{Operation: p.Command.CmdRRotateCcw90, Delay: p.Delay.CmdRRotateCcw90}
		RCw180  = Command{Operation: p.Command.CmdRRotateCw180, Delay: p.Delay.CmdRRotateCw180}
		RCcw180 = Command{Operation: p.Command.CmdRRotateCcw180, Delay: p.Delay.CmdRRotateCcw180}
		LCw180  = Command{Operation: p.Command.CmdLRotateCw180, Delay: p.Delay.CmdLRotateCw180}
		LCcw180 = Command{Operation: p.Command.CmdLRotateCcw180, Delay: p.Delay.CmdLRotateCcw180}
	)

	switch token {
	case "R": // Right
		if dir == CW {
			return []Command{
				RCw90,
				ROpen,
				RCcw90,
				RClose,
			}, nil
		} else {
			return []Command{
				RCcw90,
				ROpen,
				RCw90,
				RClose,
			}, nil
		}
	case "D": // Down
		if dir == CCW { // 左臂是 Down, 与公式相反
			return []Command{
				LCw90,
				LOpen,
				LCcw90,
				LClose,
			}, nil
		} else {
			return []Command{
				LCcw90,
				LOpen,
				LCw90,
				LClose,
			}, nil
		}
	case "F": // Front
		if dir == CW {
			return []Command{
				ROpen,
				LCcw90,
				RClose,
				LOpen,
				LCw90,
				LClose,
				RCcw90,
				ROpen,
				RCw90,
				LCw90,
				RClose,
				LOpen,
				LCcw90,
				LClose,
			}, nil
		} else {
			return []Command{
				ROpen,
				LCcw90,
				RClose,
				LOpen,
				LCw90,
				LClose,
				RCw90,
				ROpen,
				RCcw90,
				LCw90,
				RClose,
				LOpen,
				LCcw90,
				LClose,
			}, nil
		}
	case "B": // Back
		if dir == CW {
			return []Command{
				LOpen,
				RCcw90,
				LClose,
				ROpen,
				RCw90,
				RClose,
				LCcw90,
				LOpen,
				LCw90,
				RCw90,
				LClose,
				ROpen,
				RCcw90,
				RClose,
			}, nil
		} else {
			return []Command{
				LOpen,
				RCcw90,
				LClose,
				ROpen,
				RCw90,
				RClose,
				LCw90,
				LOpen,
				LCcw90,
				RCw90,
				LClose,
				ROpen,
				RCcw90,
				RClose,
			}, nil
		}
	case "U": // Up
		if dir == CW {
			return []Command{
				LOpen,
				RCw180,
				LClose,
				LCcw90,
				LOpen,
				LCw90,
				RCcw180,
				LClose,
			}, nil
		} else {
			return []Command{
				LOpen,
				RCw180,
				LClose,
				LCw90,
				LOpen,
				LCcw90,
				RCcw180,
				LClose,
			}, nil
		}
	case "L": // Left
		if dir == CW {
			return []Command{
				ROpen,
				LCw180,
				RClose,
				RCcw90,
				ROpen,
				RCw90,
				LCcw180,
				RClose,
			}, nil
		} else {
			return []Command{
				ROpen,
				LCw180,
				RClose,
				RCw90,
				ROpen,
				RCcw90,
				LCcw180,
				RClose,
			}, nil
		}
	default:
		return nil, errors.New("invalid token")
	}
}
