package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// TronLegacyPalette defines all colors used in the theme
type TronLegacyPalette struct {
	// Core colors
	Background       string
	Foreground       string
	ForegroundDim    string
	ForegroundBright string
	ForegroundLight  string  // Light blue-gray for variables

	// Accent colors
	Cyan        string
	Blue        string
	BlueBright  string
	BlueBrighter string  // Bright blue for brackets
	Green       string
	GreenDim    string
	GreenAlt    string
	GreenBright string
	GreenParam  string  // Bright green for parameters
	Yellow      string
	YellowDim   string
	YellowBright string  // Bright yellow for special highlights
	Orange      string
	OrangeLight string
	OrangeDark  string  // Dark orange for classes
	Red         string
	RedLight    string  // Light red for string escapes
	Purple      string

	// UI colors
	Border          string
	BorderSubtle    string
	Selection       string
	SelectionSubtle string
	LineHighlight   string
	Shadow          string

	// Semantic colors
	Error      string
	ErrorBg    string
	Warning    string
	Success    string
	SuccessBg  string
	Info       string

	// Syntax colors
	Comment     string
	CommentDark string  // Darker comment for better visibility
	String      string
	Number      string
	Keyword     string
	Function    string
	Type        string
	Variable    string
	Constant    string
	Operator    string
	Punctuation string
	Tag         string

	// Terminal colors
	Black      string
	White      string
}

// GetTronLegacyPalette returns the complete color palette
func GetTronLegacyPalette() TronLegacyPalette {
	return TronLegacyPalette{
		// Core colors
		Background:       "#14191fff",
		Foreground:       "#aec2e0ff",
		ForegroundDim:    "#647c9bff",
		ForegroundBright: "#dae3f1ff",
		ForegroundLight:  "#d0dfe6ff",

		// Accent colors
		Cyan:         "#6ee2ffff",
		Blue:         "#267fb5ff",
		BlueBright:   "#2b6db9ff",
		BlueBrighter: "#2f9ee2ff",
		Green:        "#C7F026ff",
		GreenDim:     "#4d5f07ff",
		GreenAlt:     "#95CC5Eff",
		GreenBright:  "#A6E22Eff",
		GreenParam:   "#95CC5Eff",
		Yellow:       "#FFE792ff",
		YellowDim:    "#ffd12cff",
		YellowBright: "#ffde10ff",
		Orange:       "#FFB20Dff",
		OrangeLight:  "#F79D1Eff",
		OrangeDark:   "#F79D1Eff",
		Red:          "#FF410Dff",
		RedLight:     "#ff6113ff",
		Purple:       "#967EFBff",

		// UI colors
		Border:          "#183c66ff",
		BorderSubtle:    "#133153ff",
		Selection:       "#183c66ff",
		SelectionSubtle: "#133153ff",
		LineHighlight:   "#28323eff",
		Shadow:          "#00000040",

		// Semantic colors
		Error:     "#F92672ff",
		ErrorBg:   "#660000ff",
		Warning:   "#FFE792ff",
		Success:   "#C7F026ff",
		SuccessBg: "#144212ff",
		Info:      "#6ee2ffff",

		// Syntax colors
		Comment:     "#6684a7ff", // Adjusted for WCAG AA contrast
		CommentDark: "#3c4b5dff",
		String:      "#FF410Dff",
		Number:      "#C7F026ff",
		Keyword:     "#748aa6ff",
		Function:    "#FFB20Dff",
		Type:        "#267fb5ff",
		Variable:    "#d0dfe6ff",
		Constant:    "#FFB20Dff",
		Operator:    "#748aa6ff",
		Punctuation: "#aec2e0ff",
		Tag:         "#267fb5ff",

		// Terminal colors
		Black: "#000000ff",
		White: "#ffffffff",
	}
}

// generateTheme creates the complete Zed theme structure
func generateTheme() map[string]any {
	palette := GetTronLegacyPalette()

	return map[string]any{
		"$schema": "https://zed.dev/schema/themes/v0.2.0.json",
		"name":    "Tron Legacy",
		"author":  "Hypermodules LLC, Bret Comnes (ported to Zed)",
		"themes": []any{
			map[string]any{
				"name":       "Tron Legacy",
				"appearance": "dark",
				"style":      generateThemeStyle(palette),
			},
		},
	}
}

// generateThemeStyle creates the style object for the theme
func generateThemeStyle(p TronLegacyPalette) map[string]any {
	return map[string]any{
		// Background and surfaces
		"background":                  p.Background,
		"surface.background":          p.Background,
		"elevated_surface.background": p.LineHighlight,
		"panel.background":            p.Background,

		// Borders
		"border":                p.Border,
		"border.variant":        p.BorderSubtle,
		"border.focused":        p.Cyan,
		"border.selected":       p.Cyan,
		"border.transparent":    p.Shadow,
		"border.disabled":       p.ForegroundDim,

		// Text colors
		"text":             p.Foreground,
		"text.muted":       p.Keyword,
		"text.placeholder": p.ForegroundDim,
		"text.disabled":    p.ForegroundDim,
		"text.accent":      p.Cyan,

		// Icons
		"icon":             p.Foreground,
		"icon.muted":       p.Keyword,
		"icon.disabled":    p.ForegroundDim,
		"icon.placeholder": p.ForegroundDim,
		"icon.accent":      p.Cyan,

		// Elements
		"element.background": p.LineHighlight,
		"element.hover":      p.Selection,
		"element.active":     p.BlueBright,
		"element.selected":   p.Selection,
		"element.disabled":   p.SelectionSubtle,

		// Ghost elements
		"ghost_element.background": p.Shadow,
		"ghost_element.hover":      p.LineHighlight,
		"ghost_element.active":     p.Selection,
		"ghost_element.selected":   p.Selection,
		"ghost_element.disabled":   p.SelectionSubtle,

		// Editor specific
		"editor.foreground":                          p.Foreground,
		"editor.background":                          p.Background,
		"editor.gutter.background":                   p.Background,
		"editor.subheader.background":                p.LineHighlight,
		"editor.active_line.background":              p.LineHighlight,
		"editor.highlighted_line.background":         p.LineHighlight,
		"editor.line_number":                         p.GreenDim,
		"editor.active_line_number":                  p.Green,
		"editor.hover_line_number":                   p.Green,
		"editor.invisible":                           p.ForegroundDim,
		"editor.wrap_guide":                          p.GreenDim,
		"editor.active_wrap_guide":                   p.Green,
		"editor.document_highlight.read_background":  p.Selection,
		"editor.document_highlight.write_background": p.BlueBright,

		// Text selection
		"selection.background": p.Selection,
		"editor.selection.background": p.Selection,

		// UI components
		"status_bar.background":         p.Background,
		"title_bar.background":          p.Background,
		"title_bar.inactive_background": p.Background,
		"toolbar.background":            p.Background,
		"tab_bar.background":            p.Background,
		"tab.inactive_background":       p.Background,
		"tab.active_background":         p.LineHighlight,
		"drop_target.background":        p.Shadow,
		"pane.focused_border":           p.Cyan,
		"panel.focused_border":          p.Cyan,

		// Scrollbars
		"scrollbar.thumb.background":       p.Selection,
		"scrollbar.thumb.hover_background": p.BlueBright,
		"scrollbar.thumb.border":           p.Selection,
		"scrollbar.track.background":       p.Shadow,
		"scrollbar.track.border":           p.Shadow,

		// Search
		"search.match_background": p.Yellow,
		"link_text.hover":         p.ForegroundDim,

		// Terminal colors
		"terminal.background":         p.Background,
		"terminal.foreground":         p.Foreground,
		"terminal.bright_foreground":  p.ForegroundBright,
		"terminal.dim_foreground":     p.ForegroundDim,
		"terminal.ansi.black":         p.Black,
		"terminal.ansi.bright_black":  p.ForegroundDim,
		"terminal.ansi.dim_black":     p.Shadow,
		"terminal.ansi.red":           p.Error,
		"terminal.ansi.bright_red":    p.Red,
		"terminal.ansi.dim_red":       p.ErrorBg,
		"terminal.ansi.green":         p.GreenAlt,
		"terminal.ansi.bright_green":  p.Green,
		"terminal.ansi.dim_green":     p.GreenDim,
		"terminal.ansi.yellow":        p.Yellow,
		"terminal.ansi.bright_yellow": p.Orange,
		"terminal.ansi.dim_yellow":    p.YellowDim,
		"terminal.ansi.blue":          p.Blue,
		"terminal.ansi.bright_blue":   p.Cyan,
		"terminal.ansi.dim_blue":      p.Selection,
		"terminal.ansi.magenta":       p.Purple,
		"terminal.ansi.bright_magenta": p.Purple,
		"terminal.ansi.dim_magenta":   p.Purple,
		"terminal.ansi.cyan":          p.Cyan,
		"terminal.ansi.bright_cyan":   p.Cyan,
		"terminal.ansi.dim_cyan":      p.Blue,
		"terminal.ansi.white":         p.Foreground,
		"terminal.ansi.bright_white":  p.White,
		"terminal.ansi.dim_white":     p.Keyword,

		// Version control
		"version_control.added":                  p.Green,
		"version_control.modified":               p.YellowDim,
		"version_control.deleted":                p.Error,
		"version_control.conflict_marker.ours":   p.Green,
		"version_control.conflict_marker.theirs": p.Error,

		// Status colors
		"error":                p.Error,
		"error.background":     p.ErrorBg,
		"error.border":         p.Error,
		"warning":              p.Warning,
		"warning.background":   p.Background,
		"warning.border":       p.Warning,
		"success":              p.Success,
		"success.background":   p.SuccessBg,
		"success.border":       p.Success,
		"info":                 p.Info,
		"info.background":      p.Background,
		"info.border":          p.Info,
		"hint":                 p.Blue,
		"hint.background":      p.Background,
		"hint.border":          p.Blue,
		"modified":             p.Yellow,
		"modified.background":  p.Background,
		"modified.border":      p.Yellow,
		"created":              p.Green,
		"created.background":   p.SuccessBg,
		"created.border":       p.Green,
		"deleted":              p.Error,
		"deleted.background":   p.ErrorBg,
		"deleted.border":       p.Error,
		"conflict":             p.Orange,
		"conflict.background":  p.OrangeLight,
		"conflict.border":      p.Orange,
		"renamed":              p.Blue,
		"renamed.background":   p.Background,
		"renamed.border":       p.Blue,
		"ignored":              p.ForegroundDim,
		"ignored.background":   p.Background,
		"ignored.border":       p.ForegroundDim,
		"hidden":               p.ForegroundDim,
		"hidden.background":    p.Background,
		"hidden.border":        p.ForegroundDim,
		"unreachable":          p.ForegroundDim,
		"unreachable.background": p.Background,
		"unreachable.border":   p.ForegroundDim,
		"predictive":           p.Purple,
		"predictive.background": p.Background,
		"predictive.border":    p.Purple,

		// Players (collaborative editing)
		"players": generatePlayers(p),

		// Syntax highlighting
		"syntax": generateSyntax(p),
	}
}

// generatePlayers creates collaborative editing colors
func generatePlayers(p TronLegacyPalette) []any {
	return []any{
		map[string]any{
			"cursor":     p.Blue,
			"background": p.Blue,
			"selection":  "#267fb580", // Semi-transparent blue
		},
		map[string]any{
			"cursor":     p.Cyan,
			"background": p.Cyan,
			"selection":  p.Selection,
		},
		map[string]any{
			"cursor":     p.Green,
			"background": p.Green,
			"selection":  p.GreenDim,
		},
		map[string]any{
			"cursor":     p.Orange,
			"background": p.Orange,
			"selection":  p.OrangeLight,
		},
		map[string]any{
			"cursor":     p.Purple,
			"background": p.Purple,
			"selection":  p.Purple,
		},
		map[string]any{
			"cursor":     p.Error,
			"background": p.Error,
			"selection":  p.ErrorBg,
		},
		map[string]any{
			"cursor":     p.Blue,
			"background": p.Blue,
			"selection":  p.Selection,
		},
		map[string]any{
			"cursor":     p.Keyword,
			"background": p.Keyword,
			"selection":  p.ForegroundDim,
		},
	}
}

// generateSyntax creates syntax highlighting rules
func generateSyntax(p TronLegacyPalette) map[string]any {
	return map[string]any{
		"comment": map[string]any{
			"color":       p.Comment,
			"font_style":  nil,
			"font_weight": nil,
		},
		"comment.doc": map[string]any{
			"color":       p.Comment,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string": map[string]any{
			"color":       p.String,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.escape": map[string]any{
			"color":       p.RedLight,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.regex": map[string]any{
			"color":       p.Cyan,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.special": map[string]any{
			"color":       p.Orange,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.special.symbol": map[string]any{
			"color":       p.Orange,
			"font_style":  nil,
			"font_weight": nil,
		},
		"number": map[string]any{
			"color":       p.Number,
			"font_style":  nil,
			"font_weight": nil,
		},
		"constant": map[string]any{
			"color":       p.Constant,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"boolean": map[string]any{
			"color":       p.Constant,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"variable": map[string]any{
			"color":       p.Variable,
			"font_style":  nil,
			"font_weight": nil,
		},
		"variable.special": map[string]any{
			"color":       p.Purple,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"variable.parameter": map[string]any{
			"color":       p.GreenParam,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"keyword": map[string]any{
			"color":       p.Keyword,
			"font_style":  nil,
			"font_weight": nil,
		},
		"operator": map[string]any{
			"color":       p.Operator,
			"font_style":  nil,
			"font_weight": nil,
		},
		"function": map[string]any{
			"color":       p.Function,
			"font_style":  nil,
			"font_weight": nil,
		},
		"function.builtin": map[string]any{
			"color":       p.Blue,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"type": map[string]any{
			"color":       p.Type,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"constructor": map[string]any{
			"color":       p.OrangeDark,
			"font_style":  "italic",
			"font_weight": 700,
		},
		"namespace": map[string]any{
			"color":       p.Variable,
			"font_style":  nil,
			"font_weight": nil,
		},
		"property": map[string]any{
			"color":       p.Variable,
			"font_style":  nil,
			"font_weight": nil,
		},
		"attribute": map[string]any{
			"color":       p.OrangeDark,
			"font_style":  nil,
			"font_weight": nil,
		},
		"tag": map[string]any{
			"color":       p.Tag,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation": map[string]any{
			"color":       p.Foreground,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.bracket": map[string]any{
			"color":       p.Foreground,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.delimiter": map[string]any{
			"color":       p.ForegroundDim,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.list_marker": map[string]any{
			"color":       p.GreenBright,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.special": map[string]any{
			"color":       p.Cyan,
			"font_style":  nil,
			"font_weight": nil,
		},
		"label": map[string]any{
			"color":       p.Variable,
			"font_style":  nil,
			"font_weight": nil,
		},
		"link_text": map[string]any{
			"color":       p.ForegroundDim,
			"font_style":  nil,
			"font_weight": nil,
		},
		"link_uri": map[string]any{
			"color":       p.ForegroundDim,
			"font_style":  nil,
			"font_weight": nil,
		},
		"emphasis": map[string]any{
			"color":       nil,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"emphasis.strong": map[string]any{
			"color":       nil,
			"font_style":  nil,
			"font_weight": 700,
		},
		"title": map[string]any{
			"color":       p.Variable,
			"font_style":  nil,
			"font_weight": 700,
		},
		"hint": map[string]any{
			"color":       p.Type,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"predictive": map[string]any{
			"color":       p.Purple,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"text.literal": map[string]any{
			"color":       p.Variable,
			"font_style":  nil,
			"font_weight": nil,
		},
		"preproc": map[string]any{
			"color":       p.Cyan,
			"font_style":  nil,
			"font_weight": nil,
		},
		"embedded": map[string]any{
			"color":       p.Variable,
			"font_style":  nil,
			"font_weight": nil,
		},
		"primary": map[string]any{
			"color":       p.White,
			"font_style":  nil,
			"font_weight": 700,
		},
		"enum": map[string]any{
			"color":       p.OrangeDark,
			"font_style":  nil,
			"font_weight": nil,
		},
		"variant": map[string]any{
			"color":       p.OrangeDark,
			"font_style":  nil,
			"font_weight": nil,
		},
		"selector": map[string]any{
			"color":       p.GreenParam,
			"font_style":  nil,
			"font_weight": nil,
		},
		"selector.pseudo": map[string]any{
			"color":       p.Cyan,
			"font_style":  nil,
			"font_weight": nil,
		},
	}
}

func main() {
	// Generate the theme
	theme := generateTheme()

	// Convert to JSON with proper formatting
	jsonData, err := json.MarshalIndent(theme, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Write to file in themes directory
	err = os.WriteFile("../themes/tron-legacy.json", jsonData, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Theme generated successfully: ../themes/tron-legacy.json")
}
