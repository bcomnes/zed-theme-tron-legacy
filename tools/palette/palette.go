package palette

// TronThemePalette contains semantic color assignments for a theme variant
type TronThemePalette struct {
	// Core colors
	Background       string
	Foreground       string
	ForegroundDim    string
	ForegroundBright string

	// UI elements
	Border        string
	BorderSubtle  string
	BorderFocused string
	Selection     string
	Highlight     string
	ElementActive string
	ElementHover  string
	Shadow        string
	Transparent   string
	DropTargetBackground string
	TabBarBackground string
	StatusBarBackground string
	EditorBackground string
	EditorSubheaderBackground string
	TitleBarInactiveBackground string
	PanelStickyEntryBackground string

	// Semantic colors
	Error   string
	Warning string
	Success string
	Info    string
	Hint    string

	// Background variants for semantic colors
	ErrorBg   string
	SuccessBg string

	// Syntax highlighting - base colors
	Comment  string
	String   string
	Number   string
	Keyword  string
	Function string
	Type     string
	Variable string
	Property string

	// Syntax highlighting - accent colors
	StringEscape       string
	Regex              string
	Decorator          string
	Punctuation        string
	PunctuationSpecial string
	Link               string

	// Terminal base colors
	TerminalBlack  string
	TerminalRed    string
	TerminalGreen  string
	TerminalYellow string
	TerminalBlue   string
	TerminalPurple string
	TerminalCyan   string
	TerminalWhite  string

	// Additional semantic mappings
	LineNumber             string
	ElementActiveBright    string
	ConflictBackground     string
	TerminalBrightWhite    string
	TerminalDimGreen       string
	TerminalDimYellow      string
	VersionControlModified string
	VariableParam          string
	Constructor            string
	ClassType              string
	EnumType               string
	AttributeType          string
	SelectorAlt            string
	AnnotationType         string
	EmbeddedType           string

	// Search
	SearchHighlight string

	// Player selection colors
	Player1Selection string
	Player3Selection string
	Player4Selection string

	// Scrollbar alpha variants
	ScrollbarThumb      string
	ScrollbarThumbHover string
	ScrollbarTrackBorder string

	// Editor transparency variants
	ActiveLineBackground string
	WrapGuide string
	ActiveWrapGuide string
	DocumentHighlightRead string
	DocumentHighlightWrite string

	// Standardized player selection alphas
	Player1Alpha string
	Player3Alpha string
	Player4Alpha string

	// Frosted glass properties
	BackgroundAppearance string // "opaque" or "blurred"
	BackgroundFrosted string // Semi-transparent background for frosted effect
	TextFrosted string // Higher opacity text for frosted backgrounds
}
