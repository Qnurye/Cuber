package config

import (
	"encoding/json"
	"os"
)

// CommandConfig represents the command configuration structure.
type CommandConfig struct {
	CMD_R_GRIP_CLOSE     string `json:"CMD_R_GRIP_CLOSE"`
	CMD_R_GRIP_OPEN      string `json:"CMD_R_GRIP_OPEN"`
	CMD_L_GRIP_CLOSE     string `json:"CMD_L_GRIP_CLOSE"`
	CMD_L_GRIP_OPEN      string `json:"CMD_L_GRIP_OPEN"`
	CMD_R_ROTATE_CW_90   string `json:"CMD_R_ROTATE_CW_90"`
	CMD_R_ROTATE_CCW_90  string `json:"CMD_R_ROTATE_CCW_90"`
	CMD_R_ROTATE_CW_180  string `json:"CMD_R_ROTATE_CW_180"`
	CMD_L_ROTATE_CW_90   string `json:"CMD_L_ROTATE_CW_90"`
	CMD_L_ROTATE_CCW_90  string `json:"CMD_L_ROTATE_CCW_90"`
	CMD_L_ROTATE_CW_180  string `json:"CMD_L_ROTATE_CW_180"`
	CMD_L_ROTATE_CCW_180 string `json:"CMD_L_ROTATE_CCW_180"`
	CMD_R_ROTATE_CCW_180 string `json:"CMD_R_ROTATE_CCW_180"`
}

// Load loads the command configuration from a JSON file.
func LoadCmd() (*CommandConfig, error) {
	content, err := os.ReadFile("./config/cmd.json")
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
