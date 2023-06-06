package base64Captcha

import "embed"

// defaultEmbeddedFontsFS Built-in font storage.
//
//go:embed fonts/*.ttf
var defaultEmbeddedFontsFS embed.FS

var DefaultEmbeddedFonts = NewEmbeddedFontsStorage(defaultEmbeddedFontsFS)
