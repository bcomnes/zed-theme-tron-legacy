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
		BorderDisabled:    p.ForegroundDim,

		// Surfaces
		ElevatedSurfaceBackground: p.Highlight,
		SurfaceBackground:         p.Background,
		Background:                p.Background,

		// Elements
		ElementBackground:    p.Highlight,
		ElementHover:         p.ElementHover,
		ElementActive:        p.ElementHover,
		ElementSelected:      p.Border,
		ElementDisabled:      p.BorderSubtle,
		DropTargetBackground: p.DropTargetBackground,

		// Ghost elements
		GhostElementBackground: p.Transparent,
		GhostElementHover:      p.Highlight,
		GhostElementActive:     p.ElementHover,
		GhostElementSelected:   p.Border,
		GhostElementDisabled:   p.BorderSubtle,

		// Text
		Text:            p.Foreground,
		TextMuted:       p.ForegroundDim,
		TextPlaceholder: p.ForegroundDim,
		TextDisabled:    p.ForegroundDim,
		TextAccent:      p.BorderFocused,

		// Icons
		Icon:            p.Foreground,
		IconMuted:       p.ForegroundDim,
		IconDisabled:    p.ForegroundDim,
		IconPlaceholder: p.ForegroundDim,
		IconAccent:      p.BorderFocused,

		// Bars
		StatusBarBackground:        p.StatusBarBackground,
		TitleBarBackground:         p.StatusBarBackground,
		TitleBarInactiveBackground: p.TitleBarInactiveBackground,
		ToolbarBackground:          p.Background,
		TabBarBackground:           p.TabBarBackground,
		TabInactiveBackground:      p.TabBarBackground,
		TabActiveBackground:        p.Background,

		// Search
		SearchMatchBackground: p.SearchHighlight,

		// Panels
		PanelBackground:    p.TabBarBackground,
		PanelFocusedBorder: p.Success,
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
		EditorSubheaderBackground:              p.EditorSubheaderBackground,
		EditorActiveLineBackground:             p.ActiveLineBackground,
		EditorHighlightedLineBackground:        p.Highlight,
		EditorLineNumber:                       p.LineNumber,
		EditorActiveLineNumber:                 p.Success,
		EditorHoverLineNumber:                  p.Success,
		EditorInvisible:                        p.LineNumber,
		EditorWrapGuide:                        p.WrapGuide,
		EditorActiveWrapGuide:                  p.ActiveWrapGuide,
		EditorDocumentHighlightReadBackground:  p.DocumentHighlightRead,
		EditorDocumentHighlightWriteBackground: p.DocumentHighlightWrite,

		// Terminal
		TerminalBackground:        p.Background,
		TerminalForeground:        p.Foreground,
		TerminalBrightForeground:  p.ForegroundBright,
		TerminalDimForeground:     p.ForegroundDim,
		TerminalAnsiBlack:         p.TerminalBlack,
		TerminalAnsiBrightBlack:   p.ForegroundDim,
		TerminalAnsiDimBlack:      p.Shadow,
		TerminalAnsiRed:           p.TerminalRed,
		TerminalAnsiBrightRed:     p.String,
		TerminalAnsiDimRed:        p.ErrorBg,
		TerminalAnsiGreen:         p.TerminalGreen,
		TerminalAnsiBrightGreen:   p.Success,
		TerminalAnsiDimGreen:      p.TerminalDimGreen,
		TerminalAnsiYellow:        p.TerminalYellow,
		TerminalAnsiBrightYellow:  p.Function,
		TerminalAnsiDimYellow:     p.TerminalDimYellow,
		TerminalAnsiBlue:          p.TerminalBlue,
		TerminalAnsiBrightBlue:    p.Info,
		TerminalAnsiDimBlue:       p.Border,
		TerminalAnsiMagenta:       p.TerminalPurple,
		TerminalAnsiBrightMagenta: p.TerminalPurple,
		TerminalAnsiDimMagenta:    p.TerminalPurple,
		TerminalAnsiCyan:          p.TerminalCyan,
		TerminalAnsiBrightCyan:    p.Info,
		TerminalAnsiDimCyan:       p.Property,
		TerminalAnsiWhite:         p.TerminalWhite,
		TerminalAnsiBrightWhite:   p.TerminalBrightWhite,
		TerminalAnsiDimWhite:      p.Keyword,

		// Links
		LinkTextHover: p.ForegroundDim,

		// Version control
		VersionControlAdded:    p.Success,
		VersionControlModified: p.VersionControlModified,
		VersionControlDeleted:  p.Error,

		// Status colors
		Conflict:           p.Function,
		ConflictBackground: p.ConflictBackground,
		ConflictBorder:     p.Function,

		Created:           p.Success,
		CreatedBackground: p.SuccessBg,
		CreatedBorder:     p.Success,

		Deleted:           p.Error,
		DeletedBackground: p.ErrorBg,
		DeletedBorder:     p.Error,

		Error:           p.Error,
		ErrorBackground: p.ErrorBg,
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
		ModifiedBackground: p.Background,
		ModifiedBorder:     p.Warning,

		Predictive:           p.Comment,
		PredictiveBackground: p.Background,
		PredictiveBorder:     p.TerminalPurple,

		Renamed:           p.Type,
		RenamedBackground: p.Background,
		RenamedBorder:     p.Type,

		Success:           p.Success,
		SuccessBackground: p.SuccessBg,
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

	if p.PanelOverlayBackground != "" {
		style.PanelOverlayBackground = p.PanelOverlayBackground
		style.PanelOverlayHover = p.PanelOverlayHover
	}

	if p.Selection != "" {
		style.EditorSelectionBackground = p.Selection
	}

	if p.Foreground != "" {
		style.Foreground = p.Foreground
	}

	style.VersionControlConflictMarkerOurs = p.Success
	style.VersionControlConflictMarkerTheirs = p.Error

	return Style{
		Name:       name,
		Appearance: appearance,
		Style:      style,
	}
}

// generatePlayers generates the multiplayer cursor colors
func generatePlayers(p TronThemePalette) []map[string]any {
	return []map[string]any{
		{
			"cursor":     p.Type,
			"background": p.Type,
			"selection":  p.Player1Alpha,
		},
		{
			"cursor":     p.Info,
			"background": p.Info,
			"selection":  p.Border,
		},
		{
			"cursor":     p.Success,
			"background": p.Success,
			"selection":  p.Player3Alpha,
		},
		{
			"cursor":     p.Function,
			"background": p.Function,
			"selection":  p.Player4Alpha,
		},
		{
			"cursor":     p.TerminalPurple,
			"background": p.TerminalPurple,
			"selection":  p.TerminalPurple,
		},
		{
			"cursor":     p.Error,
			"background": p.Error,
			"selection":  p.ErrorBg,
		},
		{
			"cursor":     p.Type,
			"background": p.Type,
			"selection":  p.Border,
		},
		{
			"cursor":     p.Keyword,
			"background": p.Keyword,
			"selection":  p.ForegroundDim,
		},
	}
}

// generateSyntax generates the syntax highlighting rules
func generateSyntax(p TronThemePalette) map[string]any {
	syntax := map[string]any{
		// Comments
		"comment": map[string]any{
			"color":       p.Comment,
			"font_style":  nil,
			"font_weight": nil,
		},
		"comment.doc": map[string]any{
			"color":       p.Comment,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Literals
		"string": map[string]any{
			"color":       p.String,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.escape": map[string]any{
			"color":       p.StringEscape,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.regex": map[string]any{
			"color":       p.Regex,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.special": map[string]any{
			"color":       p.Decorator,
			"font_style":  nil,
			"font_weight": nil,
		},
		"string.special.symbol": map[string]any{
			"color":       p.Function,
			"font_style":  nil,
			"font_weight": nil,
		},
		"number": map[string]any{
			"color":       p.Number,
			"font_style":  nil,
			"font_weight": nil,
		},
		"boolean": map[string]any{
			"color":       p.Function,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"constant": map[string]any{
			"color":       p.Function,
			"font_style":  "italic",
			"font_weight": nil,
		},

		// Identifiers
		"variable": map[string]any{
			"color":       p.Variable,
			"font_style":  nil,
			"font_weight": nil,
		},
		"variable.builtin": map[string]any{
			"color":       p.TerminalPurple,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"variable.parameter": map[string]any{
			"color":       p.VariableParam,
			"font_style":  nil,
			"font_weight": nil,
		},
		"variable.special": map[string]any{
			"color":       p.TerminalPurple,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"field": map[string]any{
			"color":       p.Property,
			"font_style":  nil,
			"font_weight": nil,
		},
		"property": map[string]any{
			"color":       p.Property,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Functions
		"function": map[string]any{
			"color":       p.Function,
			"font_style":  nil,
			"font_weight": nil,
		},
		"function.builtin": map[string]any{
			"color":       p.Keyword,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"method": map[string]any{
			"color":       p.Function,
			"font_style":  nil,
			"font_weight": nil,
		},
		"constructor": map[string]any{
			"color":       p.Constructor,
			"font_style":  "italic",
			"font_weight": 700,
		},

		// Keywords
		"keyword": map[string]any{
			"color":       p.Keyword,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"keyword.control": map[string]any{
			"color":       p.Keyword,
			"font_style":  nil,
			"font_weight": nil,
		},
		"operator": map[string]any{
			"color":       p.Keyword,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Types
		"type": map[string]any{
			"color":       p.Type,
			"font_style":  "italic",
			"font_weight": 700,
		},
		"class": map[string]any{
			"color":       p.ClassType,
			"font_style":  "italic",
			"font_weight": 700,
		},
		"interface": map[string]any{
			"color":       p.Property,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"enum": map[string]any{
			"color":       p.EnumType,
			"font_style":  nil,
			"font_weight": nil,
		},
		"variant": map[string]any{
			"color":       p.EnumType,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Markup
		"tag": map[string]any{
			"color":       p.Type,
			"font_style":  nil,
			"font_weight": nil,
		},
		"attribute": map[string]any{
			"color":       p.AttributeType,
			"font_style":  nil,
			"font_weight": nil,
		},
		"selector": map[string]any{
			"color":       p.SelectorAlt,
			"font_style":  nil,
			"font_weight": nil,
		},
		"selector.pseudo": map[string]any{
			"color":       p.Decorator,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Punctuation
		"punctuation": map[string]any{
			"color":       p.Punctuation,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.bracket": map[string]any{
			"color":       p.Punctuation,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.delimiter": map[string]any{
			"color":       p.PunctuationSpecial,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.special": map[string]any{
			"color":       p.Info,
			"font_style":  nil,
			"font_weight": nil,
		},
		"punctuation.list_marker": map[string]any{
			"color":       p.Success,
			"font_style":  nil,
			"font_weight": nil,
		},

		// Misc
		"label": map[string]any{
			"color":       p.Decorator,
			"font_style":  nil,
			"font_weight": nil,
		},
		"namespace": map[string]any{
			"color":       p.Link,
			"font_style":  nil,
			"font_weight": nil,
		},
		"module": map[string]any{
			"color":       p.Link,
			"font_style":  nil,
			"font_weight": nil,
		},
		"decorator": map[string]any{
			"color":       p.Decorator,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"macro": map[string]any{
			"color":       p.Decorator,
			"font_style":  nil,
			"font_weight": 700,
		},
		"parameter": map[string]any{
			"color":       p.VariableParam,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"annotation": map[string]any{
			"color":       p.AnnotationType,
			"font_style":  "italic",
			"font_weight": nil,
		},

		// Special
		"emphasis": map[string]any{
			"color":       nil,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"emphasis.strong": map[string]any{
			"color":       nil,
			"font_style":  nil,
			"font_weight": 700,
		},
		"title": map[string]any{
			"color":       p.Variable,
			"font_style":  nil,
			"font_weight": 700,
		},
		"link_uri": map[string]any{
			"color":       p.Link,
			"font_style":  nil,
			"font_weight": nil,
		},
		"link_text": map[string]any{
			"color":       p.Link,
			"font_style":  nil,
			"font_weight": nil,
		},
		"embedded": map[string]any{
			"color":       p.EmbeddedType,
			"font_style":  nil,
			"font_weight": nil,
		},
		"primary": map[string]any{
			"color":       p.Foreground,
			"font_style":  nil,
			"font_weight": 700,
		},
		"hint": map[string]any{
			"color":       p.Hint,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"predictive": map[string]any{
			"color":       p.TerminalPurple,
			"font_style":  "italic",
			"font_weight": nil,
		},
		"preproc": map[string]any{
			"color":       p.Info,
			"font_style":  nil,
			"font_weight": nil,
		},
		"regex": map[string]any{
			"color":       p.Regex,
			"font_style":  nil,
			"font_weight": nil,
		},
		"text.literal": map[string]any{
			"color":       p.EmbeddedType,
			"font_style":  nil,
			"font_weight": nil,
		},
	}

	return syntax
}
