package goimage

import (
	"image"
	"image/color"
	"io/ioutil"

	"github.com/golang/freetype"
	"golang.org/x/image/font/gofont/goregular"
)

// WriteText writes the given text using the Go font
func (g Image) WriteText(txt string, color color.Color, size float64, x, y int) error {
	// Parse Go font
	font, err := freetype.ParseFont(goregular.TTF)

	if err != nil {
		return err
	}

	// Create freetype context
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(size)
	c.SetClip(g.rgba.Bounds())
	c.SetDst(g.rgba)
	c.SetSrc(image.NewUniform(color))

	// Draw text
	_, err = c.DrawString(txt, freetype.Pt(x, y))

	return err
}

// WriteTextFont writes the given text using the given font
func (g Image) WriteTextFont(fontFile, txt string, color color.Color, size float64, x, y int) error {
	// Load font
	data, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return err
	}

	// Parse font
	font, err := freetype.ParseFont(data)

	if err != nil {
		return err
	}

	// Create freetype context
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(size)
	c.SetClip(g.rgba.Bounds())
	c.SetDst(g.rgba)
	c.SetSrc(image.NewUniform(color))

	// Draw text
	_, err = c.DrawString(txt, freetype.Pt(x, y))

	return err
}
