# Tron Legacy Theme Conversion Plan

## Color Conversion Requirements

### 1. HSL to Hex Conversions
- **white**: `hsl(0, 0%, 100%)` → needs conversion to hex

### 2. Alpha Channel Conversions
- **black-shadow**: `color(var(black) alpha(0.25))`
  - Base: `#000000`
  - Alpha: 0.25 (25%)
  - Result: `#00000040`

### 3. Lightness Adjustments
These require HSL manipulation:

- **green-light-subdue**: `color(var(green-light) l(20%))`
  - Base: `#C7F026`
  - Set lightness to 20%

- **blue-bg-highlight**: `color(var(blue-bg) l(20%))`
  - Base: `#14191f`
  - Set lightness to 20%

- **blue-fg-white**: `color(var(blue-fg) l(90%))`
  - Base: `#aec2e0`
  - Set lightness to 90%

- **blue-hilight-subdue**: `color(var(blue-highlight) l(20%))`
  - Base: `#183c66`
  - Set lightness to 20%

- **gray-copy-dark**: `color(var(gray-copy) l(50%))`
  - Base: `#748aa6`
  - Set lightness to 50%

- **yellow-highlight-subdue**: `color(var(yellow-highlight) l(- 20%))`
  - Base: `#FFE792`
  - Decrease lightness by 20%

### 4. Contrast Adjustments
These require WCAG contrast ratio calculations:

- **blue-highlight-contrast**: `color(var(blue-highlight) min-contrast(var(blue-highlight) 2.5))`
  - Base: `#183c66`
  - Ensure 2.5:1 contrast ratio against itself

- **gray-comment**: `color(#324357 min-contrast(var(blue-bg) 3))`
  - Base: `#324357`
  - Ensure 3:1 contrast ratio against `#14191f`

## Conversion Process

### Step 1: Simple Conversions
1. Convert HSL white to hex
2. Apply alpha to black-shadow

### Step 2: Lightness Adjustments
1. Convert hex to HSL
2. Adjust lightness component
3. Convert back to hex

### Step 3: Contrast Calculations
1. Calculate current contrast ratio
2. Adjust lightness until target ratio is met
3. Convert final color to hex

### Step 4: Generate Final Color Map
Create a complete mapping of all variable names to their final hex values.

## Tools Needed
- HSL to Hex converter
- Hex to HSL converter
- Alpha channel calculator
- WCAG contrast ratio calculator
- Lightness adjustment calculator

## Complete Color Inventory

### All Variables from Sublime Theme
```json
{
  "variables": {
    // Direct hex values (no conversion needed)
    "black": "#000000",
    "red-string": "#FF410D",
    "red-error": "#F92672",
    "red-error-background": "#660000",
    "green": "#95CC5E",
    "green-light": "#C7F026",
    "green-light-boost": "#A6E22E",
    "green-light-boost-background": "#144212",
    "yellow-highlight": "#FFE792",
    "orange-entity": "#FFB20D",
    "orange-entity-dark": "#F79D1E",
    "blue-pure": "#2F33AB",
    "blue-bg": "#14191f",
    "blue-fg": "#aec2e0",
    "blue-accent": "#6ee2ff",
    "blue-storage": "#267fb5",
    "blue-highlight": "#183c66",
    "gray-copy": "#748aa6",
    "purple-accent": "#967EFB",

    // Needs conversion
    "white": "hsl(0, 0%, 100%)",
    "black-shadow": "color(var(black) alpha(0.25))",
    "green-light-subdue": "color(var(green-light) l(20%))",
    "blue-bg-highlight": "color(var(blue-bg) l(20%))",
    "blue-fg-white": "color(var(blue-fg) l(90%))",
    "blue-highlight-contrast": "color(var(blue-highlight) min-contrast(var(blue-highlight) 2.5))",
    "blue-hilight-subdue": "color(var(blue-highlight) l(20%))",
    "gray-copy-dark": "color(var(gray-copy) l(50%))",
    "gray-comment": "color(#324357 min-contrast(var(blue-bg) 3))",
    "yellow-highlight-subdue": "color(var(yellow-highlight) l(- 20%))"
  }
}
```

### Usage in Sublime Theme
The following shows where each color variable is used:

**UI Elements (globals)**:
- `foreground`: var(blue-fg)
- `background`: var(blue-bg)
- `accent`: var(blue-accent)
- `caret`: var(yellow-highlight)
- `line_highlight`: var(blue-bg-highlight)
- `selection`: var(blue-highlight)
- `selection_border`: var(blue-highlight-contrast)
- `inactive_selection`: var(blue-hilight-subdue)
- `misspelling`: var(red-error)
- `shadow`: var(black-shadow)
- `active_guide`: var(green-light)
- `stack_guide`: var(green-light-subdue)
- `highlight`: var(yellow-highlight-subdue)
- `find_highlight_foreground`: var(blue-bg)
- `find_highlight`: var(yellow-highlight)
- `brackets_foreground`: var(blue-accent)
- `bracket_contents_foreground`: var(green-light)
- `tags_foreground`: var(green-light)
- `line_diff_added`: var(green-light)
- `line_diff_modified`: var(yellow-highlight-subdue)
- `line_diff_deleted`: var(red-error)

**Syntax Highlighting**:
- Comments: var(gray-comment)
- Strings: var(red-string)
- Numbers: var(green-light)
- Constants: var(orange-entity)
- Variables: var(blue-fg-white)
- Keywords: var(gray-copy)
- Operators: var(gray-copy)
- Functions: var(orange-entity)
- Classes: var(orange-entity), var(orange-entity-dark)
- Storage: var(white), var(blue-storage)
- Language variables: var(purple-accent)
- And many more...

## Final Output Format
```json
{
  "color_map": {
    "white": "#ffffff",
    "black": "#000000",
    "black-shadow": "#00000040",
    "red-string": "#FF410D",
    // ... all other colors with final hex values
  }
}
```
