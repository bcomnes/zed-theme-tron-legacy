package light

import "github.com/bcomnes/zed-theme-tron-legacy/tools/palette"

// GetPalette returns the light theme palette
func GetPalette() palette.TronThemePalette {
	colors := GetColors()

	return palette.TronThemePalette{
		// Core colors
		Background:       colors["gray50"],
		Foreground:       colors["gray700"],
		ForegroundDim:    colors["gray400"],
		ForegroundBright: colors["gray900"],

		// UI elements
		Border:        colors["gray300"],
		BorderSubtle:  colors["gray200"],
		BorderFocused: colors["blue400"],
		Selection:     colors["gray300"],
		Highlight:     colors["gray100"],
		ElementActive: colors["blue600"],
		Shadow:        colors["shadow"],

		// Semantic colors
		Error:   colors["red700"],
		Warning: colors["yellow700"],
		Success: colors["green700"],
		Info:    colors["blue400"],

		// Background variants
		ErrorBg:   colors["red100"],
		SuccessBg: colors["green100"],

		// Syntax highlighting - base colors
		Comment:  colors["gray500"],
		String:   colors["red600"],
		Number:   colors["green700"],
		Keyword:  colors["blue700"],
		Function: colors["orange700"],
		Type:     colors["blue700"],
		Variable: colors["gray900"],
		Property: colors["green600"],

		// Syntax highlighting - accent colors
		StringEscape:       colors["red500"],
		Regex:              colors["blue400"],
		Decorator:          colors["pink700"],
		Punctuation:        colors["gray700"],
		PunctuationSpecial: colors["gray400"],
		Link:               colors["blue500"],

		// Terminal colors
		TerminalBlack:  colors["black"],
		TerminalRed:    colors["red700"],
		TerminalGreen:  colors["green600"],
		TerminalYellow: colors["yellow700"],
		TerminalBlue:   colors["blue700"],
		TerminalPurple: colors["purple700"],
		TerminalCyan:   colors["blue400"],
		TerminalWhite:  colors["gray700"],

		// Additional semantic mappings
		LineNumber:             colors["gray400"],
		ElementActiveBright:    colors["blue600"],
		ConflictBackground:     colors["orange600"],
		TerminalBrightWhite:    colors["gray900"],
		TerminalDimGreen:       colors["green200"],
		TerminalDimYellow:      colors["yellow600"],
		VersionControlModified: colors["yellow600"],
		VariableParam:          colors["green600"],
		Constructor:            colors["orange600"],
		ClassType:              colors["orange600"],
		EnumType:               colors["orange600"],
		AttributeType:          colors["orange600"],
		SelectorAlt:            colors["green600"],
		AnnotationType:         colors["yellow600"],
		EmbeddedType:           colors["yellow600"],

		// Search
		SearchHighlight: colors["blue400alpha"],

		// Player selection colors
		Player1Selection: colors["blue700alpha"],
		Player3Selection: colors["green200"],
		Player4Selection: colors["orange600"],
	}
}
