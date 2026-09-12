cd ~/TITAN

cat > cmd/titan/main.go <<'EOF'
package main

import (
	"fmt"
	"os"

	"titan/internal/core"
	"titan/internal/ui"
)

func main() {
	ui.Clear()
	ui.Logo()

	core.Startup()

	for core.Runtime.Running {
		ui.DrawMenu()

		choice := ui.Prompt()
		item := ui.MenuChoice(choice)

		if item == nil {
			ui.Error("Invalid menu selection.")
			ui.Pause()
			continue
		}

		if item.ID == "00" {
			core.Shutdown()
			break
		}

		ui.Clear()
		ui.Logo()

		ui.Header(item.Name)
		fmt.Printf("\n  %s\n", item.Description)

		ui.Success("Module selected: " + item.Name)

		ui.Pause()
		ui.Clear()
		ui.Logo()
	}

	fmt.Println()
	os.Exit(0)
}
EOF

gofmt -w cmd/titan/main.go

go test ./...
go build -o titan ./cmd/titan