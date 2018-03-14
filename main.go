package main

import (
	"bufio"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	// open "test.jpg"
	if _, err := os.Stat("./input"); os.IsNotExist(err) {
		err = os.Mkdir("input", 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		err = os.Mkdir("output", 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("- OPTIMAZE - Image optimizer by Sharif")
	fmt.Println("Put your JPG image/s in 'input' folder")
	fmt.Println("width,height,quality,")
	text, _ := reader.ReadString('\n')

	p := strings.Split(text, ",")

	w64, _ := strconv.ParseUint(string(p[0]), 10, 64)
	h64, _ := strconv.ParseUint(string(p[1]), 10, 64)
	q64, _ := strconv.ParseInt(string(p[2]), 10, 64)

	width := uint(w64)
	height := uint(h64)
	quality := int(q64)

	files, err := ioutil.ReadDir("./input/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		//fmt.Println(f.Name())
		input := "./input/"
		input += f.Name()
		fmt.Println(f.Name())

		file, err := os.Open(input)
		if err != nil {
			log.Fatal(err)
		}

		// decode jpeg into image.Image
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		var opt jpeg.Options

		opt.Quality = quality

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(width, height, img, resize.Lanczos3)

		output := "./output/"
		output += f.Name()

		out, err := os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		// write new image to file
		jpeg.Encode(out, m, &opt)
	}

}
