package dark

import (
	_ "embed"
	"log"

	"github.com/bcomnes/zed-theme-tron-legacy/tools/csscolors"
	"github.com/bcomnes/zed-theme-tron-legacy/tools/palette"
)

//go:embed colors.css
var colorsCSS []byte

// GetPalette returns the dark theme palette
func GetPalette() palette.TronThemePalette {
	// Load colors from CSS
	colors, err := csscolors.LoadColors(colorsCSS)
	if err != nil {
		log.Fatalf("Failed to load dark colors: %v", err)
	}

	return palette.TronThemePalette{
		// Core text colors
		Foreground:       colors.MustGet("gray200"),
		ForegroundDim:    colors.MustGet("gray500"),
		ForegroundBright: colors.MustGet("gray50"),

		// Core backgrounds
		Background:       colors.MustGet("gray900"),
		EditorBackground: colors.MustGet("gray900"),

		// UI backgrounds
		TabBarBackground:           colors.MustGet("gray800"),
		StatusBarBackground:        colors.MustGet("gray750"),
		EditorSubheaderBackground:  colors.MustGet("gray800"),
		TitleBarInactiveBackground: colors.MustGet("gray800"),
		PanelOverlayBackground:     colors.MustGet("gray850"),
		PanelOverlayHover:          colors.MustGet("gray700"),

		// UI elements
		Border:               colors.MustGet("neutral600"),
		BorderSubtle:         colors.MustGet("gray700"),
		BorderFocused:        colors.MustGet("blue200"),
		Selection:            colors.MustGet("gray700"),
		Highlight:            colors.MustGet("neutral800"),
		ElementHover:         colors.MustGet("gray700"),
		Transparent:          colors.MustGet("transparent"),
		DropTargetBackground: colors.MustGet("blue200Alpha18"),

		// Semantic status colors
		Error:     colors.MustGet("red500"),
		ErrorBg:   colors.MustGet("red700"),
		Warning:   colors.MustGet("yellow500"),
		Success:   colors.MustGet("green300"),
		SuccessBg: colors.MustGet("green700"),
		Info:      colors.MustGet("blue200"),
		Hint:      colors.MustGet("gray500"),

		// Editor specific
		LineNumber:             colors.MustGet("gray500"),
		ActiveLineBackground:   colors.MustGet("gray800Alpha75"),
		WrapGuide:              colors.MustGet("gray500Alpha30"),
		ActiveWrapGuide:        colors.MustGet("gray400Alpha50"),
		DocumentHighlightRead:  colors.MustGet("blue200Alpha10"),
		DocumentHighlightWrite: colors.MustGet("blue200Alpha40"),
		SearchHighlight:        colors.MustGet("blue200Alpha"),

		// Syntax - basic types
		Comment:       colors.MustGet("gray400"),
		String:        colors.MustGet("red400"),
		StringEscape:  colors.MustGet("red300"),
		Number:        colors.MustGet("green300"),
		Keyword:       colors.MustGet("blue500"),
		Function:  colors.MustGet("orange500"),
		Variable:  colors.MustGet("blue200Bright"),
		Property:  colors.MustGet("green500"),
		Namespace: colors.MustGet("blue300"),

		// Syntax - advanced types
		Type:          colors.MustGet("blue500"),
		Constructor:   colors.MustGet("orange400"),
		EnumType:      colors.MustGet("orange400"),
		AttributeType: colors.MustGet("orange400"),
		EmbeddedType:  colors.MustGet("yellow400"),
		SpecialSyntax: colors.MustGet("pink500"),
		Regex:         colors.MustGet("blue200"),
		Selector:      colors.MustGet("green500"),

		// Syntax - punctuation
		Punctuation:        colors.MustGet("gray200"),
		PunctuationSpecial: colors.MustGet("gray500"),

		// Terminal colors
		TerminalBlack:       colors.MustGet("black"),
		TerminalRed:         colors.MustGet("red500"),
		TerminalGreen:       colors.MustGet("green500"),
		TerminalYellow:      colors.MustGet("yellow500"),
		TerminalBlue:        colors.MustGet("blue500"),
		TerminalPurple:      colors.MustGet("purple500"),
		TerminalCyan:        colors.MustGet("blue200"),
		TerminalWhite:       colors.MustGet("gray200"),
		TerminalBrightWhite: colors.MustGet("pureWhite"),
		TerminalDimGreen:    colors.MustGet("green600"),
		TerminalDimYellow:   colors.MustGet("yellow400"),
		TerminalDimBlack:    colors.MustGet("shadow"),

		// Version control
		VersionControlModified: colors.MustGet("yellow400"),
		ConflictBackground:     colors.MustGet("orange400"),

		// Scrollbar
		ScrollbarThumb:       colors.MustGet("neutral600Alpha"),
		ScrollbarThumbHover:  colors.MustGet("blue200Alpha50"),
		ScrollbarThumbActive: colors.MustGet("blue200Alpha80"),
		ScrollbarTrackBorder: colors.MustGet("gray650"),

		// Player/collaboration colors
		Player1Alpha:     colors.MustGet("blue500Alpha24"),
		Player2Alpha:     colors.MustGet("gray700"),
		Player3Alpha:     colors.MustGet("green600Alpha24"),
		Player4Alpha:     colors.MustGet("orange400Alpha24"),

		// Accent colors (from Gruvbox theme) - Tron-themed selection
		Accents: []string{
			colors.MustGet("blue200"),    // Tron blue (primary)
			colors.MustGet("orange500"),  // Tron orange
			colors.MustGet("green300"),   // Tron green
			colors.MustGet("red500"),     // Red accent
			colors.MustGet("purple500"),  // Purple accent
			colors.MustGet("yellow500"),  // Yellow accent
			colors.MustGet("blue500"),    // Darker blue variant
		},
	}
}

// GetFrostedPalette returns the dark frosted glass theme palette
func GetFrostedPalette() palette.TronThemePalette {
	// Load colors from CSS
	colors, err := csscolors.LoadColors(colorsCSS)
	if err != nil {
		log.Fatalf("Failed to load dark colors: %v", err)
	}

	// Start with the regular palette
	p := GetPalette()

	// Override for frosted glass effect
	p.BackgroundAppearance = "blurred"
	p.Background = colors.MustGet("gray900Frosted") // 67% opacity like Glazier

	// Make UI elements transparent
	p.TabBarBackground = colors.MustGet("transparent")
	p.StatusBarBackground = colors.MustGet("gray900Frosted")
	p.Border = colors.MustGet("gray600Alpha67")
	p.BorderSubtle = colors.MustGet("gray600Alpha40")
	p.Selection = colors.MustGet("gray700Alpha40")

	// Adjust text for better contrast on blurred background
	p.Foreground = colors.MustGet("gray200Frosted") // Slightly higher opacity
	p.ForegroundDim = colors.MustGet("gray500Frosted")

	// Keep elevated surfaces opaque for readability
	p.Highlight = colors.MustGet("gray800")

	// Use semi-opaque backgrounds for title bar (uses StatusBarBackground)
	// Title bar uses p.StatusBarBackground, so it will get gray900Frosted

	// Set editor background with slight transparency
	p.EditorBackground = colors.MustGet("gray900Alpha93")
	p.EditorSubheaderBackground = colors.MustGet("gray800")
	p.TitleBarInactiveBackground = colors.MustGet("gray800Alpha80")

	// Make drop target more visible on frosted background
	p.DropTargetBackground = colors.MustGet("blue200Alpha20")

	// Adjust other UI elements for frosted effect
	p.ElementHover = colors.MustGet("gray700")

	// Editor specific adjustments
	p.ActiveLineBackground = colors.MustGet("gray800Alpha33")
	p.DocumentHighlightRead = colors.MustGet("blue200Alpha13")
	p.DocumentHighlightWrite = colors.MustGet("blue200Alpha27")

	// Overlay backgrounds with transparency
	p.PanelOverlayBackground = colors.MustGet("gray800Alpha80")
	p.PanelOverlayHover = colors.MustGet("gray750Alpha80")

	// Keep the same accent colors for frosted variant
	// (Accents are already set from base palette)

	return p
}
