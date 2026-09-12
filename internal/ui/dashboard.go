package ui

import (
	"fmt"
	"time"

	"titan/internal/core"
)

func Dashboard() {
	Clear()
	Logo()
	Header("Dashboard")

	fmt.Println()
	fmt.Printf("  Core Status     : %s\n", core.Status())
	fmt.Printf("  Version         : %s\n", core.Version)
	fmt.Printf("  Mode            : %s\n", core.Mode)
	fmt.Printf("  Runtime         : %s\n", time.Since(core.Runtime.StartTime).Round(time.Second))
	fmt.Printf("  Modules         : %d\n", len(core.Registry))

	enabled := 0
	for _, module := range core.Registry {
		if module.Enabled {
			enabled++
		}
	}

	fmt.Printf("  Enabled Modules  : %d\n", enabled)
	fmt.Printf("  UI Theme         : RED\n")

	fmt.Println()
	red.Println("  ----------------------------------------")
	bright.Println("           TITAN SYSTEM ONLINE")
	red.Println("  ----------------------------------------")

	Pause()
}