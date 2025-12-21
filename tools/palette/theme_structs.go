package palette

// Theme represents the complete theme structure
type Theme struct {
	Schema string  `json:"$schema"`
	Author string  `json:"author"`
	Name   string  `json:"name"`
	Themes []Style `json:"themes"`
}

// Style represents a single theme style variant
type Style struct {
	Name       string      `json:"name"`
	Appearance string      `json:"appearance"`
	Accents    []string    `json:"accents,omitempty"`
	Style      *ThemeStyle `json:"style"`
}

// ThemeStyle represents the style properties in the exact order of official themes
type ThemeStyle struct {
	// Borders (exact order from official themes)
	Border            string `json:"border"`
	BorderVariant     string `json:"border.variant"`
	BorderFocused     string `json:"border.focused"`
	BorderSelected    string `json:"border.selected"`
	BorderTransparent string `json:"border.transparent"`
	BorderDisabled    string `json:"border.disabled"`

	// Surfaces
	ElevatedSurfaceBackground string `json:"elevated_surface.background"`
	SurfaceBackground         string `json:"surface.background"`
	Background                string `json:"background"`
	BackgroundAppearance      string `json:"background.appearance,omitempty"`

	// Elements
	ElementBackground    string `json:"element.background"`
	ElementHover         string `json:"element.hover"`
	ElementActive        string `json:"element.active"`
	ElementSelected      string `json:"element.selected"`
	ElementDisabled      string `json:"element.disabled"`
	DropTargetBackground string `json:"drop_target.background"`

	// Ghost elements
	GhostElementBackground string `json:"ghost_element.background"`
	GhostElementHover      string `json:"ghost_element.hover"`
	GhostElementActive     string `json:"ghost_element.active"`
	GhostElementSelected   string `json:"ghost_element.selected"`
	GhostElementDisabled   string `json:"ghost_element.disabled"`

	// Text
	Text            string `json:"text"`
	TextMuted       string `json:"text.muted"`
	TextPlaceholder string `json:"text.placeholder"`
	TextDisabled    string `json:"text.disabled"`
	TextAccent      string `json:"text.accent"`

	// Icons
	Icon            string `json:"icon"`
	IconMuted       string `json:"icon.muted"`
	IconDisabled    string `json:"icon.disabled"`
	IconPlaceholder string `json:"icon.placeholder"`
	IconAccent      string `json:"icon.accent"`

	// Bars
	StatusBarBackground        string `json:"status_bar.background"`
	TitleBarBackground         string `json:"title_bar.background"`
	TitleBarInactiveBackground string `json:"title_bar.inactive_background"`
	ToolbarBackground          string `json:"toolbar.background"`
	TabBarBackground           string `json:"tab_bar.background"`
	TabInactiveBackground      string `json:"tab.inactive_background"`
	TabActiveBackground        string `json:"tab.active_background"`

	// Search
	SearchMatchBackground string `json:"search.match_background"`

	// Panels
	PanelBackground        string `json:"panel.background"`
	PanelFocusedBorder     string `json:"panel.focused_border"`
	PanelOverlayBackground string `json:"panel.overlay_background,omitempty"`
	PanelOverlayHover      string `json:"panel.overlay_hover,omitempty"`
	PaneFocusedBorder      string `json:"pane.focused_border"`

	// Scrollbar
	ScrollbarThumbBackground      string `json:"scrollbar.thumb.background"`
	ScrollbarThumbHoverBackground string `json:"scrollbar.thumb.hover_background"`
	ScrollbarThumbBorder          string `json:"scrollbar.thumb.border"`
	ScrollbarTrackBackground      string `json:"scrollbar.track.background"`
	ScrollbarTrackBorder          string `json:"scrollbar.track.border"`

	// Editor
	EditorForeground                       string `json:"editor.foreground"`
	EditorBackground                       string `json:"editor.background"`
	EditorGutterBackground                 string `json:"editor.gutter.background"`
	EditorSubheaderBackground              string `json:"editor.subheader.background"`
	EditorActiveLineBackground             string `json:"editor.active_line.background"`
	EditorHighlightedLineBackground        string `json:"editor.highlighted_line.background"`
	EditorLineNumber                       string `json:"editor.line_number"`
	EditorActiveLineNumber                 string `json:"editor.active_line_number"`
	EditorHoverLineNumber                  string `json:"editor.hover_line_number"`
	EditorSelectionBackground              string `json:"editor.selection.background,omitempty"`
	EditorInvisible                        string `json:"editor.invisible"`
	EditorWrapGuide                        string `json:"editor.wrap_guide"`
	EditorActiveWrapGuide                  string `json:"editor.active_wrap_guide"`
	EditorDocumentHighlightReadBackground  string `json:"editor.document_highlight.read_background"`
	EditorDocumentHighlightWriteBackground string `json:"editor.document_highlight.write_background"`

	// Terminal
	TerminalBackground        string `json:"terminal.background"`
	TerminalForeground        string `json:"terminal.foreground"`
	TerminalBrightForeground  string `json:"terminal.bright_foreground"`
	TerminalDimForeground     string `json:"terminal.dim_foreground"`
	TerminalAnsiBlack         string `json:"terminal.ansi.black"`
	TerminalAnsiBrightBlack   string `json:"terminal.ansi.bright_black"`
	TerminalAnsiDimBlack      string `json:"terminal.ansi.dim_black"`
	TerminalAnsiRed           string `json:"terminal.ansi.red"`
	TerminalAnsiBrightRed     string `json:"terminal.ansi.bright_red"`
	TerminalAnsiDimRed        string `json:"terminal.ansi.dim_red"`
	TerminalAnsiGreen         string `json:"terminal.ansi.green"`
	TerminalAnsiBrightGreen   string `json:"terminal.ansi.bright_green"`
	TerminalAnsiDimGreen      string `json:"terminal.ansi.dim_green"`
	TerminalAnsiYellow        string `json:"terminal.ansi.yellow"`
	TerminalAnsiBrightYellow  string `json:"terminal.ansi.bright_yellow"`
	TerminalAnsiDimYellow     string `json:"terminal.ansi.dim_yellow"`
	TerminalAnsiBlue          string `json:"terminal.ansi.blue"`
	TerminalAnsiBrightBlue    string `json:"terminal.ansi.bright_blue"`
	TerminalAnsiDimBlue       string `json:"terminal.ansi.dim_blue"`
	TerminalAnsiMagenta       string `json:"terminal.ansi.magenta"`
	TerminalAnsiBrightMagenta string `json:"terminal.ansi.bright_magenta"`
	TerminalAnsiDimMagenta    string `json:"terminal.ansi.dim_magenta"`
	TerminalAnsiCyan          string `json:"terminal.ansi.cyan"`
	TerminalAnsiBrightCyan    string `json:"terminal.ansi.bright_cyan"`
	TerminalAnsiDimCyan       string `json:"terminal.ansi.dim_cyan"`
	TerminalAnsiWhite         string `json:"terminal.ansi.white"`
	TerminalAnsiBrightWhite   string `json:"terminal.ansi.bright_white"`
	TerminalAnsiDimWhite      string `json:"terminal.ansi.dim_white"`

	// Links
	LinkTextHover string `json:"link_text.hover"`

	// Version control
	VersionControlAdded                string `json:"version_control.added"`
	VersionControlModified             string `json:"version_control.modified"`
	VersionControlDeleted              string `json:"version_control.deleted"`
	VersionControlConflictMarkerOurs   string `json:"version_control.conflict_marker.ours,omitempty"`
	VersionControlConflictMarkerTheirs string `json:"version_control.conflict_marker.theirs,omitempty"`

	// Status colors (in exact order)
	Conflict           string `json:"conflict"`
	ConflictBackground string `json:"conflict.background"`
	ConflictBorder     string `json:"conflict.border"`

	Created           string `json:"created"`
	CreatedBackground string `json:"created.background"`
	CreatedBorder     string `json:"created.border"`

	Deleted           string `json:"deleted"`
	DeletedBackground string `json:"deleted.background"`
	DeletedBorder     string `json:"deleted.border"`

	Error           string `json:"error"`
	ErrorBackground string `json:"error.background"`
	ErrorBorder     string `json:"error.border"`

	Foreground string `json:"foreground,omitempty"`

	Hidden           string `json:"hidden"`
	HiddenBackground string `json:"hidden.background"`
	HiddenBorder     string `json:"hidden.border"`

	Hint           string `json:"hint"`
	HintBackground string `json:"hint.background"`
	HintBorder     string `json:"hint.border"`

	Ignored           string `json:"ignored"`
	IgnoredBackground string `json:"ignored.background"`
	IgnoredBorder     string `json:"ignored.border"`

	Info           string `json:"info"`
	InfoBackground string `json:"info.background"`
	InfoBorder     string `json:"info.border"`

	Modified           string `json:"modified"`
	ModifiedBackground string `json:"modified.background"`
	ModifiedBorder     string `json:"modified.border"`

	Predictive           string `json:"predictive"`
	PredictiveBackground string `json:"predictive.background"`
	PredictiveBorder     string `json:"predictive.border"`

	Renamed           string `json:"renamed"`
	RenamedBackground string `json:"renamed.background"`
	RenamedBorder     string `json:"renamed.border"`

	Success           string `json:"success"`
	SuccessBackground string `json:"success.background"`
	SuccessBorder     string `json:"success.border"`

	Unreachable           string `json:"unreachable"`
	UnreachableBackground string `json:"unreachable.background"`
	UnreachableBorder     string `json:"unreachable.border"`

	Warning           string `json:"warning"`
	WarningBackground string `json:"warning.background"`
	WarningBorder     string `json:"warning.border"`

	// Players and Syntax
	Players []Player     `json:"players"`
	Syntax  SyntaxStyles `json:"syntax"`

	// ========== NEWLY ADDED FIELDS (2024) ==========
	// These fields were added to match the latest Zed theme schema
	// They are optional and will be omitted if empty
	// Move them to their correct position if they are found in official themes one day

	// Selection Background
	ElementSelectionBackground string `json:"element.selection_background,omitempty"`

	// Indent Guides
	PanelIndentGuide        string `json:"panel.indent_guide,omitempty"`
	PanelIndentGuideHover   string `json:"panel.indent_guide_hover,omitempty"`
	PanelIndentGuideActive  string `json:"panel.indent_guide_active,omitempty"`
	EditorIndentGuide       string `json:"editor.indent_guide,omitempty"`
	EditorIndentGuideActive string `json:"editor.indent_guide_active,omitempty"`

	// Editor Debugger
	EditorDebuggerActiveLineBackground       string `json:"editor.debugger_active_line.background,omitempty"`
	EditorDocumentHighlightBracketBackground string `json:"editor.document_highlight.bracket_background,omitempty"`

	// Scrollbar Active State
	ScrollbarThumbActiveBackground string `json:"scrollbar.thumb.active_background,omitempty"`

	// Minimap
	MinimapThumbBackground       string `json:"minimap.thumb.background,omitempty"`
	MinimapThumbHoverBackground  string `json:"minimap.thumb.hover_background,omitempty"`
	MinimapThumbActiveBackground string `json:"minimap.thumb.active_background,omitempty"`
	MinimapThumbBorder           string `json:"minimap.thumb.border,omitempty"`

	// Terminal ANSI Background
	TerminalAnsiBackground string `json:"terminal.ansi.background,omitempty"`

	// Additional Version Control
	VersionControlRenamed  string `json:"version_control.renamed,omitempty"`
	VersionControlConflict string `json:"version_control.conflict,omitempty"`
	VersionControlIgnored  string `json:"version_control.ignored,omitempty"`

	// Pane Group
	PaneGroupBorder string `json:"pane_group.border,omitempty"`

	// Debugger Accent
	DebuggerAccent string `json:"debugger.accent,omitempty"`
	// ========== END NEWLY ADDED FIELDS ==========
}

// Player represents a multiplayer cursor color scheme
type Player struct {
	Cursor     string `json:"cursor"`
	Background string `json:"background"`
	Selection  string `json:"selection"`
}

// SyntaxStyle represents a single syntax highlighting style
type SyntaxStyle struct {
	Color      string  `json:"color"`
	FontStyle  *string `json:"font_style"`
	FontWeight *int    `json:"font_weight"`
}

// SyntaxStyles represents all syntax highlighting styles
type SyntaxStyles struct {
	Attribute             SyntaxStyle `json:"attribute"`
	Boolean               SyntaxStyle `json:"boolean"`
	Comment               SyntaxStyle `json:"comment"`
	CommentDoc            SyntaxStyle `json:"comment.doc"`
	Constant              SyntaxStyle `json:"constant"`
	Constructor           SyntaxStyle `json:"constructor"`
	Embedded              SyntaxStyle `json:"embedded"`
	Emphasis              SyntaxStyle `json:"emphasis"`
	EmphasisStrong        SyntaxStyle `json:"emphasis.strong"`
	Enum                  SyntaxStyle `json:"enum"`
	Function              SyntaxStyle `json:"function"`
	FunctionBuiltin       SyntaxStyle `json:"function.builtin"` // From Gruvbox theme
	Hint                  SyntaxStyle `json:"hint"`
	Keyword               SyntaxStyle `json:"keyword"`
	Label                 SyntaxStyle `json:"label"`
	LinkText              SyntaxStyle `json:"link_text"`
	LinkURI               SyntaxStyle `json:"link_uri"`
	Namespace             SyntaxStyle `json:"namespace"`
	Number                SyntaxStyle `json:"number"`
	Operator              SyntaxStyle `json:"operator"`
	Predictive            SyntaxStyle `json:"predictive"`
	Preproc               SyntaxStyle `json:"preproc"`
	Primary               SyntaxStyle `json:"primary"`
	Property              SyntaxStyle `json:"property"`
	Punctuation           SyntaxStyle `json:"punctuation"`
	PunctuationBracket    SyntaxStyle `json:"punctuation.bracket"`
	PunctuationDelimiter  SyntaxStyle `json:"punctuation.delimiter"`
	PunctuationListMarker SyntaxStyle `json:"punctuation.list_marker"`
	PunctuationSpecial    SyntaxStyle `json:"punctuation.special"`
	Selector              SyntaxStyle `json:"selector"`
	SelectorPseudo        SyntaxStyle `json:"selector.pseudo"`
	String                SyntaxStyle `json:"string"`
	StringEscape          SyntaxStyle `json:"string.escape"`
	StringRegex           SyntaxStyle `json:"string.regex"`
	StringSpecial         SyntaxStyle `json:"string.special"`
	StringSpecialSymbol   SyntaxStyle `json:"string.special.symbol"`
	Tag                   SyntaxStyle `json:"tag"`
	TextLiteral           SyntaxStyle `json:"text.literal"`
	Title                 SyntaxStyle `json:"title"`
	Type                  SyntaxStyle `json:"type"`
	Variable              SyntaxStyle `json:"variable"`
	VariableSpecial       SyntaxStyle `json:"variable.special"`
	Variant               SyntaxStyle `json:"variant"`

	// Diff/Patch syntax highlighting
	DiffPlus       SyntaxStyle `json:"diff.plus,omitempty"`
	DiffMinus      SyntaxStyle `json:"diff.minus,omitempty"`
	DiffDelta      SyntaxStyle `json:"diff.delta,omitempty"`
	DiffDeltaMoved SyntaxStyle `json:"diff.delta.moved,omitempty"`
}
