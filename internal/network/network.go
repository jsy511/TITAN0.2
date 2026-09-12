package network

import (
	"fmt"
	"net"
)

type InterfaceInfo struct {
	Name   string
	MTU    int
	Status string
}

func GetInterfaces() []InterfaceInfo {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	result := make([]InterfaceInfo, 0, len(interfaces))

	for _, iface := range interfaces {
		status := "DOWN"

		if iface.Flags&net.FlagUp != 0 {
			status = "UP"
		}

		result = append(result, InterfaceInfo{
			Name:   iface.Name,
			MTU:    iface.MTU,
			Status: status,
		})
	}

	return result
}

func Display() {
	fmt.Println()
	fmt.Println("  [ NETWORK INFORMATION ]")
	fmt.Println("  -----------------------")

	interfaces := GetInterfaces()

	if len(interfaces) == 0 {
		fmt.Println("  No network interfaces detected.")
		return
	}

	for _, iface := range interfaces {
		fmt.Printf(
			"  %-12s MTU: %-5d Status: %s\n",
			iface.Name,
			iface.MTU,
			iface.Status,
		)
	}
}