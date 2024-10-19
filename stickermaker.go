package main

import (
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize" // External library for image resizing
)

const (
	targetWidth  = 512
	targetHeight = 512
)

func main() {
	// Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Scan directory for jpg and png files and run processImage if it is an image
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
	// Open the image file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Resize the image to fit sticker specifications
	resizedImg := resize.Resize(targetWidth, targetHeight, img, resize.Lanczos3)

	// Create a new file for the resized image
	outFile, err := os.Create(path + "_resized.png")
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Encode the resized image as PNG
	err = png.Encode(outFile, resizedImg)
	if err != nil {
		return err
	}

	fmt.Println("Resized image saved to:", outFile.Name())
	return nil
}
