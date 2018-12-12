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
		m: map[xy][]int{},
	}
	for scanner.Scan() {
		var id, x, y, w, h int
		fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		f.claim(id, x, y, w, h)
	}

	fmt.Printf("No overlapping claim for ID %d", f.safe())
}

type xy struct{ x, y int }
type fabric struct {
	m map[xy][]int
}

func (f *fabric) claim(id, x, y, w, h int) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			f.m[xy{i, j}] = append(f.m[xy{i, j}], id)
		}
	}
}

func (f *fabric) safe() int {
	duplicates := map[int]bool{}
	for _, ids := range f.m {
		switch len(ids) {
		case 0:
			log.Fatalf("Claim has no claim ID")
			break
		case 1:
			if _, ok := duplicates[ids[0]]; !ok {
				duplicates[ids[0]] = true
			}
			break
		default:
			for _, id := range ids {
				duplicates[id] = false
			}
		}
	}

	for id, valid := range duplicates {
		if valid {
			return id
		}
	}
	return -1
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
				if len(val) == 1 {
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
