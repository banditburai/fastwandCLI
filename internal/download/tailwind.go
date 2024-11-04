package download

import (
	"encoding/json"
	"fastwand/internal/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type GithubRelease struct {
	TagName string `json:"tag_name"`
}

func getLatestVersion(repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "fastwand")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var release GithubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}

	return release.TagName, nil
}

func DownloadTailwind(directory string, useDaisy bool) (string, error) {
	sysInfo := utils.GetSystemInfo()

	var version string
	var err error
	var filename string
	var urlBase string

	if useDaisy {
		version, err = getLatestVersion("dobicinaitis/tailwind-cli-extra")
		if err != nil {
			return "", fmt.Errorf("failed to get DaisyUI version: %w", err)
		}
		filename = fmt.Sprintf("tailwindcss-extra-%s-%s", sysInfo.FormattedOS, sysInfo.Arch)
		urlBase = "https://github.com/dobicinaitis/tailwind-cli-extra/releases/download"
	} else {
		version, err = getLatestVersion("tailwindlabs/tailwindcss")
		if err != nil {
			return "", fmt.Errorf("failed to get Tailwind version: %w", err)
		}
		filename = fmt.Sprintf("tailwindcss-%s-%s", sysInfo.FormattedOS, sysInfo.Arch)
		urlBase = "https://github.com/tailwindlabs/tailwindcss/releases/download"
	}

	// Add .exe extension for Windows
	if sysInfo.OS == "windows" {
		filename += ".exe"
	}

	// Ensure directory exists
	if err := os.MkdirAll(directory, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	// Construct download URL
	url := fmt.Sprintf("%s/%s/%s", urlBase, version, filename)

	// Create custom client with timeouts
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set User-Agent
	req.Header.Set("User-Agent", "fastwand")

	// Make request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to download Tailwind: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download Tailwind: HTTP %d", resp.StatusCode)
	}

	// Create temporary file
	tempFile, err := os.CreateTemp(directory, "tailwind-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up temp file on error

	// Copy downloaded content to temp file
	if _, err := io.Copy(tempFile, resp.Body); err != nil {
		tempFile.Close()
		return "", fmt.Errorf("failed to write file: %w", err)
	}
	tempFile.Close()

	// Move to final location
	finalPath := filepath.Join(directory, "tailwindcss")
	if sysInfo.OS == "windows" {
		finalPath += ".exe"
	}

	// Remove existing file if it exists
	if _, err := os.Stat(finalPath); err == nil {
		if err := os.Remove(finalPath); err != nil {
			return "", fmt.Errorf("failed to remove existing tailwind: %w", err)
		}
	}

	// Move temp file to final location
	if err := os.Rename(tempFile.Name(), finalPath); err != nil {
		return "", fmt.Errorf("failed to move file to final location: %w", err)
	}

	// Make executable after moving to final location
	if err := os.Chmod(finalPath, 0755); err != nil {
		return "", fmt.Errorf("failed to make executable: %w", err)
	}

	return finalPath, nil
}
