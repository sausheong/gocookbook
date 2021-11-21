package image

import (
	"errors"
	"os"
	"testing"
)

func TestLoadImage(t *testing.T) {
	testImageFile("monalisa.png", t)
}

func TestSaveImage(t *testing.T) {
	savedFileName := "testSaveImage.png"
	grid := load("monalisa.png")
	save(savedFileName, grid)
	testFileExists(savedFileName, t)
	testImageFile(savedFileName, t)
	// cleanup
	t.Cleanup(func() {
		os.Remove(savedFileName)
	})
}

func TestFlip(t *testing.T) {
	grid := load("monalisa.png")
	p1 := grid[0][0]
	flip(grid)
	p2 := grid[0][479]
	if p1 != p2 {
		t.Fatal("Pixels not flipped")
	}
	save("flipped.png", grid)
	testImageFile("flipped.png", t)
	t.Cleanup(func() {
		os.Remove("flipped.png")
	})
}

func TestGrayscale(t *testing.T) {
	grid := load("monalisa.png")
	gray := grayscale(grid)

	save("grayscale.png", gray)
	testImageFile("grayscale.png", t)
	t.Cleanup(func() {
		os.Remove("grayscale.png")
	})
}

func TestResize(t *testing.T) {
	var scale float64 = 2
	grid := load("monalisa.png")
	resized := resize(grid, scale)
	if len(resized) != int(321*scale) || len(resized[0]) != int(480*scale) {
		t.Error("Resized grid is wrong size", "width:", len(resized), "length:", len(resized[0]))
	}
	save("resized.png", resized)
	t.Cleanup(func() {
		os.Remove("resized.png")
	})
}

// test if image is loaded and is correct size
func testImageFile(path string, t *testing.T) {
	grid := load(path)
	if len(grid) != 321 || len(grid[0]) != 480 {
		t.Error("Grid is wrong size", "width:", len(grid), "length:", len(grid[0]))
	}
}

// test if file exists
func testFileExists(path string, t *testing.T) {
	_, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		t.Fatal("Image not created")
	}
}
