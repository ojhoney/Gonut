package main

import "math"

func dotProduct(c1 Coordinate, c2 Coordinate) float64 {
	return c1.x*c2.x + c1.y*c2.y + c1.z*c2.z
}

func distanceSquared(c1 Coordinate, c2 Coordinate) float64 {
	return (c1.x-c2.x)*(c1.x-c2.x) + (c1.y-c2.y)*(c1.y-c2.y) + (c1.z-c2.z)*(c1.z-c2.z)
}

func rotate(c Coordinate, yaw, pitch, roll float64) Coordinate {
	sinA, cosA := math.Sin(yaw), math.Cos(yaw)
	sinB, cosB := math.Sin(pitch), math.Cos(pitch)
	sinC, cosC := math.Sin(roll), math.Cos(roll)

	x := cosA*cosB*c.x + (cosA*sinB*sinC-sinA*cosC)*c.y + (cosA*sinB*cosC-sinA*sinC)*c.z
	y := sinA*cosB*c.x + (sinA*sinB*sinC+cosA*cosC)*c.y + (sinA*sinB*cosC-cosA*sinC)*c.z
	z := -sinB*c.x + cosB*sinC*c.y + cosB*cosC*c.z

	return Coordinate{x, y, z}
}
