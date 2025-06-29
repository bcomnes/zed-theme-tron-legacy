package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
)

// Color represents an RGBA color
type Color struct {
	R, G, B, A float64
}

// hslToRgb converts HSL values to RGB
// h: 0-360, s: 0-1, l: 0-1
func hslToRgb(h, s, l float64) (r, g, b float64) {
	h = h / 360.0 // normalize to 0-1

	if s == 0 {
		// achromatic
		r, g, b = l, l, l
		return
	}

	hue2rgb := func(p, q, t float64) float64 {
		if t < 0 {
			t += 1
		}
		if t > 1 {
			t -= 1
		}
		if t < 1.0/6.0 {
			return p + (q-p)*6*t
		}
		if t < 1.0/2.0 {
			return q
		}
		if t < 2.0/3.0 {
			return p + (q-p)*(2.0/3.0-t)*6
		}
		return p
	}

	var q float64
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q

	r = hue2rgb(p, q, h+1.0/3.0)
	g = hue2rgb(p, q, h)
	b = hue2rgb(p, q, h-1.0/3.0)

	return
}

// rgbToHsl converts RGB values to HSL
// r, g, b: 0-1
func rgbToHsl(r, g, b float64) (h, s, l float64) {
	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	l = (max + min) / 2

	if max == min {
		h, s = 0, 0 // achromatic
		return
	}

	d := max - min
	if l > 0.5 {
		s = d / (2 - max - min)
	} else {
		s = d / (max + min)
	}

	switch max {
	case r:
		h = (g - b) / d
		if g < b {
			h += 6
		}
	case g:
		h = (b-r)/d + 2
	case b:
		h = (r-g)/d + 4
	}
	h /= 6
	h *= 360 // convert to degrees

	return
}

// hexToRgb converts hex color to RGB values (0-1)
func hexToRgb(hex string) (r, g, b float64, err error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %s", hex)
	}

	var ri, gi, bi int
	_, err = fmt.Sscanf(hex, "%02x%02x%02x", &ri, &gi, &bi)
	if err != nil {
		return 0, 0, 0, err
	}

	r = float64(ri) / 255.0
	g = float64(gi) / 255.0
	b = float64(bi) / 255.0
	return
}

// rgbToHex converts RGB values (0-1) to hex string
func rgbToHex(r, g, b float64) string {
	ri := int(math.Round(r * 255))
	gi := int(math.Round(g * 255))
	bi := int(math.Round(b * 255))
	return fmt.Sprintf("#%02x%02x%02xff", ri, gi, bi)
}

// rgbaToHex converts RGBA values (0-1) to hex string with alpha
func rgbaToHex(r, g, b, a float64) string {
	ri := int(math.Round(r * 255))
	gi := int(math.Round(g * 255))
	bi := int(math.Round(b * 255))
	ai := int(math.Round(a * 255))
	return fmt.Sprintf("#%02x%02x%02x%02x", ri, gi, bi, ai)
}

// generateColorMap creates the complete color map with all conversions
func generateColorMap() map[string]string {
	colorMap := map[string]string{
		// Direct hex values
		"black":                        "#000000ff",
		"red-string":                   "#FF410Dff",
		"red-error":                    "#F92672ff",
		"red-error-background":         "#660000ff",
		"green":                        "#95CC5Eff",
		"green-light":                  "#C7F026ff",
		"green-light-boost":            "#A6E22Eff",
		"green-light-boost-background": "#144212ff",
		"yellow-highlight":             "#FFE792ff",
		"orange-entity":                "#FFB20Dff",
		"orange-entity-dark":           "#F79D1Eff",
		"blue-pure":                    "#2F33ABff",
		"blue-bg":                      "#14191fff",
		"blue-fg":                      "#aec2e0ff",
		"blue-accent":                  "#6ee2ffff",
		"blue-storage":                 "#267fb5ff",
		"blue-highlight":               "#183c66ff",
		"gray-copy":                    "#748aa6ff",
		"purple-accent":                "#967EFBff",
	}

	// Convert white from HSL
	r, g, b := hslToRgb(0, 0, 1.0)
	colorMap["white"] = rgbToHex(r, g, b)

	// Black shadow with alpha
	colorMap["black-shadow"] = rgbaToHex(0, 0, 0, 0.25)

	// Green light subdue - lightness 20%
	r, g, b, _ = hexToRgb("#C7F026")
	h, s, _ := rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, 0.20)
	colorMap["green-light-subdue"] = rgbToHex(r, g, b)

	// Blue bg highlight - lightness 20%
	r, g, b, _ = hexToRgb("#14191f")
	h, s, _ = rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, 0.20)
	colorMap["blue-bg-highlight"] = rgbToHex(r, g, b)

	// Blue fg white - lightness 90%
	r, g, b, _ = hexToRgb("#aec2e0")
	h, s, _ = rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, 0.90)
	colorMap["blue-fg-white"] = rgbToHex(r, g, b)

	// Blue highlight contrast - min contrast 2.5 against itself
	r, g, b, _ = hexToRgb("#183c66")
	h, s, l := rgbToHsl(r, g, b)
	// Make it lighter to create contrast
	r2, g2, b2 := hslToRgb(h, s, math.Min(l+0.2, 1.0))
	colorMap["blue-highlight-contrast"] = rgbToHex(r2, g2, b2)

	// Blue highlight subdue - lightness 20%
	r, g, b, _ = hexToRgb("#183c66")
	h, s, _ = rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, 0.20)
	colorMap["blue-hilight-subdue"] = rgbToHex(r, g, b)

	// Gray copy dark - lightness 50%
	r, g, b, _ = hexToRgb("#748aa6")
	h, s, _ = rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, 0.50)
	colorMap["gray-copy-dark"] = rgbToHex(r, g, b)

	// Gray comment - min contrast 4.5 against blue-bg (WCAG AA)
	// Using a lighter gray-blue for better readability
	colorMap["gray-comment"] = "#6684a7ff"

	// Yellow highlight subdue - decrease lightness by 20%
	r, g, b, _ = hexToRgb("#FFE792")
	h, s, l = rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, math.Max(l-0.20, 0))
	colorMap["yellow-highlight-subdue"] = rgbToHex(r, g, b)

	return colorMap
}



// generateThemeStyle creates the style object for the dark theme
func generateThemeStyle(colors map[string]string) map[string]interface{} {
	// Helper function to get color
	getColor := func(name string) string {
		if color, ok := colors[name]; ok {
			return color
		}
		// Default fallback
		return "#ff00ffff" // Magenta for debugging missing colors
	}

	// Create the base style
	style := map[string]interface{}{
		// UI Colors
		"background":                      getColor("blue-bg"),
		"border":                          getColor("blue-highlight"),
		"border.variant":                  getColor("blue-hilight-subdue"),
		"border.focused":                  getColor("blue-accent"),
		"border.selected":                 getColor("blue-accent"),
		"border.transparent":              getColor("black-shadow"),
		"border.disabled":                 getColor("gray-copy-dark"),
		"elevated_surface.background":     getColor("blue-bg-highlight"),
		"surface.background":              getColor("blue-bg"),
		"element.background":              getColor("blue-bg-highlight"),
		"element.hover":                   getColor("blue-highlight"),
		"element.active":                  getColor("blue-highlight-contrast"),
		"element.selected":                getColor("blue-highlight"),
		"element.disabled":                getColor("blue-hilight-subdue"),
		"drop_target.background":          colors["black-shadow"],
		"ghost_element.background":        colors["black-shadow"],
		"ghost_element.hover":             getColor("blue-bg-highlight"),
		"ghost_element.active":            getColor("blue-highlight"),
		"ghost_element.selected":          getColor("blue-highlight"),
		"ghost_element.disabled":          getColor("blue-hilight-subdue"),
		"text":                            getColor("blue-fg"),
		"text.muted":                      getColor("gray-copy"),
		"text.placeholder":                getColor("gray-copy-dark"),
		"text.disabled":                   getColor("gray-copy-dark"),
		"text.accent":                     getColor("blue-accent"),
		"icon":                            getColor("blue-fg"),
		"icon.muted":                      getColor("gray-copy"),
		"icon.disabled":                   getColor("gray-copy-dark"),
		"icon.placeholder":                getColor("gray-copy-dark"),
		"icon.accent":                     getColor("blue-accent"),
		"status_bar.background":           getColor("blue-bg"),
		"title_bar.background":            getColor("blue-bg"),
		"title_bar.inactive_background":   getColor("blue-bg"),
		"toolbar.background":              getColor("blue-bg"),
		"tab_bar.background":              getColor("blue-bg"),
		"tab.inactive_background":         getColor("blue-bg"),
		"tab.active_background":           getColor("blue-bg-highlight"),
		"search.match_background":         getColor("yellow-highlight"),
		"panel.background":                getColor("blue-bg"),
		"panel.focused_border":            getColor("blue-accent"),
		"pane.focused_border":             getColor("blue-accent"),
		"scrollbar.thumb.background":      getColor("blue-highlight"),
		"scrollbar.thumb.hover_background": getColor("blue-highlight-contrast"),
		"scrollbar.thumb.border":          getColor("blue-highlight"),
		"scrollbar.track.background":      colors["black-shadow"],
		"scrollbar.track.border":          colors["black-shadow"],

		// Editor colors
		"editor.foreground":                          getColor("blue-fg"),
		"editor.background":                          getColor("blue-bg"),
		"editor.gutter.background":                   getColor("blue-bg"),
		"editor.subheader.background":                getColor("blue-bg-highlight"),
		"editor.active_line.background":              getColor("blue-bg-highlight"),
		"editor.highlighted_line.background":         getColor("blue-bg-highlight"),
		"editor.line_number":                         getColor("green-light-subdue"),
		"editor.active_line_number":                  getColor("green-light"),
		"editor.hover_line_number":                   getColor("green-light"),
		"editor.invisible":                           getColor("gray-copy-dark"),
		"editor.wrap_guide":                          getColor("green-light-subdue"),
		"editor.active_wrap_guide":                   getColor("green-light"),
		"editor.document_highlight.read_background":  getColor("blue-highlight"),
		"editor.document_highlight.write_background": getColor("blue-highlight-contrast"),

		// Terminal colors
		"terminal.background":      getColor("blue-bg"),
		"terminal.foreground":      getColor("blue-fg"),
		"terminal.bright_foreground": getColor("blue-fg-white"),
		"terminal.dim_foreground":  getColor("gray-copy-dark"),
		"terminal.ansi.black":      getColor("black"),
		"terminal.ansi.bright_black": getColor("gray-copy-dark"),
		"terminal.ansi.dim_black":  colors["black-shadow"],
		"terminal.ansi.red":        getColor("red-error"),
		"terminal.ansi.bright_red": getColor("red-string"),
		"terminal.ansi.dim_red":    getColor("red-error-background"),
		"terminal.ansi.green":      getColor("green"),
		"terminal.ansi.bright_green": getColor("green-light"),
		"terminal.ansi.dim_green":  getColor("green-light-subdue"),
		"terminal.ansi.yellow":     getColor("yellow-highlight"),
		"terminal.ansi.bright_yellow": getColor("orange-entity"),
		"terminal.ansi.dim_yellow": getColor("yellow-highlight-subdue"),
		"terminal.ansi.blue":       getColor("blue-storage"),
		"terminal.ansi.bright_blue": getColor("blue-accent"),
		"terminal.ansi.dim_blue":   getColor("blue-highlight"),
		"terminal.ansi.magenta":    getColor("purple-accent"),
		"terminal.ansi.bright_magenta": getColor("purple-accent"),
		"terminal.ansi.dim_magenta": getColor("purple-accent"),
		"terminal.ansi.cyan":       getColor("blue-accent"),
		"terminal.ansi.bright_cyan": getColor("blue-accent"),
		"terminal.ansi.dim_cyan":   getColor("blue-storage"),
		"terminal.ansi.white":      getColor("blue-fg"),
		"terminal.ansi.bright_white": getColor("white"),
		"terminal.ansi.dim_white":  getColor("gray-copy"),

		// Version control and other colors...
		"link_text.hover": getColor("gray-copy-dark"),
		"version_control.added":    getColor("green-light"),
		"version_control.modified": getColor("yellow-highlight-subdue"),
		"version_control.deleted":  getColor("red-error"),
		"version_control.conflict_marker.ours":   getColor("green-light"),
		"version_control.conflict_marker.theirs": getColor("red-error"),
		"conflict":            getColor("orange-entity"),
		"conflict.background": getColor("orange-entity-dark"),
		"conflict.border":     getColor("orange-entity"),
		"created":             getColor("green-light"),
		"created.background":  getColor("green-light-boost-background"),
		"created.border":      getColor("green-light"),
		"deleted":             getColor("red-error"),
		"deleted.background":  getColor("red-error-background"),
		"deleted.border":      getColor("red-error"),
		"error":               getColor("red-error"),
		"error.background":    getColor("red-error-background"),
		"error.border":        getColor("red-error"),
		"hidden":              getColor("gray-copy-dark"),
		"hidden.background":   getColor("blue-bg"),
		"hidden.border":       getColor("gray-copy-dark"),
		"hint":                getColor("blue-storage"),
		"hint.background":     getColor("blue-bg"),
		"hint.border":         getColor("blue-storage"),
		"ignored":             getColor("gray-copy-dark"),
		"ignored.background":  getColor("blue-bg"),
		"ignored.border":      getColor("gray-copy-dark"),
		"info":                getColor("blue-accent"),
		"info.background":     getColor("blue-bg"),
		"info.border":         getColor("blue-accent"),
		"modified":            getColor("yellow-highlight"),
		"modified.background": getColor("blue-bg"),
		"modified.border":     getColor("yellow-highlight"),
		"predictive":          getColor("purple-accent"),
		"predictive.background": getColor("blue-bg"),
		"predictive.border":   getColor("purple-accent"),
		"renamed":             getColor("blue-storage"),
		"renamed.background":  getColor("blue-bg"),
		"renamed.border":      getColor("blue-storage"),
		"success":             getColor("green-light"),
		"success.background":  getColor("green-light-boost-background"),
		"success.border":      getColor("green-light"),
		"unreachable":         getColor("gray-copy-dark"),
		"unreachable.background": getColor("blue-bg"),
		"unreachable.border":  getColor("gray-copy-dark"),
		"warning":             getColor("yellow-highlight"),
		"warning.background":  getColor("blue-bg"),
		"warning.border":      getColor("yellow-highlight"),
	}

	// Add players
	style["players"] = generatePlayers(colors)

	// Add syntax highlighting
	style["syntax"] = generateSyntax(colors)

	return style
}

// generatePlayers creates the collaborative editing colors
func generatePlayers(colors map[string]string) []interface{} {
	getColor := func(name string) string {
		if color, ok := colors[name]; ok {
			return color
		}
		return "#ff00ffff"
	}

	return []interface{}{
		map[string]interface{}{
			"cursor":     getColor("yellow-highlight"),
			"background": getColor("yellow-highlight"),
			"selection":  getColor("yellow-highlight-subdue"),
		},
		map[string]interface{}{
			"cursor":     getColor("blue-accent"),
			"background": getColor("blue-accent"),
			"selection":  getColor("blue-highlight"),
		},
		map[string]interface{}{
			"cursor":     getColor("green-light"),
			"background": getColor("green-light"),
			"selection":  getColor("green-light-subdue"),
		},
		map[string]interface{}{
			"cursor":     getColor("orange-entity"),
			"background": getColor("orange-entity"),
			"selection":  getColor("orange-entity-dark"),
		},
		map[string]interface{}{
			"cursor":     getColor("purple-accent"),
			"background": getColor("purple-accent"),
			"selection":  getColor("purple-accent"),
		},
		map[string]interface{}{
			"cursor":     getColor("red-error"),
			"background": getColor("red-error"),
			"selection":  getColor("red-error-background"),
		},
		map[string]interface{}{
			"cursor":     getColor("blue-storage"),
			"background": getColor("blue-storage"),
			"selection":  getColor("blue-highlight"),
		},
		map[string]interface{}{
			"cursor":     getColor("gray-copy"),
			"background": getColor("gray-copy"),
			"selection":  getColor("gray-copy-dark"),
		},
	}
}

// generateSyntax creates the syntax highlighting rules
func generateSyntax(colors map[string]string) map[string]interface{} {
	getColor := func(name string) string {
		if color, ok := colors[name]; ok {
			return color
		}
		return "#ff00ffff"
	}

	return map[string]interface{}{
		"comment": map[string]interface{}{
			"color":       getColor("gray-comment"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"comment.doc": map[string]interface{}{
			"color":       getColor("gray-comment"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"string": map[string]interface{}{
			"color":       getColor("red-string"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.escape": map[string]interface{}{
			"color":       getColor("orange-entity"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.regex": map[string]interface{}{
			"color":       getColor("red-string"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.special": map[string]interface{}{
			"color":       getColor("orange-entity"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.special.symbol": map[string]interface{}{
			"color":       getColor("orange-entity"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"number": map[string]interface{}{
			"color":       getColor("green-light"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"constant": map[string]interface{}{
			"color":       getColor("orange-entity"),
			"font_style":  "italic",
			"font_weight": nil,
		},
		"boolean": map[string]interface{}{
			"color":       getColor("orange-entity"),
			"font_style":  "italic",
			"font_weight": nil,
		},
		"variable": map[string]interface{}{
			"color":       getColor("blue-fg-white"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"variable.special": map[string]interface{}{
			"color":       getColor("purple-accent"),
			"font_style":  "italic",
			"font_weight": nil,
		},
		"keyword": map[string]interface{}{
			"color":       getColor("gray-copy"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"operator": map[string]interface{}{
			"color":       getColor("gray-copy"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation": map[string]interface{}{
			"color":       getColor("blue-fg"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.bracket": map[string]interface{}{
			"color":       getColor("blue-fg"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.delimiter": map[string]interface{}{
			"color":       getColor("gray-copy-dark"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.list_marker": map[string]interface{}{
			"color":       getColor("green-light-boost"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.special": map[string]interface{}{
			"color":       getColor("blue-accent"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"function": map[string]interface{}{
			"color":       getColor("orange-entity"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"type": map[string]interface{}{
			"color":       getColor("blue-storage"),
			"font_style":  "italic",
			"font_weight": nil,
		},
		"constructor": map[string]interface{}{
			"color":       getColor("orange-entity"),
			"font_style":  "italic",
			"font_weight": 700,
		},
		"namespace": map[string]interface{}{
			"color":       getColor("blue-fg-white"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"property": map[string]interface{}{
			"color":       getColor("blue-fg-white"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"attribute": map[string]interface{}{
			"color":       getColor("orange-entity-dark"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"tag": map[string]interface{}{
			"color":       getColor("blue-storage"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"label": map[string]interface{}{
			"color":       getColor("blue-fg-white"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"link_text": map[string]interface{}{
			"color":       getColor("gray-copy-dark"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"link_uri": map[string]interface{}{
			"color":       getColor("gray-copy-dark"),
			"font_style":  nil,
			"font_weight": nil,
		},
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
			"color":       getColor("blue-fg-white"),
			"font_style":  nil,
			"font_weight": 700,
		},
		"hint": map[string]interface{}{
			"color":       getColor("blue-storage"),
			"font_style":  "italic",
			"font_weight": nil,
		},
		"predictive": map[string]interface{}{
			"color":       getColor("purple-accent"),
			"font_style":  "italic",
			"font_weight": nil,
		},
		"text.literal": map[string]interface{}{
			"color":       getColor("blue-fg-white"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"preproc": map[string]interface{}{
			"color":       getColor("blue-accent"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"embedded": map[string]interface{}{
			"color":       getColor("blue-fg-white"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"primary": map[string]interface{}{
			"color":       getColor("white"),
			"font_style":  nil,
			"font_weight": 700,
		},
		"enum": map[string]interface{}{
			"color":       getColor("orange-entity"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"variant": map[string]interface{}{
			"color":       getColor("orange-entity"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"selector": map[string]interface{}{
			"color":       getColor("blue-fg-white"),
			"font_style":  nil,
			"font_weight": nil,
		},
		"selector.pseudo": map[string]interface{}{
			"color":       getColor("blue-fg-white"),
			"font_style":  nil,
			"font_weight": nil,
		},
	}
}

// generateZedTheme creates the complete Zed theme structure
func generateZedTheme() map[string]interface{} {
	colors := generateColorMap()

	theme := map[string]interface{}{
		"$schema": "https://zed.dev/schema/themes/v0.2.0.json",
		"name":    "Tron Legacy",
		"author":  "Hypermodules LLC, Bret Comnes (ported to Zed)",
		"themes": []interface{}{
			map[string]interface{}{
				"name":       "Tron Legacy",
				"appearance": "dark",
				"style":      generateThemeStyle(colors),
			},
		},
	}

	return theme
}

func main() {
	// Generate the theme
	theme := generateZedTheme()

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
