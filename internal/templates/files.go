package templates

import (
	"os"
	"path/filepath"
)

type TemplateFiles struct {
	Path    string
	Content []byte
}

func getTemplateFiles(templateType string) ([]TemplateFiles, error) {
	// Get all files from shared directory
	shared, err := getFilesFromDir("shared")
	if err != nil {
		return nil, err
	}

	// Get template-specific files
	template, err := getFilesFromDir(templateType)
	if err != nil {
		return nil, err
	}

	return append(shared, template...), nil
}

func CreateProjectFiles(directory string, useDaisy bool) error {
	templateType := "vanilla"
	if useDaisy {
		templateType = "daisy"
	}

	files, err := getTemplateFiles(templateType)
	if err != nil {
		return err
	}

	for _, file := range files {
		// Create directory if it doesn't exist
		dir := filepath.Dir(filepath.Join(directory, file.Path))
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		// Write file
		if err := os.WriteFile(
			filepath.Join(directory, file.Path),
			file.Content,
			0644,
		); err != nil {
			return err
		}
	}

	return nil
}
