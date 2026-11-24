package main

import "fmt"

// FloodFill recolors the region of same-colored, 4-connected pixels around
// (sr, sc) to the new color, in place, and returns the image. The early
// return when the color is unchanged prevents infinite recursion.
func FloodFill(image [][]int, sr, sc, color int) [][]int {
	if len(image) == 0 || sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[sr]) {
		return image
	}
	old := image[sr][sc]
	if old == color {
		return image
	}
	var fill func(r, c int)
	fill = func(r, c int) {
		if r < 0 || r >= len(image) || c < 0 || c >= len(image[r]) || image[r][c] != old {
			return
		}
		image[r][c] = color
		fill(r+1, c)
		fill(r-1, c)
		fill(r, c+1)
		fill(r, c-1)
	}
	fill(sr, sc)
	return image
}

func main() {
	img := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(FloodFill(img, 1, 1, 2))
}
