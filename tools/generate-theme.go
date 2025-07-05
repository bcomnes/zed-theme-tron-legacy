package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bcomnes/zed-theme-tron-legacy/tools/dark"
	"github.com/bcomnes/zed-theme-tron-legacy/tools/light"
	"github.com/bcomnes/zed-theme-tron-legacy/tools/palette"
)

func main() {
	// Get palettes from the respective packages
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

	// Generate the complete theme using the palette package
	theme := palette.GenerateTheme(
		"Tron Legacy",
		"Bret Comnes",
		variants...,
	)

	// Marshal to JSON with indentation
	jsonData, err := json.MarshalIndent(theme, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling theme: %v\n", err)
		os.Exit(1)
	}

	// Write to output file
	outputPath := "../themes/tron-legacy.json"
	err = os.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Theme generated successfully!")
}
