package dark

import "github.com/bcomnes/zed-theme-tron-legacy/tools/palette"

// GetPalette returns the dark theme palette
func GetPalette() palette.TronThemePalette {
	return palette.TronThemePalette{
		// Core colors
		Background:       gray900,
		Foreground:       gray200,
		ForegroundDim:    gray500,
		ForegroundBright: gray50,

		// UI elements
		Border:               neutral600,
		BorderSubtle:         gray700,
		BorderFocused:        blue200,
		Selection:            gray700,
		Highlight:            neutral800,
		ElementActive:        blue400,
		ElementHover:         gray700,
		Shadow:               shadow,
		Transparent:          transparent,
		DropTargetBackground: dropTarget,
		TabBarBackground:     gray800,
		StatusBarBackground:  gray750,
		EditorBackground:     gray900,
		EditorSubheaderBackground: gray800,
		TitleBarInactiveBackground: gray800,

		// Semantic colors
		Error:   red500,
		Warning: yellow500,
		Success: green300,
		Info:    blue200,

		// Background variants
		ErrorBg:   red700,
		SuccessBg: green700,

		// Syntax highlighting - base colors
		Comment:  gray400,
		String:   red400,
		Number:   green300,
		Keyword:  blue500,
		Function: orange500,
		Type:     blue500,
		Variable: blue200Bright,
		Property: green500,

		// Syntax highlighting - accent colors
		StringEscape:       red300,
		Regex:              blue200,
		Decorator:          pink500,
		Punctuation:        gray200,
		PunctuationSpecial: gray500,
		Link:               blue300,

		// Terminal colors
		TerminalBlack:  black,
		TerminalRed:    red500,
		TerminalGreen:  green500,
		TerminalYellow: yellow500,
		TerminalBlue:   blue500,
		TerminalPurple: purple500,
		TerminalCyan:   blue200,
		TerminalWhite:  gray200,

		// Additional semantic mappings
		LineNumber:             gray500,
		ElementActiveBright:    blue400,
		ConflictBackground:     orange400,
		TerminalBrightWhite:    pureWhite,
		TerminalDimGreen:       green600,
		TerminalDimYellow:      yellow400,
		VersionControlModified: yellow400,
		VariableParam:          green500,
		Constructor:            orange400,
		ClassType:              orange400,
		EnumType:               orange400,
		AttributeType:          orange400,
		SelectorAlt:            green500,
		AnnotationType:         yellow400,
		EmbeddedType:           yellow400,

		// Search
		SearchHighlight: blue200Alpha,

		// Player selection colors
		Player1Selection: blue500Alpha,
		Player3Selection: green600,
		Player4Selection: orange400,

		// Scrollbar alpha variants
		ScrollbarThumb:       neutral600Alpha,
		ScrollbarThumbHover:  blue200Alpha50,
		ScrollbarTrackBorder: oneTrackBorder,

		// Editor transparency variants
		ActiveLineBackground:   gray800Alpha75,
		WrapGuide:              gray500Alpha30,
		ActiveWrapGuide:        gray400Alpha50,
		DocumentHighlightRead:  blue200Alpha10,
		DocumentHighlightWrite: blue200Alpha40,

		// Standardized player selection alphas
		Player1Alpha: player1Alpha,
		Player3Alpha: player3Alpha,
		Player4Alpha: player4Alpha,
	}
}

// GetFrostedPalette returns the dark frosted glass theme palette
func GetFrostedPalette() palette.TronThemePalette {
	// Start with the regular palette
	p := GetPalette()

	// Override for frosted glass effect
	p.BackgroundAppearance = "blurred"
	p.Background = gray900Frosted // 67% opacity like Glazier
	p.BackgroundFrosted = gray900Frosted

	// Make UI elements transparent
	p.TabBarBackground = transparent
	p.StatusBarBackground = statusBarFrosted
	p.Border = borderFrosted
	p.BorderSubtle = borderSubtleFrosted
	p.Selection = selectionFrosted

	// Adjust text for better contrast on blurred background
	p.Foreground = gray200Frosted // Slightly higher opacity
	p.ForegroundDim = gray500Frosted
	p.TextFrosted = gray200Frosted

	// Keep elevated surfaces opaque for readability
	p.Highlight = elevatedFrosted

	// Use semi-opaque backgrounds for title bar (uses StatusBarBackground)
	// Title bar uses p.StatusBarBackground, so it will get statusBarFrosted

	// Set editor background with slight transparency
	p.EditorBackground = editorFrosted
	p.EditorSubheaderBackground = gray800
	p.TitleBarInactiveBackground = titleBarInactiveFrosted

	// Make drop target more visible on frosted background
	p.DropTargetBackground = dropTargetFrosted

	// Adjust other UI elements for frosted effect
	p.ElementHover = gray700

	// Editor specific adjustments
	p.ActiveLineBackground = activeLineFrosted
	p.DocumentHighlightRead = docHighlightReadFrosted
	p.DocumentHighlightWrite = docHighlightWriteFrosted

	return p
}
