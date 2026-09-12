cd ~/TITAN

cat > internal/core/config_test.go <<'EOF'
package core

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigSaveAndLoad(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "titan.json")

	original := DefaultConfig()

	if err := SaveConfig(configPath, original); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	loaded, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if loaded.Name != original.Name {
		t.Fatalf("name mismatch: got %q, want %q", loaded.Name, original.Name)
	}

	if loaded.Version != original.Version {
		t.Fatalf("version mismatch: got %q, want %q", loaded.Version, original.Version)
	}

	if loaded.Mode != original.Mode {
		t.Fatalf("mode mismatch: got %q, want %q", loaded.Mode, original.Mode)
	}

	if loaded.Interface.Theme != original.Interface.Theme {
		t.Fatalf(
			"theme mismatch: got %q, want %q",
			loaded.Interface.Theme,
			original.Interface.Theme,
		)
	}

	if len(loaded.Modules) != len(original.Modules) {
		t.Fatalf(
			"module count mismatch: got %d, want %d",
			len(loaded.Modules),
			len(original.Modules),
		)
	}

	for name, enabled := range original.Modules {
		if loaded.Modules[name] != enabled {
			t.Fatalf(
				"module %q mismatch: got %v, want %v",
				name,
				loaded.Modules[name],
			)
		}
	}

	if _, err := os.Stat(configPath); err != nil {
		t.Fatalf("config file does not exist: %v", err)
	}
}

func TestLoadMissingConfig(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "missing.json")

	_, err := LoadConfig(configPath)

	if err == nil {
		t.Fatal("expected error when loading missing config")
	}
}
EOF

gofmt -w internal/core/config_test.go
go test ./internal/core