# Tron Legacy Theme for Zed

## Overview

This repository contains a theme generator for the Tron Legacy theme family for the Zed editor. The generator creates multiple theme variants from a unified color system.

## Theme Variants

The theme includes four variants:
- **Tron Legacy** - Dark theme with opaque backgrounds
- **Tron Legacy Light** - Light theme with opaque backgrounds
- **Tron Legacy Frosted** - Dark theme with translucent backgrounds for glass effects
- **Tron Legacy Light Frosted** - Light theme with translucent backgrounds for glass effects

## Architecture

### Directory Structure
```
tools/
├── generate-theme.go      # Main theme generator
├── generate-theme_test.go # Tests for theme generation
├── palette/
│   ├── palette.go        # TronThemePalette struct definition
│   ├── generator.go      # Theme JSON generation logic
│   └── theme_structs.go  # Zed theme JSON structure definitions
├── dark/
│   ├── colors.css        # Dark color definitions
│   └── palette.go        # Dark palette mapping
├── light/
│   ├── colors.css        # Light color definitions
│   └── palette.go        # Light palette mapping
└── csscolors/
    └── loader.go         # CSS color parser and loader
```

### Three-Layer Architecture

The theme generator uses a three-layer architecture to separate concerns:

1. **Raw Colors Layer** (`colors.css`)
   - Non-semantic color values (e.g., `--gray900: #14191fff`)
   - Platform-specific color definitions
   - Alpha variants for transparency effects
   - No knowledge of how colors will be used

2. **Semantic Palette Layer** (`TronThemePalette`)
   - Maps raw colors to semantic purposes
   - Platform-agnostic color assignments
   - Examples: `Foreground`, `Background`, `Error`, `Success`
   - Provides consistent interface across all theme variants

3. **Theme Structure Layer** (`theme_structs.go`)
   - Defines exact Zed JSON structure
   - Enforces field ordering to match official themes
   - Maps semantic colors to specific Zed properties
   - Handles Zed-specific features (accents, syntax highlighting)

### Color System

1. **CSS Color Definitions** (`colors.css`)
   - Define raw color values as CSS variables
   - Colors are non-semantic (e.g., `--gray900`, `--blue200`)
   - Frosted variants include semi-transparent versions
   - Special alpha variants for UI effects
   - Embedded in Go files using `//go:embed` directive

2. **Semantic Palette** (`TronThemePalette`)
   - Central struct containing all semantic color assignments
   - Groups colors by purpose:
     - Core colors (text, backgrounds)
     - UI elements (borders, selections, highlights)
     - Semantic status colors (error, success, warning)
     - Syntax highlighting colors
     - Terminal colors
     - Special features (accents, player colors)
   - Each variant (dark/light/frosted) creates its own palette instance

3. **Theme Generation Process**
   ```
   colors.css → csscolors.LoadColors() → TronThemePalette → generator.GenerateTheme() → theme JSON
   ```
   - CSS colors are parsed into a ColorMap
   - Palette maps colors to semantic purposes
   - Generator transforms palette into Zed's exact JSON structure
   - Struct tags ensure proper JSON field ordering

### Key Design Decisions

1. **Separation of Concerns**
   - Colors (what values) separated from usage (where applied)
   - Platform-specific (Zed) separated from generic theme data
   - Visual variants share structure but not values

2. **Type Safety**
   - Go structs enforce valid theme properties
   - Compile-time validation of color assignments
   - No arbitrary fields allowed in output

3. **Official Theme Compliance**
   - Matches exact field ordering of Zed's official themes
   - Supports all standard syntax highlighting fields
   - Includes Gruvbox-specific features (accents, function.builtin)

4. **Maintainability**
   - Single source of truth for each color value
   - Easy to add new variants or adjust colors
   - Clear mapping from design intent to implementation

### Building

You can build the theme using either:
- Direct Go commands: `cd tools && go run generate-theme.go`
- Makefile: `make generate` (if available)

### Key Principles

- **Semantic Abstraction**: The `TronThemePalette` struct provides a semantic layer between raw colors and their usage
- **Variant Consistency**: All variants share the same semantic structure but with different color mappings
- **Frosted Glass Support**: Frosted variants use semi-transparent colors while maintaining readability
- **Color Pairing**: Frosted colors should visually pair with their opaque counterparts

### Testing

Run tests with: `cd tools && go test ./...`

The test suite includes:
- Unit tests for theme generation (`generate-theme_test.go`)
- Validation test (`palette/validate_structs_test.go`) that:
  - Performs a shallow clone of the official Zed repository
  - Validates our theme structs against One, Gruvbox, and Ayu themes
  - Ensures we have all required fields and no unexpected extras
  - Handles known exceptions for optional/theme-specific fields

The generated theme is output to: `themes/tron-legacy.json`

### Official Theme References

The official Zed themes can be found in the Zed repository at:
- **Repository**: `https://github.com/zed-industries/zed`
- **Theme Directory**: `assets/themes/`
- **Key Official Themes**:
  - `assets/themes/one/one.json` - One Dark/Light themes
  - `assets/themes/gruvbox/gruvbox.json` - Gruvbox variants (includes accents)
  - `assets/themes/ayu/ayu.json` - Ayu theme (different field ordering)

These themes serve as the reference for:
- Field ordering and structure
- Standard syntax highlighting fields
- Optional features (accents, function.builtin)
- Theme variant organization

**Important Notes on Theme Differences**:
- Official themes have slight structural differences
- **One/Gruvbox Pattern**: Includes `version_control` fields after `link_text.hover`
- **Ayu Pattern**: Omits `version_control` fields entirely, status colors immediately follow `link_text.hover`
- **Gruvbox Exclusives**: `accents` array and `function.builtin` syntax field
- **Priority**: Follow One/Gruvbox field ordering as the primary pattern, supporting all features found across official themes
