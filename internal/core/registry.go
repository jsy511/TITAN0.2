cd ~/TITAN

cat > internal/core/registry.go <<'EOF'
package core

type Module struct {
	Name        string
	Description string
	Enabled     bool
}

var Registry = []Module{
	{
		Name:        "dashboard",
		Description: "TITAN system dashboard",
		Enabled:     true,
	},
	{
		Name:        "terminal",
		Description: "Integrated terminal interface",
		Enabled:     true,
	},
	{
		Name:        "system",
		Description: "System information and diagnostics",
		Enabled:     true,
	},
	{
		Name:        "network",
		Description: "Network diagnostics and authorized testing",
		Enabled:     true,
	},
	{
		Name:        "dns",
		Description: "DNS diagnostics and analysis",
		Enabled:     true,
	},
	{
		Name:        "tools",
		Description: "Security research tool management",
		Enabled:     true,
	},
	{
		Name:        "modules",
		Description: "TITAN module management",
		Enabled:     true,
	},
	{
		Name:        "lab",
		Description: "Local security laboratory environment",
		Enabled:     true,
	},
	{
		Name:        "ctf",
		Description: "CTF and educational security exercises",
		Enabled:     true,
	},
	{
		Name:        "reports",
		Description: "Security assessment reports",
		Enabled:     true,
	},
	{
		Name:        "storage",
		Description: "TITAN data and storage management",
		Enabled:     true,
	},
	{
		Name:        "processes",
		Description: "Process monitoring",
		Enabled:     true,
	},
	{
		Name:        "packages",
		Description: "Package and dependency management",
		Enabled:     true,
	},
	{
		Name:        "logs",
		Description: "System and TITAN logs",
		Enabled:     true,
	},
	{
		Name:        "configuration",
		Description: "TITAN configuration management",
		Enabled:     true,
	},
	{
		Name:        "diagnostics",
		Description: "TITAN diagnostics and health checks",
		Enabled:     true,
	},
}

func GetModule(name string) *Module {
	for i := range Registry {
		if Registry[i].Name == name {
			return &Registry[i]
		}
	}

	return nil
}

func EnableModule(name string) bool {
	module := GetModule(name)

	if module == nil {
		return false
	}

	module.Enabled = true
	return true
}

func DisableModule(name string) bool {
	module := GetModule(name)

	if module == nil {
		return false
	}

	module.Enabled = false
	return true
}

func IsModuleEnabled(name string) bool {
	module := GetModule(name)

	if module == nil {
		return false
	}

	return module.Enabled
}
EOF

gofmt -w internal/core/registry.go
go test ./internal/core