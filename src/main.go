package main

import (
	"fmt"
	"log"

	"main/src/color"
)

func main() {
	imageWidth, imageHeight := 256.0, 256.0

	fmt.Println("P3")
	fmt.Println(imageWidth, imageHeight)
	fmt.Println("255")

	for i := 0.0; i < imageHeight; i++ {
		log.Println("Scanlines remaining:", (imageHeight - i))
		for j := 0.0; j < imageWidth; j++ {
			pixelColor := color.New([3]float64{i / (imageWidth - 1.0), 0.0, j / (imageHeight - 1.0)})
			color.WriteColor(pixelColor)
		}
	}

	log.Println("Done.")
}
