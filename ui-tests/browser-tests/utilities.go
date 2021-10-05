package uitest

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

func generateByteFile(name string, size int) string {
	// convert size from MB to bytes
	size *= 1000000
	data := make([]byte, size)

	filePath := fmt.Sprintf("./testdata/%s", name)
	f, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	defer f.Close()
	_, err = f.Write(data)

	if err != nil {
		panic(err)
	}

	file, _ := filepath.Abs(filePath)
	return file
}

func generateImage(name string) string {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	shade := color.RGBA{200, 80, 188, 0x29}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2:
				img.Set(x, y, shade)
			case x >= width/2 && y >= height/2:
				img.Set(x, y, color.White)
			default:
			}
		}
	}

	filePath := fmt.Sprintf("./testdata/%s", name)
	f, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	png.Encode(f, img)

	file, _ := filepath.Abs(filePath)
	return file
}

func removeFiles(files []string) {
	for _, file := range files {
		err := os.Remove(fmt.Sprintf("./testdata/%s", file))

		if err != nil {
			panic(err)
		}
	}
}
