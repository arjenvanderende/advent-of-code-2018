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

	for i, a := range ids {
		for _, b := range ids[i+1:] {
			if common, ok := compare(a, b); ok {
				fmt.Printf("Common id: %s", common)
				return
			}
		}
	}
	fmt.Println("No common id found")
}

func compare(a, b string) (string, bool) {
	common := make([]byte, 0)
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			common = append(common, a[i])
		}
	}
	return string(common), len(a) == len(common)+1
}
