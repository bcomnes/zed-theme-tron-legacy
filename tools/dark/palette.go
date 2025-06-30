package dark

import "github.com/bcomnes/zed-theme-tron-legacy/tools/palette"

// GetPalette returns the dark theme palette
func GetPalette() palette.TronThemePalette {
	colors := GetColors()

	return palette.TronThemePalette{
		// Core colors
		Background:       colors["gray900"],
		Foreground:       colors["gray200"],
		ForegroundDim:    colors["gray500"],
		ForegroundBright: colors["gray50"],

		// UI elements
		Border:        colors["neutral600"],
		BorderSubtle:  colors["gray700"],
		BorderFocused: colors["blue200"],
		Selection:     colors["gray700"],
		Highlight:     colors["neutral800"],
		ElementActive: colors["blue400"],
		ElementHover:  colors["gray700"],
		Shadow:        colors["shadow"],
		Transparent:   colors["transparent"],
		DropTargetBackground: colors["droptarget"],
		TabBarBackground: colors["gray800"],
		StatusBarBackground: colors["gray750"],

		// Semantic colors
		Error:   colors["red500"],
		Warning: colors["yellow500"],
		Success: colors["green300"],
		Info:    colors["blue200"],

		// Background variants
		ErrorBg:   colors["red700"],
		SuccessBg: colors["green700"],

		// Syntax highlighting - base colors
		Comment:  colors["gray400"],
		String:   colors["red400"],
		Number:   colors["green300"],
		Keyword:  colors["blue500"],
		Function: colors["orange500"],
		Type:     colors["blue500"],
		Variable: colors["blue200bright"],
		Property: colors["green500"],

		// Syntax highlighting - accent colors
		StringEscape:       colors["red300"],
		Regex:              colors["blue200"],
		Decorator:          colors["pink500"],
		Punctuation:        colors["gray200"],
		PunctuationSpecial: colors["gray500"],
		Link:               colors["blue300"],

		// Terminal colors
		TerminalBlack:  colors["black"],
		TerminalRed:    colors["red500"],
		TerminalGreen:  colors["green500"],
		TerminalYellow: colors["yellow500"],
		TerminalBlue:   colors["blue500"],
		TerminalPurple: colors["purple500"],
		TerminalCyan:   colors["blue200"],
		TerminalWhite:  colors["gray200"],

		// Additional semantic mappings
		LineNumber:             colors["gray500"],
		ElementActiveBright:    colors["blue400"],
		ConflictBackground:     colors["orange400"],
		TerminalBrightWhite:    colors["purewhite"],
		TerminalDimGreen:       colors["green600"],
		TerminalDimYellow:      colors["yellow400"],
		VersionControlModified: colors["yellow400"],
		VariableParam:          colors["green500"],
		Constructor:            colors["orange400"],
		ClassType:              colors["orange400"],
		EnumType:               colors["orange400"],
		AttributeType:          colors["orange400"],
		SelectorAlt:            colors["green500"],
		AnnotationType:         colors["yellow400"],
		EmbeddedType:           colors["yellow400"],

		// Search
		SearchHighlight: colors["blue200alpha"],

		// Player selection colors
		Player1Selection: colors["blue500alpha"],
		Player3Selection: colors["green600"],
		Player4Selection: colors["orange400"],
	}
}
