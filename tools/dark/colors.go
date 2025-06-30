package dark

// GetColors returns a map of all unique colors used in the dark theme
func GetColors() map[string]string {
	return map[string]string{
		// Monochrome
		"black":       "#000000ff",
		"purewhite":   "#ffffffff",
		"shadow":      "#00000040",
		"transparent": "#00000000",

		// Gray scale (dark to light)
		"gray900": "#14191fff", // Background
		"gray800": "#1c2128ff", // LineHighlight (more neutral)
		"gray750": "#23282fff", // Status bar (neutral dark gray)
		"gray700": "#2a3039ff", // BorderSubtle (neutral)
		"gray600": "#323842ff", // Border, Selection (neutral)
		"gray500": "#647c9bff", // ForegroundDim
		"gray400": "#586676ff", // Comment
		"gray200": "#aec2e0ff", // Foreground
		"gray50":  "#dae3f1ff", // ForegroundBright

		// Neutral UI grays (for non-blue UI elements)
		"neutral800": "#1a1d23ff", // Neutral highlight
		"neutral600": "#2d3139ff", // Neutral border

		// Blues
		"blue500":       "#267fb5ff", // Primary blue
		"blue400":       "#2b6db9ff", // Bright blue
		"blue300":       "#4a95b3ff", // Cyan dim
		"blue200":       "#6ee2ffff", // Primary cyan
		"blue200bright": "#c8d9e8ff", // Bright blue-fg (90% lightness)
		"blue200alpha":  "#6ee2ff40", // Primary cyan with 25% opacity for search

		// Greens
		"green700": "#144212ff", // Success bg
		"green600": "#4d5f07ff", // Green dim
		"green500": "#95CC5Eff", // Alt green
		"green300": "#C7F026ff", // Primary green

		// Yellows/Oranges
		"yellow500": "#FFE792ff", // Primary yellow
		"yellow400": "#ffd12cff", // Dim yellow
		"orange500": "#FFB20Dff", // Primary orange
		"orange400": "#F79D1Eff", // Light orange

		// Reds
		"red700": "#660000ff", // Error bg
		"red500": "#F92672ff", // Error red
		"red400": "#FF410Dff", // Primary red (string color)
		"red300": "#FF5F52ff", // Light red

		// Purples
		"purple500": "#967EFBff", // Primary purple
		"pink500":   "#ff79c6ff", // Primary pink

		// Special alpha variants
		"blue500alpha": "#267fb580", // Selection with alpha
		"droptarget":   "#6ee2ff18", // Very low opacity cyan for drop targets
	}
}
