package ai

import (
	"os"
	"strings"
)

const ConfigPath = "/etc/wtf/config"

func LoadAPIKey() string {
	// env override (CI / debug)
	if v := os.Getenv("WTF_AI_API_KEY"); v != "" {
		return v
	}

	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		return ""
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "WTF_AI_API_KEY=") {
			return strings.TrimPrefix(line, "WTF_AI_API_KEY=")
		}
	}

	return ""
}

