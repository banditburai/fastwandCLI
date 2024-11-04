package templates

import (
	"os"
	"path/filepath"
)

func CreateProjectFiles(directory string, useDaisy bool) error {
	// Create assets directory
	assetsDir := filepath.Join(directory, "assets")
	if err := os.MkdirAll(assetsDir, 0755); err != nil {
		return err
	}

	// Create input.css
	inputCSSPath := filepath.Join(assetsDir, "input.css")
	if err := os.WriteFile(inputCSSPath, []byte(InputCSS), 0644); err != nil {
		return err
	}

	// Create tailwind.config.js
	configContent := TailwindConfigVanilla
	if useDaisy {
		configContent = TailwindConfigDaisy
	}
	configPath := filepath.Join(directory, "tailwind.config.js")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		return err
	}

	// Create main.py with appropriate template
	mainContent := MainPYVanilla
	if useDaisy {
		mainContent = MainPYDaisy
	}
	mainPath := filepath.Join(directory, "main.py")
	if err := os.WriteFile(mainPath, []byte(mainContent), 0644); err != nil {
		return err
	}

	return nil
}
