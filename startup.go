package goconsole

import (
	"fmt"
	"os"
)

/*
A startup function that should be called on the main thread

that provides a developer with useful information
*/
func Startup() {
	env := os.Getenv("GO_ENV")
	const protocol = "http://"
	if env == "" {
		env = "development"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = fmt.Sprintf(":%v", port)
	name, _ := readModuleInfo()
	ip := getLocalIp()

	PrintColour("brightWhite", "\n  > %s\n", name) 					// Package name
	PrintColour("brightBlack", "\tEnvironments: .env (%s)", env)	// Current environment
	PrintColour("brightMagenta", "\n\t- Local:")					// Localhost address label
	PrintColour("cyan", "\t  %v%v%v", protocol, "localhost", port)	// Localhost address value
	if ip != "" {
		PrintColour("brightMagenta", "\n\t- Network:")				// Network address label
		PrintColour("cyan", "\t  %v%v%v\n", protocol,ip, port)		// Network address value
	}

	PrintColour("green", "\n  ✓ Server Ready\n\n")					// Server ready
}