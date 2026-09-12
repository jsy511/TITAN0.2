package tools

import "fmt"

type Tool struct {
	Name        string
	Description string
	Category    string
	Available   bool
}

var Registry = []Tool{
	{
		Name:        "System Inspector",
		Description: "Inspect local system information",
		Category:    "Diagnostics",
		Available:   true,
	},
	{
		Name:        "Network Inspector",
		Description: "Inspect local network interfaces",
		Category:    "Network",
		Available:   true,
	},
	{
		Name:        "DNS Resolver",
		Description: "Perform DNS diagnostics",
		Category:    "DNS",
		Available:   true,
	},
	{
		Name:        "Local Port Viewer",
		Description: "View locally available network socket information",
		Category:    "Network",
		Available:   true,
	},
	{
		Name:        "Hash Utility",
		Description: "Calculate file hashes for verification",
		Category:    "Integrity",
		Available:   true,
	},
}

func List() []Tool {
	return Registry
}

func Display() {
	fmt.Println()
	fmt.Println("  [ TITAN TOOLS ]")
	fmt.Println("  ----------------------------------------")

	for index, tool := range Registry {
		status := "READY"

		if !tool.Available {
			status = "UNAVAILABLE"
		}

		fmt.Printf(
			"  [%02d] %-22s %-14s %s\n",
			index+1,
			tool.Name,
			tool.Category,
			status,
		)

		fmt.Printf("       %s\n", tool.Description)
	}
}

func Get(name string) *Tool {
	for index := range Registry {
		if Registry[index].Name == name {
			return &Registry[index]
		}
	}

	return nil
}