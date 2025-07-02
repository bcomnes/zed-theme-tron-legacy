package light

// Monochrome
const (
	black       string = "#000000ff"
	pureWhite   string = "#ffffffff"
	shadow      string = "#00000020"
	transparent string = "#00000000"
)

// Gray scale (light to dark) - inverted from dark theme
const (
	gray50  string = "#f5f7faff" // Original gray background
	gray100 string = "#e8ecf2ff" // LineHighlight
	gray150 string = "#dfe5edff" // Status bar
	gray200 string = "#d1dae6ff" // BorderSubtle
	gray300 string = "#b8c5d6ff" // Border, Selection
	gray400 string = "#8a9db5ff" // Line numbers, etc
	gray500 string = "#6b7e96ff" // Comment
	gray600 string = "#526073ff" // ForegroundDim (sidebar text)
	gray700 string = "#2d3e4fff" // Foreground
	gray900 string = "#14191fff" // ForegroundBright
)

// Blues - adjusted for light mode
const (
	blue600      string = "#1a5f8aff" // Primary blue (darker for light bg)
	blue500      string = "#267fb5ff" // Standard blue
	blue400      string = "#3988c0ff" // Medium blue
	blue300      string = "#4a95b3ff" // Light blue
	blue200      string = "#0099ccff" // Primary cyan (darker for light bg)
	blue200Alpha string = "#0099cc30" // Primary cyan with alpha for search
)

// Greens - adjusted for light mode
const (
	green100 string = "#e6f7e3ff" // Success bg
	green600 string = "#3a5f00ff" // Dark green
	green500 string = "#5a8b2cff" // Alt green
	green400 string = "#7aad3aff" // Primary green (darker for light bg)
)

// Yellows/Oranges - adjusted for light mode
const (
	yellow600 string = "#c9a000ff" // Primary yellow (darker for light bg)
	yellow500 string = "#dbb200ff" // Medium yellow
	orange600 string = "#cc7700ff" // Primary orange (darker for light bg)
	orange500 string = "#e68a00ff" // Light orange
	orange700 string = "#b35900ff" // Even darker orange for better contrast on modified
)

// Reds - adjusted for light mode
const (
	red100 string = "#ffe6e6ff" // Error bg
	red600 string = "#cc0033ff" // Error red (darker for light bg)
	red500 string = "#d91e18ff" // Primary red (string color)
	red400 string = "#e74c3cff" // Light red
)

// Purples - adjusted for light mode
const (
	purple600 string = "#6a56ccff" // Primary purple (darker for light bg)
	pink600   string = "#d1459aff" // Primary pink (darker for light bg)
)

// Special alpha variants
const (
	blue300Alpha    string = "#4a95b340" // Selection with alpha
	dropTarget      string = "#0099cc18" // Very low opacity cyan for drop targets
	gray300Alpha    string = "#6b7e964d" // Gray500 with 30% opacity for better contrast on scrollbar
	blue400Alpha    string = "#0099ccb3" // Primary cyan with 70% opacity for scrollbar hover
	oneTrackBorder  string = "#2e333cff" // One theme's scrollbar track border color
	gray300AlphaDark string = "#6b7e9666" // Gray500 with 40% opacity for better balance on frosted
)

// Additional alpha variants for better transparency
const (
	gray100Alpha75 string = "#e8ecf2bf" // Active line background with 75% opacity
	gray500Alpha30 string = "#6b7e9640" // Wrap guide with 25% opacity using darker gray
	gray600Alpha50 string = "#52607359" // Active wrap guide with 35% opacity using even darker gray
	blue200Alpha10 string = "#0099cc1a" // Document highlight read with 10% opacity
	blue200Alpha40 string = "#0099cc66" // Document highlight write with 40% opacity
)

// Standardized player selection opacity (24% like other themes)
const (
	player1Alpha string = "#4a95b33d" // Blue with 24% opacity
	player3Alpha string = "#3a5f003d" // Green with 24% opacity
	player4Alpha string = "#e68a003d" // Orange with 24% opacity
)

// Frosted glass variants (85% opacity for subtle transparency)
const (
	gray50Frosted  string = "#f5f7fad9" // Frosted background - 85% opacity
	gray700Frosted string = "#2d3e4fee" // Higher opacity text for frosted
	gray600Frosted string = "#526073dd" // Higher opacity dim text
)

// UI element transparency for frosted
const (
	elevatedFrosted   string = "#e8ecf2ff" // Keep elevated surfaces opaque
	borderFrosted     string = "#b8c5d666" // Semi-transparent border
	selectionFrosted  string = "#d1dae666" // Semi-transparent selection
	statusBarFrosted  string = "#e8ecf2dd" // 87% opacity for status bar
	editorFrosted     string = "#f5f7faf2" // 95% opacity for editor background
)

// Additional frosted UI adjustments
const (
	dropTargetFrosted      string = "#0099cc33" // Slightly more opaque drop target
	activeLineFrosted      string = "#e8ecf255" // More transparent active line
	docHighlightReadFrosted  string = "#0099cc22" // Slightly more visible read highlight
	docHighlightWriteFrosted string = "#0099cc44" // Slightly more visible write highlight
	titleBarInactiveFrosted string = "#e8ecf2cc" // Inactive title bar with 80% opacity
)
