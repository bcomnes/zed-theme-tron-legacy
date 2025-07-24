package light

import (
	_ "embed"
	"log"

	"github.com/bcomnes/zed-theme-tron-legacy/tools/csscolors"
	"github.com/bcomnes/zed-theme-tron-legacy/tools/palette"
)

//go:embed colors.css
var colorsCSS []byte

// GetPalette returns the light theme palette
func GetPalette() palette.TronThemePalette {
	// Load colors from CSS
	colors, err := csscolors.LoadColors(colorsCSS)
	if err != nil {
		log.Fatalf("Failed to load light colors: %v", err)
	}

	return palette.TronThemePalette{
		// Visual Hierarchy Layers
		Background:             colors.MustGet("gray50"),
		EditorBackground:       colors.MustGet("gray50"),
		BackgroundElevated:     colors.MustGet("gray100"),
		BackgroundOverlay:      colors.MustGet("gray125"),
		BackgroundOverlayHover: colors.MustGet("gray200"),

		// Midground Layer
		Surface:           colors.MustGet("gray100"),
		SurfaceHighlight:  colors.MustGet("gray50"),
		EditorSubheader:   colors.MustGet("gray100"),
		Statusbar:         colors.MustGet("gray150"),
		StatusbarInactive: colors.MustGet("gray100"),

		// Foreground Layer
		Foreground:       colors.MustGet("gray700"),
		ForegroundMuted:  colors.MustGet("gray600"),
		ForegroundStrong: colors.MustGet("gray900"),

		// Interactive Elements
		Border:        colors.MustGet("gray300"),
		BorderSubtle:  colors.MustGet("gray200"),
		BorderFocused: colors.MustGet("blue200"),

		// Selection & Highlights
		Selection:              colors.MustGet("gray200"),
		// SelectionAlpha:         colors.MustGet("gray200Alpha40"), // element.selection_background is broken in Zed
		ActiveLine:             colors.MustGet("gray100Alpha75"),
		MatchHighlight:         colors.MustGet("blue200Alpha"),
		DocumentHighlight:      colors.MustGet("blue200Alpha10"),
		DocumentHighlightWrite: colors.MustGet("blue200Alpha40"),

		// Interactive States
		Interactive: colors.MustGet("gray200"),
		DropTarget:  colors.MustGet("blue200Alpha18"),
		Transparent: colors.MustGet("transparent"),

		// Semantic Colors
		Error:          colors.MustGet("red600"),
		ErrorSurface:   colors.MustGet("red100"),
		Warning:        colors.MustGet("yellow600"),
		Success:        colors.MustGet("green400"),
		SuccessSurface: colors.MustGet("green100"),
		Info:           colors.MustGet("blue200"),
		Hint:           colors.MustGet("gray400"),
		Accent:         colors.MustGet("orange600"),
		UIAccent:       colors.MustGet("green400"),  // Use primary green for UI accents

		// Editor Guidelines
		GuideNormal: colors.MustGet("gray500Alpha30"),
		GuideActive: colors.MustGet("gray600Alpha50"),
		LineNumber:  colors.MustGet("gray400"),

		// Syntax Highlighting
		Comment:      colors.MustGet("gray500"),
		String:       colors.MustGet("red500"),
		StringEscape: colors.MustGet("red400"),
		Number:       colors.MustGet("green400"),
		Keyword:      colors.MustGet("blue600"),
		Function:     colors.MustGet("orange600"),
		Variable:     colors.MustGet("blue500"),
		Type:         colors.MustGet("blue600"),
		Property:     colors.MustGet("green500"),
		Namespace:    colors.MustGet("blue400"),

		// Special Syntax Elements
		Constructor: colors.MustGet("orange500"),
		Enum:        colors.MustGet("orange500"),
		Attribute:   colors.MustGet("orange500"),
		Embedded:    colors.MustGet("yellow500"),
		Decorator:   colors.MustGet("pink600"),
		Regex:       colors.MustGet("blue200"),
		Tag:         colors.MustGet("blue600"),
		SpecialVariable: colors.MustGet("purple600"),

		// Punctuation
		Punctuation:      colors.MustGet("gray700"),
		PunctuationMuted: colors.MustGet("gray400"),

		// Terminal Colors
		TerminalBlack:        colors.MustGet("black"),          // 0 - Black
		TerminalRed:          colors.MustGet("red500"),         // 1 - Red (string color)
		TerminalGreen:        colors.MustGet("green400"),       // 2 - Green (number color)
		TerminalYellow:       colors.MustGet("yellow500"),      // 3 - Yellow (embedded color, dimmer)
		TerminalBlue:         colors.MustGet("blue600"),        // 4 - Blue (keyword color)
		TerminalPurple:       colors.MustGet("pink600"),        // 5 - Magenta (decorator color)
		TerminalCyan:         colors.MustGet("blue200"),        // 6 - Cyan (signature Tron cyan)
		TerminalWhite:        colors.MustGet("gray800"),        // 7 - White (terminal white)
		TerminalBrightBlack:  colors.MustGet("gray600"),        // 8 - Bright Black
		TerminalBrightRed:    colors.MustGet("red400"),         // 9 - Bright Red (string escape color)
		TerminalBrightGreen:  colors.MustGet("green500"),       // 10 - Bright Green (property color)
		TerminalBrightYellow: colors.MustGet("yellow600"),      // 11 - Bright Yellow (warning color)
		TerminalBrightBlue:   colors.MustGet("blue500"),        // 12 - Bright Blue (variable color)
		TerminalBrightPurple: colors.MustGet("pink500"),        // 13 - Bright Magenta
		TerminalBrightCyan:   colors.MustGet("blue400"),        // 14 - Bright Cyan (namespace color)
		TerminalBrightWhite:  colors.MustGet("white"),          // 15 - Bright White
		TerminalDimGreen:     colors.MustGet("green600"),
		TerminalDimYellow:    colors.MustGet("yellow500"),
		TerminalDimBlack:     colors.MustGet("shadow"),
		TerminalDimRed:       colors.MustGet("red500"),     // Same as regular red
		TerminalDimBlue:      colors.MustGet("gray300"),    // Use border color
		TerminalDimCyan:      colors.MustGet("green500"),   // Use property color
		TerminalDimWhite:     colors.MustGet("blue600"),    // Use keyword color
		TerminalDimMagenta:   colors.MustGet("pink600"),    // Same as regular magenta

		// Version Control
		VCSModified: colors.MustGet("orange700"),
		VCSConflict: colors.MustGet("orange500"),

		// UI Components
		ScrollbarThumb:       colors.MustGet("gray300Alpha"),
		ScrollbarThumbHover:  colors.MustGet("blue400Alpha"),
		ScrollbarThumbActive: colors.MustGet("blue400Alpha80"),
		ScrollbarTrackBorder: colors.MustGet("gray200"),

		// Collaboration/Players
		Player1: colors.MustGet("blue300Alpha24"),
		Player2: colors.MustGet("gray700"),
		Player3: colors.MustGet("green600Alpha24"),
		Player4: colors.MustGet("orange500Alpha24"),

		// Theme Properties
		BackgroundAppearance: "opaque",
		Accents: []string{
			colors.MustGet("blue200"),   // Tron blue (primary cyan)
			colors.MustGet("orange500"), // Tron orange
			colors.MustGet("green400"),  // Tron green
			colors.MustGet("red500"),    // Red accent
			colors.MustGet("pink600"),   // Pink accent (replaced purple)
			colors.MustGet("yellow500"), // Yellow accent
			colors.MustGet("blue500"),   // Darker blue variant
		},
	}
}

// GetFrostedPalette returns the light frosted glass theme palette
func GetFrostedPalette() palette.TronThemePalette {
	// Load colors from CSS
	colors, err := csscolors.LoadColors(colorsCSS)
	if err != nil {
		log.Fatalf("Failed to load light colors: %v", err)
	}

	// Start with the regular palette
	p := GetPalette()

	// Override for frosted glass effect
	p.BackgroundAppearance = "blurred"
	p.Background = colors.MustGet("gray50Frosted") // 67% opacity like Glazier

	// Make UI elements transparent
	p.Surface = colors.MustGet("transparent")
	p.Statusbar = colors.MustGet("gray100Alpha87")
	p.Border = colors.MustGet("gray300Alpha60")
	p.BorderSubtle = colors.MustGet("transparent")
	p.Selection = colors.MustGet("gray200Alpha40")
	// p.SelectionAlpha = colors.MustGet("gray300AlphaDark") // element.selection_background is broken in Zed

	// Adjust text for better contrast on blurred background
	p.Foreground = colors.MustGet("gray700Frosted")   // Slightly higher opacity
	p.ForegroundMuted = colors.MustGet("gray600Frosted")

	// Keep elevated surfaces opaque for readability
	p.BackgroundElevated = colors.MustGet("gray100")

	// Set editor background with slight transparency
	p.EditorBackground = colors.MustGet("gray50Alpha95")
	p.EditorSubheader = colors.MustGet("gray100")
	p.StatusbarInactive = colors.MustGet("gray100Alpha80")

	// Use darker orange for better contrast on modified files
	p.VCSModified = colors.MustGet("orange700")

	// Make drop target more visible on frosted background
	p.DropTarget = colors.MustGet("blue200Alpha20")

	// Adjust other UI elements for frosted effect
	p.Interactive = colors.MustGet("gray200")

	// Editor specific adjustments
	p.ActiveLine = colors.MustGet("gray100Alpha33")
	p.DocumentHighlight = colors.MustGet("blue200Alpha13")
	p.DocumentHighlightWrite = colors.MustGet("blue200Alpha27")

	// Use darker scrollbar thumb for better visibility on frosted
	p.ScrollbarThumb = colors.MustGet("gray300AlphaDark")

	// Override terminal dim blue to use transparent border color
	p.TerminalDimBlue = colors.MustGet("gray300Alpha60")

	// Overlay backgrounds with transparency
	p.BackgroundOverlay = colors.MustGet("gray100Alpha80")
	p.BackgroundOverlayHover = colors.MustGet("gray150Alpha80")

	// Keep the same accent colors for frosted variant
	// (Accents are already set from base palette)

	return p
}
