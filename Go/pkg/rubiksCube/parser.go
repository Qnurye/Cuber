package rubiksCube

import (
	"errors"
	"fmt"
	"log"
	"qnurye/Cuber/pkg/config"
	"strings"
)

type Direction int

const (
	CCW Direction = 1 << iota // Counter-Clock Wise
	CW                        // Clock Wise
)

type Command struct {
	Operation string
	Delay     int
}

type ListNode struct {
	Command
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

type CubeStatus struct {
	Down    string
	Right   string
	Front   string
	Back    string
	Left    string
	Up      string
	LDegree int
	RDegree int
	LStatus GripStatus
	RStatus GripStatus
}

// type Hand int

type Hand string

const (
	//HandL Hand = iota
	//HandR
	HandL Hand = "HandL"
	HandR Hand = "HandR"
)

// type GripStatus int

type GripStatus string

const (
	Open  GripStatus = "Open"
	Close GripStatus = "Close"
)

type CubeParser struct {
	Command    *config.CommandConfig
	Delay      *config.CommandDelayConfig
	CubeStatus *CubeStatus
}

func NewCubeParser(cmd *config.CommandConfig, delay *config.CommandDelayConfig) *CubeParser {
	return &CubeParser{Command: cmd, Delay: delay, CubeStatus: &CubeStatus{
		Down:    "D",
		Right:   "R",
		Front:   "F",
		Back:    "B",
		Left:    "L",
		Up:      "U",
		LStatus: Close,
		RStatus: Close,
	}}
}

func (p *CubeParser) ParseFormula(formula string) (*LinkedList, error) {
	ts := strings.Fields(formula) // tokens
	cmdList := &LinkedList{}      // commands linked list

	for _, t := range ts {
		if t[0] == '(' {
			break
		}
		c, err := p.parseToken(t)
		if err != nil {
			return nil, err
		}
		ConnectLists(cmdList, c)
	}

	return p.optimize(cmdList), nil
}

func (list *LinkedList) Append(command Command) {
	newNode := &ListNode{Command: command, Next: nil}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func ConnectLists(list1, list2 *LinkedList) {
	if list1.Head == nil {
		list1.Head = list2.Head
	} else {
		current := list1.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = list2.Head
	}
}

func (list *LinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Print(current.Command.Operation, " ")
		current = current.Next
	}
	fmt.Println()
}

func (p *CubeParser) optimize(commands *LinkedList) *LinkedList {
	commands.Print()

	current := commands.Head
	previous := commands.Head
	for current != nil && current.Next != nil {
		if (current.Command.Operation == p.Command.CmdRGripClose &&
			current.Next.Command.Operation == p.Command.CmdRGripOpen) ||
			(current.Command.Operation == p.Command.CmdLGripClose &&
				current.Next.Command.Operation == p.Command.CmdLGripOpen) {
			previous.Next = current.Next.Next
			current = previous.Next
		} else if current.Command.Operation == current.Next.Command.Operation {
			switch current.Command.Operation {
			case p.Command.CmdLRotateCw90:
				current.Command.Operation = p.Command.CmdLRotateCw180
				current.Command.Delay = p.Delay.CmdLRotateCw180
			case p.Command.CmdLRotateCcw90:
				current.Command.Operation = p.Command.CmdLRotateCcw180
				current.Command.Delay = p.Delay.CmdLRotateCcw180
			case p.Command.CmdRRotateCw90:
				current.Command.Operation = p.Command.CmdRRotateCw180
				current.Command.Delay = p.Delay.CmdRRotateCw180
			case p.Command.CmdRRotateCcw90:
				current.Command.Operation = p.Command.CmdRRotateCcw180
				current.Command.Delay = p.Delay.CmdRRotateCw180
			}
			current.Next = current.Next.Next
			previous = current
			current = current.Next
		} else {
			previous = current
			current = current.Next
		}
	}

	ConnectLists(commands, p.gripCmd(HandL, Open))
	ConnectLists(commands, p.gripCmd(HandR, Open))

	commands.Print()
	return commands
}

func (p *CubeParser) parseToken(token string) (*LinkedList, error) {
	log.Printf("parsing: %v", token)
	if len(token) < 1 {
		return &LinkedList{}, errors.New("invalid token length")
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
		c1, nil := p.getCommand(token[:len(token)-1], CW)
		c2, nil := p.getCommand(token[:len(token)-1], CW)
		ConnectLists(c1, c2)
		return c1, nil
	default:
		return p.getCommand(token, CW)
	}
}
