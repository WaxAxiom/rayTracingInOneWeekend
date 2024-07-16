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
	oc := center.AddVec3(r.Origin().ScaleFloat(-1.0))
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

	return color.SumColor(
		color.New([3]float64{1.0, 1.0, 1.0}).ScaleFloat(1.0-a),
		color.New([3]float64{0.5, 0.7, 1.0}).ScaleFloat(a))
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

	pixelDeltaU := viewportU.ScaleFloat(1 / float64(imageWidth))
	pixelDeltaV := viewportV.ScaleFloat(1 / float64(imageHeight))

	viewportUpperLeft := cameraCenter.AddVec3(*vec3.New([3]float64{0, 0, focalLength})).ScaleFloat(-1.0)
	viewportUpperLeft = viewportUpperLeft.AddVec3(viewportU.ScaleFloat(-1.0 / 2.0)).AddVec3(viewportV.ScaleFloat(-1.0 / 2.0))

	pixel00Loc := viewportUpperLeft.AddVec3(vec3.SumVec3(pixelDeltaU, pixelDeltaV).ScaleFloat(0.5))

	fmt.Println("P3")
	fmt.Println(imageWidth, imageHeight)
	fmt.Println("255")

	for i := 0; i < int(imageHeight); i++ {
		log.Println("Scanlines remaining:", (imageHeight - i))
		for j := 0; j < imageWidth; j++ {
			pixelCenter := pixel00Loc.AddVec3(vec3.SumVec3(pixelDeltaU.ScaleFloat(float64(j)), pixelDeltaV.ScaleFloat(float64(i))))
			rayDirection := pixelCenter.AddVec3(cameraCenter.ScaleFloat(-1.0))

			r := ray.New(*cameraCenter, rayDirection)

			pixelColor := rayColor(*r)

			color.WriteColor(pixelColor)
		}
	}

	log.Println("Done.")
}
