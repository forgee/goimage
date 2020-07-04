package goimage

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"os"

	"golang.org/x/image/draw"
)

type Image struct {
	rgba *image.RGBA
}

// NewImage creates and returns a new goimage image
func NewImage(width, height int) Image {
	return Image{
		rgba: image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

// AddImage adds an image to the main goimage
func (g Image) AddImage(path string, x, y int) error {
	// Open image
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	// Close file handle
	defer file.Close()

	// Decode image
	img, err := png.Decode(file)

	if err != nil {
		return err
	}

	// Draw given image
	draw.Draw(g.rgba, image.Rect(x, y, x+img.Bounds().Dx(), y+img.Bounds().Dy()), img, image.ZP, draw.Src)

	return nil
}

// AddUniformImage adds an uniform image with the given color and size
func (g Image) AddUniformImage(width, height int, color color.Color, x, y int) {
	draw.Draw(
		g.rgba,
		image.Rect(x, y, x+width, y+height),
		image.NewUniform(color),
		image.ZP,
		draw.Src,
	)
}

// Save saves the goImage to the given path
func (g Image) Save(path string) error {
	// Open the destination file
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	// Close file handle
	defer file.Close()

	// Encode goImage to the file
	return png.Encode(file, g.rgba)
}

// Encode encodes the image into the given writer
func (g Image) Encode(dst io.Writer) error {
	return png.Encode(dst, g.rgba)
}
