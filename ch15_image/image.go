package image

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

// save the image
func save(filePath string, grid [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	img := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			img.Set(x, y, grid[x][y])
		}
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer file.Close()
	png.Encode(file, img.SubImage(img.Rect))
}

// load the image
func load(filePath string) (grid [][]color.Color) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot read file:", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}
	size := img.Bounds().Size()
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		grid = append(grid, y)
	}
	return
}

//flip the image
func flip(grid [][]color.Color) {
	for x := 0; x < len(grid); x++ {
		col := grid[x]
		for y := 0; y < len(col)/2; y++ {
			k := len(col) - y - 1
			col[y], col[k] = col[k], col[y]
		}
	}
}

// convert image to grayscale
func grayscale(grid [][]color.Color) (grayImg [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	grayImg = make([][]color.Color, xlen)
	for i := 0; i < len(grayImg); i++ {
		grayImg[i] = make([]color.Color, ylen)
	}

	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			pix := grid[x][y].(color.NRGBA)
			gray := uint8(float64(pix.R)/3.0 + float64(pix.G)/3.0 + float64(pix.B)/3.0)
			grayImg[x][y] = color.NRGBA{gray, gray, gray, pix.A}
		}
	}
	return
}

// resize the image
func resize(grid [][]color.Color, scale float64) (resized [][]color.Color) {
	xlen, ylen := int(float64(len(grid))*scale), int(float64(len(grid[0]))*scale)
	resized = make([][]color.Color, xlen)
	for i := 0; i < len(resized); i++ {
		resized[i] = make([]color.Color, ylen)
	}
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			xp := int(math.Floor(float64(x) / scale))
			yp := int(math.Floor(float64(y) / scale))
			resized[x][y] = grid[xp][yp]
		}
	}
	return
}
