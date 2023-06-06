package base64Captcha

import (
	"math/rand"

	"github.com/golang/freetype/truetype"
)

var fontsSimple = DefaultEmbeddedFonts.LoadFontsByNames([]string{
	"fonts/3Dumb.ttf",
	"fonts/ApothecaryFont.ttf",
	"fonts/Comismsh.ttf",
	"fonts/DENNEthree-dee.ttf",
	"fonts/DeborahFancyDress.ttf",
	"fonts/Flim-Flam.ttf",
	"fonts/RitaSmith.ttf",
	"fonts/actionj.ttf",
	"fonts/chromohv.ttf",
})

// var fontemoji = loadFontByName("fonts/seguiemj.ttf")
var fontsAll = fontsSimple

// randFontFrom choose random font family.选择随机的字体
func randFontFrom(fonts []*truetype.Font) *truetype.Font {
	fontCount := len(fonts)

	if fontCount == 0 {
		//loading default fonts
		fonts = fontsAll
		fontCount = len(fontsAll)
	}
	index := rand.Intn(fontCount)
	return fonts[index]
}
