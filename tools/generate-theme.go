package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// TronThemePalette contains semantic color assignments for a theme variant
type TronThemePalette struct {
	// Core colors
	Background       string
	Foreground       string
	ForegroundDim    string
	ForegroundBright string

	// UI elements
	Border        string
	BorderSubtle  string
	BorderFocused string
	Selection     string
	Highlight     string
	ElementActive string
	Shadow        string

	// Semantic colors
	Error   string
	Warning string
	Success string
	Info    string

	// Background variants for semantic colors
	ErrorBg   string
	SuccessBg string

	// Syntax highlighting - base colors
	Comment  string
	String   string
	Number   string
	Keyword  string
	Function string
	Type     string
	Variable string
	Property string

	// Syntax highlighting - accent colors
	StringEscape string
	Regex        string
	Decorator    string
	Punctuation  string
	PunctuationSpecial string
	Link        string

	// Terminal base colors (reuse main palette where possible)
	TerminalBlack  string
	TerminalRed    string
	TerminalGreen  string
	TerminalYellow string
	TerminalBlue   string
	TerminalPurple string
	TerminalCyan   string
	TerminalWhite  string

	// Additional semantic mappings
	LineNumber           string
	ElementActiveBright  string
	ConflictBackground   string
	TerminalBrightWhite  string
	TerminalDimGreen     string
	TerminalDimYellow    string
	VersionControlModified string
	VariableParam        string
	Constructor          string
	ClassType            string
	EnumType             string
	AttributeType        string
	SelectorAlt          string
	AnnotationType       string
	EmbeddedType         string
	// Player selection colors
	Player1Selection     string
	Player3Selection     string
	Player4Selection     string
}

// GetTronColors returns a map of all unique colors used in the theme
func GetTronColors() map[string]string {
	return map[string]string{
		// Monochrome
		"black":  "#000000ff",
		"white":  "#ffffffff",
		"shadow": "#00000040",

		// Gray scale (dark to light)
		"gray900": "#14191fff", // Background
		"gray800": "#28323eff", // LineHighlight
		"gray700": "#133153ff", // BorderSubtle
		"gray600": "#183c66ff", // Border, Selection
		"gray500": "#647c9bff", // ForegroundDim
		"gray400": "#586676ff", // Comment
		"gray300": "#748aa6ff", // Keyword, Operator
		"gray200": "#aec2e0ff", // Foreground
		"gray100": "#d0dfe6ff", // Variable
		"gray50":  "#dae3f1ff", // ForegroundBright

		// Blues
		"blue700": "#1a5278ff", // Link color
		"blue500": "#267fb5ff", // Primary blue
		"blue400": "#2b6db9ff", // Bright blue
		"blue300": "#4a95b3ff", // Cyan dim
		"blue200": "#6ee2ffff", // Primary cyan

		// Greens
		"green700": "#144212ff", // Success bg
		"green600": "#4d5f07ff", // Green dim
		"green500": "#95CC5Eff", // Alt green
		"green400": "#A6E22Eff", // Bright green alt
		"green300": "#C7F026ff", // Primary green

		// Yellows/Oranges
		"yellow500": "#FFE792ff", // Primary yellow
		"yellow400": "#ffd12cff", // Dim yellow
		"orange500": "#FFB20Dff", // Primary orange
		"orange400": "#F79D1Eff", // Light orange
		"orange600": "#cc8f0aff", // Dim orange

		// Reds
		"red700": "#660000ff", // Error bg
		"red600": "#cc3309ff", // Dark red
		"red500": "#F92672ff", // Error red
		"red400": "#FF410Dff", // Primary red
		"red300": "#ff6113ff", // Light red

		// Purples
		"purple500": "#967EFBff", // Primary purple
		"purple600": "#7865c9ff", // Dim purple
		"pink500":   "#ff79c6ff", // Primary pink

		// Special alpha variants
		"blue500alpha": "#267fb580", // Selection with alpha

		// Pure white (different from gray scale white)
		"purewhite": "#ffffffff",

		// Player selection colors
		"greendim": "#4d5f07ff",
		"orangelight": "#F79D1Eff",
	}
}

// GetTronLegacyPalette returns the dark theme palette
func GetTronLegacyPalette() TronThemePalette {
	colors := GetTronColors()

	return TronThemePalette{
		// Core colors
		Background:       colors["gray900"],
		Foreground:       colors["gray200"],
		ForegroundDim:    colors["gray500"],
		ForegroundBright: colors["gray50"],

		// UI elements
		Border:        colors["gray600"],
		BorderSubtle:  colors["gray700"],
		BorderFocused: colors["blue200"],
		Selection:     colors["gray600"],
		Highlight:     colors["gray800"],
		ElementActive: colors["blue400"],
		Shadow:        colors["shadow"],

		// Semantic colors
		Error:   colors["red500"],
		Warning: colors["yellow500"],
		Success: colors["green300"],
		Info:    colors["blue200"],

		// Background variants
		ErrorBg:   colors["red700"],
		SuccessBg: colors["green700"],

		// Syntax highlighting - base colors
		Comment:  colors["gray400"],
		String:   colors["red400"],
		Number:   colors["green300"],
		Keyword:  colors["gray300"],
		Function: colors["orange500"],
		Type:     colors["blue500"],
		Variable: colors["gray100"],
		Property: colors["green500"],

		// Syntax highlighting - accent colors
		StringEscape:       colors["red300"],
		Regex:              colors["blue200"],
		Decorator:          colors["pink500"],
		Punctuation:        colors["gray200"],
		PunctuationSpecial: colors["gray500"],
		Link:               colors["blue700"],

		// Terminal colors
		TerminalBlack:  colors["black"],
		TerminalRed:    colors["red500"],
		TerminalGreen:  colors["green500"],
		TerminalYellow: colors["yellow500"],
		TerminalBlue:   colors["blue500"],
		TerminalPurple: colors["purple500"],
		TerminalCyan:   colors["blue200"],
		TerminalWhite:  colors["gray200"],

		// Additional semantic mappings
		LineNumber:           colors["green600"],
		ElementActiveBright:  colors["blue400"],
		ConflictBackground:   colors["orange400"],
		TerminalBrightWhite:  colors["purewhite"],
		TerminalDimGreen:     colors["green600"],
		TerminalDimYellow:    colors["yellow400"],
		VersionControlModified: colors["yellow400"],
		VariableParam:        colors["green500"],
		Constructor:          colors["orange400"],
		ClassType:            colors["orange400"],
		EnumType:             colors["orange400"],
		AttributeType:        colors["orange400"],
		SelectorAlt:          colors["green500"],
		AnnotationType:       colors["yellow400"],
		EmbeddedType:         colors["yellow400"],
		// Player selection colors
		Player1Selection:     colors["blue500alpha"],
		Player3Selection:     colors["greendim"],
		Player4Selection:     colors["orangelight"],
	}
}

// generateTheme generates the complete theme JSON structure
func generateTheme(palette TronThemePalette) map[string]interface{} {
	return map[string]interface{}{
		"$schema": "https://zed.dev/schema/themes/v0.2.0.json",
		"author":  "Hypermodules LLC, Bret Comnes (ported to Zed)",
		"name":    "Tron Legacy",
		"themes": []interface{}{
			generateThemeStyle("Tron Legacy", "dark", palette),
		},
	}
}

// generateThemeStyle generates the style portion of the theme
// This function maps semantic palette values to theme properties
func generateThemeStyle(name string, appearance string, p TronThemePalette) map[string]interface{} {
	return map[string]interface{}{
		"name":       name,
		"appearance": appearance,
		"style": map[string]interface{}{
			// Core colors
			"background":       p.Background,
			"foreground":       p.Foreground,
			"text":             p.Foreground,
			"text.muted":       p.ForegroundDim,
			"text.placeholder": p.ForegroundDim,
			"text.disabled":    p.ForegroundDim,
			"text.accent":      p.BorderFocused,

			// Icon colors
			"icon":             p.Foreground,
			"icon.muted":       p.ForegroundDim,
			"icon.disabled":    p.ForegroundDim,
			"icon.placeholder": p.ForegroundDim,
			"icon.accent":      p.BorderFocused,

			// Borders
			"border":             p.Border,
			"border.variant":     p.BorderSubtle,
			"border.focused":     p.BorderFocused,
			"border.selected":    p.BorderFocused,
			"border.transparent": p.Shadow,
			"border.disabled":    p.ForegroundDim,

			// UI surfaces
			"surface.background":          p.Background,
			"elevated_surface.background": p.Highlight,
			"element.background":          p.Highlight,
			"element.hover":               p.Border,
			"element.active":              p.ElementActive,
			"element.selected":            p.Border,
			"element.disabled":            p.BorderSubtle,
			"drop_target.background":      p.Shadow,

			// Ghost elements
			"ghost_element.background": p.Shadow,
			"ghost_element.hover":      p.Highlight,
			"ghost_element.active":     p.Border,
			"ghost_element.selected":   p.Border,
			"ghost_element.disabled":   p.BorderSubtle,

			// Editor specific
			"editor.background":                          p.Background,
			"editor.foreground":                          p.Foreground,
			"editor.gutter.background":                   p.Background,
			"editor.active_line.background":              p.Highlight,
			"editor.highlighted_line.background":         p.Highlight,
			"editor.line_number":                         p.LineNumber,
			"editor.active_line_number":                  p.Success,
			"editor.hover_line_number":                   p.Success,
			"editor.invisible":                           p.ForegroundDim,
			"editor.wrap_guide":                          p.LineNumber,
			"editor.active_wrap_guide":                   p.Success,
			"editor.selection.background":                p.Selection,
			"editor.document_highlight.read_background":  p.Selection,
			"editor.document_highlight.write_background": p.ElementActiveBright,
			"editor.subheader.background":                p.Highlight,

			// Search
			"search.match_background": p.Warning,

			// Panels and tabs
			"panel.background":              p.Background,
			"panel.focused_border":          p.BorderFocused,
			"pane.focused_border":           p.BorderFocused,
			"status_bar.background":         p.Background,
			"title_bar.background":          p.Background,
			"title_bar.inactive_background": p.Background,
			"toolbar.background":            p.Background,
			"tab_bar.background":            p.Background,
			"tab.inactive_background":       p.Background,
			"tab.active_background":         p.Highlight,

			// Scrollbar
			"scrollbar.thumb.background":       p.Border,
			"scrollbar.thumb.hover_background": p.ElementActiveBright,
			"scrollbar.thumb.border":           p.Border,
			"scrollbar.track.background":       p.Shadow,
			"scrollbar.track.border":           p.Shadow,

			// Terminal
			"terminal.background":          p.Background,
			"terminal.foreground":          p.Foreground,
			"terminal.bright_foreground":   p.ForegroundBright,
			"terminal.dim_foreground":      p.ForegroundDim,
			"terminal.ansi.black":          p.TerminalBlack,
			"terminal.ansi.red":            p.TerminalRed,
			"terminal.ansi.green":          p.TerminalGreen,
			"terminal.ansi.yellow":         p.TerminalYellow,
			"terminal.ansi.blue":           p.TerminalBlue,
			"terminal.ansi.magenta":        p.TerminalPurple,
			"terminal.ansi.cyan":           p.TerminalCyan,
			"terminal.ansi.white":          p.TerminalWhite,
			"terminal.ansi.bright_black":   p.ForegroundDim,
			"terminal.ansi.bright_red":     p.String, // Reuse string color
			"terminal.ansi.bright_green":   p.Success,
			"terminal.ansi.bright_yellow":  p.Function, // Reuse function/orange
			"terminal.ansi.bright_blue":    p.Info,
			"terminal.ansi.bright_magenta": p.TerminalPurple,
			"terminal.ansi.bright_cyan":    p.Info,
			"terminal.ansi.bright_white":   p.TerminalBrightWhite,
			"terminal.ansi.dim_black":      p.Shadow,
			"terminal.ansi.dim_red":        p.ErrorBg,
			"terminal.ansi.dim_green":      p.TerminalDimGreen,
			"terminal.ansi.dim_yellow":     p.TerminalDimYellow,
			"terminal.ansi.dim_blue":       p.Border,
			"terminal.ansi.dim_magenta":    p.TerminalPurple,
			"terminal.ansi.dim_cyan":       p.Property, // Cyan dim
			"terminal.ansi.dim_white":      p.Keyword,

			// Links
			"link_text.hover": p.ForegroundDim,

			// Version control
			"version_control.added":                  p.Success,
			"version_control.modified":               p.VersionControlModified,
			"version_control.deleted":                p.Error,
			"version_control.conflict_marker.ours":   p.Success,
			"version_control.conflict_marker.theirs": p.Error,

			// Diagnostics
			"error":              p.Error,
			"error.background":   p.ErrorBg,
			"error.border":       p.Error,
			"warning":            p.Warning,
			"warning.background": p.Background,
			"warning.border":     p.Warning,
			"info":               p.Info,
			"info.background":    p.Background,
			"info.border":        p.Info,
			"hint":               p.Type,
			"hint.background":    p.Background,
			"hint.border":        p.Type,
			"success":            p.Success,
			"success.background": p.SuccessBg,
			"success.border":     p.Success,

			// Git status
			"created":              p.Success,
			"created.background":   p.SuccessBg,
			"created.border":       p.Success,
			"deleted":              p.Error,
			"deleted.background":   p.ErrorBg,
			"deleted.border":       p.Error,
			"modified":             p.Warning,
			"modified.background":  p.Background,
			"modified.border":      p.Warning,
			"renamed":              p.Type,
			"renamed.background":   p.Background,
			"renamed.border":       p.Type,
			"conflict":             p.Function,
			"conflict.background":  p.ConflictBackground,
			"conflict.border":      p.Function,

			// Other states
			"ignored":                p.ForegroundDim,
			"ignored.background":     p.Background,
			"ignored.border":         p.ForegroundDim,
			"hidden":                 p.ForegroundDim,
			"hidden.background":      p.Background,
			"hidden.border":          p.ForegroundDim,
			"unreachable":            p.ForegroundDim,
			"unreachable.background": p.Background,
			"unreachable.border":     p.ForegroundDim,
			"predictive":             p.TerminalPurple,
			"predictive.background":  p.Background,
			"predictive.border":      p.TerminalPurple,

			// Players
			"players": generatePlayers(p),

			// Syntax highlighting
			"syntax": generateSyntax(p),
		},
	}
}

// generatePlayers generates the multiplayer cursor colors
func generatePlayers(p TronThemePalette) []map[string]interface{} {
	return []map[string]interface{}{
		{
			"cursor":     p.Type,
			"background": p.Type,
			"selection":  p.Player1Selection,
		},
		{
			"cursor":     p.Info,
			"background": p.Info,
			"selection":  p.Border,
		},
		{
			"cursor":     p.Success,
			"background": p.Success,
			"selection":  p.Player3Selection,
		},
		{
			"cursor":     p.Function,
			"background": p.Function,
			"selection":  p.Player4Selection,
		},
		{
			"cursor":     p.TerminalPurple,
			"background": p.TerminalPurple,
			"selection":  p.TerminalPurple,
		},
		{
			"cursor":     p.Error,
			"background": p.Error,
			"selection":  p.ErrorBg,
		},
		{
			"cursor":     p.Type,
			"background": p.Type,
			"selection":  p.Border,
		},
		{
			"cursor":     p.Keyword,
			"background": p.Keyword,
			"selection":  p.ForegroundDim,
		},
	}
}

// generateSyntax generates the syntax highlighting rules
func generateSyntax(p TronThemePalette) map[string]interface{} {
	syntax := map[string]interface{}{
		// Comments
		"comment": map[string]interface{}{
			"color":       p.Comment,
			"font_style":  nil,
			"font_weight": nil,
		},
		"comment.doc": map[string]interface{}{
			"color":       p.Comment,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Literals
		"string": map[string]interface{}{
			"color":       p.String,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.escape": map[string]interface{}{
			"color":       p.StringEscape,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.regex": map[string]interface{}{
			"color":       p.Regex,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.special": map[string]interface{}{
			"color":       p.Decorator,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.special.symbol": map[string]interface{}{
			"color":       p.Function,
			"font_style":  nil,
			"font_weight": nil,
		},
		"number": map[string]interface{}{
			"color":       p.Number,
			"font_style":  nil,
			"font_weight": nil,
		},
		"boolean": map[string]interface{}{
			"color":       p.Function,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"constant": map[string]interface{}{
			"color":       p.Function,
			"font_style":  "italic",
			"font_weight": nil,
		},

		// Identifiers
		"variable": map[string]interface{}{
			"color":       p.Variable,
			"font_style":  nil,
			"font_weight": nil,
		},
		"variable.builtin": map[string]interface{}{
			"color":       p.PunctuationSpecial,
			"font_style":  nil,
			"font_weight": nil,
		},
		"variable.parameter": map[string]interface{}{
			"color":       p.VariableParam,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"variable.special": map[string]interface{}{
			"color":       p.TerminalPurple,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"field": map[string]interface{}{
			"color":       p.Property,
			"font_style":  nil,
			"font_weight": nil,
		},
		"property": map[string]interface{}{
			"color":       p.Property,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Functions
		"function": map[string]interface{}{
			"color":       p.Function,
			"font_style":  nil,
			"font_weight": nil,
		},
		"function.builtin": map[string]interface{}{
			"color":       p.Type,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"method": map[string]interface{}{
			"color":       p.Function,
			"font_style":  nil,
			"font_weight": nil,
		},
		"constructor": map[string]interface{}{
			"color":       p.Constructor,
			"font_style":  "italic",
			"font_weight": 700,
		},

		// Keywords
		"keyword": map[string]interface{}{
			"color":       p.Keyword,
			"font_style":  nil,
			"font_weight": nil,
		},
		"operator": map[string]interface{}{
			"color":       p.Keyword,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Types
		"type": map[string]interface{}{
			"color":       p.Type,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"class": map[string]interface{}{
			"color":       p.ClassType,
			"font_style":  nil,
			"font_weight": 700,
		},
		"interface": map[string]interface{}{
			"color":       p.Property,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"enum": map[string]interface{}{
			"color":       p.EnumType,
			"font_style":  nil,
			"font_weight": nil,
		},
		"variant": map[string]interface{}{
			"color":       p.EnumType,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Markup
		"tag": map[string]interface{}{
			"color":       p.Type,
			"font_style":  nil,
			"font_weight": nil,
		},
		"attribute": map[string]interface{}{
			"color":       p.AttributeType,
			"font_style":  nil,
			"font_weight": nil,
		},
		"selector": map[string]interface{}{
			"color":       p.SelectorAlt,
			"font_style":  nil,
			"font_weight": nil,
		},
		"selector.pseudo": map[string]interface{}{
			"color":       p.Decorator,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Punctuation
		"punctuation": map[string]interface{}{
			"color":       p.Punctuation,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.bracket": map[string]interface{}{
			"color":       p.Punctuation,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.delimiter": map[string]interface{}{
			"color":       p.PunctuationSpecial,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.special": map[string]interface{}{
			"color":       p.Info,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.list_marker": map[string]interface{}{
			"color":       p.Success,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Misc
		"label": map[string]interface{}{
			"color":       p.Decorator,
			"font_style":  nil,
			"font_weight": nil,
		},
		"namespace": map[string]interface{}{
			"color":       p.Link,
			"font_style":  nil,
			"font_weight": nil,
		},
		"module": map[string]interface{}{
			"color":       p.Link,
			"font_style":  nil,
			"font_weight": nil,
		},
		"decorator": map[string]interface{}{
			"color":       p.Decorator,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"macro": map[string]interface{}{
			"color":       p.Decorator,
			"font_style":  nil,
			"font_weight": 700,
		},
		"parameter": map[string]interface{}{
			"color":       p.VariableParam,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"annotation": map[string]interface{}{
			"color":       p.AnnotationType,
			"font_style":  "italic",
			"font_weight": nil,
		},

		// Special
		"emphasis": map[string]interface{}{
			"color":       nil,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"emphasis.strong": map[string]interface{}{
			"color":       nil,
			"font_style":  nil,
			"font_weight": 700,
		},
		"title": map[string]interface{}{
			"color":       p.Variable,
			"font_style":  nil,
			"font_weight": 700,
		},
		"link_uri": map[string]interface{}{
			"color":       p.Link,
			"font_style":  "underline",
			"font_weight": nil,
		},
		"link_text": map[string]interface{}{
			"color":       p.Link,
			"font_style":  nil,
			"font_weight": nil,
		},
		"embedded": map[string]interface{}{
			"color":       p.EmbeddedType,
			"font_style":  nil,
			"font_weight": nil,
		},
		"primary": map[string]interface{}{
			"color":       p.Foreground,
			"font_style":  nil,
			"font_weight": 700,
		},
		"hint": map[string]interface{}{
			"color":       p.Type,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"predictive": map[string]interface{}{
			"color":       p.TerminalPurple,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"preproc": map[string]interface{}{
			"color":       p.Info,
			"font_style":  nil,
			"font_weight": nil,
		},
		"regex": map[string]interface{}{
			"color":       p.Regex,
			"font_style":  nil,
			"font_weight": nil,
		},
		"text.literal": map[string]interface{}{
			"color":       p.EmbeddedType,
			"font_style":  nil,
			"font_weight": nil,
		},
	}

	return syntax
}

func main() {
	palette := GetTronLegacyPalette()
	theme := generateTheme(palette)

	jsonData, err := json.MarshalIndent(theme, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling theme: %v\n", err)
		os.Exit(1)
	}

	outputPath := "../themes/tron-legacy.json"
	err = os.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Theme generated successfully!")
}
