package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strconv"

	"github.com/masonelmore/quad/quad"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: quad <filename>.png tolerance")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	img, err := png.Decode(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tolerance, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	out := quad.Quadify(img, tolerance)
	saveImage(out, "out.png")
}

func saveImage(i *image.RGBA, filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	png.Encode(f, i)
}
