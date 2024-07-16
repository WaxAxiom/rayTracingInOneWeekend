package main

import (
	"fmt"
	"log"

	"main/src/color"
	"main/src/point3"
	"main/src/ray"
	"main/src/vec3"
)

func hitSphere(center point3.Point3, radius float64, r ray.Ray) bool {
	oc := vec3.Sum(center, vec3.ScaleFloat(r.Origin(), -1.0))
	a := vec3.Dot(r.Direction(), r.Direction())
	b := -2.0 * vec3.Dot(r.Direction(), oc)
	c := vec3.Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant >= 0
}

func rayColor(r ray.Ray) color.Color {
	if hitSphere(*vec3.New([3]float64{0, 0, -1}), 0.5, r) {
		return color.New([3]float64{1, 0, 0})
	}

	unitDirection := vec3.UnitVector(r.Direction())
	a := 0.5 * (unitDirection.Y() + 1.0)

	return color.Sum(color.ScaleFloat(color.New([3]float64{1.0, 1.0, 1.0}), (1.0-a)),
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
	viewportV := vec3.New([3]float64{0, -viewportHeight, 0})

	pixelDeltaU := vec3.ScaleFloat(*viewportU, 1/float64(imageWidth))
	pixelDeltaV := vec3.ScaleFloat(*viewportV, 1/float64(imageHeight))

	viewportUpperLeft := vec3.Sum(*cameraCenter, vec3.ScaleFloat(*vec3.New([3]float64{0, 0, focalLength}), -1.0))
	viewportUpperLeft = vec3.Sum(viewportUpperLeft, vec3.ScaleFloat(*viewportU, -1.0/2.0))
	viewportUpperLeft = vec3.Sum(viewportUpperLeft, vec3.ScaleFloat(*viewportV, -1.0/2.0))

	pixel00Loc := vec3.Sum(vec3.ScaleFloat(vec3.Sum(pixelDeltaU, pixelDeltaV), 0.5), viewportUpperLeft)

	fmt.Println("P3")
	fmt.Println(imageWidth, imageHeight)
	fmt.Println("255")

	for i := 0; i < int(imageHeight); i++ {
		log.Println("Scanlines remaining:", (imageHeight - i))
		for j := 0; j < imageWidth; j++ {
			pixelCenter := vec3.Sum(pixel00Loc, vec3.Sum(vec3.ScaleFloat(pixelDeltaU, float64(j)), vec3.ScaleFloat(pixelDeltaV, float64(i))))

			rayDirection := vec3.Sum(pixelCenter, vec3.ScaleFloat(*cameraCenter, -1.0))
			r := ray.New(*cameraCenter, rayDirection)

			pixelColor := rayColor(*r)

			color.WriteColor(pixelColor)
		}
	}

	log.Println("Done.")
}
