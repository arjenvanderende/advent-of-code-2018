package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open input.txt: %v", err)
	}

	scanner := bufio.NewScanner(file)
	f := fabric{
		m: map[xy]int{},
	}
	for scanner.Scan() {
		var id, x, y, w, h int
		fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		f.claim(x, y, w, h)
	}

	fmt.Printf("Multiple claims on %d square inches", f.duplicate())
}

type xy struct{ x, y int }
type fabric struct {
	m map[xy]int
}

func (f *fabric) claim(x, y, w, h int) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			f.m[xy{i, j}]++
		}
	}
}

func (f *fabric) duplicate() int {
	duplicate := 0
	for _, val := range f.m {
		if val > 1 {
			duplicate++
		}
	}
	return duplicate
}

func (f fabric) String() string {
	var maxX, maxY int
	for xy := range f.m {
		if xy.x > maxX {
			maxX = xy.x
		}
		if xy.y > maxY {
			maxY = xy.y
		}
	}

	result := ""
	for y := 0; y <= maxY+1; y++ {
		for x := 0; x <= maxX+1; x++ {
			if val, ok := f.m[xy{x, y}]; ok {
				if val == 1 {
					result += "#"
				} else {
					result += "X"
				}
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
}
