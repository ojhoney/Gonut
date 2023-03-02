package main

import (
	"fmt"
	"math"
)

var (
	PI               = math.Pi
	MAX_GRAY  int    = 11
	GRAYSCALE []rune = []rune(".,-~:;=!*#$@") // furthest to closest
)

type Canvas struct {
	center       Coordinate
	width        int
	height       int
	ratio        float64 // grid 한칸의 길이
	grid         [][]int
	depth_buffer [][]float64
}

func InitCanvas(width, height int, ratio float64) *Canvas {
	grid := make([][]int, height)
	depth_buffer := make([][]float64, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
		depth_buffer[i] = make([]float64, width)
	}
	return &Canvas{
		width:        width,
		height:       height,
		ratio:        ratio,
		grid:         grid,
		depth_buffer: depth_buffer,
	}
}

type Env struct {
	//Eye    Coordinate // (d, 0, 0) for now
	Light  Coordinate // light source coordinate (unit)
	Canvas *Canvas
	d1     float64 // eye to canvas distance
	d2     float64 // canvas to donut center distance
}

func InitEnv(canvas *Canvas, d1 float64, d2 float64) *Env {
	// eye := Coordinate{d1, 0, 0}
	// canvas.center
	// distance :=
	return &Env{
		// Eye:      eye,
		Light:  Coordinate{1, 0, 0},
		Canvas: canvas,
		d1:     d1,
		d2:     d2,
	}
}

// given eye, canvas and 3d point, find 3d point's projection to the canvas

// func (env *Env) project(c Coordinate) (float64, float64) {
// 	u := c.y * env.d1 / (c.x + env.d1 + env.d2)
// 	v := c.z * env.d1 / (c.x + env.d1 + env.d2)
// 	return u, v
// }

func (env *Env) Draw(xyz Coordinate, light_val float64) {
	canvas := env.Canvas

	u := xyz.y * env.d1 / (xyz.x + env.d1 + env.d2)
	v := xyz.z * env.d1 / (xyz.x + env.d1 + env.d2)

	row := int(math.Floor(u/canvas.ratio)) + canvas.width/2
	if !(0 <= row && row < canvas.width) {
		fmt.Printf("Out of bound %v -> %v\n", u, row)
		return
	}

	col := int(math.Floor(v/canvas.ratio)) + canvas.width/2
	if !(0 <= col && col < canvas.width) {
		fmt.Printf("Out of bound %v -> %v\n", v, col)
		return
	}

	depth := distanceSquared(canvas.center, xyz) // distance from canvas to point

	// if light_val < 0 && depth < canvas.depth_buffer[row][col] {
	// 	fmt.Println(depth, canvas.depth_buffer[row][col], light_val)
	// 	return
	// }
	if depth < canvas.depth_buffer[row][col] || canvas.depth_buffer[row][col] == 0 {
		canvas.depth_buffer[row][col] = depth
		canvas.grid[row][col] = int((light_val * 0.9 * float64(MAX_GRAY)))
	}

}
func main() {
	canvas := InitCanvas(40, 40, 0.4)
	d1 := 200.0 // eye to canvas distance
	d2 := 10.0  // canvas to donut center distance
	env := InitEnv(canvas, d1, d2)
	donut := Donut{
		tubeRadius:   2,
		centerRadius: 4}

	yaw := 0.0 //math.Pi / 3
	pitch := 0.7
	roll := 0.0

	for theta := 0.0; theta < 2*math.Pi; theta += 0.05 {
		for phi := 0.0; phi < 2*math.Pi; phi += 0.05 {
			xyz := donut.getCoordinate(theta, phi)
			xyz = rotate(xyz, yaw, pitch, roll)

			normal := donut.getNormal(theta, phi)
			normal = rotate(normal, yaw, pitch, roll)
			light_val := dotProduct(env.Light, normal)

			if light_val < 0 {
				continue

			}
			env.Draw(xyz, light_val)
		}
		// fmt.Println(u, v)

	}
	canvas.show()
	fmt.Println(string(GRAYSCALE[11]))

}
