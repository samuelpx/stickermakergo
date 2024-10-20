package main

import (
	"fmt"
	"image"
	 "image/png"
    _ "image/jpeg"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

const (
	targetWidth  = 512
	targetHeight = 512
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

    
	err = filepath.WalkDir(dir, checkImage)
	if err != nil {
		fmt.Println("Error walking directory:", err)
	}
}

func checkImage(path string, info fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() && !strings.Contains(info.Name(), "resized") && (filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".jpeg" || filepath.Ext(path) == ".png") {
		fmt.Println("Processing:", path)
		err := processImage(path)
		if err != nil {
			fmt.Println("Error processing image:", err)
		}
	}
	return nil
}

func processImage(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

    
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
    w, h := getRatio(img)

    fmt.Printf("%d is the width of %s !\n%d is the height of %s\n",
    w, path, h, path)


	resizedImg := resize.Resize(w, h, img, resize.Lanczos3)


	outFile, err := os.Create(path + "_resized.png")
	if err != nil {
		return err
	}
	defer outFile.Close()


	err = png.Encode(outFile, resizedImg)
	if err != nil {
		return err
	}

	fmt.Println("Resized image saved to:", outFile.Name())
	return nil
}

func getRatio(img image.Image) (uint, uint) {
    w := uint(img.Bounds().Dx())
    h := uint(img.Bounds().Dy())
    if w == h {
        w = 512
        h = 512
    } else if w > h  {
        ratio := 512 / float64(w)
        w = uint(float64(w) * ratio)
        h = uint(float64(h) * ratio)
    } else if h > w  {
        ratio := 512 / float64(h)
        w = uint(float64(w) * ratio)
        h = uint(float64(h) * ratio)
    }

    return w, h
}
