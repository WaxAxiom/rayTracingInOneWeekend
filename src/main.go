package main

import (
	"main/src/camera"
	"main/src/hittable"
	"main/src/sphere"
	"main/src/vec3"
)

func main() {
	var world hittable.Hittables
	world.Add(sphere.New(*vec3.New([3]float64{0, 0, -1}), 0.5))
	world.Add(sphere.New(*vec3.New([3]float64{0, -100.5, -1}), 100))

	var cam = new(camera.Camera)
	cam.AspectRatio = 16.0 / 9.0
	cam.ImageWidth = 400

	cam.Render(world)
}
