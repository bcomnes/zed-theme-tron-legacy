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
├── palette/
│   ├── palette.go        # TronThemePalette struct definition
│   └── generator.go      # Theme JSON generation logic
├── dark/
│   ├── colors.css        # Dark color definitions
│   └── palette.go        # Dark palette mapping
└── light/
    ├── colors.css        # Light color definitions
    └── palette.go        # Light palette mapping
```

### Color System

1. **CSS Color Definitions** (`colors.css`)
   - Define raw color values as CSS variables
   - Colors are non-semantic (e.g., `--gray900`, `--blue200`)
   - Frosted variants include semi-transparent versions
   - Special alpha variants for UI effects

2. **Semantic Palette** (`TronThemePalette`)
   - Maps CSS colors to semantic meanings
   - Provides consistent interface across all variants
   - Examples: `Background`, `Foreground`, `Error`, `Success`

3. **Theme Generation**
   - CSS colors are loaded and parsed
   - Mapped to semantic palette values
   - Generated into Zed-compatible JSON format

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

The generated theme is output to: `themes/tron-legacy.json`
