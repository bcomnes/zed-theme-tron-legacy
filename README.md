# Tron Legacy Theme for Zed

A dark theme for Zed editor inspired by the visual aesthetics of Tron Legacy, originally created for Sublime Text by Hypermodules LLC and Bret Comnes.

![Theme Preview](./screenshots/tron-dark.webp)
![Theme Preview](./screenshots/tron-light.webp)

## Features

- Dark blue background with cyan accents reminiscent of the Tron universe
- High contrast syntax highlighting optimized for readability
- Carefully selected colors that maintain WCAG contrast ratios
- Support for all Zed UI elements and syntax tokens

## Color Palette

### Primary Colors
- **Background**: `#14191f` - Deep blue-black
- **Foreground**: `#aec2e0` - Light blue-gray
- **Accent**: `#6ee2ff` - Bright cyan

### Syntax Colors
- **Strings**: `#FF410D` - Bright red-orange
- **Numbers**: `#C7F026` - Bright green
- **Functions**: `#FFB20D` - Golden orange
- **Keywords**: `#748aa6` - Muted blue-gray
- **Comments**: `#6684a7` - Light gray-blue (adjusted for readability)
- **Types**: `#267fb5` - Medium blue
- **Variables**: `#dae3f1` - Very light blue

## Installation

### From Zed Extensions

1. Open Zed
2. Press `Cmd+Shift+P` (macOS) or `Ctrl+Shift+P` (Linux/Windows)
3. Type "extensions" and select "zed: extensions"
4. Search for "Tron Legacy"
5. Click "Install"

### Manual Installation

1. Clone this repository
2. Open Zed's extensions directory:
   - macOS: `~/Library/Application Support/Zed/extensions/`
   - Linux: `~/.config/zed/extensions/`
3. Create a symlink or copy this folder to the extensions directory
4. Restart Zed
5. Select "Tron Legacy" from the theme picker (Cmd/Ctrl+K, Cmd/Ctrl+T)

## Development

This theme was generated using a custom Go tool that converts Sublime Text color schemes to Zed's theme format.

### Building from Source

```bash
cd tools

# Generate the theme
go run generate-theme.go
```

### Project Structure
```
zed-theme-tron-legacy/
├── README.md                    # This file
├── extension.toml              # Extension metadata
├── LICENSE                     # MIT License
├── themes/
│   └── tron-legacy.json        # The Zed theme file
└── tools/                      # Theme generation tools
    ├── generate-theme.go       # Theme generator
    └── go.mod                  # Go module file
```

## Color Palette

The theme uses a carefully crafted color palette inspired by the Tron Legacy aesthetic:

- **Background**: Deep blue-black (`#14191f`)
- **Foreground**: Light blue-gray (`#aec2e0`)
- **Accent**: Bright cyan (`#6ee2ff`)
- **Strings**: Bright red-orange (`#FF410D`)
- **Numbers**: Bright green (`#C7F026`)
- **Functions**: Golden orange (`#FFB20D`)
- **Keywords**: Muted blue-gray (`#748aa6`)
- **Comments**: Light gray-blue (`#6684a7`) - Adjusted for WCAG AA contrast
- **Types**: Medium blue (`#267fb5`)
- **Variables**: Very light blue (`#dae3f1`)

The comment color was specifically adjusted from the original to ensure proper readability with a 4.58:1 contrast ratio against the background.

## Credits

- Original Sublime Text theme by [Hypermodules LLC](https://github.com/hypermodules) and [Bret Comnes](https://github.com/bcomnes)
- Ported to Zed by [Bret Comnes](https://github.com/bcomnes)

## License

MIT License - See [LICENSE](LICENSE) file for details
