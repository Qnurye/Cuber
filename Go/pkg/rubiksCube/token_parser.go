package rubiksCube

import (
	"errors"
	"fmt"
)

var (
	Down    = "D"
	Right   = "R"
	Front   = "F"
	Back    = "B"
	Left    = "L"
	Up      = "U"
	LDegree = 0
	RDegree = 0
	LStatus = Close
	RStatus = Close
)

type Hand int

const (
	HandL Hand = iota
	HandR
)

type GripStatus int

const (
	Open GripStatus = iota
	Close
)

func (p *CubeParser) gripCmd(hand Hand, status GripStatus) *LinkedList {
	cmdList := &LinkedList{}
	if hand == HandL {
		if status == Open {
			cmdList.Append(Command{p.Command.CmdLGripOpen, p.Delay.CmdLGripOpen})
		} else {
			cmdList.Append(Command{p.Command.CmdLGripClose, p.Delay.CmdLGripClose})
		}
		LStatus = status
	} else {
		if status == Open {
			cmdList.Append(Command{p.Command.CmdRGripOpen, p.Delay.CmdRGripOpen})
		} else {
			cmdList.Append(Command{p.Command.CmdRGripClose, p.Delay.CmdRGripClose})
		}
		RStatus = status
	}
	fmt.Printf("after grip %v to %v,\n", hand, status)
	getStatus()
	return cmdList
}

func (p *CubeParser) getCommand(token string, dir Direction) (*LinkedList, error) {
	cmdList := &LinkedList{}

	switch token {
	case Right: // Right
		if dir == CW {
			ConnectLists(cmdList, p.rotateCmd(HandR, -90))
		} else {
			ConnectLists(cmdList, p.rotateCmd(HandR, 90))
		}
	case Down: // Left Hand
		if dir == CW {
			ConnectLists(cmdList, p.rotateCmd(HandL, -90))
		} else {
			ConnectLists(cmdList, p.rotateCmd(HandL, 90))
		}
	case Front:
		ConnectLists(cmdList, p.gripCmd(HandL, Open))
		ConnectLists(cmdList, p.rotateCmd(HandR, 90))
		ConnectLists(cmdList, p.gripCmd(HandL, Close))
		if dir == CW {
			ConnectLists(cmdList, p.rotateCmd(HandL, -90))
		} else {
			ConnectLists(cmdList, p.rotateCmd(HandL, 90))
		}
	case Back:
		ConnectLists(cmdList, p.gripCmd(HandL, Open))
		ConnectLists(cmdList, p.rotateCmd(HandR, -90))
		ConnectLists(cmdList, p.gripCmd(HandL, Close))
		if dir == CW {
			ConnectLists(cmdList, p.rotateCmd(HandL, -90))
		} else {
			ConnectLists(cmdList, p.rotateCmd(HandL, 90))
		}
	case Up:
		ConnectLists(cmdList, p.gripCmd(HandL, Open))
		ConnectLists(cmdList, p.rotateCmd(HandR, 180))
		ConnectLists(cmdList, p.gripCmd(HandL, Close))
		if dir == CW {
			ConnectLists(cmdList, p.rotateCmd(HandL, -90))
		} else {
			ConnectLists(cmdList, p.rotateCmd(HandL, 90))
		}
	case Left:
		ConnectLists(cmdList, p.gripCmd(HandR, Open))
		ConnectLists(cmdList, p.rotateCmd(HandL, 180))
		ConnectLists(cmdList, p.gripCmd(HandR, Close))
		if dir == CW {
			ConnectLists(cmdList, p.rotateCmd(HandR, -90))
		} else {
			ConnectLists(cmdList, p.rotateCmd(HandR, 90))
		}
	default:
		return nil, errors.New("invalid token")
	}
	return cmdList, nil
}

func (p *CubeParser) transformFaces(hand Hand, dir Direction, degree int) error {
	if dir == CCW {
		degree = 360 - degree
	}

	if hand == HandL {
		if degree%360 == 0 {
		} else if degree%270 == 0 {
			Front, Right, Back, Left = Left, Front, Right, Back
		} else if degree%180 == 0 {
			Front, Right, Back, Left = Back, Left, Front, Right
		} else if degree%90 == 0 {
			Front, Right, Back, Left = Right, Back, Left, Front
		}
	} else {
		if degree%360 == 0 {
		} else if degree%270 == 0 {
			Front, Down, Back, Up = Down, Back, Up, Front
		} else if degree%180 == 0 {
			Front, Down, Back, Up = Back, Up, Front, Down
		} else if degree%90 == 0 {
			Front, Down, Back, Up = Up, Front, Down, Back
		}
	}

	fmt.Printf("after transforming %v by %v of %v,\n", hand, degree, dir)
	getStatus()
	return nil
}

func (p *CubeParser) fixRotate(hand Hand, degree int) *LinkedList {
	cmdList := &LinkedList{}

	if hand == HandL {
		if LStatus == Open {
			ConnectLists(cmdList, p.rotateCmd(HandL, degree))
		} else {
			ConnectLists(cmdList, p.gripCmd(HandL, Open))
			ConnectLists(cmdList, p.rotateCmd(HandL, degree))
			ConnectLists(cmdList, p.gripCmd(HandL, Close))
		}
	} else {
		if RStatus == Open {
			ConnectLists(cmdList, p.rotateCmd(HandR, degree))
		} else {
			ConnectLists(cmdList, p.gripCmd(HandR, Open))
			ConnectLists(cmdList, p.rotateCmd(HandR, degree))
			ConnectLists(cmdList, p.gripCmd(HandR, Close))
		}
	}

	return cmdList
}

func (p *CubeParser) rotateCmd(hand Hand, degree int) *LinkedList {
	cmdList := &LinkedList{}

	var dir Direction
	if degree < 0 {
		dir = CCW
	} else {
		dir = CW
	}

	if degree == 180 || degree == -180 {
		// 180度防缠绕
		if (hand == HandL && LDegree != 0) || (hand == HandR && RDegree != 0) {
			degree = -degree
		}
	} else {
		// 90 度防缠绕
		if hand == HandL {
			if LDegree > 180 {
				ConnectLists(cmdList, p.fixRotate(HandL, LDegree))
			}
		} else {
			if RDegree > 180 {
				ConnectLists(cmdList, p.fixRotate(HandR, RDegree))
			}
		}
	}

	// L
	if hand == HandL {
		// 先归位另一边
		if RDegree%180 != 0 {
			ConnectLists(cmdList, p.fixRotate(HandR, -RDegree))
		}
		// 换面了
		if RStatus == Open {
			_ = p.transformFaces(hand, dir, degree)
		}

		if dir == CW {
			if degree == 90 || degree == -90 {
				cmdList.Append(Command{p.Command.CmdLRotateCw90, p.Delay.CmdLRotateCw90})
				LDegree += 90
			} else if degree == 180 || degree == -180 {
				cmdList.Append(Command{p.Command.CmdLRotateCw180, p.Delay.CmdLRotateCw180})
				LDegree += 180
			}
		} else {
			if degree == 90 || degree == -90 {
				cmdList.Append(Command{p.Command.CmdLRotateCcw90, p.Delay.CmdLRotateCcw90})
				LDegree -= 90
			} else if degree == 180 || degree == -180 {
				cmdList.Append(Command{p.Command.CmdLRotateCcw180, p.Delay.CmdLRotateCcw180})
				LDegree -= 180
			}
		}
	} else {
		// 先归位另一边
		if LDegree%180 != 0 {
			ConnectLists(cmdList, p.fixRotate(HandL, -LDegree))
		}
		// 换面了
		if LStatus == Open {
			_ = p.transformFaces(hand, dir, degree)
		}

		if dir == CW {
			if degree == 90 || degree == -90 {
				cmdList.Append(Command{p.Command.CmdRRotateCw90, p.Delay.CmdRRotateCw90})
				RDegree += 90
			} else if degree == 180 || degree == -180 {
				cmdList.Append(Command{p.Command.CmdRRotateCw180, p.Delay.CmdRRotateCw180})
				RDegree += 180
			}
		} else {
			if degree == 90 || degree == -90 {
				cmdList.Append(Command{p.Command.CmdRRotateCcw90, p.Delay.CmdRRotateCcw90})
				RDegree -= 90
			} else if degree == 180 || degree == -180 {
				cmdList.Append(Command{p.Command.CmdRRotateCcw180, p.Delay.CmdRRotateCcw180})
				RDegree -= 180
			}
		}
	}

	fmt.Printf("after rotating %v by %v,\n", hand, degree)
	getStatus()
	return cmdList
}

func getStatus() {
	fmt.Printf("Left: %v @ %v\tRight: %v @ %v\n\n", LStatus, LDegree, RStatus, RDegree)
}
