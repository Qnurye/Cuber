package rubiksCube

import (
	"errors"
	"fmt"
)

func (p *CubeParser) gripCmd(hand Hand, status GripStatus) *LinkedList {
	cmdList := &LinkedList{}
	if hand == HandL {
		if status == Open {
			cmdList.Append(Command{p.Command.CmdLGripOpen, p.Delay.CmdLGripOpen})
		} else {
			cmdList.Append(Command{p.Command.CmdLGripClose, p.Delay.CmdLGripClose})
		}
		p.CubeStatus.LStatus = status
	} else {
		if status == Open {
			cmdList.Append(Command{p.Command.CmdRGripOpen, p.Delay.CmdRGripOpen})
		} else {
			cmdList.Append(Command{p.Command.CmdRGripClose, p.Delay.CmdRGripClose})
		}
		p.CubeStatus.RStatus = status
	}
	fmt.Printf("[GRIP] %v to %v,\n", hand, status)
	p.getStatus()
	return cmdList
}

func (p *CubeParser) getCommand(token string, dir Direction) (*LinkedList, error) {
	cmdList := &LinkedList{}

	switch token {
	case p.CubeStatus.Right: // Right
		if dir == CW {
			ConnectLists(cmdList, p.rotateCmd(HandR, -90))
		} else {
			ConnectLists(cmdList, p.rotateCmd(HandR, 90))
		}
	case p.CubeStatus.Down: // Left Hand
		if dir == CW {
			ConnectLists(cmdList, p.rotateCmd(HandL, -90))
		} else {
			ConnectLists(cmdList, p.rotateCmd(HandL, 90))
		}
	case p.CubeStatus.Front:
		ConnectLists(cmdList, p.gripCmd(HandL, Open))
		ConnectLists(cmdList, p.rotateCmd(HandR, 90))
		ConnectLists(cmdList, p.gripCmd(HandL, Close))
		if dir == CW {
			ConnectLists(cmdList, p.rotateCmd(HandL, -90))
		} else {
			ConnectLists(cmdList, p.rotateCmd(HandL, 90))
		}
	case p.CubeStatus.Back:
		ConnectLists(cmdList, p.gripCmd(HandL, Open))
		ConnectLists(cmdList, p.rotateCmd(HandR, -90))
		ConnectLists(cmdList, p.gripCmd(HandL, Close))
		if dir == CW {
			ConnectLists(cmdList, p.rotateCmd(HandL, -90))
		} else {
			ConnectLists(cmdList, p.rotateCmd(HandL, 90))
		}
	case p.CubeStatus.Up:
		ConnectLists(cmdList, p.gripCmd(HandL, Open))
		ConnectLists(cmdList, p.rotateCmd(HandR, 180))
		ConnectLists(cmdList, p.gripCmd(HandL, Close))
		if dir == CW {
			ConnectLists(cmdList, p.rotateCmd(HandL, -90))
		} else {
			ConnectLists(cmdList, p.rotateCmd(HandL, 90))
		}
	case p.CubeStatus.Left:
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

func (p *CubeParser) transformFaces(hand Hand, degree int) error {
	if degree < 0 {
		degree += 360
	}

	if hand == HandL {
		if degree%360 == 0 {
		} else if degree%270 == 0 {
			p.CubeStatus.Front, p.CubeStatus.Right, p.CubeStatus.Back, p.CubeStatus.Left = p.CubeStatus.Left, p.CubeStatus.Front, p.CubeStatus.Right, p.CubeStatus.Back
		} else if degree%180 == 0 {
			p.CubeStatus.Front, p.CubeStatus.Right, p.CubeStatus.Back, p.CubeStatus.Left = p.CubeStatus.Back, p.CubeStatus.Left, p.CubeStatus.Front, p.CubeStatus.Right
		} else if degree%90 == 0 {
			p.CubeStatus.Front, p.CubeStatus.Right, p.CubeStatus.Back, p.CubeStatus.Left = p.CubeStatus.Right, p.CubeStatus.Back, p.CubeStatus.Left, p.CubeStatus.Front
		}
	} else {
		if degree%360 == 0 {
		} else if degree%270 == 0 {
			p.CubeStatus.Front, p.CubeStatus.Down, p.CubeStatus.Back, p.CubeStatus.Up = p.CubeStatus.Down, p.CubeStatus.Back, p.CubeStatus.Up, p.CubeStatus.Front
		} else if degree%180 == 0 {
			p.CubeStatus.Front, p.CubeStatus.Down, p.CubeStatus.Back, p.CubeStatus.Up = p.CubeStatus.Back, p.CubeStatus.Up, p.CubeStatus.Front, p.CubeStatus.Down
		} else if degree%90 == 0 {
			p.CubeStatus.Front, p.CubeStatus.Down, p.CubeStatus.Back, p.CubeStatus.Up = p.CubeStatus.Up, p.CubeStatus.Front, p.CubeStatus.Down, p.CubeStatus.Back
		}
	}

	fmt.Printf("[TRANSFORM] %v by %v,\n", hand, degree)
	p.getStatus()
	return nil
}

func (p *CubeParser) fixRotate(hand Hand, degree int) *LinkedList {
	cmdList := &LinkedList{}

	if hand == HandL {
		if p.CubeStatus.LStatus == Open {
			ConnectLists(cmdList, p.rotateCmd(HandL, degree))
		} else {
			ConnectLists(cmdList, p.gripCmd(HandL, Open))
			ConnectLists(cmdList, p.rotateCmd(HandL, degree))
			ConnectLists(cmdList, p.gripCmd(HandL, Close))
		}
	} else {
		if p.CubeStatus.RStatus == Open {
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
		if (hand == HandL && p.CubeStatus.LDegree != 0) || (hand == HandR && p.CubeStatus.RDegree != 0) {
			degree = -degree
		}
	} else {
		// 90 度防缠绕
		if hand == HandL {
			if p.CubeStatus.LDegree > 180 {
				ConnectLists(cmdList, p.fixRotate(HandL, p.CubeStatus.LDegree))
			}
		} else {
			if p.CubeStatus.RDegree > 180 {
				ConnectLists(cmdList, p.fixRotate(HandR, p.CubeStatus.RDegree))
			}
		}
	}

	// L
	if hand == HandL {
		// 先归位另一边
		if p.CubeStatus.RDegree%180 != 0 {
			ConnectLists(cmdList, p.fixRotate(HandR, -p.CubeStatus.RDegree))
		}
		// 换面了
		if p.CubeStatus.RStatus == Open {
			_ = p.transformFaces(hand, degree)
		}

		if dir == CW {
			if degree == 90 || degree == -90 {
				cmdList.Append(Command{p.Command.CmdLRotateCw90, p.Delay.CmdLRotateCw90})
				p.CubeStatus.LDegree += 90
			} else if degree == 180 || degree == -180 {
				cmdList.Append(Command{p.Command.CmdLRotateCw180, p.Delay.CmdLRotateCw180})
				p.CubeStatus.LDegree += 180
			}
		} else {
			if degree == 90 || degree == -90 {
				cmdList.Append(Command{p.Command.CmdLRotateCcw90, p.Delay.CmdLRotateCcw90})
				p.CubeStatus.LDegree -= 90
			} else if degree == 180 || degree == -180 {
				cmdList.Append(Command{p.Command.CmdLRotateCcw180, p.Delay.CmdLRotateCcw180})
				p.CubeStatus.LDegree -= 180
			}
		}
	} else {
		// 先归位另一边
		if p.CubeStatus.LDegree%180 != 0 {
			ConnectLists(cmdList, p.fixRotate(HandL, -p.CubeStatus.LDegree))
		}
		// 换面了
		if p.CubeStatus.LStatus == Open {
			_ = p.transformFaces(hand, degree)
		}

		if dir == CW {
			if degree == 90 || degree == -90 {
				cmdList.Append(Command{p.Command.CmdRRotateCw90, p.Delay.CmdRRotateCw90})
				p.CubeStatus.RDegree += 90
			} else if degree == 180 || degree == -180 {
				cmdList.Append(Command{p.Command.CmdRRotateCw180, p.Delay.CmdRRotateCw180})
				p.CubeStatus.RDegree += 180
			}
		} else {
			if degree == 90 || degree == -90 {
				cmdList.Append(Command{p.Command.CmdRRotateCcw90, p.Delay.CmdRRotateCcw90})
				p.CubeStatus.RDegree -= 90
			} else if degree == 180 || degree == -180 {
				cmdList.Append(Command{p.Command.CmdRRotateCcw180, p.Delay.CmdRRotateCcw180})
				p.CubeStatus.RDegree -= 180
			}
		}
	}

	fmt.Printf("[ROTATE] %v by %v,\n", hand, degree)
	p.getStatus()
	return cmdList
}

func (p *CubeParser) getStatus() {
	fmt.Printf("Status: %v\n\n", p.CubeStatus)
}
