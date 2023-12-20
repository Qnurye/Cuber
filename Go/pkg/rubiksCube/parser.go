package rubiksCube

import (
	"errors"
	"qnurye/Cuber/pkg/config"
	"strings"
)

// CubeParser represents a Rubik's Cube formula parser.
type CubeParser struct {
	Config *config.CommandConfig
}

// NewCubeParser creates a new CubeParser with the specified configuration.
func NewCubeParser(config *config.CommandConfig) *CubeParser {
	return &CubeParser{Config: config}
}

// ParseFormula parses a Rubik's Cube formula and returns the sequence of commands.
func (p *CubeParser) ParseFormula(formula string) ([]string, error) {
	ts := strings.Fields(formula) // tokens
	cmd := make([]string, 0)      // commands

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

func (p *CubeParser) parseToken(token string) ([]string, error) {
	if len(token) < 1 {
		return make([]string, 0), errors.New("invalid token length")
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

func (p *CubeParser) getCommand(token string, dir Direction) ([]string, error) {
	var (
		LClose  = p.Config.CmdLGripClose
		RClose  = p.Config.CmdRGripClose
		LOpen   = p.Config.CmdLGripOpen
		ROpen   = p.Config.CmdRGripOpen
		LCw90   = p.Config.CmdLRotateCw90
		RCw90   = p.Config.CmdRRotateCw90
		LCcw90  = p.Config.CmdLRotateCcw90
		RCcw90  = p.Config.CmdRRotateCcw90
		RCw180  = p.Config.CmdRRotateCw180
		RCcw180 = p.Config.CmdRRotateCcw180
		LCw180  = p.Config.CmdLRotateCw180
		LCcw180 = p.Config.CmdLRotateCcw180
	)

	switch token {
	case "R": // Right
		if dir == CW {
			return []string{
				RCw90,
				ROpen,
				RCcw90,
				RClose,
			}, nil
		} else {
			return []string{
				RCcw90,
				ROpen,
				RCw90,
				RClose,
			}, nil
		}
	case "D": // Down
		if dir == CCW { // 左臂是 Down, 与公式相反
			return []string{
				LCw90,
				LOpen,
				LCcw90,
				LClose,
			}, nil
		} else {
			return []string{
				LCcw90,
				LOpen,
				LCw90,
				LClose,
			}, nil
		}
	case "F": // Front
		if dir == CW {
			return []string{
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
			return []string{
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
			return []string{
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
			return []string{
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
			return []string{
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
			return []string{
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
			return []string{
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
			return []string{
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
