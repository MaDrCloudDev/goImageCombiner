package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

type Args struct {
	Image1Path string
	Image2Path string
	OutputPath string
}

func NewArgs() *Args {
	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments. Usage: program <image1_path> <image2_path> <output_filename>")
		os.Exit(1)
	}

	return &Args{
		Image1Path: os.Args[1],
		Image2Path: os.Args[2],
		OutputPath: os.Args[3],
	}
}

func main() {
	args := NewArgs()
	fmt.Printf("%+v\n", args)

	img1 := loadImage(args.Image1Path)
	img2 := loadImage(args.Image2Path)

	imgCombined := combineImages(img1, img2)

	outputPath := filepath.Join("images", args.OutputPath)

	saveImage(outputPath, imgCombined)
}

func loadImage(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening image: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		os.Exit(1)
	}

	return img
}

func combineImages(img1, img2 image.Image) image.Image {
	bounds1 := img1.Bounds()
	bounds2 := img2.Bounds()

	minWidth := bounds1.Dx()
	if bounds2.Dx() < minWidth {
		minWidth = bounds2.Dx()
	}

	minHeight := bounds1.Dy()
	if bounds2.Dy() < minHeight {
		minHeight = bounds2.Dy()
	}

	combined := image.NewRGBA(image.Rect(0, 0, minWidth, minHeight))

	for y := 0; y < minHeight; y++ {
		for x := 0; x < minWidth; x++ {
			if x%2 == 0 {
				combined.Set(x, y, img1.At(x, y))
			} else {
				combined.Set(x, y, img2.At(x, y))
			}
		}
	}

	return combined
}

func saveImage(path string, img image.Image) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error combining images: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		fmt.Printf("Error encoding image: %v\n", err)
		os.Exit(1)
	}
}