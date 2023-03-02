package main

import (
	"fmt"
	"io"
	"time"

	"github.com/gosuri/uilive"
)

func demo() {
	writer := uilive.New()

	// start listening for updates and render
	writer.Start()

	for _, f := range [][]string{{"Foo.zip", "Bar.iso"}, {"Baz.tar.gz", "Qux.img"}} {
		for i := 0; i <= 50; i++ {
			_, _ = fmt.Fprintf(writer, "Downloading %s.. (%d/%d) GB\n", f[0], i, 50)
			_, _ = fmt.Fprintf(writer.Newline(), "Downloading %s.. (%d/%d) GB\n", f[1], i, 50)
			time.Sleep(time.Millisecond * 25)
		}
		_, _ = fmt.Fprintf(writer.Bypass(), "Downloaded %s\n", f[0])
		_, _ = fmt.Fprintf(writer.Bypass(), "Downloaded %s\n", f[1])
	}
	_, _ = fmt.Fprintln(writer, "Finished: Downloaded 150GB")
	writer.Stop() // flush and stop rendering

}
func main2() {
	writer := uilive.New()
	// writer.RefreshInterval = time.Hour
	writer.Start()

	width := 10
	height := 4

	grid := make([][]rune, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]rune, width)
		for j := 0; j < width; j++ {
			grid[i][j] = '.'
		}
	}

	writers := make([]io.Writer, height)
	for i := 0; i < height; i++ {
		writers[i] = writer.Newline()
	}

	r := '$'

	for {

		for idx := 0; idx < width*height; idx++ {
			h := idx / width
			W := idx % width
			grid[h][W] = r
			for h := 0; h < height; h++ {
				fmt.Fprintf(writers[h], "%v \n", string(grid[h]))
			}
			fmt.Fprintf(writers[0], "%v \n", string(grid[h]))
			time.Sleep(time.Second)
		}

		r += 1
	}

	writer.Stop()

}
