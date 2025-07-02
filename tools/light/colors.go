package light

// GetColors returns a map of all unique colors used in the light theme
func GetColors() map[string]string {
	return map[string]string{
		// Monochrome
		"black":       "#000000ff",
		"purewhite":   "#ffffffff",
		"shadow":      "#00000020",
		"transparent": "#00000000",

		// Gray scale (light to dark) - inverted from dark theme
		"gray50":  "#f5f7faff", // Original gray background
		"blue50":  "#c8d9e8ff", // Light blue background based on cyan
		"gray100": "#e8ecf2ff", // LineHighlight
		"gray150": "#dfe5edff", // Status bar
		"gray200": "#d1dae6ff", // BorderSubtle
		"gray300": "#b8c5d6ff", // Border, Selection
		"gray400": "#8a9db5ff", // Line numbers, etc
		"gray500": "#6b7e96ff", // Comment
		"gray600": "#526073ff", // ForegroundDim (sidebar text)
		"gray700": "#2d3e4fff", // Foreground
		"gray900": "#14191fff", // ForegroundBright

		// Blues - adjusted for light mode
		"blue600":       "#1a5f8aff", // Primary blue (darker for light bg)
		"blue500":       "#267fb5ff", // Standard blue
		"blue400":       "#3988c0ff", // Medium blue
		"blue300":       "#4a95b3ff", // Light blue
		"blue200":       "#0099ccff", // Primary cyan (darker for light bg)
		"blue100":       "#c8e6f5ff", // Very light blue for backgrounds
		"blue200alpha":  "#0099cc30", // Primary cyan with alpha for search

		// Greens - adjusted for light mode
		"green100": "#e6f7e3ff", // Success bg
		"green600": "#3a5f00ff", // Dark green
		"green500": "#5a8b2cff", // Alt green
		"green400": "#7aad3aff", // Primary green (darker for light bg)

		// Yellows/Oranges - adjusted for light mode
		"yellow600": "#c9a000ff", // Primary yellow (darker for light bg)
		"yellow500": "#dbb200ff", // Medium yellow
		"orange600": "#cc7700ff", // Primary orange (darker for light bg)
		"orange500": "#e68a00ff", // Light orange

		// Reds - adjusted for light mode
		"red100": "#ffe6e6ff", // Error bg
		"red600": "#cc0033ff", // Error red (darker for light bg)
		"red500": "#d91e18ff", // Primary red (string color)
		"red400": "#e74c3cff", // Light red

		// Purples - adjusted for light mode
		"purple600": "#6a56ccff", // Primary purple (darker for light bg)
		"pink600":   "#d1459aff", // Primary pink (darker for light bg)

		// Special alpha variants
		"blue300alpha":  "#4a95b340", // Selection with alpha
		"droptarget":    "#0099cc18", // Very low opacity cyan for drop targets
		"gray300alpha":  "#6b7e964d", // Gray500 with 30% opacity for better contrast on scrollbar
		"blue400alpha":  "#0099ccb3", // Primary cyan with 70% opacity for scrollbar hover
		"oneTrackBorder": "#2e333cff", // One theme's scrollbar track border color

		// Additional alpha variants for better transparency
		"gray100alpha75": "#e8ecf2bf", // Active line background with 75% opacity
		"gray500alpha30": "#6b7e9640", // Wrap guide with 25% opacity using darker gray
		"gray600alpha50": "#52607359", // Active wrap guide with 35% opacity using even darker gray
		"blue200alpha10": "#0099cc1a", // Document highlight read with 10% opacity
		"blue200alpha40": "#0099cc66", // Document highlight write with 40% opacity

		// Standardized player selection opacity (24% like other themes)
		"player1alpha": "#4a95b33d", // Blue with 24% opacity
		"player3alpha": "#3a5f003d", // Green with 24% opacity
		"player4alpha": "#e68a003d", // Orange with 24% opacity
	}
}
