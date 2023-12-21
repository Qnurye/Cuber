package config

import (
	"encoding/json"
	"os"
)

// CommandDelayConfig represents the command configuration structure.
type CommandDelayConfig struct {
	CmdRGripClose    int `json:"CMD_R_GRIP_CLOSE"`
	CmdRGripOpen     int `json:"CMD_R_GRIP_OPEN"`
	CmdLGripClose    int `json:"CMD_L_GRIP_CLOSE"`
	CmdLGripOpen     int `json:"CMD_L_GRIP_OPEN"`
	CmdRRotateCw90   int `json:"CMD_R_ROTATE_CW_90"`
	CmdRRotateCcw90  int `json:"CMD_R_ROTATE_CCW_90"`
	CmdRRotateCw180  int `json:"CMD_R_ROTATE_CW_180"`
	CmdLRotateCw90   int `json:"CMD_L_ROTATE_CW_90"`
	CmdLRotateCcw90  int `json:"CMD_L_ROTATE_CCW_90"`
	CmdLRotateCw180  int `json:"CMD_L_ROTATE_CW_180"`
	CmdLRotateCcw180 int `json:"CMD_L_ROTATE_CCW_180"`
	CmdRRotateCcw180 int `json:"CMD_R_ROTATE_CCW_180"`
}

// LoadCmdDelay loads the command configuration from a JSON file.
func LoadCmdDelay(filePath string) (*CommandDelayConfig, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var commandConfig CommandDelayConfig
	err = json.Unmarshal(content, &commandConfig)
	if err != nil {
		return nil, err
	}

	return &commandConfig, nil
}
