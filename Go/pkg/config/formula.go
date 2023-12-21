package config

import (
	"encoding/json"
	"os"
)

// FormulaConfig represents the command configuration structure.
type FormulaConfig struct {
	Formula string `json:"formula"`
}

// LoadFormula loads the command configuration from a JSON file.
func LoadFormula() (*FormulaConfig, error) {
	content, err := os.ReadFile("./config/formula.json")
	if err != nil {
		return nil, err
	}

	var commandConfig FormulaConfig
	err = json.Unmarshal(content, &commandConfig)
	if err != nil {
		return nil, err
	}

	return &commandConfig, nil
}
