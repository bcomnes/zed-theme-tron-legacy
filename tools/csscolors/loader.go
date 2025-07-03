package csscolors

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

// ColorMap holds CSS variable names mapped to their color values
type ColorMap map[string]string

// LoadColors parses an embedded CSS file and extracts all CSS custom properties (variables)
// It returns a map of variable names (without the -- prefix) to their color values
func LoadColors(cssContent []byte) (ColorMap, error) {
	colors := make(ColorMap)

	// Create a new CSS parser
	input := parse.NewInputBytes(cssContent)
	p := css.NewParser(input, false)

	// Track if we're inside a :root rule
	inRoot := false
	currentSelector := ""

	for {
		gt, _, data := p.Next()
		if gt == css.ErrorGrammar {
			if p.Err() != nil && p.Err().Error() != "EOF" && p.Err().Error() != "unexpected EOF" {
				return nil, fmt.Errorf("CSS parse error: %w", p.Err())
			}
			break
		}

		switch gt {
		case css.BeginRulesetGrammar:
			// The selector comes from the values, not the data
			currentSelector = ""
			values := p.Values()
			for _, v := range values {
				currentSelector += string(v.Data)
			}
			currentSelector = strings.TrimSpace(currentSelector)
			if currentSelector == ":root" {
				inRoot = true
			}

		case css.EndRulesetGrammar:
			inRoot = false
			currentSelector = ""

		case css.DeclarationGrammar:
			if inRoot {
				// Property name is in data
				propName := string(data)
				if strings.HasPrefix(propName, "--") {
					// Remove the -- prefix
					varName := propName[2:]

					// Get the value tokens
					values := p.Values()
					if len(values) > 0 {
						// Concatenate all value tokens, handling whitespace properly
						var value strings.Builder
						for i, v := range values {
							if i > 0 && v.TokenType == css.WhitespaceToken {
								// Skip whitespace between tokens
								continue
							}
							value.Write(v.Data)
						}

						// Clean up the value
						colorValue := strings.TrimSpace(value.String())

						// Remove trailing semicolon if present
						colorValue = strings.TrimSuffix(colorValue, ";")

						// Skip comment-only values
						if strings.HasPrefix(colorValue, "/*") && strings.HasSuffix(colorValue, "*/") {
							continue
						}

						// Store in map
						colors[varName] = colorValue
					}
				}
			}
		case css.CustomPropertyGrammar:
			// Also handle CustomPropertyGrammar type
			if inRoot {
				propName := string(data)
				if strings.HasPrefix(propName, "--") {
					varName := propName[2:]
					values := p.Values()
					if len(values) > 0 {
						var value strings.Builder
						for i, v := range values {
							if i > 0 && v.TokenType == css.WhitespaceToken {
								continue
							}
							value.Write(v.Data)
						}
						colorValue := strings.TrimSpace(value.String())
						colorValue = strings.TrimSuffix(colorValue, ";")
						if !strings.HasPrefix(colorValue, "/*") {
							colors[varName] = colorValue
						}
					}
				}
			}
		case css.TokenGrammar:
			// Skip tokens
			continue
		}
	}

	return colors, nil
}

// Get retrieves a color value by its variable name
func (cm ColorMap) Get(name string) (string, bool) {
	color, ok := cm[name]
	return color, ok
}

// MustGet retrieves a color value by its variable name, panics if not found
func (cm ColorMap) MustGet(name string) string {
	color, ok := cm[name]
	if !ok {
		panic(fmt.Sprintf("color variable '%s' not found in CSS", name))
	}
	return color
}
