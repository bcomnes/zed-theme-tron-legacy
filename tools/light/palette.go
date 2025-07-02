package light

import "github.com/bcomnes/zed-theme-tron-legacy/tools/palette"

// GetPalette returns the light theme palette
func GetPalette() palette.TronThemePalette {
	colors := GetColors()

	return palette.TronThemePalette{
		// Core colors
		Background:       colors["gray50"],
		Foreground:       colors["gray700"],
		ForegroundDim:    colors["gray600"],
		ForegroundBright: colors["gray900"],

		// UI elements
		Border:        colors["gray300"],
		BorderSubtle:  colors["gray200"],
		BorderFocused: colors["blue200"],
		Selection:     colors["gray200"],
		Highlight:     colors["gray100"],
		ElementActive: colors["blue500"],
		ElementHover:  colors["gray200"],
		Shadow:        colors["shadow"],
		Transparent:   colors["transparent"],
		DropTargetBackground: colors["droptarget"],
		TabBarBackground: colors["gray100"],
		StatusBarBackground: colors["gray800"],

		// Semantic colors
		Error:   colors["red600"],
		Warning: colors["yellow600"],
		Success: colors["green400"],
		Info:    colors["blue200"],

		// Background variants
		ErrorBg:   colors["red100"],
		SuccessBg: colors["green100"],

		// Syntax highlighting - base colors
		Comment:  colors["gray500"],
		String:   colors["red500"],
		Number:   colors["green400"],
		Keyword:  colors["blue600"],
		Function: colors["orange600"],
		Type:     colors["blue600"],
		Variable: colors["blue500"],
		Property: colors["green500"],

		// Syntax highlighting - accent colors
		StringEscape:       colors["red400"],
		Regex:              colors["blue200"],
		Decorator:          colors["pink600"],
		Punctuation:        colors["gray700"],
		PunctuationSpecial: colors["gray400"],
		Link:               colors["blue400"],

		// Terminal colors
		TerminalBlack:  colors["black"],
		TerminalRed:    colors["red600"],
		TerminalGreen:  colors["green500"],
		TerminalYellow: colors["yellow600"],
		TerminalBlue:   colors["blue600"],
		TerminalPurple: colors["purple600"],
		TerminalCyan:   colors["blue200"],
		TerminalWhite:  colors["gray700"],

		// Additional semantic mappings
		LineNumber:             colors["gray400"],
		ElementActiveBright:    colors["blue500"],
		ConflictBackground:     colors["orange500"],
		TerminalBrightWhite:    colors["gray900"],
		TerminalDimGreen:       colors["green600"],
		TerminalDimYellow:      colors["yellow500"],
		VersionControlModified: colors["orange500"],
		VariableParam:          colors["green500"],
		Constructor:            colors["orange500"],
		ClassType:              colors["orange500"],
		EnumType:               colors["orange500"],
		AttributeType:          colors["orange500"],
		SelectorAlt:            colors["green500"],
		AnnotationType:         colors["yellow500"],
		EmbeddedType:           colors["yellow500"],

		// Search
		SearchHighlight: colors["blue200alpha"],

		// Player selection colors
		Player1Selection: colors["blue300alpha"],
		Player3Selection: colors["green600"],
		Player4Selection: colors["orange500"],

		// Scrollbar alpha variants
		ScrollbarThumb:      colors["gray300alpha"],
		ScrollbarThumbHover: colors["blue400alpha"],
		ScrollbarTrackBorder: colors["oneTrackBorder"],

		// Editor transparency variants
		ActiveLineBackground: colors["gray100alpha75"],
		WrapGuide: colors["gray500alpha30"],
		ActiveWrapGuide: colors["gray600alpha50"],
		DocumentHighlightRead: colors["blue200alpha10"],
		DocumentHighlightWrite: colors["blue200alpha40"],

		// Standardized player selection alphas
		Player1Alpha: colors["player1alpha"],
		Player3Alpha: colors["player3alpha"],
		Player4Alpha: colors["player4alpha"],
	}
}
