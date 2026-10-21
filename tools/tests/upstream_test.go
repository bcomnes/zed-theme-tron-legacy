package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"

	"github.com/bcomnes/zed-theme-tron-legacy/tools/dark"
	"github.com/bcomnes/zed-theme-tron-legacy/tools/light"
	"github.com/bcomnes/zed-theme-tron-legacy/tools/palette"
)

func readThemeDocument(t *testing.T, path string) map[string]any {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	var document map[string]any
	if err := json.Unmarshal(data, &document); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	return document
}

func generatedThemeDocument(t *testing.T) map[string]any {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot locate test file")
	}
	return readThemeDocument(t, filepath.Join(filepath.Dir(filename), "..", "..", "themes", "tron-legacy.json"))
}

func themeObject(t *testing.T, value any, label string) map[string]any {
	t.Helper()
	object, ok := value.(map[string]any)
	if !ok || len(object) == 0 {
		t.Fatalf("%s must be a nonempty object", label)
	}
	return object
}

func themeVariants(t *testing.T, document map[string]any) []any {
	t.Helper()
	variants, ok := document["themes"].([]any)
	if !ok || len(variants) == 0 {
		t.Fatal("themes must be a nonempty array")
	}
	return variants
}

func generatedVariants(t *testing.T) []any {
	t.Helper()
	variants := themeVariants(t, generatedThemeDocument(t))
	if len(variants) != 4 {
		t.Fatalf("expected four generated variants, got %d", len(variants))
	}
	expected := map[string]bool{
		"Tron Legacy": true, "Tron Legacy Frosted": true,
		"Tron Legacy Light": true, "Tron Legacy Light Frosted": true,
	}
	for _, value := range variants {
		variant := themeObject(t, value, "theme")
		name, _ := variant["name"].(string)
		if !expected[name] {
			t.Fatalf("unexpected or duplicate variant %q", name)
		}
		delete(expected, name)
	}
	return variants
}

func requireThemeColor(t *testing.T, object map[string]any, key string) {
	t.Helper()
	color, ok := object[key].(string)
	if !ok || strings.TrimSpace(color) == "" {
		t.Errorf("%s must contain a nonempty color string, got %v", key, object[key])
	}
}

func TestGeneratedThemeProperties(t *testing.T) {
	for _, value := range generatedVariants(t) {
		variant := themeObject(t, value, "theme")
		t.Run(variant["name"].(string), func(t *testing.T) {
			style := themeObject(t, variant["style"], "style")
			if _, exists := variant["accents"]; exists {
				t.Error("accents must be inside style, not on the outer variant")
			}
			accents, ok := style["accents"].([]any)
			if !ok || len(accents) == 0 {
				t.Error("style.accents must be a nonempty array")
			}
			for _, key := range []string{"foreground", "editor.selection.background"} {
				if _, exists := style[key]; exists {
					t.Errorf("unsupported style property %q must be absent", key)
				}
			}
			for _, key := range []string{"drop_target.border", "version_control.word_added", "version_control.word_deleted", "search.active_match_background"} {
				requireThemeColor(t, style, key)
			}
			syntax := themeObject(t, style["syntax"], "style.syntax")
			for _, selector := range []string{
				"variable.parameter", "punctuation.markup", "keyword.modifier", "constant.builtin",
				"lifetime", "concept", "module", "diff.delta", "markup.heading", "markup.link.url", "string.regexp", "text",
			} {
				t.Run(selector, func(t *testing.T) {
					highlight := themeObject(t, syntax[selector], "syntax."+selector)
					requireThemeColor(t, highlight, "color")
				})
			}
		})
	}
}

func TestGeneratedThemeMatchesGenerator(t *testing.T) {
	regenerated := palette.GenerateTheme("Tron Legacy", "Bret Comnes",
		palette.ThemeVariant{Name: "Tron Legacy", Appearance: "dark", Palette: dark.GetPalette()},
		palette.ThemeVariant{Name: "Tron Legacy Frosted", Appearance: "dark", Palette: dark.GetFrostedPalette()},
		palette.ThemeVariant{Name: "Tron Legacy Light", Appearance: "light", Palette: light.GetPalette()},
		palette.ThemeVariant{Name: "Tron Legacy Light Frosted", Appearance: "light", Palette: light.GetFrostedPalette()},
	)
	data, err := json.Marshal(regenerated)
	if err != nil {
		t.Fatalf("marshal regenerated theme: %v", err)
	}
	var expected map[string]any
	if err := json.Unmarshal(data, &expected); err != nil {
		t.Fatalf("parse regenerated theme: %v", err)
	}
	if !reflect.DeepEqual(generatedThemeDocument(t), expected) {
		t.Error("themes/tron-legacy.json differs from regenerated palettes; run go run generate-theme.go from tools")
	}
}

func TestUpstreamThemeCoverage(t *testing.T) {
	source, set := os.LookupEnv("ZED_SOURCE_DIR")
	if !set {
		t.Skip("set ZED_SOURCE_DIR to an absolute Zed checkout path to check official theme coverage")
	}
	if !filepath.IsAbs(source) {
		t.Fatalf("ZED_SOURCE_DIR must be an absolute path, got %q", source)
	}
	info, err := os.Stat(source)
	if err != nil {
		t.Fatalf("invalid ZED_SOURCE_DIR %q: %v", source, err)
	}
	if !info.IsDir() {
		t.Fatalf("ZED_SOURCE_DIR is not a directory: %s", source)
	}
	paths, err := filepath.Glob(filepath.Join(source, "assets", "themes", "*", "*.json"))
	if err != nil {
		t.Fatalf("find official themes: %v", err)
	}
	if len(paths) == 0 {
		t.Fatalf("no assets/themes/*/*.json found in ZED_SOURCE_DIR %q", source)
	}
	styleKeys := make(map[string]bool)
	syntaxKeys := make(map[string]bool)
	for _, path := range paths {
		for _, value := range themeVariants(t, readThemeDocument(t, path)) {
			variant := themeObject(t, value, path+" theme")
			style := themeObject(t, variant["style"], path+" style")
			for key := range style {
				// Background appearance is optional and depends on the variant's transparency.
				if key != "background.appearance" {
					styleKeys[key] = true
				}
			}
			if value, exists := style["syntax"]; exists {
				for key := range themeObject(t, value, path+" syntax") {
					syntaxKeys[key] = true
				}
			}
		}
	}
	for _, value := range generatedVariants(t) {
		variant := themeObject(t, value, "theme")
		t.Run(variant["name"].(string), func(t *testing.T) {
			style := themeObject(t, variant["style"], "style")
			syntax := themeObject(t, style["syntax"], "syntax")
			var missing []string
			for key := range styleKeys {
				if _, exists := style[key]; !exists {
					missing = append(missing, "style."+key)
				}
			}
			for key := range syntaxKeys {
				if _, exists := syntax[key]; !exists {
					missing = append(missing, "style.syntax."+key)
				}
			}
			sort.Strings(missing)
			if len(missing) > 0 {
				t.Errorf("missing official theme properties:\n%s", strings.Join(missing, "\n"))
			}
		})
	}
}
