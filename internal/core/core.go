package core

import (
	"fmt"
	"time"
)

const (
	Name    = "TITAN"
	Version = "0.1.0"
	Mode    = "RESEARCH"
)

type State struct {
	Name      string
	Version   string
	Mode      string
	StartTime time.Time
	Running   bool
}

var Runtime = State{
	Name:      Name,
	Version:   Version,
	Mode:      Mode,
	StartTime: time.Now(),
	Running:   true,
}

func Startup() {
	fmt.Println("TITAN Core initializing...")
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Mode: %s\n", Mode)
	fmt.Println("Core status: ONLINE")
}

func Shutdown() {
	Runtime.Running = false
	fmt.Println("TITAN Core shutting down...")
}

func Status() string {
	if Runtime.Running {
		return "ONLINE"
	}

	return "OFFLINE"
}