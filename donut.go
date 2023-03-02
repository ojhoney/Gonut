package main

import "math"

type Coordinate struct {
	x float64
	y float64
	z float64
}

type Donut struct {
	center       Coordinate
	tubeRadius   float64 // radius of the tube
	centerRadius float64 // radius from the center of the hole to the center of the tube
}

func (d *Donut) getCoordinate(theta float64, phi float64) Coordinate {

	z := d.centerRadius*math.Cos(phi) + d.tubeRadius*math.Cos(theta)*math.Cos(phi)
	y := d.centerRadius*math.Sin(phi) + d.tubeRadius*math.Cos(theta)*math.Sin(phi)
	x := d.tubeRadius * math.Sin(theta)

	return Coordinate{x, y, z}
}

func (d *Donut) getNormal(theta float64, phi float64) Coordinate {
	z := math.Cos(theta) * math.Cos(phi)
	y := math.Cos(theta) * math.Sin(phi)
	x := math.Sin(theta)

	return Coordinate{x, y, z}
}
