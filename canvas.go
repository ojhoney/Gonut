package main

import (
	"fmt"
	"math"
)

func (canvas *Canvas) Draw(u, v float64, light_val float64) {
	r := int(math.Floor(u/canvas.ratio)) + canvas.width/2
	if !(0 <= r && r < canvas.width) {
		fmt.Printf("Out of bound %v -> %v\n", u, r)
		return
	}

	c := int(math.Floor(v/canvas.ratio)) + canvas.width/2
	if !(0 <= c && c < canvas.width) {
		fmt.Printf("Out of bound %v -> %v\n", v, c)
		return
	}

	// note that light_val is [-1, 1]
	// filter to [0, 1]
	// we should convert this to [0, MAX_GRAY]
	val := int((light_val * float64(MAX_GRAY)))

	fmt.Println(u, v)
	fmt.Println(light_val, val)

	if val > canvas.grid[r][c] {
		canvas.grid[r][c] = val
	}

}

func (canvas *Canvas) show() {

	for i := 0; i < canvas.height; i++ {
		for j := 0; j < canvas.width; j++ {
			fmt.Printf("%s ", string(GRAYSCALE[canvas.grid[i][j]]))
			//fmt.Printf("%s ", string(GRAYSCALE[MAX_GRAY-min(s.grid[i][j])]))
		}
		fmt.Println()
	}

}
