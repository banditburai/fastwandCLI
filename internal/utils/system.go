package utils

import (
	"fmt"
	"os"
	"path/filepath"
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

func GetTailwindPath(directory string) (string, error) {
	// Get path in same format as download function
	tailwindPath := filepath.Join(directory, "tailwindcss")
	if runtime.GOOS == "windows" {
		tailwindPath += ".exe"
	}

	// Get absolute path first, just like in init
	absPath, err := filepath.Abs(tailwindPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Check if file exists and is executable, matching init's check
	if _, err := os.Stat(absPath); err != nil {
		return "", fmt.Errorf("tailwind binary not found in %s: %w", directory, err)
	}

	// Check if file is executable on Unix systems
	if runtime.GOOS != "windows" {
		info, err := os.Stat(absPath)
		if err != nil {
			return "", fmt.Errorf("failed to check file permissions: %w", err)
		}
		if info.Mode()&0111 == 0 {
			return "", fmt.Errorf("tailwind binary is not executable")
		}
	}

	return absPath, nil
}
