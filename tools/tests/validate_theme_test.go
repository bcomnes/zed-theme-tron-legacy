package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

func TestValidateThemeJSON(t *testing.T) {
	// Get the directory of the current test file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Failed to get current file path")
	}

	// Build paths relative to the test file location
	testDir := filepath.Dir(filename)
	projectRoot := filepath.Join(testDir, "..", "..")

	// Path to the generated theme file
	themePath := filepath.Join(projectRoot, "themes", "tron-legacy.json")

	// Read the theme file to extract the $schema URL
	themeData, err := os.ReadFile(themePath)
	if err != nil {
		t.Fatalf("Failed to read theme file: %v", err)
	}

	// Parse the theme to get the $schema field
	var theme map[string]any
	if err := json.Unmarshal(themeData, &theme); err != nil {
		t.Fatalf("Failed to parse theme JSON: %v", err)
	}

	schemaURL, ok := theme["$schema"].(string)
	if !ok || schemaURL == "" {
		t.Fatal("Theme file does not contain a valid $schema field")
	}

	// Create loaders for both schema and theme
	themeLoader := gojsonschema.NewReferenceLoader("file://" + themePath)
	schemaLoader := gojsonschema.NewReferenceLoader(schemaURL)

	// Validate the theme against the schema
	result, err := gojsonschema.Validate(schemaLoader, themeLoader)
	if err != nil {
		t.Fatalf("Error validating theme: %v", err)
	}

	// Check if the document is valid
	if !result.Valid() {
		t.Errorf("The theme file is not valid according to the schema")
		for _, desc := range result.Errors() {
			t.Errorf("- %s", desc)
		}
	}
}

func TestThemeFileExists(t *testing.T) {
	// Get the directory of the current test file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Failed to get current file path")
	}

	// Build path to theme file
	testDir := filepath.Dir(filename)
	projectRoot := filepath.Join(testDir, "..", "..")
	themePath := filepath.Join(projectRoot, "themes", "tron-legacy.json")

	// Check if file exists by trying to load it
	loader := gojsonschema.NewReferenceLoader("file://" + themePath)
	_, err := loader.LoadJSON()
	if err != nil {
		t.Errorf("Failed to load theme file: %v", err)
	}
}
