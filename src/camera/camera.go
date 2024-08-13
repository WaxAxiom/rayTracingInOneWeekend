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
	"main/src/vec3"
	"math"
)

type Camera struct {
	AspectRatio float64
	ImageWidth  int
	imageHeight int
	center      point3.Point3
	pixel00Loc  point3.Point3
	pixelDeltaU vec3.Vec3
	pixelDeltaV vec3.Vec3
}

func (c *Camera) Render(world hittable.Hittables) {
	c.Initialize()

	fmt.Println("P3")
	fmt.Println(c.ImageWidth, c.imageHeight)
	fmt.Println("255")

	for i := 0; i < int(c.imageHeight); i++ {
		log.Println("Scanlines remaining:", (c.imageHeight - i))
		for j := 0; j < c.ImageWidth; j++ {
			pixelCenter := c.pixel00Loc.AddVec3(vec3.SumVec3(c.pixelDeltaU.ScaleFloat(float64(j)), c.pixelDeltaV.ScaleFloat(float64(i))))
			rayDirection := pixelCenter.AddVec3(c.center.ScaleFloat(-1.0))

			r := ray.New(c.center, rayDirection)

			pixelColor := c.rayColor(*r, world)

			color.WriteColor(pixelColor)
		}
	}

	log.Println("Done.")
}

func (c *Camera) Initialize() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)
	if c.imageHeight < 1 {
		c.imageHeight = 1
	}

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

func (c *Camera) getRay() {

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
