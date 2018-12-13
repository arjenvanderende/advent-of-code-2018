package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open input.txt: %v", err)
	}
	defer file.Close()

	var p polymer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p = newPolymer(scanner.Text())
		break
	}

	for p.collapse() {
	}

	fmt.Println(p.length())
}

type polymer struct {
	elements []rune
}

func newPolymer(value string) polymer {
	elements := make([]rune, 0)
	for _, r := range value {
		elements = append(elements, r)
	}

	return polymer{
		elements: elements,
	}
}

func (p *polymer) collapse() bool {
	for i := 0; i < len(p.elements)-1; i++ {
		a := p.elements[i]
		b := p.elements[i+1]
		if a != b &&
			((unicode.IsLower(a) && a == unicode.ToLower(b)) ||
				(unicode.IsUpper(a) && a == unicode.ToUpper(b))) {
			p.elements = append(p.elements[0:i], p.elements[i+2:]...)
			return true
		}
	}
	return false
}

func (p *polymer) value() string {
	return string(p.elements)
}

func (p *polymer) length() int {
	return len(p.elements)
}
