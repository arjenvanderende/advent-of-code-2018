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

	// parse numbers from file
	scanner := bufio.NewScanner(file)
	numbers := make([]int, 0)
	for scanner.Scan() {
		var num int
		fmt.Sscanf(scanner.Text(), "%d", &num)
		numbers = append(numbers, num)
	}

	// find duplicate frequency
	frequencies := map[int]bool{
		0: true,
	}
	total := 0
	for {
		for _, num := range numbers {
			total += num
			if frequencies[total] == true {
				fmt.Printf("Duplicate freq: %d", total)
				return
			}
			frequencies[total] = true
		}
	}
}
