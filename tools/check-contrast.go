package main

import (
	"fmt"
	"math"
	"strings"
)

// hexToRgb converts hex color to RGB values (0-1)
func hexToRgb(hex string) (r, g, b float64, err error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 && len(hex) != 8 {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %s", hex)
	}

	var ri, gi, bi int
	_, err = fmt.Sscanf(hex[:6], "%02x%02x%02x", &ri, &gi, &bi)
	if err != nil {
		return 0, 0, 0, err
	}

	r = float64(ri) / 255.0
	g = float64(gi) / 255.0
	b = float64(bi) / 255.0
	return
}

// rgbToHsl converts RGB values to HSL
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

// hslToRgb converts HSL values to RGB
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

// rgbToHex converts RGB values (0-1) to hex string
func rgbToHex(r, g, b float64) string {
	ri := int(math.Round(r * 255))
	gi := int(math.Round(g * 255))
	bi := int(math.Round(b * 255))
	return fmt.Sprintf("#%02x%02x%02xff", ri, gi, bi)
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

func main() {
	// Colors to check
	bgColor := "#14191f"
	commentColor := "#324357"
	originalCommentColor := "#324357"

	// Get RGB values
	bgR, bgG, bgB, _ := hexToRgb(bgColor)
	commentR, commentG, commentB, _ := hexToRgb(commentColor)

	// Calculate current contrast ratio
	currentRatio := contrastRatio(commentR, commentG, commentB, bgR, bgG, bgB)

	fmt.Printf("Current contrast ratio: %.2f:1\n", currentRatio)
	fmt.Printf("Background: %s\n", bgColor)
	fmt.Printf("Original comment: %s\n", originalCommentColor)
	fmt.Printf("Current comment: %s\n", commentColor)

	// WCAG requirements:
	// AA normal text: 4.5:1
	// AA large text: 3:1
	// AAA normal text: 7:1
	// AAA large text: 4.5:1

	targetRatio := 4.5 // AA normal text

	if currentRatio < targetRatio {
		fmt.Printf("\nContrast too low! Need at least %.1f:1 for WCAG AA\n", targetRatio)

		// Try to fix by adjusting lightness
		h, s, l := rgbToHsl(commentR, commentG, commentB)
		fmt.Printf("\nOriginal HSL: H=%.0f, S=%.2f, L=%.2f\n", h, s, l)

		// Try increasing lightness
		bestL := l
		bestRatio := currentRatio

		for testL := l; testL <= 1.0; testL += 0.01 {
			testR, testG, testB := hslToRgb(h, s, testL)
			ratio := contrastRatio(testR, testG, testB, bgR, bgG, bgB)

			if ratio >= targetRatio {
				bestL = testL
				bestRatio = ratio
				break
			}
		}

		// Generate the improved color
		newR, newG, newB := hslToRgb(h, s, bestL)
		newHex := rgbToHex(newR, newG, newB)

		fmt.Printf("\nImproved comment color: %s\n", newHex)
		fmt.Printf("New HSL: H=%.0f, S=%.2f, L=%.2f\n", h, s, bestL)
		fmt.Printf("New contrast ratio: %.2f:1\n", bestRatio)

		// Also try some pre-defined good comment colors for dark themes
		fmt.Println("\nAlternative comment colors with good contrast:")

		alternatives := []string{
			"#5c6873", // Slightly lighter gray-blue
			"#6b7a8a", // Medium gray-blue
			"#7a8b9d", // Light gray-blue
			"#8899aa", // Even lighter
			"#889cb1", // VSCode-like comment
		}

		for _, alt := range alternatives {
			altR, altG, altB, _ := hexToRgb(alt)
			altRatio := contrastRatio(altR, altG, altB, bgR, bgG, bgB)
			fmt.Printf("%s - Contrast: %.2f:1\n", alt, altRatio)
		}

		// Try adjusting saturation too
		fmt.Println("\nTrying with reduced saturation:")
		for satMult := 1.0; satMult >= 0.3; satMult -= 0.1 {
			testS := s * satMult
			testR, testG, testB := hslToRgb(h, testS, bestL)
			testHex := rgbToHex(testR, testG, testB)
			testRatio := contrastRatio(testR, testG, testB, bgR, bgG, bgB)
			fmt.Printf("S=%.2f: %s - Contrast: %.2f:1\n", testS, testHex, testRatio)
		}

	} else {
		fmt.Printf("\nContrast is good! Meets WCAG AA standards.\n")
	}

	// Also check against the lighter background color
	bgHighlight := "#28323e"
	bgHighR, bgHighG, bgHighB, _ := hexToRgb(bgHighlight)
	highlightRatio := contrastRatio(commentR, commentG, commentB, bgHighR, bgHighG, bgHighB)
	fmt.Printf("\nContrast against highlighted line background (%s): %.2f:1\n", bgHighlight, highlightRatio)
}
