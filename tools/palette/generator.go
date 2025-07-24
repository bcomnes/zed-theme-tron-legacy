package palette

// ThemeVariant represents a single theme variant with its metadata
type ThemeVariant struct {
	Name       string
	Appearance string
	Palette    TronThemePalette
}

// GenerateTheme generates the complete theme JSON structure with any number of variants
func GenerateTheme(name string, author string, variants ...ThemeVariant) Theme {
	themes := make([]Style, 0, len(variants))
	for _, variant := range variants {
		themes = append(themes, GenerateThemeStyle(variant.Name, variant.Appearance, variant.Palette))
	}

	return Theme{
		Schema: "https://zed.dev/schema/themes/v0.2.0.json",
		Author: author,
		Name:   name,
		Themes: themes,
	}
}

// GenerateThemeStyle generates the style portion of the theme
// This function maps semantic palette values to theme properties
func GenerateThemeStyle(name string, appearance string, p TronThemePalette) Style {
	style := &ThemeStyle{
		// Borders
		Border:            p.Border,
		BorderVariant:     p.BorderSubtle,
		BorderFocused:     p.BorderFocused,
		BorderSelected:    p.BorderFocused,
		BorderTransparent: p.Transparent,
		BorderDisabled:    p.ForegroundMuted,

		// Surfaces
		ElevatedSurfaceBackground: p.BackgroundElevated,
		SurfaceBackground:         p.Background,
		Background:                p.Background,

		// Elements
		ElementBackground:    p.BackgroundElevated,
		ElementHover:         p.Interactive,
		ElementActive:        p.Interactive,
		ElementSelected:      p.Border,
		ElementDisabled:      p.BorderSubtle,
		DropTargetBackground: p.DropTarget,

		// Ghost elements
		GhostElementBackground: p.Transparent,
		GhostElementHover:      p.Interactive,
		GhostElementActive:     p.Interactive,
		GhostElementSelected:   p.Border,
		GhostElementDisabled:   p.BorderSubtle,

		// Text
		Text:            p.Foreground,
		TextMuted:       p.ForegroundMuted,
		TextPlaceholder: p.ForegroundMuted,
		TextDisabled:    p.ForegroundMuted,
		TextAccent:      p.BorderFocused,

		// Icons
		Icon:            p.Foreground,
		IconMuted:       p.ForegroundMuted,
		IconDisabled:    p.ForegroundMuted,
		IconPlaceholder: p.ForegroundMuted,
		IconAccent:      p.BorderFocused,

		// Bars
		StatusBarBackground:        p.Statusbar,
		TitleBarBackground:         p.Statusbar,
		TitleBarInactiveBackground: p.StatusbarInactive,
		ToolbarBackground:          p.Background,
		TabBarBackground:           p.Surface,
		TabInactiveBackground:      p.Surface,
		TabActiveBackground:        p.Background,

		// Search
		SearchMatchBackground: p.MatchHighlight,

		// Panels
		PanelBackground:    p.Surface,
		PanelFocusedBorder: p.UIAccent,
		PaneFocusedBorder:  p.BorderFocused,

		// Scrollbar
		ScrollbarThumbBackground:      p.ScrollbarThumb,
		ScrollbarThumbHoverBackground: p.ScrollbarThumbHover,
		ScrollbarThumbBorder:          p.Border,
		ScrollbarTrackBackground:      p.Transparent,
		ScrollbarTrackBorder:          p.ScrollbarTrackBorder,

		// Editor
		EditorForeground:                       p.Foreground,
		EditorBackground:                       p.EditorBackground,
		EditorGutterBackground:                 p.EditorBackground,
		EditorSubheaderBackground:              p.EditorSubheader,
		EditorActiveLineBackground:             p.ActiveLine,
		EditorHighlightedLineBackground:        p.BackgroundElevated,
		EditorLineNumber:                       p.LineNumber,
		EditorActiveLineNumber:                 p.UIAccent,
		EditorHoverLineNumber:                  p.UIAccent,
		EditorInvisible:                        p.LineNumber,
		EditorWrapGuide:                        p.GuideNormal,
		EditorActiveWrapGuide:                  p.GuideActive,
		EditorDocumentHighlightReadBackground:  p.DocumentHighlight,
		EditorDocumentHighlightWriteBackground: p.DocumentHighlightWrite,

		// Terminal
		TerminalBackground:        p.Background,
		TerminalForeground:        p.Foreground,
		TerminalBrightForeground:  p.ForegroundStrong,
		TerminalDimForeground:     p.ForegroundMuted,
		TerminalAnsiBlack:         p.TerminalBlack,
		TerminalAnsiBrightBlack:   p.TerminalBrightBlack,
		TerminalAnsiDimBlack:      p.TerminalDimBlack,
		TerminalAnsiRed:           p.TerminalRed,
		TerminalAnsiBrightRed:     p.TerminalBrightRed,
		TerminalAnsiDimRed:        p.TerminalDimRed,
		TerminalAnsiGreen:         p.TerminalGreen,
		TerminalAnsiBrightGreen:   p.TerminalBrightGreen,
		TerminalAnsiDimGreen:      p.TerminalDimGreen,
		TerminalAnsiYellow:        p.TerminalYellow,
		TerminalAnsiBrightYellow:  p.TerminalBrightYellow,
		TerminalAnsiDimYellow:     p.TerminalDimYellow,
		TerminalAnsiBlue:          p.TerminalBlue,
		TerminalAnsiBrightBlue:    p.TerminalBrightBlue,
		TerminalAnsiDimBlue:       p.TerminalDimBlue,
		TerminalAnsiMagenta:       p.TerminalPurple,
		TerminalAnsiBrightMagenta: p.TerminalBrightPurple,
		TerminalAnsiDimMagenta:    p.TerminalDimMagenta,
		TerminalAnsiCyan:          p.TerminalCyan,
		TerminalAnsiBrightCyan:    p.TerminalBrightCyan,
		TerminalAnsiDimCyan:       p.TerminalDimCyan,
		TerminalAnsiWhite:         p.TerminalWhite,
		TerminalAnsiBrightWhite:   p.TerminalBrightWhite,
		TerminalAnsiDimWhite:      p.TerminalDimWhite,

		// Links
		LinkTextHover: p.ForegroundMuted,

		// Version control
		VersionControlAdded:    p.Success,
		VersionControlModified: p.VCSModified,
		VersionControlDeleted:  p.Error,

		// Status colors
		Conflict:           p.Accent,
		ConflictBackground: p.VCSConflict,
		ConflictBorder:     p.Accent,

		Created:           p.Success,
		CreatedBackground: p.SuccessSurface,
		CreatedBorder:     p.Success,

		Deleted:           p.Error,
		DeletedBackground: p.ErrorSurface,
		DeletedBorder:     p.Error,

		Error:           p.Error,
		ErrorBackground: p.ErrorSurface,
		ErrorBorder:     p.Error,

		Hidden:           p.Comment,
		HiddenBackground: p.Background,
		HiddenBorder:     p.Comment,

		Hint:           p.Hint,
		HintBackground: p.Background,
		HintBorder:     p.Hint,

		Ignored:           p.Comment,
		IgnoredBackground: p.Background,
		IgnoredBorder:     p.Comment,

		Info:           p.Info,
		InfoBackground: p.Background,
		InfoBorder:     p.Info,

		Modified:           p.Warning,
		ModifiedBackground: p.BackgroundElevated,
		ModifiedBorder:     p.Warning,

		Predictive:           p.Comment,
		PredictiveBackground: p.Background,
		PredictiveBorder:     p.TerminalPurple,

		Renamed:           p.Type,
		RenamedBackground: p.Background,
		RenamedBorder:     p.Type,

		Success:           p.Success,
		SuccessBackground: p.SuccessSurface,
		SuccessBorder:     p.Success,

		Unreachable:           p.Comment,
		UnreachableBackground: p.Background,
		UnreachableBorder:     p.Comment,

		Warning:           p.Warning,
		WarningBackground: p.Background,
		WarningBorder:     p.Warning,

		// Players and Syntax
		Players: generatePlayers(p),
		Syntax:  generateSyntax(p),
	}

	// Add optional fields
	if p.BackgroundAppearance != "" && p.BackgroundAppearance != "opaque" {
		style.BackgroundAppearance = p.BackgroundAppearance
	}

	if p.BackgroundOverlay != "" {
		style.PanelOverlayBackground = p.BackgroundOverlay
		style.PanelOverlayHover = p.BackgroundOverlayHover
	}

	if p.Selection != "" {
		style.EditorSelectionBackground = p.Selection
	}

	if p.Foreground != "" {
		style.Foreground = p.Foreground
	}

	style.VersionControlConflictMarkerOurs = p.UIAccent
	style.VersionControlConflictMarkerTheirs = p.Error

	// ========== NEWLY ADDED FIELDS (2024) ==========
	// Map values for newly added fields (2024)
	// These fields are optional in Zed themes and use fallback values when not specified

	// Selection background - semi-transparent selection color
	// NOTE: element.selection_background is broken in Zed, so we don't set it
	// if p.SelectionAlpha != "" {
	// 	style.ElementSelectionBackground = p.SelectionAlpha
	// }

	// Indent guides - use subtle borders/muted colors
	style.PanelIndentGuide = p.Border  // Better contrast than BorderSubtle
	style.PanelIndentGuideHover = p.BorderFocused
	style.PanelIndentGuideActive = p.UIAccent  // Use accent color
	style.EditorIndentGuide = p.Border  // Better contrast than WrapGuide
	style.EditorIndentGuideActive = p.UIAccent  // Use accent color

	// Editor debugger - use error/warning colors
	style.EditorDebuggerActiveLineBackground = p.ErrorSurface
	style.EditorDocumentHighlightBracketBackground = p.DocumentHighlight

	// Scrollbar active state - use accent color with transparency
	style.ScrollbarThumbActiveBackground = p.ScrollbarThumbActive  // Darker than hover with transparency

	// Minimap - similar to scrollbar but more transparent
	style.MinimapThumbBackground = p.ScrollbarThumb  // Already has good transparency
	style.MinimapThumbHoverBackground = p.ScrollbarThumbHover  // Use consistent hover state
	style.MinimapThumbActiveBackground = p.ScrollbarThumbActive  // Use darker active color
	style.MinimapThumbBorder = p.Transparent

	// Terminal ANSI background - same as terminal background
	style.TerminalAnsiBackground = p.Background

	// Additional version control colors
	style.VersionControlRenamed = p.VCSModified
	style.VersionControlConflict = p.Accent
	style.VersionControlIgnored = p.Comment

	// Pane group border
	style.PaneGroupBorder = p.Border

	// Debugger accent - use error color for breakpoints
	style.DebuggerAccent = p.Error
	// ========== END NEWLY ADDED FIELDS ==========


	return Style{
		Name:       name,
		Appearance: appearance,
		Accents:    p.Accents, // From Gruvbox theme - optional accent colors for UI highlights
		Style:      style,
	}
}

// generatePlayers generates the multiplayer cursor colors
func generatePlayers(p TronThemePalette) []Player {
	return []Player{
		{
			Cursor:     p.Type,
			Background: p.Type,
			Selection:  p.Player1,
		},
		{
			Cursor:     p.Info,
			Background: p.Info,
			Selection:  p.Player2,
		},
		{
			Cursor:     p.Success,
			Background: p.Success,
			Selection:  p.Player3,
		},
		{
			Cursor:     p.Accent,
			Background: p.Accent,
			Selection:  p.Player4,
		},
		{
			Cursor:     p.TerminalPurple,
			Background: p.TerminalPurple,
			Selection:  p.TerminalPurple,
		},
		{
			Cursor:     p.Error,
			Background: p.Error,
			Selection:  p.ErrorSurface,
		},
		{
			Cursor:     p.Type,
			Background: p.Type,
			Selection:  p.Border,
		},
		{
			Cursor:     p.Keyword,
			Background: p.Keyword,
			Selection:  p.ForegroundMuted,
		},
	}
}

// generateSyntax generates the syntax highlighting rules
func generateSyntax(p TronThemePalette) SyntaxStyles {
	// Helper function to create string pointers
	str := func(s string) *string { return &s }
	num := func(n int) *int { return &n }

	return SyntaxStyles{
		// Official theme field order
		Attribute:             SyntaxStyle{Color: p.Attribute, FontStyle: nil, FontWeight: nil},
		Boolean:               SyntaxStyle{Color: p.Accent, FontStyle: str("italic"), FontWeight: nil},
		Comment:               SyntaxStyle{Color: p.Comment, FontStyle: nil, FontWeight: nil},
		CommentDoc:            SyntaxStyle{Color: p.Comment, FontStyle: nil, FontWeight: nil},
		Constant:              SyntaxStyle{Color: p.Accent, FontStyle: str("italic"), FontWeight: nil},
		Constructor:           SyntaxStyle{Color: p.Constructor, FontStyle: str("italic"), FontWeight: num(700)},
		Embedded:              SyntaxStyle{Color: p.Embedded, FontStyle: nil, FontWeight: nil},
		Emphasis:              SyntaxStyle{Color: p.Info, FontStyle: str("italic"), FontWeight: nil},
		EmphasisStrong:        SyntaxStyle{Color: p.Accent, FontStyle: nil, FontWeight: num(700)},
		Enum:                  SyntaxStyle{Color: p.Enum, FontStyle: nil, FontWeight: nil},
		Function:              SyntaxStyle{Color: p.Function, FontStyle: nil, FontWeight: nil},
		FunctionBuiltin:       SyntaxStyle{Color: p.Type, FontStyle: str("italic"), FontWeight: nil}, // From Gruvbox theme
		Hint:                  SyntaxStyle{Color: p.Hint, FontStyle: str("italic"), FontWeight: nil},
		Keyword:               SyntaxStyle{Color: p.Keyword, FontStyle: str("italic"), FontWeight: nil},
		Label:                 SyntaxStyle{Color: p.Decorator, FontStyle: nil, FontWeight: nil},
		LinkText:              SyntaxStyle{Color: p.Namespace, FontStyle: nil, FontWeight: nil},
		LinkURI:               SyntaxStyle{Color: p.Namespace, FontStyle: nil, FontWeight: nil},
		Namespace:             SyntaxStyle{Color: p.Namespace, FontStyle: nil, FontWeight: nil},
		Number:                SyntaxStyle{Color: p.Number, FontStyle: nil, FontWeight: nil},
		Operator:              SyntaxStyle{Color: p.Keyword, FontStyle: nil, FontWeight: nil},
		Predictive:            SyntaxStyle{Color: p.TerminalPurple, FontStyle: str("italic"), FontWeight: nil},
		Preproc:               SyntaxStyle{Color: p.Info, FontStyle: nil, FontWeight: nil},
		Primary:               SyntaxStyle{Color: p.Foreground, FontStyle: nil, FontWeight: num(700)},
		Property:              SyntaxStyle{Color: p.Property, FontStyle: nil, FontWeight: nil},
		Punctuation:           SyntaxStyle{Color: p.Punctuation, FontStyle: nil, FontWeight: nil},
		PunctuationBracket:    SyntaxStyle{Color: p.Punctuation, FontStyle: nil, FontWeight: nil},
		PunctuationDelimiter:  SyntaxStyle{Color: p.PunctuationMuted, FontStyle: nil, FontWeight: nil},
		PunctuationListMarker: SyntaxStyle{Color: p.UIAccent, FontStyle: nil, FontWeight: nil},
		PunctuationSpecial:    SyntaxStyle{Color: p.Info, FontStyle: nil, FontWeight: nil},
		Selector:              SyntaxStyle{Color: p.Property, FontStyle: nil, FontWeight: nil},
		SelectorPseudo:        SyntaxStyle{Color: p.Decorator, FontStyle: nil, FontWeight: nil},
		String:                SyntaxStyle{Color: p.String, FontStyle: nil, FontWeight: nil},
		StringEscape:          SyntaxStyle{Color: p.StringEscape, FontStyle: nil, FontWeight: nil},
		StringRegex:           SyntaxStyle{Color: p.Regex, FontStyle: nil, FontWeight: nil},
		StringSpecial:         SyntaxStyle{Color: p.Decorator, FontStyle: nil, FontWeight: nil},
		StringSpecialSymbol:   SyntaxStyle{Color: p.Accent, FontStyle: nil, FontWeight: nil},
		Tag:                   SyntaxStyle{Color: p.Tag, FontStyle: nil, FontWeight: nil},
		TextLiteral:           SyntaxStyle{Color: p.Embedded, FontStyle: nil, FontWeight: nil},
		Title:                 SyntaxStyle{Color: p.Variable, FontStyle: nil, FontWeight: num(700)},
		Type:                  SyntaxStyle{Color: p.Type, FontStyle: str("italic"), FontWeight: num(700)},
		Variable:              SyntaxStyle{Color: p.Variable, FontStyle: nil, FontWeight: nil},
		VariableSpecial:       SyntaxStyle{Color: p.SpecialVariable, FontStyle: str("italic"), FontWeight: nil},
		Variant:               SyntaxStyle{Color: p.Enum, FontStyle: nil, FontWeight: nil},
	}
}
