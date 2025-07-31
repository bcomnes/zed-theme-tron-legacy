package palette

// TronThemePalette contains semantic color assignments for a theme variant
type TronThemePalette struct {
	// ==========================================================================
	// Visual Hierarchy Layers
	// ==========================================================================

	// Background Layer (furthest back)
	Background         string // Primary background for editor and terminals
	EditorBackground   string // Editor-specific background (can differ for frosted variants)
	BackgroundElevated string // Elevated surfaces (was Highlight)
	BackgroundOverlay  string // Overlay/popup backgrounds (was PanelOverlayBackground)
	BackgroundOverlayHover string // Hover state for overlays (was PanelOverlayHover)

	// Midground Layer (UI chrome)
	Surface           string // UI surface (was TabBarBackground)
	SurfaceHighlight  string // Active/selected surface (was Background again)
	EditorSubheader   string // Editor subheader background (breadcrumbs, etc.)
	Statusbar         string // Status/title bars (was StatusBarBackground)
	StatusbarInactive string // Inactive title bar (was TitleBarInactiveBackground)

	// Foreground Layer (content)
	Foreground       string // Primary text and icons
	ForegroundMuted  string // Secondary/disabled text (was ForegroundDim)
	ForegroundStrong string // Emphasized text (was ForegroundBright)

	// ==========================================================================
	// Interactive Elements
	// ==========================================================================

	// Borders
	Border        string // Standard borders
	BorderSubtle  string // Subtle/variant borders
	BorderFocused string // Focused/accent borders

	// Selection & Highlights
	Selection              string // Text selection background
	SelectionAlpha         string // Semi-transparent selection (was ElementSelectionBackground)
	ActiveLine             string // Active line highlight (was ActiveLineBackground)
	MatchHighlight         string // Search/find matches (was SearchHighlight)
	DocumentHighlight      string // Symbol references (was DocumentHighlightRead)
	DocumentHighlightWrite string // Symbol write references

	// Interactive States
	Interactive       string // Interactive element background (was ElementHover)
	DropTarget        string // Drop target indicators (was DropTargetBackground)
	Transparent       string // Fully transparent

	// ==========================================================================
	// Semantic Colors
	// ==========================================================================

	// Status Colors
	// Pattern: Base colors = FOREGROUND colors for text/borders, Surface colors = BACKGROUND colors for surfaces
	Error          string // Errors, deletions - FOREGROUND color for text/borders
	ErrorSurface   string // Error backgrounds - BACKGROUND color for surfaces/panels
	Warning        string // Warnings, modifications - FOREGROUND color for text/borders
	Success        string // Success, additions - FOREGROUND color for text/borders
	SuccessSurface string // Success backgrounds - BACKGROUND color for surfaces/panels
	Info           string // Information - FOREGROUND color for text/borders
	Hint           string // Hint color (subtle info) - FOREGROUND color
	Accent         string // Primary accent - FOREGROUND color for text/borders
	UIAccent       string // UI accent for active elements, conflict markers - FOREGROUND color

	// Editor Guidelines
	GuideNormal string // Wrap guide, indent guide (was WrapGuide)
	GuideActive string // Active wrap guide (was ActiveWrapGuide)
	LineNumber  string // Line numbers, invisibles

	// ==========================================================================
	// Syntax Highlighting
	// ==========================================================================

	// Comments & Documentation
	Comment string

	// Literals
	String       string
	StringEscape string
	Number       string

	// Language Constructs
	Keyword   string
	Function  string
	Variable  string
	Type      string
	Property  string
	Namespace string

	// Special Syntax Elements
	Constructor   string
	Enum          string // Was EnumType
	Attribute     string // Was AttributeType
	Embedded      string // Was EmbeddedType
	Decorator     string // Was SpecialSyntax - used for labels, decorators
	Regex         string
	Tag           string // Was under Type, now separate
	SpecialVariable string // For self/this/super keywords

	// Punctuation
	Punctuation        string
	PunctuationMuted   string // Was PunctuationSpecial

	// ==========================================================================
	// Terminal Colors (ANSI)
	// ==========================================================================

	// Standard 16 colors
	TerminalBlack        string
	TerminalRed          string
	TerminalGreen        string
	TerminalYellow       string
	TerminalBlue         string
	TerminalPurple       string
	TerminalCyan         string
	TerminalWhite        string
	TerminalBrightBlack  string
	TerminalBrightRed    string
	TerminalBrightGreen  string
	TerminalBrightYellow string
	TerminalBrightBlue   string
	TerminalBrightPurple string
	TerminalBrightCyan   string
	TerminalBrightWhite  string

	// Dim variants (less common)
	TerminalDimGreen  string
	TerminalDimYellow string
	TerminalDimBlack  string
	TerminalDimRed    string
	TerminalDimBlue   string
	TerminalDimCyan   string
	TerminalDimWhite  string
	TerminalDimMagenta string

	// ==========================================================================
	// Version Control
	// ==========================================================================
	// Note: Git diffs use FOREGROUND colors because they overlay on syntax-highlighted text

	VCSModified string // Git diff modified lines - used as FOREGROUND color
	VCSConflict string // Git conflict backgrounds - used as BACKGROUND color

	// ==========================================================================
	// UI Components
	// ==========================================================================

	// Scrollbar
	ScrollbarThumb       string
	ScrollbarThumbHover  string
	ScrollbarThumbActive string
	ScrollbarTrackBorder string

	// Collaboration/Players
	Player1       string // Was Player1Alpha
	Player2       string // Was Player2Alpha
	Player3       string // Was Player3Alpha
	Player4       string // Was Player4Alpha

	// ==========================================================================
	// Theme Properties
	// ==========================================================================

	BackgroundAppearance string // "opaque" or "blurred"
	Accents              []string // User-selectable accent colors
}
