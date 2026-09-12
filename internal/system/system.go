package system

import (
	"fmt"
	"runtime"
)

type Info struct {
	OS           string
	Architecture string
	CPUCount     int
	GoVersion    string
}

func GetInfo() Info {
	return Info{
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		CPUCount:     runtime.NumCPU(),
		GoVersion:    runtime.Version(),
	}
}

func Display() {
	info := GetInfo()

	fmt.Println()
	fmt.Println("  [ SYSTEM INFORMATION ]")
	fmt.Println("  ----------------------")
	fmt.Printf("  OS           : %s\n", info.OS)
	fmt.Printf("  Architecture : %s\n", info.Architecture)
	fmt.Printf("  CPU Cores    : %d\n", info.CPUCount)
	fmt.Printf("  Go Runtime   : %s\n", info.GoVersion)
}