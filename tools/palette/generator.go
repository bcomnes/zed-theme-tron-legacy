package palette

// ThemeVariant represents a single theme variant with its metadata
type ThemeVariant struct {
	Name       string
	Appearance string
	Palette    TronThemePalette
}

// GenerateTheme generates the complete theme JSON structure with any number of variants
func GenerateTheme(name string, author string, variants ...ThemeVariant) map[string]any {
	themes := make([]any, 0, len(variants))
	for _, variant := range variants {
		themes = append(themes, GenerateThemeStyle(variant.Name, variant.Appearance, variant.Palette))
	}

	return map[string]any{
		"$schema": "https://zed.dev/schema/themes/v0.2.0.json",
		"author":  author,
		"name":    name,
		"themes":  themes,
	}
}

// GenerateThemeStyle generates the style portion of the theme
// This function maps semantic palette values to theme properties
func GenerateThemeStyle(name string, appearance string, p TronThemePalette) map[string]any {
	style := map[string]any{
		"name":       name,
		"appearance": appearance,
		"style": map[string]any{
			// Core colors
			"background":       p.Background,
			"foreground":       p.Foreground,
			"text":             p.Foreground,
			"text.muted":       p.ForegroundDim,
			"text.placeholder": p.ForegroundDim,
			"text.disabled":    p.ForegroundDim,
			"text.accent":      p.BorderFocused,

			// Icon colors
			"icon":             p.Foreground,
			"icon.muted":       p.ForegroundDim,
			"icon.disabled":    p.ForegroundDim,
			"icon.placeholder": p.ForegroundDim,
			"icon.accent":      p.BorderFocused,

			// Borders
			"border":             p.Border,
			"border.variant":     p.BorderSubtle,
			"border.focused":     p.BorderFocused,
			"border.selected":    p.BorderFocused,
			"border.transparent": p.Transparent,
			"border.disabled":    p.ForegroundDim,

			// UI surfaces
			"surface.background":          p.Background,
			"elevated_surface.background": p.Highlight,
			"element.background":          p.Highlight,
			"element.hover":               p.ElementHover,
			"element.active":              p.ElementHover,
			"element.selected":            p.Border,
			"element.disabled":            p.BorderSubtle,
			"drop_target.background":      p.DropTargetBackground,

			// Ghost elements
			"ghost_element.background": p.Transparent,
			"ghost_element.hover":      p.Highlight,
			"ghost_element.active":     p.ElementHover,
			"ghost_element.selected":   p.Border,
			"ghost_element.disabled":   p.BorderSubtle,

			// Editor specific
			"editor.background":                          p.EditorBackground,
			"editor.foreground":                          p.Foreground,
			"editor.gutter.background":                   p.EditorBackground,
			"editor.active_line.background":              p.ActiveLineBackground,
			"editor.highlighted_line.background":         p.Highlight,
			"editor.line_number":                         p.LineNumber,
			"editor.active_line_number":                  p.Success,
			"editor.hover_line_number":                   p.Success,
			"editor.invisible":                           p.LineNumber,
			"editor.wrap_guide":                          p.WrapGuide,
			"editor.active_wrap_guide":                   p.ActiveWrapGuide,
			"editor.selection.background":                p.Selection,
			"editor.document_highlight.read_background":  p.DocumentHighlightRead,
			"editor.document_highlight.write_background": p.DocumentHighlightWrite,
			"editor.subheader.background":                p.EditorSubheaderBackground,

			// Search
			"search.match_background": p.SearchHighlight,

			// Panels and tabs
			"panel.background":              p.TabBarBackground,
			"panel.focused_border":          p.Success,
			"pane.focused_border":           p.BorderFocused,
			"status_bar.background":         p.StatusBarBackground,
			"title_bar.background":          p.StatusBarBackground,
			"title_bar.inactive_background": p.TitleBarInactiveBackground,
			"toolbar.background":            p.Background,
			"tab_bar.background":            p.TabBarBackground,
			"tab.inactive_background":       p.TabBarBackground,
			"tab.active_background":         p.Background,

			// Scrollbar
			"scrollbar.thumb.background":       p.ScrollbarThumb,
			"scrollbar.thumb.hover_background": p.ScrollbarThumbHover,
			"scrollbar.thumb.border":           p.Border,
			"scrollbar.track.background":       p.Transparent,
			"scrollbar.track.border":           p.ScrollbarTrackBorder,

			// Terminal
			"terminal.background":          p.Background,
			"terminal.foreground":          p.Foreground,
			"terminal.bright_foreground":   p.ForegroundBright,
			"terminal.dim_foreground":      p.ForegroundDim,
			"terminal.ansi.black":          p.TerminalBlack,
			"terminal.ansi.red":            p.TerminalRed,
			"terminal.ansi.green":          p.TerminalGreen,
			"terminal.ansi.yellow":         p.TerminalYellow,
			"terminal.ansi.blue":           p.TerminalBlue,
			"terminal.ansi.magenta":        p.TerminalPurple,
			"terminal.ansi.cyan":           p.TerminalCyan,
			"terminal.ansi.white":          p.TerminalWhite,
			"terminal.ansi.bright_black":   p.ForegroundDim,
			"terminal.ansi.bright_red":     p.String, // Reuse string color
			"terminal.ansi.bright_green":   p.Success,
			"terminal.ansi.bright_yellow":  p.Function, // Reuse function/orange
			"terminal.ansi.bright_blue":    p.Info,
			"terminal.ansi.bright_magenta": p.TerminalPurple,
			"terminal.ansi.bright_cyan":    p.Info,
			"terminal.ansi.bright_white":   p.TerminalBrightWhite,
			"terminal.ansi.dim_black":      p.Shadow,
			"terminal.ansi.dim_red":        p.ErrorBg,
			"terminal.ansi.dim_green":      p.TerminalDimGreen,
			"terminal.ansi.dim_yellow":     p.TerminalDimYellow,
			"terminal.ansi.dim_blue":       p.Border,
			"terminal.ansi.dim_magenta":    p.TerminalPurple,
			"terminal.ansi.dim_cyan":       p.Property, // Cyan dim
			"terminal.ansi.dim_white":      p.Keyword,

			// Links
			"link_text.hover": p.ForegroundDim,

			// Version control
			"version_control.added":                  p.Success,
			"version_control.modified":               p.VersionControlModified,
			"version_control.deleted":                p.Error,
			"version_control.conflict_marker.ours":   p.Success,
			"version_control.conflict_marker.theirs": p.Error,

			// Diagnostics
			"error":              p.Error,
			"error.background":   p.ErrorBg,
			"error.border":       p.Error,
			"warning":            p.Warning,
			"warning.background": p.Background,
			"warning.border":     p.Warning,
			"info":               p.Info,
			"info.background":    p.Background,
			"info.border":        p.Info,
			"hint":               p.Type,
			"hint.background":    p.Background,
			"hint.border":        p.Type,
			"success":            p.Success,
			"success.background": p.SuccessBg,
			"success.border":     p.Success,

			// Git status
			"created":              p.Success,
			"created.background":   p.SuccessBg,
			"created.border":       p.Success,
			"deleted":              p.Error,
			"deleted.background":   p.ErrorBg,
			"deleted.border":       p.Error,
			"modified":             p.Warning,
			"modified.background":  p.Background,
			"modified.border":      p.Warning,
			"renamed":              p.Type,
			"renamed.background":   p.Background,
			"renamed.border":       p.Type,
			"conflict":             p.Function,
			"conflict.background":  p.ConflictBackground,
			"conflict.border":      p.Function,

			// Other states
			"ignored":                p.Comment,
			"ignored.background":     p.Background,
			"ignored.border":         p.Comment,
			"hidden":                 p.Comment,
			"hidden.background":      p.Background,
			"hidden.border":          p.Comment,
			"unreachable":            p.Comment,
			"unreachable.background": p.Background,
			"unreachable.border":     p.Comment,
			"predictive":             p.Comment,
			"predictive.background":  p.Background,
			"predictive.border":      p.TerminalPurple,

			// Players
			"players": generatePlayers(p),

			// Syntax highlighting
			"syntax": generateSyntax(p),
		},
	}

	// Add background appearance if specified
	if p.BackgroundAppearance != "" && p.BackgroundAppearance != "opaque" {
		style["style"].(map[string]any)["background.appearance"] = p.BackgroundAppearance
	}

	return style
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
			"color":       p.Type,
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
