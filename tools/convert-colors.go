package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
)

// Color represents an RGBA color
type Color struct {
	R, G, B, A float64
}

// ColorConversion holds the conversion data
type ColorConversion struct {
	Name     string `json:"name"`
	Original string `json:"original"`
	Hex      string `json:"hex"`
}

// hslToRgb converts HSL values to RGB
// h: 0-360, s: 0-1, l: 0-1
func hslToRgb(h, s, l float64) (r, g, b float64) {
	h = h / 360.0 // normalize to 0-1

	if s == 0 {
		// achromatic
		r, g, b = l, l, l
		return
	}

	hue2rgb := func(p, q, t float64) float64 {
		if t < 0 {
			t += 1
		}
		if t > 1 {
			t -= 1
		}
		if t < 1.0/6.0 {
			return p + (q-p)*6*t
		}
		if t < 1.0/2.0 {
			return q
		}
		if t < 2.0/3.0 {
			return p + (q-p)*(2.0/3.0-t)*6
		}
		return p
	}

	var q float64
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q

	r = hue2rgb(p, q, h+1.0/3.0)
	g = hue2rgb(p, q, h)
	b = hue2rgb(p, q, h-1.0/3.0)

	return
}

// rgbToHsl converts RGB values to HSL
// r, g, b: 0-1
func rgbToHsl(r, g, b float64) (h, s, l float64) {
	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	l = (max + min) / 2

	if max == min {
		h, s = 0, 0 // achromatic
		return
	}

	d := max - min
	if l > 0.5 {
		s = d / (2 - max - min)
	} else {
		s = d / (max + min)
	}

	switch max {
	case r:
		h = (g - b) / d
		if g < b {
			h += 6
		}
	case g:
		h = (b-r)/d + 2
	case b:
		h = (r-g)/d + 4
	}
	h /= 6
	h *= 360 // convert to degrees

	return
}

// hexToRgb converts hex color to RGB values (0-1)
func hexToRgb(hex string) (r, g, b float64, err error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %s", hex)
	}

	var ri, gi, bi int
	_, err = fmt.Sscanf(hex, "%02x%02x%02x", &ri, &gi, &bi)
	if err != nil {
		return 0, 0, 0, err
	}

	r = float64(ri) / 255.0
	g = float64(gi) / 255.0
	b = float64(bi) / 255.0
	return
}

// rgbToHex converts RGB values (0-1) to hex string
func rgbToHex(r, g, b float64) string {
	ri := int(math.Round(r * 255))
	gi := int(math.Round(g * 255))
	bi := int(math.Round(b * 255))
	return fmt.Sprintf("#%02x%02x%02x", ri, gi, bi)
}

// rgbaToHex converts RGBA values (0-1) to hex string with alpha
func rgbaToHex(r, g, b, a float64) string {
	ri := int(math.Round(r * 255))
	gi := int(math.Round(g * 255))
	bi := int(math.Round(b * 255))
	ai := int(math.Round(a * 255))
	return fmt.Sprintf("#%02x%02x%02x%02x", ri, gi, bi, ai)
}

// relativeLuminance calculates the relative luminance of a color
func relativeLuminance(r, g, b float64) float64 {
	// Convert to linear RGB
	linearize := func(c float64) float64 {
		if c <= 0.03928 {
			return c / 12.92
		}
		return math.Pow((c+0.055)/1.055, 2.4)
	}

	r = linearize(r)
	g = linearize(g)
	b = linearize(b)

	return 0.2126*r + 0.7152*g + 0.0722*b
}

// contrastRatio calculates the contrast ratio between two colors
func contrastRatio(r1, g1, b1, r2, g2, b2 float64) float64 {
	l1 := relativeLuminance(r1, g1, b1)
	l2 := relativeLuminance(r2, g2, b2)

	lighter := math.Max(l1, l2)
	darker := math.Min(l1, l2)

	return (lighter + 0.05) / (darker + 0.05)
}

// adjustColorForContrast adjusts a color to meet minimum contrast ratio against background
func adjustColorForContrast(r, g, b, bgR, bgG, bgB, minRatio float64) (float64, float64, float64) {
	h, s, l := rgbToHsl(r, g, b)

	// Try adjusting lightness in both directions
	bestL := l
	bestRatio := contrastRatio(r, g, b, bgR, bgG, bgB)

	// Try lighter values
	for testL := l; testL <= 1.0; testL += 0.01 {
		testR, testG, testB := hslToRgb(h, s, testL)
		ratio := contrastRatio(testR, testG, testB, bgR, bgG, bgB)
		if ratio >= minRatio && ratio < bestRatio*1.5 { // Don't go too far
			bestL = testL
			bestRatio = ratio
			break
		}
	}

	// If we didn't find a good lighter value, try darker
	if bestRatio < minRatio {
		for testL := l; testL >= 0.0; testL -= 0.01 {
			testR, testG, testB := hslToRgb(h, s, testL)
			ratio := contrastRatio(testR, testG, testB, bgR, bgG, bgB)
			if ratio >= minRatio {
				bestL = testL
				break
			}
		}
	}

	return hslToRgb(h, s, bestL)
}

func main() {
	conversions := []ColorConversion{}

	// Convert white from HSL
	r, g, b := hslToRgb(0, 0, 1.0)
	conversions = append(conversions, ColorConversion{
		Name:     "white",
		Original: "hsl(0, 0%, 100%)",
		Hex:      rgbToHex(r, g, b),
	})

	// Black shadow with alpha
	conversions = append(conversions, ColorConversion{
		Name:     "black-shadow",
		Original: "color(var(black) alpha(0.25))",
		Hex:      rgbaToHex(0, 0, 0, 0.25),
	})

	// Green light subdue - lightness 20%
	r, g, b, _ = hexToRgb("#C7F026")
	h, s, _ := rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, 0.20)
	conversions = append(conversions, ColorConversion{
		Name:     "green-light-subdue",
		Original: "color(var(green-light) l(20%))",
		Hex:      rgbToHex(r, g, b),
	})

	// Blue bg highlight - lightness 20%
	r, g, b, _ = hexToRgb("#14191f")
	h, s, _ = rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, 0.20)
	conversions = append(conversions, ColorConversion{
		Name:     "blue-bg-highlight",
		Original: "color(var(blue-bg) l(20%))",
		Hex:      rgbToHex(r, g, b),
	})

	// Blue fg white - lightness 90%
	r, g, b, _ = hexToRgb("#aec2e0")
	h, s, _ = rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, 0.90)
	conversions = append(conversions, ColorConversion{
		Name:     "blue-fg-white",
		Original: "color(var(blue-fg) l(90%))",
		Hex:      rgbToHex(r, g, b),
	})

	// Blue highlight contrast - min contrast 2.5 against itself
	// This is a special case - we need to make it lighter/darker than itself
	r, g, b, _ = hexToRgb("#183c66")
	h, s, l := rgbToHsl(r, g, b)
	// Make it lighter to create contrast
	r2, g2, b2 := hslToRgb(h, s, math.Min(l+0.2, 1.0))
	conversions = append(conversions, ColorConversion{
		Name:     "blue-highlight-contrast",
		Original: "color(var(blue-highlight) min-contrast(var(blue-highlight) 2.5))",
		Hex:      rgbToHex(r2, g2, b2),
	})

	// Blue highlight subdue - lightness 20%
	r, g, b, _ = hexToRgb("#183c66")
	h, s, _ = rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, 0.20)
	conversions = append(conversions, ColorConversion{
		Name:     "blue-hilight-subdue",
		Original: "color(var(blue-highlight) l(20%))",
		Hex:      rgbToHex(r, g, b),
	})

	// Gray copy dark - lightness 50%
	r, g, b, _ = hexToRgb("#748aa6")
	h, s, _ = rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, 0.50)
	conversions = append(conversions, ColorConversion{
		Name:     "gray-copy-dark",
		Original: "color(var(gray-copy) l(50%))",
		Hex:      rgbToHex(r, g, b),
	})

	// Gray comment - min contrast 3.0 against blue-bg
	r, g, b, _ = hexToRgb("#324357")
	bgR, bgG, bgB, _ := hexToRgb("#14191f")
	r, g, b = adjustColorForContrast(r, g, b, bgR, bgG, bgB, 3.0)
	conversions = append(conversions, ColorConversion{
		Name:     "gray-comment",
		Original: "color(#324357 min-contrast(var(blue-bg) 3))",
		Hex:      rgbToHex(r, g, b),
	})

	// Yellow highlight subdue - decrease lightness by 20%
	r, g, b, _ = hexToRgb("#FFE792")
	h, s, l = rgbToHsl(r, g, b)
	r, g, b = hslToRgb(h, s, math.Max(l-0.20, 0))
	conversions = append(conversions, ColorConversion{
		Name:     "yellow-highlight-subdue",
		Original: "color(var(yellow-highlight) l(- 20%))",
		Hex:      rgbToHex(r, g, b),
	})

	// Create the final color map
	colorMap := map[string]string{
		// Direct hex values
		"black":                        "#000000",
		"red-string":                   "#FF410D",
		"red-error":                    "#F92672",
		"red-error-background":         "#660000",
		"green":                        "#95CC5E",
		"green-light":                  "#C7F026",
		"green-light-boost":            "#A6E22E",
		"green-light-boost-background": "#144212",
		"yellow-highlight":             "#FFE792",
		"orange-entity":                "#FFB20D",
		"orange-entity-dark":           "#F79D1E",
		"blue-pure":                    "#2F33AB",
		"blue-bg":                      "#14191f",
		"blue-fg":                      "#aec2e0",
		"blue-accent":                  "#6ee2ff",
		"blue-storage":                 "#267fb5",
		"blue-highlight":               "#183c66",
		"gray-copy":                    "#748aa6",
		"purple-accent":                "#967EFB",
	}

	// Add converted values
	for _, conv := range conversions {
		colorMap[conv.Name] = conv.Hex
		fmt.Printf("%s: %s -> %s\n", conv.Name, conv.Original, conv.Hex)
	}

	// Output as JSON
	output := map[string]interface{}{
		"conversions": conversions,
		"color_map":   colorMap,
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	// Write to file
	err = os.WriteFile("converted-colors.json", jsonData, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nConverted colors saved to converted-colors.json")
}
