package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/nfnt/resize"
)

func main() {
	files := getAllPngsInDir(filepath.Join("static", "images", "full"))
	fmt.Println("found ", len(files), " files")
	for _, filename := range files {
		fmt.Println("resizing file ", filename)
		file, err := os.Open(filepath.Join("static", "images", "full", filename))
		if err != nil {
			log.Fatal(err)
		}

		// decode png into image.Image
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// resize to width 350 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(350, 0, img, resize.Lanczos3)

		out, err := os.Create(filepath.Join("static", "images", "thumbs", filename))
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		png.Encode(out, m)
	}

}

func getAllPngsInDir(dirname string) []string {
	var files []string
	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()
	fi, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, fi := range fi {
		if fi.Mode().IsRegular() {
			if path.Ext(fi.Name()) == ".png" {
				files = append(files, fi.Name())
			}
		}
	}
	return files
}
