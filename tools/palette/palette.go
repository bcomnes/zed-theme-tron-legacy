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
	ElementActive        string
	ElementActiveBright  string
	ElementHover         string
	Shadow               string
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
	VariableParam string
	Property    string
	Link        string

	// Syntax - advanced types
	Type           string
	Constructor    string
	ClassType      string
	EnumType       string
	AttributeType  string
	AnnotationType string
	EmbeddedType   string
	Decorator      string
	Regex          string
	SelectorAlt    string

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

	// Version control
	VersionControlModified string
	ConflictBackground     string

	// Scrollbar
	ScrollbarThumb       string
	ScrollbarThumbHover  string
	ScrollbarTrackBorder string

	// Player/collaboration colors
	Player1Selection string
	Player1Alpha     string
	Player3Selection string
	Player3Alpha     string
	Player4Selection string
	Player4Alpha     string

	// Frosted glass properties
	BackgroundAppearance string // "opaque" or "blurred"
	BackgroundFrosted    string // Semi-transparent background for frosted effect
	TextFrosted          string // Higher opacity text for frosted backgrounds
}
