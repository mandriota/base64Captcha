package base64Captcha

import (
	"embed"
	"testing"
)

// sources:
// fonts/3Dumb.ttf (142.224kB)
// fonts/ApothecaryFont.ttf (62.08kB)
// fonts/Comismsh.ttf (80.132kB)
// fonts/DENNEthree-dee.ttf (83.188kB)
// fonts/DeborahFancyDress.ttf (32.52kB)
// fonts/Flim-Flam.ttf (140.576kB)
// fonts/RitaSmith.ttf (31.24kB)
// fonts/actionj.ttf (34.944kB)
// fonts/chromohv.ttf (45.9kB)
// fonts/readme.md (162B)

//go:embed fonts/*.ttf
var fonts embed.FS

var embeddedFonts = NewEmbeddedFontsStorage(fonts)

func Test_loadFontByName(t *testing.T) {
	f := embeddedFonts.LoadFontByName("fonts/Comismsh.ttf")
	if f == nil {
		t.Error("failed")
	}

	defer recoverPanic(t)
	_ = embeddedFonts.LoadFontByName("fonts/readme.md")
}
func recoverPanic(t *testing.T) {
	r := recover()
	if r == nil {
		t.Error("not trigger panic")
	}
}

func Test_loadFontsByNames(t *testing.T) {

	fs := embeddedFonts.LoadFontsByNames([]string{"fonts/Comismsh.ttf"})
	if len(fs) != 1 {
		t.Error("failed")
	}
	defer recoverPanic(t)
	embeddedFonts.LoadFontsByNames([]string{"fonts/actionj.txxxxxtf"})
}
