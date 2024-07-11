package main

import (
	"fmt"
	"log"

	"main/src/vector"
)

func main() {
	imageWidth, imageHeight := 256.0, 256.0

	fmt.Println("P3\n", imageWidth, imageHeight, "\n255")

	for i := 0.0; i < imageHeight; i++ {
		log.Println("Scanlines remaining:", (imageHeight - i))
		for j := 0.0; j < imageWidth; j++ {
			pixelColor := vector.New([3]float64{i / (imageWidth - 1.0), 0.0, j / (imageHeight - 1.0)})
			vector.WriteColor(pixelColor)
		}
	}

	log.Println("Done.")
}
