package config

import (
	"encoding/json"
	"os"
)

// CommandConfig represents the command configuration structure.
type CommandConfig struct {
	CmdRGripClose    string `json:"CMD_R_GRIP_CLOSE"`
	CmdRGripOpen     string `json:"CMD_R_GRIP_OPEN"`
	CmdLGripClose    string `json:"CMD_L_GRIP_CLOSE"`
	CmdLGripOpen     string `json:"CMD_L_GRIP_OPEN"`
	CmdRRotateCw90   string `json:"CMD_R_ROTATE_CW_90"`
	CmdRRotateCcw90  string `json:"CMD_R_ROTATE_CCW_90"`
	CmdRRotateCw180  string `json:"CMD_R_ROTATE_CW_180"`
	CmdLRotateCw90   string `json:"CMD_L_ROTATE_CW_90"`
	CmdLRotateCcw90  string `json:"CMD_L_ROTATE_CCW_90"`
	CmdLRotateCw180  string `json:"CMD_L_ROTATE_CW_180"`
	CmdLRotateCcw180 string `json:"CMD_L_ROTATE_CCW_180"`
	CmdRRotateCcw180 string `json:"CMD_R_ROTATE_CCW_180"`
}

// LoadCmd loads the command configuration from a JSON file.
func LoadCmd(filePath string) (*CommandConfig, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var commandConfig CommandConfig
	err = json.Unmarshal(content, &commandConfig)
	if err != nil {
		return nil, err
	}

	return &commandConfig, nil
}
