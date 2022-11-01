package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load() Settings {
	var settings Settings
	filePath := fmt.Sprintf("./.config/%s.json", "local")
	fileBytes, err := os.ReadFile(filePath)

	if err != nil {
		return Settings{}
	}

	if err := json.Unmarshal(fileBytes, &settings); err != nil {
		return Settings{}
	}

	return settings
}
