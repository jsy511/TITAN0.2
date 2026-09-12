package dns

import (
	"fmt"
	"net"
)

type LookupResult struct {
	Hostnames []string
	Addresses []string
}

func Lookup(host string) LookupResult {
	result := LookupResult{}

	names, err := net.LookupAddr(host)
	if err == nil {
		result.Hostnames = names
	}

	addresses, err := net.LookupHost(host)
	if err == nil {
		result.Addresses = addresses
	}

	return result
}

func Display(host string) {
	fmt.Println()
	fmt.Println("  [ DNS DIAGNOSTICS ]")
	fmt.Println("  ------------------")
	fmt.Printf("  Host: %s\n\n", host)

	result := Lookup(host)

	if len(result.Addresses) == 0 {
		fmt.Println("  No addresses resolved.")
	} else {
		fmt.Println("  Addresses:")

		for _, address := range result.Addresses {
			fmt.Printf("    %s\n", address)
		}
	}

	if len(result.Hostnames) > 0 {
		fmt.Println("\n  Reverse names:")

		for _, hostname := range result.Hostnames {
			fmt.Printf("    %s\n", hostname)
		}
	}
}