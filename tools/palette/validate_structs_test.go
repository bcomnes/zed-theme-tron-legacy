package palette

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// TestValidateThemeStructsAgainstOfficial checks that our theme structs match the official Zed themes.
// This test performs a shallow clone of the Zed repository to fetch the latest official themes
// and validates that our ThemeStyle struct contains all necessary fields without extra ones.
func TestValidateThemeStructsAgainstOfficial(t *testing.T) {
	// Create a temporary directory for the shallow checkout
	tempDir, err := os.MkdirTemp("", "zed-themes-validation")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Initialize a new git repository
	cmd := exec.Command("git", "init")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}

	// Add the remote
	cmd = exec.Command("git", "remote", "add", "origin", "https://github.com/zed-industries/zed.git")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to add remote: %v", err)
	}

	// Enable sparse checkout
	cmd = exec.Command("git", "config", "core.sparseCheckout", "true")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to enable sparse checkout: %v", err)
	}

	// Configure sparse checkout to only get the themes directory
	sparseCheckoutPath := filepath.Join(tempDir, ".git", "info", "sparse-checkout")
	if err := os.WriteFile(sparseCheckoutPath, []byte("assets/themes/\n"), 0644); err != nil {
		t.Fatalf("Failed to write sparse-checkout: %v", err)
	}

	// Perform shallow fetch of just the latest commit
	cmd = exec.Command("git", "fetch", "--depth=1", "origin", "main")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to fetch: %v", err)
	}

	// Checkout the files
	cmd = exec.Command("git", "checkout", "main")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to checkout: %v", err)
	}

	// Define the themes to check
	themesToCheck := []struct {
		name string
		path string
	}{
		{"One", "assets/themes/one/one.json"},
		{"Gruvbox", "assets/themes/gruvbox/gruvbox.json"},
		{"Ayu", "assets/themes/ayu/ayu.json"},
	}

	// Known exceptions - fields that are in our structs but not in all official themes
	// These are documented features that are only present in specific themes or are optional
	knownExceptions := map[string]bool{
		// Theme-specific features
		"accents":              true, // Only in Gruvbox - array of accent colors for UI customization
		"function.builtin":     true, // Only in Gruvbox - special syntax highlighting for built-in functions

		// Optional appearance features
		"background.appearance":      true, // For themes with blur/transparency effects (frosted variants)
		"panel.overlay_background":   true, // Custom panel overlay backgrounds (not in all themes)
		"panel.overlay_hover":        true, // Custom panel overlay hover states (not in all themes)

		// Optional editor features
		"foreground":                 true, // Global foreground color (not always specified)
		"editor.selection.background": true, // Optional selection background override

		// Newly added fields (2024) - These are optional fields not present in most themes
		"element.selection_background":               true, // Selection background for UI elements
		"panel.indent_guide":                         true, // Panel indent guide lines
		"panel.indent_guide_hover":                   true, // Panel indent guide hover state
		"panel.indent_guide_active":                  true, // Panel indent guide active state
		"editor.indent_guide":                        true, // Editor indent guide lines
		"editor.indent_guide_active":                 true, // Editor indent guide active state
		"editor.debugger_active_line.background":     true, // Debugger active line highlighting
		"editor.document_highlight.bracket_background": true, // Bracket matching highlight
		"scrollbar.thumb.active_background":          true, // Scrollbar thumb when dragging
		"minimap.thumb.background":                   true, // Minimap thumb color
		"minimap.thumb.hover_background":             true, // Minimap thumb hover state
		"minimap.thumb.active_background":            true, // Minimap thumb active state
		"minimap.thumb.border":                       true, // Minimap thumb border
		"terminal.ansi.background":                   true, // Terminal ANSI background
		"version_control.renamed":                    true, // Version control renamed files
		"version_control.conflict":                   true, // Version control conflicts
		"version_control.ignored":                    true, // Version control ignored files
		"pane_group.border":                          true, // Border between pane groups
		"debugger.accent":                            true, // Debugger accent color

		// Version control markers (not in all themes)
		"version_control.conflict_marker.ours":   true, // Git conflict marker for "ours" section
		"version_control.conflict_marker.theirs": true, // Git conflict marker for "theirs" section
	}

	allOfficialFields := make(map[string]bool)
	officialFieldsByTheme := make(map[string]map[string]bool)

	// Load and analyze each official theme
	for _, theme := range themesToCheck {
		themePath := filepath.Join(tempDir, theme.path)
		data, err := os.ReadFile(themePath)
		if err != nil {
			t.Fatalf("Failed to read %s: %v", theme.name, err)
		}

		var themeData map[string]interface{}
		if err := json.Unmarshal(data, &themeData); err != nil {
			t.Fatalf("Failed to parse %s: %v", theme.name, err)
		}

		// Get the first theme variant's style
		themes, ok := themeData["themes"].([]interface{})
		if !ok || len(themes) == 0 {
			t.Fatalf("%s has no themes array", theme.name)
		}

		firstTheme, ok := themes[0].(map[string]interface{})
		if !ok {
			t.Fatalf("%s first theme is not an object", theme.name)
		}

		style, ok := firstTheme["style"].(map[string]interface{})
		if !ok {
			t.Fatalf("%s has no style object", theme.name)
		}

		// Collect all fields from this theme
		themeFields := make(map[string]bool)
		collectFields(style, "", themeFields)
		officialFieldsByTheme[theme.name] = themeFields

		// Add to the combined set
		for field := range themeFields {
			allOfficialFields[field] = true
		}
	}

	// Get fields from our ThemeStyle struct
	ourFields := getStructFields(reflect.TypeOf(ThemeStyle{}), "")

	// Check for missing fields (in official themes but not in our struct)
	missingFields := []string{}
	for field := range allOfficialFields {
		if !ourFields[field] && !knownExceptions[field] {
			missingFields = append(missingFields, field)
		}
	}

	// Check for extra fields (in our struct but not in any official theme)
	extraFields := []string{}
	for field := range ourFields {
		if !allOfficialFields[field] && !knownExceptions[field] {
			extraFields = append(extraFields, field)
		}
	}

	// Report results
	if len(missingFields) > 0 {
		t.Errorf("Missing fields in our struct that exist in official themes:\n")
		for _, field := range missingFields {
			// Show which themes have this field
			themesWithField := []string{}
			for themeName, fields := range officialFieldsByTheme {
				if fields[field] {
					themesWithField = append(themesWithField, themeName)
				}
			}
			t.Errorf("  - %s (found in: %s)", field, strings.Join(themesWithField, ", "))
		}
	}

	if len(extraFields) > 0 {
		t.Errorf("\nExtra fields in our struct that don't exist in any official theme:\n")
		for _, field := range extraFields {
			t.Errorf("  - %s", field)
		}
	}

	if len(missingFields) == 0 && len(extraFields) == 0 {
		t.Logf("✓ Theme struct validation passed! All fields match official themes.")
	}

	// Also validate the syntax fields
	t.Run("ValidateSyntaxFields", func(t *testing.T) {
		validateSyntaxFields(t, tempDir, themesToCheck)
	})
}

// collectFields recursively collects all field paths from a map
func collectFields(data interface{}, prefix string, fields map[string]bool) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			fieldPath := key
			if prefix != "" {
				fieldPath = prefix + "." + key
			}

			// Special handling for syntax fields - they have color/font_style/font_weight sub-properties
			if strings.HasPrefix(fieldPath, "syntax.") && !strings.Contains(fieldPath[7:], ".") {
				// This is a direct syntax field like "syntax.comment"
				// Don't add sub-properties as separate fields
				fields[fieldPath] = true
			} else if fieldPath == "syntax" {
				// Handle the syntax object itself
				if syntaxMap, ok := value.(map[string]interface{}); ok {
					for syntaxKey := range syntaxMap {
						fields["syntax."+syntaxKey] = true
					}
				}
			} else {
				// Add the field itself
				fields[fieldPath] = true

				// Recurse into nested objects (but not syntax style properties)
				if nestedMap, ok := value.(map[string]interface{}); ok {
					collectFields(nestedMap, fieldPath, fields)
				}
			}
		}
	}
}

// getStructFields gets all JSON field names from a struct type
func getStructFields(t reflect.Type, prefix string) map[string]bool {
	fields := make(map[string]bool)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		// Get the JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// Parse the JSON tag
		tagParts := strings.Split(jsonTag, ",")
		fieldName := tagParts[0]

		// Skip omitempty fields for this check if needed
		fullPath := fieldName
		if prefix != "" {
			fullPath = prefix + "." + fieldName
		}

		// Special handling for different types
		if fieldName == "syntax" && field.Type.Kind() == reflect.Struct {
			// For syntax, we want to get the field names from SyntaxStyles struct
			syntaxType := field.Type
			for i := 0; i < syntaxType.NumField(); i++ {
				syntaxField := syntaxType.Field(i)
				if syntaxField.PkgPath == "" {
					syntaxJsonTag := syntaxField.Tag.Get("json")
					if syntaxJsonTag != "" && syntaxJsonTag != "-" {
						syntaxTagParts := strings.Split(syntaxJsonTag, ",")
						fields["syntax."+syntaxTagParts[0]] = true
					}
				}
			}
		} else if field.Type.Kind() == reflect.Slice {
			// Skip slice fields like Players
			fields[fullPath] = true
		} else if field.Type.Kind() == reflect.Struct && fieldName != "syntax" {
			// For other structs, recurse
			fields[fullPath] = true
			nestedFields := getStructFields(field.Type, fullPath)
			for nestedField := range nestedFields {
				fields[nestedField] = true
			}
		} else {
			fields[fullPath] = true
		}
	}

	return fields
}

// validateSyntaxFields specifically validates syntax highlighting fields.
// This is a separate validation to ensure our SyntaxStyles struct matches the official themes.
func validateSyntaxFields(t *testing.T, tempDir string, themes []struct{ name, path string }) {
	allSyntaxFields := make(map[string]bool)
	syntaxFieldsByTheme := make(map[string]map[string]bool)

	// Known syntax exceptions - these are valid fields that don't appear in all themes
	syntaxExceptions := map[string]bool{
		"function.builtin": true, // Only in Gruvbox - highlights built-in functions differently
	}

	// Collect syntax fields from official themes
	for _, theme := range themes {
		themePath := filepath.Join(tempDir, theme.path)
		data, err := os.ReadFile(themePath)
		if err != nil {
			continue
		}

		var themeData map[string]interface{}
		if err := json.Unmarshal(data, &themeData); err != nil {
			continue
		}

		themes := themeData["themes"].([]interface{})
		firstTheme := themes[0].(map[string]interface{})
		style := firstTheme["style"].(map[string]interface{})

		if syntax, ok := style["syntax"].(map[string]interface{}); ok {
			themeFields := make(map[string]bool)
			for field := range syntax {
				themeFields[field] = true
				allSyntaxFields[field] = true
			}
			syntaxFieldsByTheme[theme.name] = themeFields
		}
	}

	// Get our syntax fields
	ourSyntaxType := reflect.TypeOf(SyntaxStyles{})
	ourSyntaxFields := make(map[string]bool)

	for i := 0; i < ourSyntaxType.NumField(); i++ {
		field := ourSyntaxType.Field(i)
		if field.PkgPath == "" {
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" && jsonTag != "-" {
				tagParts := strings.Split(jsonTag, ",")
				ourSyntaxFields[tagParts[0]] = true
			}
		}
	}

	// Check for missing syntax fields
	missingSyntax := []string{}
	for field := range allSyntaxFields {
		if !ourSyntaxFields[field] && !syntaxExceptions[field] {
			missingSyntax = append(missingSyntax, field)
		}
	}

	if len(missingSyntax) > 0 {
		t.Errorf("Missing syntax fields:\n")
		for _, field := range missingSyntax {
			themesWithField := []string{}
			for themeName, fields := range syntaxFieldsByTheme {
				if fields[field] {
					themesWithField = append(themesWithField, themeName)
				}
			}
			t.Errorf("  - %s (in: %s)", field, strings.Join(themesWithField, ", "))
		}
	}
}
