package main

import (
	"fmt"
	"log"

	"main/src/color"
	"main/src/ray"
	"main/src/vec3"
)

func rayColor(r ray.Ray) color.Color {
	unitDirection := vec3.UnitVector(r.Direction())
	a := 0.5 * (unitDirection.Y() + 1.0)

	return color.ScaleColor(color.ScaleFloat(color.New([3]float64{1.0, 1.0, 1.0}), (1.0-a)),
		color.ScaleFloat(color.New([3]float64{0.5, 0.7, 1.0}), a))
}

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 400

	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := vec3.New([3]float64{0, 0, 0})

	viewportU := vec3.New([3]float64{viewportWidth, 0, 0})
	viewportV := vec3.New([3]float64{0, viewportHeight, 0})

	pixelDeltaU := vec3.ScaleFloat(*viewportU, 1/float64(imageWidth))
	pixelDeltaV := vec3.ScaleFloat(*viewportV, 1/float64(imageHeight))

	viewportUpperLeft := vec3.Sum(*cameraCenter, *vec3.New([3]float64{0, 0, focalLength}))
	viewportUpperLeft = vec3.Sum(viewportUpperLeft, vec3.ScaleFloat(*viewportU, 1.0/2.0))
	viewportUpperLeft = vec3.Sum(viewportUpperLeft, vec3.ScaleFloat(*viewportV, 1.0/2.0))

	pixel00Loc := vec3.ScaleVec3(vec3.Sum(viewportUpperLeft, *vec3.New([3]float64{0.5, 0.5, 0.5})), vec3.Sum(pixelDeltaU, pixelDeltaV))

	fmt.Println("P3")
	fmt.Println(imageWidth, imageHeight)
	fmt.Println("255")

	for i := 0; i < int(imageHeight); i++ {
		log.Println("Scanlines remaining:", (imageHeight - i))
		for j := 0; j < imageWidth; j++ {
			pixelCenter := vec3.Sum(vec3.Sum(pixel00Loc, vec3.ScaleFloat(pixelDeltaU, float64(i))),
				vec3.ScaleFloat(pixelDeltaV, float64(j)))

			rayDirection := vec3.Sum(pixelCenter, vec3.ScaleFloat(*cameraCenter, -1.0))
			r := ray.New(*cameraCenter, rayDirection)

			pixelColor := rayColor(*r)

			// pixelColor := color.New([3]float64{float64(i) / float64(imageWidth-1.0), 0.0, float64(j) / float64(imageHeight-1.0)})
			color.WriteColor(pixelColor)
		}
	}

	log.Println("Done.")
}
