package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bcomnes/zed-theme-tron-legacy/tools/atelier"
	"github.com/bcomnes/zed-theme-tron-legacy/tools/dark"
	"github.com/bcomnes/zed-theme-tron-legacy/tools/light"
	"github.com/bcomnes/zed-theme-tron-legacy/tools/palette"
)

func main() {
	// Get palettes from the respective packages
	darkPalette := dark.GetPalette()
	lightPalette := light.GetPalette()
	atelierPalette := atelier.GetPalette()

	// Create theme variants
	variants := []palette.ThemeVariant{
		{
			Name:       "Tron Legacy",
			Appearance: "dark",
			Palette:    darkPalette,
		},
		{
			Name:       "Tron Legacy Light (WIP)",
			Appearance: "light",
			Palette:    lightPalette,
		},
		{
			Name:       "Tron Legacy Atelier (WIP)",
			Appearance: "light",
			Palette:    atelierPalette,
		},
		// Example: Adding more variants is easy!
		// {
		//     Name:       "Tron Legacy High Contrast",
		//     Appearance: "dark",
		//     Palette:    highContrastPalette,
		// },
		// {
		//     Name:       "Tron Legacy Soft",
		//     Appearance: "light",
		//     Palette:    softPalette,
		// },
	}

	// Generate the complete theme using the palette package
	theme := palette.GenerateTheme(
		"Tron Legacy",
		"Hypermodules LLC, Bret Comnes (ported to Zed)",
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
