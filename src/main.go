package main

import (
	"fmt"
	"log"
	"math"

	"main/src/color"
	"main/src/hittable"
	"main/src/interval"
	"main/src/ray"
	"main/src/sphere"
	"main/src/vec3"
)

// func hitSphere(center point3.Point3, radius float64, r ray.Ray) float64 {
// 	oc := center.AddVec3(r.Origin().ScaleFloat(-1.0))
// 	a := r.Direction().LengthSquared()
// 	h := vec3.Dot(r.Direction(), oc)
// 	c := oc.LengthSquared() - radius*radius
// 	discriminant := h*h - a*c

// 	if discriminant < 0 {
// 		return -1.0
// 	} else {
// 		return (h - math.Sqrt(discriminant)) / a
// 	}
// }

func rayColor(r ray.Ray, world hittable.Hittables) color.Color {
	var rec hittable.HitRecord
	if world.Hit(r, *interval.New(0, math.Inf(1)), &rec) {
		return color.SumColor(color.New([3]float64{rec.Normal.X(), rec.Normal.Y(), rec.Normal.Z()}),
			color.New([3]float64{1, 1, 1})).ScaleFloat(0.5)
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

	var world hittable.Hittables
	world.Add(sphere.New(*vec3.New([3]float64{0, 0, -1}), 0.5))
	world.Add(sphere.New(*vec3.New([3]float64{0, -100.5, -1}), 100))

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

			pixelColor := rayColor(*r, world)

			color.WriteColor(pixelColor)
		}
	}

	log.Println("Done.")
}
