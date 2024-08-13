// Package camera contains the Camera type and associated methods
package camera

import (
	"fmt"
	"log"
	"main/src/color"
	"main/src/hittable"
	"main/src/interval"
	"main/src/point3"
	"main/src/ray"
	"main/src/rtweekend"
	"main/src/vec3"
	"math"
)

type Camera struct {
	AspectRatio       float64
	ImageWidth        int
	SamplesPerPixel   int
	imageHeight       int
	center            point3.Point3
	pixel00Loc        point3.Point3
	pixelDeltaU       vec3.Vec3
	pixelDeltaV       vec3.Vec3
	pixelSamplesScale float64
}

func (c *Camera) Render(world hittable.Hittables) {
	c.Initialize()

	fmt.Println("P3")
	fmt.Println(c.ImageWidth, c.imageHeight)
	fmt.Println("255")

	for j := 0; j < int(c.imageHeight); j++ {
		log.Println("Scanlines remaining:", (c.imageHeight - j))
		for i := 0; i < c.ImageWidth; i++ {
			pixelColor := color.New([3]float64{0, 0, 0})

			for sample := 0; sample < c.SamplesPerPixel; sample++ {
				r := c.getRay(i, j)
				pixelColor = pixelColor.AddColor(c.rayColor(r, world))
			}

			// pixelCenter := c.pixel00Loc.AddVec3(vec3.SumVec3(c.pixelDeltaU.ScaleFloat(float64(j)), c.pixelDeltaV.ScaleFloat(float64(i))))
			// rayDirection := pixelCenter.AddVec3(c.center.ScaleFloat(-1.0))

			// r := ray.New(c.center, rayDirection)

			// pixelColor := c.rayColor(*r, world)

			// log.Println(pixelColor)

			color.WriteColor(pixelColor.ScaleFloat(c.pixelSamplesScale))
		}
	}

	log.Println("Done.")
}

func (c *Camera) Initialize() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)
	if c.imageHeight < 1 {
		c.imageHeight = 1
	}

	c.SamplesPerPixel = 10
	c.pixelSamplesScale = 1.0 / float64(c.SamplesPerPixel)

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(c.ImageWidth) / float64(c.imageHeight))
	cameraCenter := vec3.New([3]float64{0, 0, 0})
	c.center = *cameraCenter

	viewportU := vec3.New([3]float64{viewportWidth, 0, 0})
	viewportV := vec3.New([3]float64{0, -viewportHeight, 0})

	c.pixelDeltaU = viewportU.ScaleFloat(1 / float64(c.ImageWidth))
	c.pixelDeltaV = viewportV.ScaleFloat(1 / float64(c.imageHeight))

	viewportUpperLeft := cameraCenter.AddVec3(*vec3.New([3]float64{0, 0, focalLength})).ScaleFloat(-1.0)
	viewportUpperLeft = viewportUpperLeft.AddVec3(viewportU.ScaleFloat(-1.0 / 2.0)).AddVec3(viewportV.ScaleFloat(-1.0 / 2.0))

	c.pixel00Loc = viewportUpperLeft.AddVec3(vec3.SumVec3(c.pixelDeltaU, c.pixelDeltaV).ScaleFloat(0.5))
}

func (c *Camera) getRay(i int, j int) ray.Ray {
	offset := SampleSquare()
	pixelSample := c.pixel00Loc.AddVec3(vec3.SumVec3(c.pixelDeltaU.ScaleFloat(float64(i)+offset.X()),
		c.pixelDeltaV.ScaleFloat(float64(j)+offset.Y())))

	rayOrigin := c.center
	rayDirection := pixelSample.AddVec3(rayOrigin.ScaleFloat(-1.0))

	return *ray.New(rayOrigin, rayDirection)
}

func SampleSquare() vec3.Vec3 {
	return *vec3.New([3]float64{rtweekend.RandomFloat64() - 0.5, rtweekend.RandomFloat64() - 0.5, 0})
}

func (c *Camera) rayColor(r ray.Ray, world hittable.Hittables) color.Color {
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
