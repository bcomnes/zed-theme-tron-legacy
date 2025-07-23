package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/bcomnes/zed-theme-tron-legacy/tools/dark"
	"github.com/bcomnes/zed-theme-tron-legacy/tools/light"
	"github.com/bcomnes/zed-theme-tron-legacy/tools/palette"
)

func TestPalettesCanBeGenerated(t *testing.T) {
	// Test that dark palettes can be generated without panicking
	darkPalette := dark.GetPalette()
	if darkPalette.Background == "" {
		t.Error("Dark palette background should not be empty")
	}

	darkFrostedPalette := dark.GetFrostedPalette()
	if darkFrostedPalette.Background == "" {
		t.Error("Dark frosted palette background should not be empty")
	}

	// Test that light palettes can be generated without panicking
	lightPalette := light.GetPalette()
	if lightPalette.Background == "" {
		t.Error("Light palette background should not be empty")
	}

	lightFrostedPalette := light.GetFrostedPalette()
	if lightFrostedPalette.Background == "" {
		t.Error("Light frosted palette background should not be empty")
	}
}

func TestThemeGeneration(t *testing.T) {
	// Get palettes
	darkPalette := dark.GetPalette()
	lightPalette := light.GetPalette()
	darkFrostedPalette := dark.GetFrostedPalette()
	lightFrostedPalette := light.GetFrostedPalette()

	// Create theme variants
	variants := []palette.ThemeVariant{
		{
			Name:       "Tron Legacy",
			Appearance: "dark",
			Palette:    darkPalette,
		},
		{
			Name:       "Tron Legacy Frosted",
			Appearance: "dark",
			Palette:    darkFrostedPalette,
		},
		{
			Name:       "Tron Legacy Light",
			Appearance: "light",
			Palette:    lightPalette,
		},
		{
			Name:       "Tron Legacy Light Frosted",
			Appearance: "light",
			Palette:    lightFrostedPalette,
		},
	}

	// Generate the complete theme
	theme := palette.GenerateTheme(
		"Tron Legacy",
		"Bret Comnes",
		variants...,
	)

	// Test that theme has expected structure
	if theme.Schema == "" {
		t.Error("Theme should have $schema field")
	}

	if theme.Author == "" {
		t.Error("Theme should have author field")
	}

	if theme.Name == "" {
		t.Error("Theme should have name field")
	}

	// Test that themes array has expected length
	if len(theme.Themes) != 4 {
		t.Errorf("Expected 4 theme variants, got %d", len(theme.Themes))
	}
}

func TestJSONMarshalling(t *testing.T) {
	// Get palettes
	darkPalette := dark.GetPalette()
	lightPalette := light.GetPalette()
	darkFrostedPalette := dark.GetFrostedPalette()
	lightFrostedPalette := light.GetFrostedPalette()

	// Create theme variants
	variants := []palette.ThemeVariant{
		{
			Name:       "Tron Legacy",
			Appearance: "dark",
			Palette:    darkPalette,
		},
		{
			Name:       "Tron Legacy Frosted",
			Appearance: "dark",
			Palette:    darkFrostedPalette,
		},
		{
			Name:       "Tron Legacy Light",
			Appearance: "light",
			Palette:    lightPalette,
		},
		{
			Name:       "Tron Legacy Light Frosted",
			Appearance: "light",
			Palette:    lightFrostedPalette,
		},
	}

	// Generate the complete theme
	theme := palette.GenerateTheme(
		"Tron Legacy",
		"Bret Comnes",
		variants...,
	)

	// Test that theme can be marshalled to JSON
	jsonData, err := json.MarshalIndent(theme, "", "  ")
	if err != nil {
		t.Errorf("Error marshaling theme to JSON: %v", err)
	}

	// Test that JSON is non-empty
	if len(jsonData) == 0 {
		t.Error("Generated JSON should not be empty")
	}

	// Test that JSON can be unmarshalled back
	var unmarshalled palette.Theme
	err = json.Unmarshal(jsonData, &unmarshalled)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %v", err)
	}

	// Verify unmarshalled data
	if unmarshalled.Name != "Tron Legacy" {
		t.Errorf("Expected theme name 'Tron Legacy', got '%s'", unmarshalled.Name)
	}
}

func TestOutputFileCanBeWritten(t *testing.T) {
	// Get palettes
	darkPalette := dark.GetPalette()
	lightPalette := light.GetPalette()
	darkFrostedPalette := dark.GetFrostedPalette()
	lightFrostedPalette := light.GetFrostedPalette()

	// Create theme variants
	variants := []palette.ThemeVariant{
		{
			Name:       "Tron Legacy",
			Appearance: "dark",
			Palette:    darkPalette,
		},
		{
			Name:       "Tron Legacy Frosted",
			Appearance: "dark",
			Palette:    darkFrostedPalette,
		},
		{
			Name:       "Tron Legacy Light",
			Appearance: "light",
			Palette:    lightPalette,
		},
		{
			Name:       "Tron Legacy Light Frosted",
			Appearance: "light",
			Palette:    lightFrostedPalette,
		},
	}

	// Generate the complete theme
	theme := palette.GenerateTheme(
		"Tron Legacy",
		"Bret Comnes",
		variants...,
	)

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(theme, "", "  ")
	if err != nil {
		t.Errorf("Error marshaling theme: %v", err)
	}

	// Write to temporary file to test file I/O
	tempDir := t.TempDir()
	testOutputPath := filepath.Join(tempDir, "test-theme.json")

	err = os.WriteFile(testOutputPath, jsonData, 0644)
	if err != nil {
		t.Errorf("Error writing test file: %v", err)
	}

	// Test that file exists and is readable
	_, err = os.Stat(testOutputPath)
	if err != nil {
		t.Errorf("Error checking test file: %v", err)
	}

	// Test that file can be read back
	readData, err := os.ReadFile(testOutputPath)
	if err != nil {
		t.Errorf("Error reading test file: %v", err)
	}

	if len(readData) == 0 {
		t.Error("Read file should not be empty")
	}
}
