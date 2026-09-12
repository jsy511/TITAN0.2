cd ~/TITAN

cat > internal/core/core_test.go <<'EOF'
package core

import "testing"

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config.Name != Name {
		t.Fatalf("expected config name %q, got %q", Name, config.Name)
	}

	if config.Version != Version {
		t.Fatalf("expected config version %q, got %q", Version, config.Version)
	}

	if config.Mode != Mode {
		t.Fatalf("expected config mode %q, got %q", Mode, config.Mode)
	}

	if config.Interface.Theme != "red" {
		t.Fatalf("expected red theme, got %q", config.Interface.Theme)
	}

	if len(config.Modules) == 0 {
		t.Fatal("expected modules in default configuration")
	}
}

func TestModuleRegistry(t *testing.T) {
	if len(Registry) == 0 {
		t.Fatal("module registry is empty")
	}

	for _, module := range Registry {
		if module.Name == "" {
			t.Fatal("module has empty name")
		}

		if GetModule(module.Name) == nil {
			t.Fatalf("module %q cannot be retrieved", module.Name)
		}
	}
}

func TestModuleEnableDisable(t *testing.T) {
	module := GetModule("dashboard")

	if module == nil {
		t.Fatal("dashboard module not found")
	}

	DisableModule("dashboard")

	if IsModuleEnabled("dashboard") {
		t.Fatal("dashboard should be disabled")
	}

	EnableModule("dashboard")

	if !IsModuleEnabled("dashboard") {
		t.Fatal("dashboard should be enabled")
	}
}

func TestHealthCheck(t *testing.T) {
	results := HealthCheck()

	if len(results) == 0 {
		t.Fatal("health check returned no results")
	}

	for _, result := range results {
		if result.Name == "" {
			t.Fatal("health check returned an unnamed component")
		}

		if result.Status == "" {
			t.Fatalf("health check returned empty status for %q", result.Name)
		}
	}
}

func TestCoreStatus(t *testing.T) {
	Runtime.Running = true

	if Status() != "ONLINE" {
		t.Fatalf("expected ONLINE status, got %q", Status())
	}

	Runtime.Running = false

	if Status() != "OFFLINE" {
		t.Fatalf("expected OFFLINE status, got %q", Status())
	}

	Runtime.Running = true
}
EOF

gofmt -w internal/core/core_test.go
go test ./internal/core