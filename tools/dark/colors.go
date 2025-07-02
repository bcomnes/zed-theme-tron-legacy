package dark

// Monochrome
const (
	black       string = "#000000ff"
	pureWhite   string = "#ffffffff"
	shadow      string = "#00000040"
	transparent string = "#00000000"
)

// Gray scale (dark to light)
const (
	gray900 string = "#14191fff" // Background
	gray800 string = "#1c2128ff" // LineHighlight (more neutral)
	gray750 string = "#23282fff" // Status bar (neutral dark gray)
	gray700 string = "#2a3039ff" // BorderSubtle (neutral)
	gray600 string = "#323842ff" // Border, Selection (neutral)
	gray500 string = "#647c9bff" // ForegroundDim
	gray400 string = "#586676ff" // Comment
	gray200 string = "#aec2e0ff" // Foreground
	gray50  string = "#dae3f1ff" // ForegroundBright
)

// Neutral UI grays (for non-blue UI elements)
const (
	neutral800 string = "#1a1d23ff" // Neutral highlight
	neutral600 string = "#2d3139ff" // Neutral border
)

// Blues
const (
	blue500       string = "#267fb5ff" // Primary blue
	blue400       string = "#2b6db9ff" // Bright blue
	blue300       string = "#4a95b3ff" // Cyan dim
	blue200       string = "#6ee2ffff" // Primary cyan
	blue200Bright string = "#c8d9e8ff" // Bright blue-fg (90% lightness)
	blue200Alpha  string = "#6ee2ff40" // Primary cyan with 25% opacity for search
)

// Greens
const (
	green700 string = "#144212ff" // Success bg
	green600 string = "#4d5f07ff" // Green dim
	green500 string = "#95CC5Eff" // Alt green
	green300 string = "#C7F026ff" // Primary green
)

// Yellows/Oranges
const (
	yellow500 string = "#FFE792ff" // Primary yellow
	yellow400 string = "#ffd12cff" // Dim yellow
	orange500 string = "#FFB20Dff" // Primary orange
	orange400 string = "#F79D1Eff" // Light orange
)

// Reds
const (
	red700 string = "#660000ff" // Error bg
	red500 string = "#F92672ff" // Error red
	red400 string = "#FF410Dff" // Primary red (string color)
	red300 string = "#FF5F52ff" // Light red
)

// Purples
const (
	purple500 string = "#967EFBff" // Primary purple
	pink500   string = "#ff79c6ff" // Primary pink
)

// Special alpha variants
const (
	blue500Alpha    string = "#267fb580" // Selection with alpha
	dropTarget      string = "#6ee2ff18" // Very low opacity cyan for drop targets
	neutral600Alpha string = "#647c9b4d" // Gray500 with 30% opacity for better contrast on scrollbar
	blue200Alpha50  string = "#6ee2ffb3" // Primary cyan with 70% opacity for scrollbar hover
	oneTrackBorder  string = "#2e333cff" // One theme's scrollbar track border color
)

// Additional alpha variants for better transparency
const (
	gray800Alpha75 string = "#1c2128bf" // Active line background with 75% opacity
	gray500Alpha30 string = "#647c9b40" // Wrap guide with 25% opacity using lighter gray
	gray400Alpha50 string = "#58667659" // Active wrap guide with 35% opacity using comment color
	blue200Alpha10 string = "#6ee2ff1a" // Document highlight read with 10% opacity
	blue200Alpha40 string = "#6ee2ff66" // Document highlight write with 40% opacity
)

// Standardized player selection opacity (24% like other themes)
const (
	player1Alpha string = "#267fb53d" // Blue with 24% opacity
	player3Alpha string = "#4d5f073d" // Green with 24% opacity
	player4Alpha string = "#F79D1E3d" // Orange with 24% opacity
)

// Frosted glass variants (80% opacity for subtle transparency)
const (
	gray900Frosted string = "#14191fcc" // Frosted background - 80% opacity
	gray200Frosted string = "#aec2e0ee" // Higher opacity text for frosted
	gray500Frosted string = "#647c9bf2" // Higher opacity dim text (95% for better sidebar contrast)

)

// UI element transparency for frosted
const (
	elevatedFrosted  string = "#1c2128ff" // Keep elevated surfaces opaque
	borderFrosted    string = "#323842aa" // Semi-transparent border using gray600 (67% opacity)
	borderSubtleFrosted string = "#32384266" // Subtle border for frosted (40% opacity)
	selectionFrosted string = "#2a303966" // Semi-transparent selection (40% opacity)
	statusBarFrosted string = "#14191fcc" // 80% opacity for status bar

	editorFrosted    string = "#14191fee" // 93% opacity for editor background
)

// Additional frosted UI elements
const (
	dropTargetFrosted      string = "#6ee2ff33" // Slightly more opaque drop target

	activeLineFrosted      string = "#1c212855" // More transparent active line
	docHighlightReadFrosted  string = "#6ee2ff22" // Slightly more visible read highlight
	docHighlightWriteFrosted string = "#6ee2ff44" // Slightly more visible write highlight

	titleBarInactiveFrosted string = "#1c2128cc" // Inactive title bar with 80% opacity
)
