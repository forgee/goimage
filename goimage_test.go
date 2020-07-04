package goimage

import (
	"testing"
	"image/color"
)

func TestGoImage_SetBackGroundImage(t *testing.T) {
	// Create new goimage
	img := NewImage(200, 200)

	// Set background image
	if err := img.SetBackGroundImage("test/gopher.png"); err != nil {
		t.Errorf("Cannot set background image: %v", err)
		t.Fail()
	}

	// Save image
	if err := img.Save("test/test1.png"); err != nil {
		t.Errorf("Cannot save image: %v", err)
		t.Fail()
	}
}

func TestGoImage_AddImage(t *testing.T) {
	// Create new goimage
	img := NewImage(200, 200)

	// Add image
	if err := img.AddImage("test/gopher.png", 50, 60); err != nil {
		t.Errorf("Cannot add image: %v", err)
		t.Fail()
	}

	// Save image
	if err := img.Save("test/test2.png"); err != nil {
		t.Errorf("Cannot save image: %v", err)
		t.Fail()
	}
}

func TestGoImage_WriteText(t *testing.T) {
	// Create new goimage
	img := NewImage(200, 200)

	// Write text
	if err := img.WriteText("True gopher", color.RGBA{200, 70, 70, 255}, 14, 50, 60); err != nil {
		t.Errorf("Cannot write text to image: %v", err)
		t.Fail()
	}

	// Save image
	if err := img.Save("test/test3.png"); err != nil {
		t.Errorf("Cannot save image: %v", err)
		t.Fail()
	}
}

func TestGoImage_AddUniformImage(t *testing.T) {
	// Create new goimage
	img := NewImage(200, 200)

	// Add uniform image
	img.AddUniformImage(100, 100, color.RGBA{200, 70, 70, 255}, 20, 20)

	// Save image
	if err := img.Save("test/test4.png"); err != nil {
		t.Errorf("Cannot save image: %v", err)
		t.Fail()
	}
}