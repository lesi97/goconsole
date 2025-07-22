package goconsole

import (
	"net"
	"os"
	"strings"
)

/*
Function to retrieve an IP address that can be used to access the Go server on the network
*/
func getLocalIp() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addresses {
		ipNet, isValid := address.(*net.IPNet)
		if !isValid {
			continue
		}
		ip := ipNet.IP
		if ip.IsLoopback() {
			continue
		}
		if ip.To4() == nil {
			continue
		}
		return ip.String()
	}
	return ""
}

/*
Function to read the version number and module name from the `go.mod` file
*/
func readModuleInfo() (name, version string) {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return "unknown-module", "v0.0.0"
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			name = strings.TrimSpace(strings.TrimPrefix(line, "module"))
		}
		if strings.HasPrefix(line, "// Version:") {
			version = strings.TrimSpace(strings.TrimPrefix(line, "// Version:"))
		}
	}

	if version == "" {
		version = "v0.0.0"
	}
	return name, version
}
