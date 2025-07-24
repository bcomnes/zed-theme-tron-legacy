package dark

import (
	_ "embed"
	"testing"

	"github.com/bcomnes/zed-theme-tron-legacy/tools/utils/colorvalidation"
)

//go:embed palette.go
var paletteGoContent []byte

func TestAllColorsUsed(t *testing.T) {
	colorvalidation.ValidateAllColorsUsed(t, colorsCSS, paletteGoContent)
}

func TestNoDuplicateColorValues(t *testing.T) {
	colorvalidation.ValidateNoDuplicateColorValues(t, colorsCSS)
}
