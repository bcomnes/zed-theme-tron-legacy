package light

import "github.com/bcomnes/zed-theme-tron-legacy/tools/palette"

// GetPalette returns the light theme palette
func GetPalette() palette.TronThemePalette {
	return palette.TronThemePalette{
		// Core colors
		Background:       gray50,
		Foreground:       gray700,
		ForegroundDim:    gray600,
		ForegroundBright: gray900,

		// UI elements
		Border:               gray300,
		BorderSubtle:         gray200,
		BorderFocused:        blue200,
		Selection:            gray200,
		Highlight:            gray100,
		ElementActive:        blue500,
		ElementHover:         gray200,
		Shadow:               shadow,
		Transparent:          transparent,
		DropTargetBackground: dropTarget,
		TabBarBackground:     gray100,
		StatusBarBackground:  gray150,
		EditorBackground:     gray50,
		EditorSubheaderBackground: gray100,
		TitleBarInactiveBackground: gray100,

		// Semantic colors
		Error:   red600,
		Warning: yellow600,
		Success: green400,
		Info:    blue200,

		// Background variants
		ErrorBg:   red100,
		SuccessBg: green100,

		// Syntax highlighting - base colors
		Comment:  gray500,
		String:   red500,
		Number:   green400,
		Keyword:  blue600,
		Function: orange600,
		Type:     blue600,
		Variable: blue500,
		Property: green500,

		// Syntax highlighting - accent colors
		StringEscape:       red400,
		Regex:              blue200,
		Decorator:          pink600,
		Punctuation:        gray700,
		PunctuationSpecial: gray400,
		Link:               blue400,

		// Terminal colors
		TerminalBlack:  black,
		TerminalRed:    red600,
		TerminalGreen:  green500,
		TerminalYellow: yellow600,
		TerminalBlue:   blue600,
		TerminalPurple: purple600,
		TerminalCyan:   blue200,
		TerminalWhite:  gray700,

		// Additional semantic mappings
		LineNumber:             gray400,
		ElementActiveBright:    blue500,
		ConflictBackground:     orange500,
		TerminalBrightWhite:    gray900,
		TerminalDimGreen:       green600,
		TerminalDimYellow:      yellow500,
		VersionControlModified: orange700,
		VariableParam:          green500,
		Constructor:            orange500,
		ClassType:              orange500,
		EnumType:               orange500,
		AttributeType:          orange500,
		SelectorAlt:            green500,
		AnnotationType:         yellow500,
		EmbeddedType:           yellow500,

		// Search
		SearchHighlight: blue200Alpha,

		// Player selection colors
		Player1Selection: blue300Alpha,
		Player3Selection: green600,
		Player4Selection: orange500,

		// Scrollbar alpha variants
		ScrollbarThumb:       gray300Alpha,
		ScrollbarThumbHover:  blue400Alpha,
		ScrollbarTrackBorder: gray200,

		// Editor transparency variants
		ActiveLineBackground:   gray100Alpha75,
		WrapGuide:              gray500Alpha30,
		ActiveWrapGuide:        gray600Alpha50,
		DocumentHighlightRead:  blue200Alpha10,
		DocumentHighlightWrite: blue200Alpha40,

		// Standardized player selection alphas
		Player1Alpha: player1Alpha,
		Player3Alpha: player3Alpha,
		Player4Alpha: player4Alpha,
	}
}

// GetFrostedPalette returns the light frosted glass theme palette
func GetFrostedPalette() palette.TronThemePalette {
	// Start with the regular palette
	p := GetPalette()

	// Override for frosted glass effect
	p.BackgroundAppearance = "blurred"
	p.Background = gray50Frosted // 67% opacity like Glazier
	p.BackgroundFrosted = gray50Frosted

	// Make UI elements transparent
	p.TabBarBackground = transparent
	p.StatusBarBackground = statusBarFrosted
	p.Border = borderFrosted
	p.BorderSubtle = transparent
	p.Selection = selectionFrosted

	// Adjust text for better contrast on blurred background
	p.Foreground = gray700Frosted // Slightly higher opacity
	p.ForegroundDim = gray600Frosted
	p.TextFrosted = gray700Frosted

	// Keep elevated surfaces opaque for readability
	p.Highlight = elevatedFrosted

	// Set editor background with slight transparency
	p.EditorBackground = editorFrosted
	p.EditorSubheaderBackground = gray100
	p.TitleBarInactiveBackground = titleBarInactiveFrosted

	// Use darker orange for better contrast on modified files
	p.VersionControlModified = orange700

	// Use transparent track border for frosted to avoid visual clutter
	// p.ScrollbarTrackBorder = transparent

	// Make drop target more visible on frosted background
	p.DropTargetBackground = dropTargetFrosted

	// Adjust other UI elements for frosted effect
	p.ElementHover = gray200

	// Editor specific adjustments
	p.ActiveLineBackground = activeLineFrosted
	p.DocumentHighlightRead = docHighlightReadFrosted
	p.DocumentHighlightWrite = docHighlightWriteFrosted

	// Use darker scrollbar thumb for better visibility on frosted
	p.ScrollbarThumb = gray300AlphaDark

	return p
}
