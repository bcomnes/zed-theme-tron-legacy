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
		// Core colors
		Background:       colors.MustGet("gray50"),
		Foreground:       colors.MustGet("gray700"),
		ForegroundDim:    colors.MustGet("gray600"),
		ForegroundBright: colors.MustGet("gray900"),

		// UI elements
		Border:               colors.MustGet("gray300"),
		BorderSubtle:         colors.MustGet("gray200"),
		BorderFocused:        colors.MustGet("blue200"),
		Selection:            colors.MustGet("gray200"),
		Highlight:            colors.MustGet("gray100"),
		ElementActive:        colors.MustGet("blue500"),
		ElementHover:         colors.MustGet("gray200"),
		Shadow:               colors.MustGet("shadow"),
		Transparent:          colors.MustGet("transparent"),
		DropTargetBackground: colors.MustGet("dropTarget"),
		TabBarBackground:     colors.MustGet("gray100"),
		StatusBarBackground:  colors.MustGet("gray150"),
		EditorBackground:     colors.MustGet("gray50"),
		EditorSubheaderBackground: colors.MustGet("gray100"),
		TitleBarInactiveBackground: colors.MustGet("gray100"),

		// Semantic colors
		Error:   colors.MustGet("red600"),
		Warning: colors.MustGet("yellow600"),
		Success: colors.MustGet("green400"),
		Info:    colors.MustGet("blue200"),
		Hint:    colors.MustGet("gray400"),

		// Background variants
		ErrorBg:   colors.MustGet("red100"),
		SuccessBg: colors.MustGet("green100"),

		// Syntax highlighting - base colors
		Comment:  colors.MustGet("gray500"),
		String:   colors.MustGet("red500"),
		Number:   colors.MustGet("green400"),
		Keyword:  colors.MustGet("blue600"),
		Function: colors.MustGet("orange600"),
		Type:     colors.MustGet("blue600"),
		Variable: colors.MustGet("blue500"),
		Property: colors.MustGet("green500"),

		// Syntax highlighting - accent colors
		StringEscape:       colors.MustGet("red400"),
		Regex:              colors.MustGet("blue200"),
		Decorator:          colors.MustGet("pink600"),
		Punctuation:        colors.MustGet("gray700"),
		PunctuationSpecial: colors.MustGet("gray400"),
		Link:               colors.MustGet("blue400"),

		// Terminal colors
		TerminalBlack:  colors.MustGet("black"),
		TerminalRed:    colors.MustGet("red600"),
		TerminalGreen:  colors.MustGet("green500"),
		TerminalYellow: colors.MustGet("yellow600"),
		TerminalBlue:   colors.MustGet("blue600"),
		TerminalPurple: colors.MustGet("purple600"),
		TerminalCyan:   colors.MustGet("blue200"),
		TerminalWhite:  colors.MustGet("gray700"),

		// Additional semantic mappings
		LineNumber:             colors.MustGet("gray400"),
		ElementActiveBright:    colors.MustGet("blue500"),
		ConflictBackground:     colors.MustGet("orange500"),
		TerminalBrightWhite:    colors.MustGet("gray900"),
		TerminalDimGreen:       colors.MustGet("green600"),
		TerminalDimYellow:      colors.MustGet("yellow500"),
		VersionControlModified: colors.MustGet("orange700"),
		VariableParam:          colors.MustGet("green500"),
		Constructor:            colors.MustGet("orange500"),
		ClassType:              colors.MustGet("orange500"),
		EnumType:               colors.MustGet("orange500"),
		AttributeType:          colors.MustGet("orange500"),
		SelectorAlt:            colors.MustGet("green500"),
		AnnotationType:         colors.MustGet("yellow500"),
		EmbeddedType:           colors.MustGet("yellow500"),

		// Search
		SearchHighlight: colors.MustGet("blue200Alpha"),

		// Player selection colors
		Player1Selection: colors.MustGet("blue300Alpha"),
		Player3Selection: colors.MustGet("green600"),
		Player4Selection: colors.MustGet("orange500"),

		// Scrollbar alpha variants
		ScrollbarThumb:       colors.MustGet("gray300Alpha"),
		ScrollbarThumbHover:  colors.MustGet("blue400Alpha"),
		ScrollbarTrackBorder: colors.MustGet("gray200"),

		// Editor transparency variants
		ActiveLineBackground:   colors.MustGet("gray100Alpha75"),
		WrapGuide:              colors.MustGet("gray500Alpha30"),
		ActiveWrapGuide:        colors.MustGet("gray600Alpha50"),
		DocumentHighlightRead:  colors.MustGet("blue200Alpha10"),
		DocumentHighlightWrite: colors.MustGet("blue200Alpha40"),

		// Standardized player selection alphas
		Player1Alpha: colors.MustGet("player1Alpha"),
		Player3Alpha: colors.MustGet("player3Alpha"),
		Player4Alpha: colors.MustGet("player4Alpha"),
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
	p.BackgroundFrosted = colors.MustGet("gray50Frosted")

	// Make UI elements transparent
	p.TabBarBackground = colors.MustGet("transparent")
	p.StatusBarBackground = colors.MustGet("statusBarFrosted")
	p.Border = colors.MustGet("borderFrosted")
	p.BorderSubtle = colors.MustGet("transparent")
	p.Selection = colors.MustGet("selectionFrosted")

	// Adjust text for better contrast on blurred background
	p.Foreground = colors.MustGet("gray700Frosted") // Slightly higher opacity
	p.ForegroundDim = colors.MustGet("gray600Frosted")
	p.TextFrosted = colors.MustGet("gray700Frosted")

	// Keep elevated surfaces opaque for readability
	p.Highlight = colors.MustGet("elevatedFrosted")

	// Set editor background with slight transparency
	p.EditorBackground = colors.MustGet("editorFrosted")
	p.EditorSubheaderBackground = colors.MustGet("gray100")
	p.TitleBarInactiveBackground = colors.MustGet("titleBarInactiveFrosted")

	// Use darker orange for better contrast on modified files
	p.VersionControlModified = colors.MustGet("orange700")

	// Use transparent track border for frosted to avoid visual clutter
	// p.ScrollbarTrackBorder = transparent

	// Make drop target more visible on frosted background
	p.DropTargetBackground = colors.MustGet("dropTargetFrosted")

	// Adjust other UI elements for frosted effect
	p.ElementHover = colors.MustGet("gray200")

	// Editor specific adjustments
	p.ActiveLineBackground = colors.MustGet("activeLineFrosted")
	p.DocumentHighlightRead = colors.MustGet("docHighlightReadFrosted")
	p.DocumentHighlightWrite = colors.MustGet("docHighlightWriteFrosted")

	// Use darker scrollbar thumb for better visibility on frosted
	p.ScrollbarThumb = colors.MustGet("gray300AlphaDark")

	return p
}
