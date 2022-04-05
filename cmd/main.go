package main

import (
	"fmt"
	"image-generator/pkg"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	width  = 480
	height = 480
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		img := pkg.GeneratePic(width, height)
		fileName := fmt.Sprintf("pics/pic%v.png", i)
		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		if err := png.Encode(file, img); err != nil {
			log.Fatal(err)
		}
	}
}
