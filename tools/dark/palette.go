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
		// Visual Hierarchy Layers
		Background:             colors.MustGet("gray900"),
		EditorBackground:       colors.MustGet("gray900"),
		BackgroundElevated:     colors.MustGet("neutral800"),
		BackgroundOverlay:      colors.MustGet("gray850"),
		BackgroundOverlayHover: colors.MustGet("gray700"),

		// Midground Layer
		Surface:           colors.MustGet("gray800"),
		SurfaceHighlight:  colors.MustGet("gray900"),
		EditorSubheader:   colors.MustGet("gray800"),
		Statusbar:         colors.MustGet("gray750"),
		StatusbarInactive: colors.MustGet("gray800"),

		// Foreground Layer
		Foreground:       colors.MustGet("gray200"),
		ForegroundMuted:  colors.MustGet("gray500"),
		ForegroundStrong: colors.MustGet("gray50"),

		// Interactive Elements
		Border:        colors.MustGet("neutral600"),
		BorderSubtle:  colors.MustGet("gray700"),
		BorderFocused: colors.MustGet("blue200"),

		// Selection & Highlights
		Selection:              colors.MustGet("gray700"),
		// SelectionAlpha:         colors.MustGet("gray700Alpha40"), // element.selection_background is broken in Zed
		ActiveLine:             colors.MustGet("gray800Alpha75"),
		MatchHighlight:         colors.MustGet("neonOrangeAlpha25"),
		DocumentHighlight:      colors.MustGet("blue200Alpha10"),
		DocumentHighlightWrite: colors.MustGet("blue200Alpha40"),

		// Interactive States
		Interactive: colors.MustGet("gray700"),
		DropTarget:  colors.MustGet("blue200Alpha18"),
		Transparent: colors.MustGet("transparent"),

		// Semantic Colors
		Error:          colors.MustGet("red500"),
		ErrorSurface:   colors.MustGet("red700"),
		Warning:        colors.MustGet("yellow500"),
		Success:        colors.MustGet("green300"),
		SuccessSurface: colors.MustGet("green700"),
		Info:           colors.MustGet("blue200"),
		Hint:           colors.MustGet("gray500"),
		Accent:         colors.MustGet("orange500"),
		UIAccent:       colors.MustGet("green300"),  // Use signature Tron green for UI accents

		// Editor Guidelines
		GuideNormal: colors.MustGet("gray500Alpha30"),
		GuideActive: colors.MustGet("gray400Alpha50"),
		LineNumber:  colors.MustGet("gray500"),

		// Syntax Highlighting
		Comment:      colors.MustGet("gray400"),
		String:       colors.MustGet("red400"),
		StringEscape: colors.MustGet("red300"),
		Number:       colors.MustGet("green300"),
		Keyword:      colors.MustGet("blue500"),
		Function:     colors.MustGet("orange500"),
		Variable:     colors.MustGet("blue200Bright"),
		Type:         colors.MustGet("blue500"),
		Property:     colors.MustGet("green500"),
		Namespace:    colors.MustGet("blue300"),

		// Special Syntax Elements
		Constructor: colors.MustGet("orange400"),
		Enum:        colors.MustGet("orange400"),
		Attribute:   colors.MustGet("orange400"),
		Embedded:    colors.MustGet("yellow400"),
		Decorator:   colors.MustGet("pink500"),
		Regex:       colors.MustGet("blue200"),
		Tag:         colors.MustGet("blue500"),
		SpecialVariable: colors.MustGet("purple500"),

		// Punctuation
		Punctuation:      colors.MustGet("gray200"),
		PunctuationMuted: colors.MustGet("gray500"),

		// Terminal Colors
		TerminalBlack:        colors.MustGet("black"),          // 0 - Black
		TerminalRed:          colors.MustGet("red400"),         // 1 - Red (string color)
		TerminalGreen:        colors.MustGet("green300"),       // 2 - Green (number color)
		TerminalYellow:       colors.MustGet("yellow400"),      // 3 - Yellow (embedded color, dimmer)
		TerminalBlue:         colors.MustGet("blue500"),        // 4 - Blue (keyword color)
		TerminalPurple:       colors.MustGet("pink500"),        // 5 - Magenta (decorator color)
		TerminalCyan:         colors.MustGet("blue200"),        // 6 - Cyan (signature Tron cyan)
		TerminalWhite:        colors.MustGet("gray200"),        // 7 - White (foreground color)
		TerminalBrightBlack:  colors.MustGet("gray450"),        // 8 - Bright Black
		TerminalBrightRed:    colors.MustGet("red300"),         // 9 - Bright Red (string escape color)
		TerminalBrightGreen:  colors.MustGet("green500"),       // 10 - Bright Green (property color)
		TerminalBrightYellow: colors.MustGet("yellow500"),      // 11 - Bright Yellow (warning color)
		TerminalBrightBlue:   colors.MustGet("blue200Bright"),  // 12 - Bright Blue (variable color)
		TerminalBrightPurple: colors.MustGet("pink400"),        // 13 - Bright Magenta
		TerminalBrightCyan:   colors.MustGet("blue300"),        // 14 - Bright Cyan (namespace color)
		TerminalBrightWhite:  colors.MustGet("pureWhite"),      // 15 - Bright White
		TerminalDimGreen:     colors.MustGet("green600"),
		TerminalDimYellow:    colors.MustGet("yellow400"),
		TerminalDimBlack:     colors.MustGet("shadow"),
		TerminalDimRed:       colors.MustGet("red400"),     // Same as regular red
		TerminalDimBlue:      colors.MustGet("neutral600"), // Use border color
		TerminalDimCyan:      colors.MustGet("green500"),   // Use property color
		TerminalDimWhite:     colors.MustGet("blue500"),    // Use keyword color
		TerminalDimMagenta:   colors.MustGet("pink500"),    // Same as regular magenta

		// Version Control
		VCSModified: colors.MustGet("yellow400"),
		VCSConflict: colors.MustGet("orange400"),

		// UI Components
		ScrollbarThumb:       colors.MustGet("neutral600Alpha"),
		ScrollbarThumbHover:  colors.MustGet("blue200Alpha50"),
		ScrollbarThumbActive: colors.MustGet("blue200Alpha80"),
		ScrollbarTrackBorder: colors.MustGet("gray650"),

		// Collaboration/Players
		Player1: colors.MustGet("blue500Alpha24"),
		Player2: colors.MustGet("gray700"),
		Player3: colors.MustGet("green600Alpha24"),
		Player4: colors.MustGet("orange400Alpha24"),

		// Theme Properties
		BackgroundAppearance: "opaque",
		Accents: []string{
			colors.MustGet("blue200"),   // Tron blue (primary)
			colors.MustGet("orange500"), // Tron orange
			colors.MustGet("green300"),  // Tron green
			colors.MustGet("red400"),    // Red accent (using theme red)
			colors.MustGet("pink500"),   // Pink accent (replaced purple)
			colors.MustGet("yellow500"), // Yellow accent
			colors.MustGet("blue500"),   // Darker blue variant
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
	p.Surface = colors.MustGet("transparent")
	p.Statusbar = colors.MustGet("gray900Frosted")
	p.Border = colors.MustGet("gray600Alpha67")
	p.BorderSubtle = colors.MustGet("gray600Alpha40")
	p.Selection = colors.MustGet("gray700Alpha40")
	// p.SelectionAlpha = colors.MustGet("gray600Alpha67") // element.selection_background is broken in Zed

	// Adjust text for better contrast on blurred background
	p.Foreground = colors.MustGet("gray200Frosted")   // Slightly higher opacity
	p.ForegroundMuted = colors.MustGet("gray500Frosted")

	// Keep elevated surfaces opaque for readability
	p.BackgroundElevated = colors.MustGet("gray800")

	// Use semi-opaque backgrounds for title bar (uses Statusbar)
	// Title bar uses p.Statusbar, so it will get gray900Frosted

	// Set editor background with slight transparency
	p.EditorBackground = colors.MustGet("gray900Alpha93")
	p.EditorSubheader = colors.MustGet("gray800")
	p.StatusbarInactive = colors.MustGet("gray800Alpha80")

	// Make drop target more visible on frosted background
	p.DropTarget = colors.MustGet("blue200Alpha20")

	// Adjust other UI elements for frosted effect
	p.Interactive = colors.MustGet("gray700")

	// Editor specific adjustments
	p.ActiveLine = colors.MustGet("gray800Alpha33")
	p.DocumentHighlight = colors.MustGet("blue200Alpha13")
	p.DocumentHighlightWrite = colors.MustGet("blue200Alpha27")

	// Override terminal dim blue to use transparent border color
	p.TerminalDimBlue = colors.MustGet("gray600Alpha67")

	// Overlay backgrounds with transparency
	p.BackgroundOverlay = colors.MustGet("gray800Alpha80")
	p.BackgroundOverlayHover = colors.MustGet("gray750Alpha80")

	// Keep the same accent colors for frosted variant
	// (Accents are already set from base palette)

	return p
}
