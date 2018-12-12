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
	ids := make([]string, 0)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	var twoCount, threeCount int
	for _, id := range ids {
		if has(id, 2) {
			twoCount++
		}
		if has(id, 3) {
			threeCount++
		}
	}
	fmt.Printf("%d ⨉ %d = %d", twoCount, threeCount, twoCount*threeCount)
}

func has(id string, count int) bool {
	counts := map[rune]int{}
	for _, r := range id {
		counts[r]++
	}

	for _, num := range counts {
		if num == count {
			return true
		}
	}
	return false
}
