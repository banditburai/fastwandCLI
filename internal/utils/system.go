package utils

import (
	"runtime"
	"strings"
)

type SystemInfo struct {
	OS          string
	Arch        string
	FormattedOS string
}

func GetSystemInfo() SystemInfo {
	system := strings.ToLower(runtime.GOOS)
	machine := strings.ToLower(runtime.GOARCH)

	// Map Darwin to macOS for user-facing messages
	formattedOS := system
	if system == "darwin" {
		formattedOS = "macos"
	}

	// Map architectures to consistent naming
	if machine == "amd64" {
		machine = "x64"
	}

	return SystemInfo{
		OS:          system,
		Arch:        machine,
		FormattedOS: formattedOS,
	}
}
