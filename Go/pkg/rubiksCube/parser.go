package rubiksCube

import (
	"errors"
	"fmt"
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

type CubeParser struct {
	Command *config.CommandConfig
	Delay   *config.CommandDelayConfig
}

func NewCubeParser(cmd *config.CommandConfig, delay *config.CommandDelayConfig) *CubeParser {
	return &CubeParser{Command: cmd, Delay: delay}
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

// CopyAndDouble 方法用于复制一段链表并将其长度延长一倍
func CopyAndDouble(original *LinkedList) *LinkedList {
	if original == nil || original.Head == nil {
		return nil
	}

	// 创建新的链表
	duplicated := &LinkedList{}

	// 复制原链表的每个节点并添加到新链表中
	current := original.Head
	for current != nil {
		duplicated.Append(current.Command)
		duplicated.Append(current.Command) // 复制节点

		current = current.Next
	}

	return duplicated
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
		} else {
			previous = current
			current = current.Next
		}
	}

	commands.Print()
	return commands
}

func (p *CubeParser) parseToken(token string) (*LinkedList, error) {
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
