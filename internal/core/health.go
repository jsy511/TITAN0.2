cd ~/TITAN

cat > internal/core/health.go <<'EOF'
package core

type HealthStatus struct {
	Name    string
	Status  string
	Details string
}

func HealthCheck() []HealthStatus {
	results := []HealthStatus{}

	// Core runtime
	coreStatus := "ONLINE"
	if !Runtime.Running {
		coreStatus = "OFFLINE"
	}

	results = append(results, HealthStatus{
		Name:    "Core Runtime",
		Status:  coreStatus,
		Details: "TITAN core runtime state",
	})

	// Module registry
	registryStatus := "ONLINE"
	if len(Registry) == 0 {
		registryStatus = "WARNING"
	}

	results = append(results, HealthStatus{
		Name:    "Module Registry",
		Status:  registryStatus,
		Details: "Registered TITAN modules",
	})

	// Configuration
	config := DefaultConfig()

	configStatus := "ONLINE"
	if config.Name == "" || config.Version == "" || config.Mode == "" {
		configStatus = "WARNING"
	}

	results = append(results, HealthStatus{
		Name:    "Configuration",
		Status:  configStatus,
		Details: "Default configuration validation",
	})

	// Enabled modules
	enabled := 0

	for _, module := range Registry {
		if module.Enabled {
			enabled++
		}
	}

	moduleStatus := "ONLINE"

	if enabled == 0 {
		moduleStatus = "WARNING"
	}

	results = append(results, HealthStatus{
		Name:    "Modules",
		Status:  moduleStatus,
		Details: "Enabled modules: " + itoa(enabled),
	})

	return results
}

func itoa(value int) string {
	if value == 0 {
		return "0"
	}

	digits := ""

	for value > 0 {
		digit := value % 10
		digits = string(rune('0'+digit)) + digits
		value /= 10
	}

	return digits
}
EOF

gofmt -w internal/core/health.go
go test ./internal/core