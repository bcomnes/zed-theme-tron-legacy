package colorvalidation

import (
	"strings"
	"testing"

	"github.com/bcomnes/zed-theme-tron-legacy/tools/csscolors"
)

// ValidateAllColorsUsed checks that all colors defined in the CSS are used in the palette.go file.
// It reports any unused colors as test errors.
func ValidateAllColorsUsed(t *testing.T, colorsCSS []byte, paletteGoContent []byte) {
	t.Helper()

	// Load colors from CSS
	colors, err := csscolors.LoadColors(colorsCSS)
	if err != nil {
		t.Fatalf("Failed to load colors: %v", err)
	}

	// Convert palette.go file content to string for searching
	paletteContent := string(paletteGoContent)

	// Track unused colors
	unusedColors := []string{}

	// Check each color from CSS
	for colorName := range colors {
		// Check if the color name appears in palette.go
		// We check for MustGet("colorName") pattern
		searchPattern := `MustGet("` + colorName + `")`
		if !strings.Contains(paletteContent, searchPattern) {
			unusedColors = append(unusedColors, colorName)
		}
	}

	// Report unused colors
	if len(unusedColors) > 0 {
		t.Errorf("The following colors from colors.css are not used in palette.go:")
		for _, color := range unusedColors {
			t.Errorf("  - %s", color)
		}
	}
}

// ValidateNoDuplicateColorValues checks that each color variable has a unique color value.
// It reports any duplicate color values as test errors.
func ValidateNoDuplicateColorValues(t *testing.T, colorsCSS []byte) {
	t.Helper()

	// Load colors from CSS
	colors, err := csscolors.LoadColors(colorsCSS)
	if err != nil {
		t.Fatalf("Failed to load colors: %v", err)
	}

	// Create a map to track which color values are used by which variables
	colorValueToVars := make(map[string][]string)

	// Check each color
	for colorName, colorValue := range colors {
		if existing, found := colorValueToVars[colorValue]; found {
			// Add to existing list
			colorValueToVars[colorValue] = append(existing, colorName)
		} else {
			// First occurrence
			colorValueToVars[colorValue] = []string{colorName}
		}
	}

	// Report duplicates
	hasDuplicates := false
	for colorValue, varNames := range colorValueToVars {
		if len(varNames) > 1 {
			hasDuplicates = true
			t.Errorf("Color value %s is used by multiple variables: %v", colorValue, varNames)
		}
	}

	if hasDuplicates {
		t.Error("Found duplicate color values. Each color variable should have a unique value.")
	}
}
