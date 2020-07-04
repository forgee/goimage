package goimage

import (
	"os"
	_ "image/png"
	"image"
	"golang.org/x/image/draw"
	"github.com/nfnt/resize"
)

// SetBackGroundImage adds an image to the main goImage with the
// goImage bounds
func (g *Image) SetBackGroundImage(path string) error {
	// Open background image
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	// Close image handle
	defer file.Close()

	// Decode image
	img, _, err := image.Decode(file)

	if err != nil {
		return err
	}

	// Resize image if needed
	if img.Bounds() != g.rgba.Bounds() {
		img = resize.Resize(uint(g.rgba.Bounds().Dx()), uint(g.rgba.Bounds().Dy()), img, resize.Lanczos3)
	}

	// Draw background image to goImage
	draw.Draw(g.rgba, g.rgba.Bounds(), img, image.ZP, draw.Src)

	return nil
}
