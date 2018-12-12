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
	total := 0
	for scanner.Scan() {
		var num int
		fmt.Sscanf(scanner.Text(), "%d", &num)
		total += num
	}
	fmt.Printf("Total: %d", total)
}
