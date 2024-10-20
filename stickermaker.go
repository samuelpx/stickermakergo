package main

import (
	"fmt"
	"image"
	"image/png"
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


	resizedImg := resize.Resize(targetWidth, targetHeight, img, resize.Lanczos3)


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
