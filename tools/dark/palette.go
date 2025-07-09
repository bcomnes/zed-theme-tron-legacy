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
		// Core colors
		Background:       colors.MustGet("gray900"),
		Foreground:       colors.MustGet("gray200"),
		ForegroundDim:    colors.MustGet("gray500"),
		ForegroundBright: colors.MustGet("gray50"),

		// UI elements
		Border:               colors.MustGet("neutral600"),
		BorderSubtle:         colors.MustGet("gray700"),
		BorderFocused:        colors.MustGet("blue200"),
		Selection:            colors.MustGet("gray700"),
		Highlight:            colors.MustGet("neutral800"),
		ElementActive:        colors.MustGet("blue400"),
		ElementHover:         colors.MustGet("gray700"),
		Shadow:               colors.MustGet("shadow"),
		Transparent:          colors.MustGet("transparent"),
		DropTargetBackground: colors.MustGet("dropTarget"),
		TabBarBackground:     colors.MustGet("gray800"),
		StatusBarBackground:  colors.MustGet("gray750"),
		EditorBackground:     colors.MustGet("gray900"),
		EditorSubheaderBackground: colors.MustGet("gray800"),
		TitleBarInactiveBackground: colors.MustGet("gray800"),

		// Semantic colors
		Error:   colors.MustGet("red500"),
		Warning: colors.MustGet("yellow500"),
		Success: colors.MustGet("green300"),
		Info:    colors.MustGet("blue200"),
		Hint:    colors.MustGet("gray500"),

		// Background variants
		ErrorBg:   colors.MustGet("red700"),
		SuccessBg: colors.MustGet("green700"),

		// Syntax highlighting - base colors
		Comment:  colors.MustGet("gray400"),
		String:   colors.MustGet("red400"),
		Number:   colors.MustGet("green300"),
		Keyword:  colors.MustGet("blue500"),
		Function: colors.MustGet("orange500"),
		Type:     colors.MustGet("blue500"),
		Variable: colors.MustGet("blue200Bright"),
		Property: colors.MustGet("green500"),

		// Syntax highlighting - accent colors
		StringEscape:       colors.MustGet("red300"),
		Regex:              colors.MustGet("blue200"),
		Decorator:          colors.MustGet("pink500"),
		Punctuation:        colors.MustGet("gray200"),
		PunctuationSpecial: colors.MustGet("gray500"),
		Link:               colors.MustGet("blue300"),

		// Terminal colors
		TerminalBlack:  colors.MustGet("black"),
		TerminalRed:    colors.MustGet("red500"),
		TerminalGreen:  colors.MustGet("green500"),
		TerminalYellow: colors.MustGet("yellow500"),
		TerminalBlue:   colors.MustGet("blue500"),
		TerminalPurple: colors.MustGet("purple500"),
		TerminalCyan:   colors.MustGet("blue200"),
		TerminalWhite:  colors.MustGet("gray200"),

		// Additional semantic mappings
		LineNumber:             colors.MustGet("gray500"),
		ElementActiveBright:    colors.MustGet("blue400"),
		ConflictBackground:     colors.MustGet("orange400"),
		TerminalBrightWhite:    colors.MustGet("pureWhite"),
		TerminalDimGreen:       colors.MustGet("green600"),
		TerminalDimYellow:      colors.MustGet("yellow400"),
		VersionControlModified: colors.MustGet("yellow400"),
		VariableParam:          colors.MustGet("green500"),
		Constructor:            colors.MustGet("orange400"),
		ClassType:              colors.MustGet("orange400"),
		EnumType:               colors.MustGet("orange400"),
		AttributeType:          colors.MustGet("orange400"),
		SelectorAlt:            colors.MustGet("green500"),
		AnnotationType:         colors.MustGet("yellow400"),
		EmbeddedType:           colors.MustGet("yellow400"),

		// Search
		SearchHighlight: colors.MustGet("blue200Alpha"),

		// Player selection colors
		Player1Selection: colors.MustGet("blue500Alpha"),
		Player3Selection: colors.MustGet("green600"),
		Player4Selection: colors.MustGet("orange400"),

		// Scrollbar alpha variants
		ScrollbarThumb:       colors.MustGet("neutral600Alpha"),
		ScrollbarThumbHover:  colors.MustGet("blue200Alpha50"),
		ScrollbarTrackBorder: colors.MustGet("oneTrackBorder"),

		// Editor transparency variants
		ActiveLineBackground:   colors.MustGet("gray800Alpha75"),
		WrapGuide:              colors.MustGet("gray500Alpha30"),
		ActiveWrapGuide:        colors.MustGet("gray400Alpha50"),
		DocumentHighlightRead:  colors.MustGet("blue200Alpha10"),
		DocumentHighlightWrite: colors.MustGet("blue200Alpha40"),

		// Standardized player selection alphas
		Player1Alpha: colors.MustGet("player1Alpha"),
		Player3Alpha: colors.MustGet("player3Alpha"),
		Player4Alpha: colors.MustGet("player4Alpha"),
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
	p.BackgroundFrosted = colors.MustGet("gray900Frosted")

	// Make UI elements transparent
	p.TabBarBackground = colors.MustGet("transparent")
	p.StatusBarBackground = colors.MustGet("statusBarFrosted")
	p.Border = colors.MustGet("borderFrosted")
	p.BorderSubtle = colors.MustGet("borderSubtleFrosted")
	p.Selection = colors.MustGet("selectionFrosted")

	// Adjust text for better contrast on blurred background
	p.Foreground = colors.MustGet("gray200Frosted") // Slightly higher opacity
	p.ForegroundDim = colors.MustGet("gray500Frosted")
	p.TextFrosted = colors.MustGet("gray200Frosted")

	// Keep elevated surfaces opaque for readability
	p.Highlight = colors.MustGet("elevatedFrosted")

	// Use semi-opaque backgrounds for title bar (uses StatusBarBackground)
	// Title bar uses p.StatusBarBackground, so it will get statusBarFrosted

	// Set editor background with slight transparency
	p.EditorBackground = colors.MustGet("editorFrosted")
	p.EditorSubheaderBackground = colors.MustGet("gray800")
	p.TitleBarInactiveBackground = colors.MustGet("titleBarInactiveFrosted")

	// Make drop target more visible on frosted background
	p.DropTargetBackground = colors.MustGet("dropTargetFrosted")

	// Adjust other UI elements for frosted effect
	p.ElementHover = colors.MustGet("gray700")

	// Editor specific adjustments
	p.ActiveLineBackground = colors.MustGet("activeLineFrosted")
	p.DocumentHighlightRead = colors.MustGet("docHighlightReadFrosted")
	p.DocumentHighlightWrite = colors.MustGet("docHighlightWriteFrosted")

	return p
}
