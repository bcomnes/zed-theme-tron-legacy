package atelier

import "github.com/bcomnes/zed-theme-tron-legacy/tools/palette"

// GetPalette returns the Atelier-inspired blue-tinted light theme palette
func GetPalette() palette.TronThemePalette {
	colors := GetColors()

	return palette.TronThemePalette{
		// Core colors
		Background:       colors["base00"],
		Foreground:       colors["base05"],
		ForegroundDim:    colors["base04"],
		ForegroundBright: colors["base07"],

		// UI elements
		Border:        colors["base02"],
		BorderSubtle:  colors["base01"],
		BorderFocused: colors["cyan"],
		Selection:     colors["base02"],
		Highlight:     colors["base01"],
		ElementActive: colors["blue"],
		Shadow:        colors["shadow"],

		// Semantic colors
		Error:   colors["red"],
		Warning: colors["orange"],
		Success: colors["green"],
		Info:    colors["cyan"],

		// Background variants
		ErrorBg:   colors["base00"],
		SuccessBg: colors["base00"],

		// Syntax highlighting - base colors
		Comment:  colors["base03"],
		String:   colors["green"],
		Number:   colors["orange"],
		Keyword:  colors["purple"],
		Function: colors["blue"],
		Type:     colors["yellow"],
		Variable: colors["red"],
		Property: colors["teal"],

		// Syntax highlighting - accent colors
		StringEscape:       colors["orange"],
		Regex:              colors["teal"],
		Decorator:          colors["magenta"],
		Punctuation:        colors["base05"],
		PunctuationSpecial: colors["base04"],
		Link:               colors["blue_bright"],

		// Terminal colors
		TerminalBlack:  colors["base00"],
		TerminalRed:    colors["red"],
		TerminalGreen:  colors["green"],
		TerminalYellow: colors["yellow"],
		TerminalBlue:   colors["blue"],
		TerminalPurple: colors["purple"],
		TerminalCyan:   colors["cyan"],
		TerminalWhite:  colors["base05"],

		// Additional semantic mappings
		LineNumber:             colors["base03"],
		ElementActiveBright:    colors["blue_bright"],
		ConflictBackground:     colors["base00"],
		TerminalBrightWhite:    colors["base07"],
		TerminalDimGreen:       colors["green_dim"],
		TerminalDimYellow:      colors["base02"],
		VersionControlModified: colors["blue"],
		VariableParam:          colors["teal"],
		Constructor:            colors["red"],
		ClassType:              colors["yellow"],
		EnumType:               colors["yellow"],
		AttributeType:          colors["blue"],
		SelectorAlt:            colors["teal"],
		AnnotationType:         colors["orange"],
		EmbeddedType:           colors["green"],

		// Search
		SearchHighlight: colors["highlight_alpha"],

		// Player selection colors
		Player1Selection: colors["blue_alpha"],
		Player3Selection: colors["gray_alpha"],
		Player4Selection: colors["highlight_alpha"],
	}
}
