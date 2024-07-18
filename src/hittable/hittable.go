// Package hittable contains the Hittable interface and other types that will embed Hittable
package hittable

import (
	"main/src/point3"
	"main/src/ray"
	"main/src/vec3"
)

type HitRecord struct {
	P         point3.Point3
	Normal    vec3.Vec3
	T         float64
	FrontFace bool
}

func (h *HitRecord) SetFaceNormal(r ray.Ray, outwardNormal vec3.Vec3) {
	h.FrontFace = vec3.Dot(r.Direction(), outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.ScaleFloat(-1.0)
	}
}

type Hittable interface {
	Hit(r ray.Ray, rayMinT float64, rayMaxT float64, rec *HitRecord) bool
}

type Hittables struct {
	objects []Hittable
}

func (hs *Hittables) Add(object Hittable) {
	hs.objects = append(hs.objects, object)
}

func (hs *Hittables) Hit(r ray.Ray, rayMinT float64, rayMaxT float64, rec *HitRecord) bool {
	var tempRecord HitRecord
	hitAnything := false
	closestSoFar := rayMaxT

	for _, object := range hs.objects {
		if object.Hit(r, rayMinT, closestSoFar, &tempRecord) {
			hitAnything = true
			closestSoFar = tempRecord.T
			*rec = tempRecord
		}
	}

	return hitAnything
}
