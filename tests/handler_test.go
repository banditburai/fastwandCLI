package tests

import (
	"fastwand/internal/templates"
	"fastwand/internal/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreateProjectFiles(t *testing.T) {
	// Create temp directory for test
	tmpDir, err := os.MkdirTemp("", "fastwand-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Test vanilla Tailwind setup
	t.Run("Vanilla Tailwind", func(t *testing.T) {
		err := templates.CreateProjectFiles(tmpDir, false)
		if err != nil {
			t.Fatalf("failed to create project files: %v", err)
		}

		// Verify files exist
		files := []string{
			filepath.Join(tmpDir, "assets", "input.css"),
			filepath.Join(tmpDir, "tailwind.config.js"),
			filepath.Join(tmpDir, "main.py"),
		}

		for _, file := range files {
			if _, err := os.Stat(file); os.IsNotExist(err) {
				t.Errorf("expected file %s to exist", file)
			}
		}
	})

	// Test DaisyUI setup
	t.Run("DaisyUI Setup", func(t *testing.T) {
		err := templates.CreateProjectFiles(tmpDir, true)
		if err != nil {
			t.Fatalf("failed to create project files: %v", err)
		}

		// Verify DaisyUI config
		configPath := filepath.Join(tmpDir, "tailwind.config.js")
		content, err := os.ReadFile(configPath)
		if err != nil {
			t.Fatalf("failed to read config file: %v", err)
		}

		if !strings.Contains(string(content), "daisyui") {
			t.Error("expected DaisyUI configuration in tailwind.config.js")
		}
	})
}

func TestSystemInfo(t *testing.T) {
	info := utils.GetSystemInfo()

	if info.OS == "" {
		t.Error("expected non-empty OS")
	}
	if info.Arch == "" {
		t.Error("expected non-empty Arch")
	}
	if info.FormattedOS == "" {
		t.Error("expected non-empty FormattedOS")
	}

	// Test Darwin to macOS mapping
	if info.OS == "darwin" && info.FormattedOS != "macos" {
		t.Error("expected Darwin to be mapped to macOS")
	}
}

type MockRoundTripper struct {
	Responses map[string]*http.Response
	Err       error
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.Err != nil {
		return nil, m.Err
	}

	if resp, ok := m.Responses[req.URL.String()]; ok {
		return resp, nil
	}

	return &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader("not found")),
	}, nil
}
