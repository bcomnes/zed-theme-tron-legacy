package light

// GetColors returns a map of all unique colors used in the light theme
func GetColors() map[string]string {
	return map[string]string{
		// Monochrome
		"black":     "#000000ff",
		"purewhite": "#ffffffff",
		"shadow":    "#00000020",

		// Gray scale (light to dark)
		"gray50":  "#e6f4ffff", // Background - light blue as requested
		"gray100": "#d1e9f6ff", // LineHighlight
		"gray200": "#b8dceeff", // BorderSubtle
		"gray300": "#9cc8e0ff", // Border, Selection
		"gray400": "#5a7a99ff", // ForegroundDim
		"gray500": "#456484ff", // Comment
		"gray700": "#1a3a5cff", // Foreground
		"gray900": "#0d1f33ff", // ForegroundBright

		// Blues (darkened for contrast)
		"blue700":       "#14537dff", // Primary blue - darkened
		"blue600":       "#1a6091ff", // Bright blue - darkened
		"blue500":       "#2875a0ff", // Cyan dim - darkened
		"blue400":       "#2190b8ff", // Primary cyan - darkened
		"blue700bright": "#1a3a5cff", // Bright blue-fg - darkened
		"blue400alpha":  "#2190b830", // Primary cyan with alpha for search

		// Greens (adjusted for contrast)
		"green100": "#d4f5cfff", // Success bg - light
		"green200": "#93c088ff", // Green dim - lighter
		"green600": "#4a8f2cff", // Alt green - darkened
		"green700": "#3d7a1fff", // Primary green - darkened

		// Yellows/Oranges (darkened for contrast)
		"yellow700": "#b39200ff", // Primary yellow - darkened
		"yellow600": "#9c7d00ff", // Dim yellow - darkened
		"orange700": "#c27700ff", // Primary orange - darkened
		"orange600": "#b36b00ff", // Light orange - darkened

		// Reds (darkened for contrast)
		"red100": "#ffe6e6ff", // Error bg - light
		"red700": "#c91959ff", // Error red - darkened
		"red600": "#d92807ff", // Primary red (string color) - darkened
		"red500": "#e03426ff", // Light red - darkened

		// Purples (darkened for contrast)
		"purple700": "#6952c9ff", // Primary purple - darkened
		"pink700":   "#d9439cff", // Primary pink - darkened

		// Special alpha variants
		"blue700alpha": "#14537d40", // Selection with alpha
	}
}
