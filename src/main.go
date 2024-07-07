package main

import (
	"fmt"
)

func main() {
	imageWidth, imageHeight := 256.0, 256.0

	fmt.Println("P3\n", imageWidth, imageHeight, "\n255")

	for i := 0.0; i < imageHeight; i++ {
		for j := 0.0; j < imageWidth; j++ {
			r := i / (imageWidth - 1.0)
			g := j / (imageHeight - 1.0)
			b := 0.0

			var ir = int(255.999 * r)
			var ig = int(255.999 * g)
			var ib = int(255.999 * b)

			fmt.Println(ir, ig, ib)
		}
	}
}
