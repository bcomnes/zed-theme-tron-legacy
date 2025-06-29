package atelier

// GetColors returns a map of all unique colors used in the Atelier-inspired blue-tinted light theme
func GetColors() map[string]string {
	return map[string]string{
		// Monochrome
		"black":     "#000000ff",
		"purewhite": "#ffffffff",
		"shadow":    "#00000015",

		// Base colors - blue-tinted instead of brown
		"base00": "#f0f4f8ff", // Background - cool off-white with blue tint
		"base01": "#e4ecf4ff", // Element background - light blue-gray
		"base02": "#d0dce8ff", // Selection/border - medium light blue-gray
		"base03": "#a8b8ccff", // Comments/inactive - medium blue-gray
		"base04": "#8394aeff", // Dark blue-gray
		"base05": "#4a5f7eff", // Default text - dark blue-gray
		"base06": "#2c4566ff", // Bright text - darker blue-gray
		"base07": "#1a2b3fff", // Darkest - very dark blue

		// Accent colors - adjusted for blue theme
		"red":     "#e53935ff", // Slightly warmer red
		"orange":  "#f57c00ff", // Warmer orange
		"yellow":  "#fbc02dff", // Golden yellow
		"green":   "#43a047ff", // True green
		"teal":    "#00897bff", // Teal instead of yellow-green
		"cyan":    "#00acc1ff", // Bright cyan
		"blue":    "#1e88e5ff", // Primary blue
		"purple":  "#5e35b1ff", // Deep purple
		"magenta": "#8e24aaff", // Magenta

		// Special variants
		"blue_bright": "#039be5ff", // Brighter blue
		"blue_dim":    "#5c7caaff", // Dimmed blue
		"green_dim":   "#66bb6aff", // Softer green
		"red_dim":     "#ef5350ff", // Softer red

		// Alpha variants for overlays
		"blue_alpha":     "#1e88e530", // Blue with alpha
		"gray_alpha":     "#a8b8cc40", // Gray with alpha
		"shadow_alpha":   "#00000010", // Very light shadow
		"highlight_alpha": "#039be520", // Highlight with alpha
	}
}
