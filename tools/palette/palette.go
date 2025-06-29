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
	Shadow        string

	// Semantic colors
	Error   string
	Warning string
	Success string
	Info    string

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
}
