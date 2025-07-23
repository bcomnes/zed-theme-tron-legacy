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
		// Core text colors
		Foreground:       colors.MustGet("gray700"),
		ForegroundDim:    colors.MustGet("gray600"),
		ForegroundBright: colors.MustGet("gray900"),

		// Core backgrounds
		Background:       colors.MustGet("gray50"),
		EditorBackground: colors.MustGet("gray50"),

		// UI backgrounds
		TabBarBackground:           colors.MustGet("gray100"),
		StatusBarBackground:        colors.MustGet("gray150"),
		EditorSubheaderBackground:  colors.MustGet("gray100"),
		TitleBarInactiveBackground: colors.MustGet("gray100"),
		PanelOverlayBackground:     colors.MustGet("gray125"),
		PanelOverlayHover:          colors.MustGet("gray200"),

		// UI elements
		Border:               colors.MustGet("gray300"),
		BorderSubtle:         colors.MustGet("gray200"),
		BorderFocused:        colors.MustGet("blue200"),
		Selection:            colors.MustGet("gray200"),
		Highlight:            colors.MustGet("gray100"),
		ElementHover:         colors.MustGet("gray200"),
		Transparent:          colors.MustGet("transparent"),
		DropTargetBackground: colors.MustGet("blue200Alpha18"),

		// Semantic status colors
		Error:     colors.MustGet("red600"),
		ErrorBg:   colors.MustGet("red100"),
		Warning:   colors.MustGet("yellow600"),
		Success:   colors.MustGet("green400"),
		SuccessBg: colors.MustGet("green100"),
		Info:      colors.MustGet("blue200"),
		Hint:      colors.MustGet("gray400"),

		// Editor specific
		LineNumber:             colors.MustGet("gray400"),
		ActiveLineBackground:   colors.MustGet("gray100Alpha75"),
		WrapGuide:              colors.MustGet("gray500Alpha30"),
		ActiveWrapGuide:        colors.MustGet("gray600Alpha50"),
		DocumentHighlightRead:  colors.MustGet("blue200Alpha10"),
		DocumentHighlightWrite: colors.MustGet("blue200Alpha40"),
		SearchHighlight:        colors.MustGet("blue200Alpha"),

		// Syntax - basic types
		Comment:       colors.MustGet("gray500"),
		String:        colors.MustGet("red500"),
		StringEscape:  colors.MustGet("red400"),
		Number:        colors.MustGet("green400"),
		Keyword:       colors.MustGet("blue600"),
		Function:      colors.MustGet("orange600"),
		Variable:      colors.MustGet("blue500"),
		Property:      colors.MustGet("green500"),
		Namespace:     colors.MustGet("blue400"),

		// Syntax - advanced types
		Type:          colors.MustGet("blue600"),
		Constructor:   colors.MustGet("orange500"),
		EnumType:      colors.MustGet("orange500"),
		AttributeType: colors.MustGet("orange500"),
		EmbeddedType:  colors.MustGet("yellow500"),
		SpecialSyntax: colors.MustGet("pink600"),
		Regex:         colors.MustGet("blue200"),
		Selector:      colors.MustGet("green500"),

		// Syntax - punctuation
		Punctuation:        colors.MustGet("gray700"),
		PunctuationSpecial: colors.MustGet("gray400"),

		// Terminal colors
		TerminalBlack:       colors.MustGet("black"),
		TerminalRed:         colors.MustGet("red600"),
		TerminalGreen:       colors.MustGet("green500"),
		TerminalYellow:      colors.MustGet("yellow600"),
		TerminalBlue:        colors.MustGet("blue600"),
		TerminalPurple:      colors.MustGet("purple600"),
		TerminalCyan:        colors.MustGet("blue200"),
		TerminalWhite:       colors.MustGet("gray700"),
		TerminalBrightWhite: colors.MustGet("gray900"),
		TerminalDimGreen:    colors.MustGet("green600"),
		TerminalDimYellow:   colors.MustGet("yellow500"),
		TerminalDimBlack:    colors.MustGet("shadow"),

		// Version control
		VersionControlModified: colors.MustGet("orange700"),
		ConflictBackground:     colors.MustGet("orange500"),

		// Scrollbar
		ScrollbarThumb:       colors.MustGet("gray300Alpha"),
		ScrollbarThumbHover:  colors.MustGet("blue400Alpha"),
		ScrollbarThumbActive: colors.MustGet("blue400Alpha80"),
		ScrollbarTrackBorder: colors.MustGet("gray200"),

		// Player/collaboration colors
		Player1Alpha: colors.MustGet("blue300Alpha24"),
		Player2Alpha: colors.MustGet("gray700"),
		Player3Alpha: colors.MustGet("green600Alpha24"),
		Player4Alpha: colors.MustGet("orange500Alpha24"),

		// Accent colors (from Gruvbox theme) - Tron-themed selection for light theme
		Accents: []string{
			colors.MustGet("blue300"),    // Tron blue (primary)
			colors.MustGet("orange500"),  // Tron orange
			colors.MustGet("green600"),   // Tron green
			colors.MustGet("red500"),     // Red accent
			colors.MustGet("purple600"),  // Purple accent
			colors.MustGet("yellow500"),  // Yellow accent
			colors.MustGet("blue500"),    // Darker blue variant
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
	p.TabBarBackground = colors.MustGet("transparent")
	p.StatusBarBackground = colors.MustGet("gray100Alpha87")
	p.Border = colors.MustGet("gray300Alpha60")
	p.BorderSubtle = colors.MustGet("transparent")
	p.Selection = colors.MustGet("gray200Alpha40")

	// Adjust text for better contrast on blurred background
	p.Foreground = colors.MustGet("gray700Frosted") // Slightly higher opacity
	p.ForegroundDim = colors.MustGet("gray600Frosted")

	// Keep elevated surfaces opaque for readability
	p.Highlight = colors.MustGet("gray100")

	// Set editor background with slight transparency
	p.EditorBackground = colors.MustGet("gray50Alpha95")
	p.EditorSubheaderBackground = colors.MustGet("gray100")
	p.TitleBarInactiveBackground = colors.MustGet("gray100Alpha80")

	// Use darker orange for better contrast on modified files
	p.VersionControlModified = colors.MustGet("orange700")

	// Use transparent track border for frosted to avoid visual clutter
	// p.ScrollbarTrackBorder = transparent

	// Make drop target more visible on frosted background
	p.DropTargetBackground = colors.MustGet("blue200Alpha20")

	// Adjust other UI elements for frosted effect
	p.ElementHover = colors.MustGet("gray200")

	// Editor specific adjustments
	p.ActiveLineBackground = colors.MustGet("gray100Alpha33")
	p.DocumentHighlightRead = colors.MustGet("blue200Alpha13")
	p.DocumentHighlightWrite = colors.MustGet("blue200Alpha27")

	// Use darker scrollbar thumb for better visibility on frosted
	p.ScrollbarThumb = colors.MustGet("gray300AlphaDark")

	// Overlay backgrounds with transparency
	p.PanelOverlayBackground = colors.MustGet("gray100Alpha80")
	p.PanelOverlayHover = colors.MustGet("gray150Alpha80")

	// Keep the same accent colors for frosted variant
	// (Accents are already set from base palette)

	return p
}
