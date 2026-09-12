cd ~/TITAN

cat > internal/ui/menu.go <<'EOF'
package ui

import (
	"fmt"
	"strconv"
)

type MenuItem struct {
	ID          string
	Name        string
	Description string
}

var Menu = []MenuItem{
	{"01", "Dashboard", "TITAN system dashboard"},
	{"02", "Terminal", "Integrated terminal"},
	{"03", "System", "System information"},
	{"04", "Network", "Network diagnostics"},
	{"05", "DNS", "DNS diagnostics"},
	{"06", "Tools", "Security research tools"},
	{"07", "Modules", "Module management"},
	{"08", "Lab", "Local security laboratory"},
	{"09", "CTF", "CTF and educational exercises"},
	{"10", "Reports", "Security reports"},
	{"11", "Storage", "Storage management"},
	{"12", "Processes", "Process monitoring"},
	{"13", "Packages", "Package management"},
	{"14", "Logs", "System and TITAN logs"},
	{"15", "Configuration", "Configuration management"},
	{"16", "Diagnostics", "TITAN diagnostics"},
	{"17", "About TITAN", "TITAN information"},
	{"00", "Exit", "Exit TITAN"},
}

func DrawMenu() {
	Header("Main Menu")

	for _, item := range Menu {
		red.Printf("  [%s] ", item.ID)
		fmt.Printf("%-16s ", item.Name)
		dim.Printf("%s\n", item.Description)
	}
}

func FindMenuItem(input string) *MenuItem {
	if len(input) == 1 {
		input = "0" + input
	}

	for i := range Menu {
		if Menu[i].ID == input {
			return &Menu[i]
		}
	}

	return nil
}

func MenuChoice(input string) *MenuItem {
	if _, err := strconv.Atoi(input); err != nil {
		return nil
	}

	return FindMenuItem(input)
}
EOF

gofmt -w internal/ui/menu.go

go test ./...
go build -o titan ./cmd/titan

git add internal/ui/menu.go
git commit -m "Add TITAN UI menu engine"
git push origin main