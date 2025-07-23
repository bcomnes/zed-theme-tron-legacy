package palette

// TronThemePalette contains semantic color assignments for a theme variant
type TronThemePalette struct {
	// Core text colors
	Foreground       string
	ForegroundDim    string
	ForegroundBright string

	// Core backgrounds
	Background       string
	EditorBackground string

	// UI backgrounds
	TabBarBackground           string
	StatusBarBackground        string
	EditorSubheaderBackground  string
	TitleBarInactiveBackground string
	PanelOverlayBackground     string
	PanelOverlayHover          string

	// UI elements
	Border               string
	BorderSubtle         string
	BorderFocused        string
	Selection            string
	Highlight            string
	ElementHover         string
	Transparent          string
	DropTargetBackground string

	// Semantic status colors
	Error     string
	ErrorBg   string
	Warning   string
	Success   string
	SuccessBg string
	Info      string
	Hint      string

	// Editor specific
	LineNumber             string
	ActiveLineBackground   string
	WrapGuide              string
	ActiveWrapGuide        string
	DocumentHighlightRead  string
	DocumentHighlightWrite string
	SearchHighlight        string

	// Syntax - basic types
	Comment     string
	String      string
	StringEscape string
	Number      string
	Keyword     string
	Function    string
	Variable    string
	Property  string
	Namespace string

	// Syntax - advanced types
	Type          string
	Constructor   string
	EnumType      string
	AttributeType string
	EmbeddedType  string
	// SpecialSyntax is used for multiple syntax elements:
	// - label (decorators/annotations)
	// - selector.pseudo (CSS pseudo-selectors)
	// - string.special (special string literals)
	SpecialSyntax string
	Regex         string
	Selector      string

	// Syntax - punctuation
	Punctuation        string
	PunctuationSpecial string

	// Terminal colors
	TerminalBlack       string
	TerminalRed         string
	TerminalGreen       string
	TerminalYellow      string
	TerminalBlue        string
	TerminalPurple      string
	TerminalCyan        string
	TerminalWhite       string
	TerminalBrightWhite string
	TerminalDimGreen    string
	TerminalDimYellow   string
	TerminalDimBlack    string // Used for terminal.ansi.dim_black

	// Version control
	VersionControlModified string
	ConflictBackground     string

	// Scrollbar
	ScrollbarThumb       string
	ScrollbarThumbHover  string
	ScrollbarThumbActive string
	ScrollbarTrackBorder string

	// Player/collaboration colors
	Player1Alpha string
	Player2Alpha string
	Player3Alpha string
	Player4Alpha string

	// Frosted glass properties
	BackgroundAppearance string // "opaque" or "blurred"

	// Accent colors (from Gruvbox theme)
	// These are optional accent colors that can be selected by users for UI highlights
	Accents []string
}
