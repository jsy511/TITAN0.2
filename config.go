cd ~/TITAN

cat > internal/core/config.go <<'EOF'
package core

import (
	"encoding/json"
	"os"
)

type Config struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Mode    string `json:"mode"`

	Interface struct {
		Theme string `json:"theme"`
	} `json:"interface"`

	Modules map[string]bool `json:"modules"`
}

func DefaultConfig() Config {
	config := Config{
		Name:    Name,
		Version: Version,
		Mode:    Mode,
	}

	config.Interface.Theme = "red"

	config.Modules = map[string]bool{
		"dashboard":     true,
		"terminal":      true,
		"system":        true,
		"network":       true,
		"dns":           true,
		"tools":         true,
		"modules":       true,
		"lab":           true,
		"ctf":           true,
		"reports":       true,
		"storage":       true,
		"processes":     true,
		"packages":      true,
		"logs":          true,
		"configuration": true,
		"diagnostics":   true,
	}

	return config
}

func SaveConfig(path string, config Config) error {
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
EOF

gofmt -w internal/core/config.go
go test ./internal/core