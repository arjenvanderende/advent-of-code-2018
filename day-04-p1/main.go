package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open input.txt: %v", err)
	}
	defer file.Close()

	// read input and sort it
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sort.Strings(lines)

	// find laziest guard
	schedule := newSchedule()
	var start int
	var s *shift
	for _, line := range lines {
		tp := line[1:17]
		t, err := time.Parse("2006-01-02 15:04", tp)
		if err != nil {
			log.Fatalf("Cannot parse time %s: %v", tp, err)
		}

		a := line[19:]
		switch a {
		case "falls asleep":
			s.setDate(t.Month(), t.Day())
			start = t.Minute()
		case "wakes up":
			s.setAsleep(start, t.Minute())
		default:
			var guard int
			_, err := fmt.Sscanf(a, "Guard #%d begins shift", &guard)
			if err != nil {
				log.Fatalf("Unable to parse action %s as guard rotation: %v", a, err)
			}

			s = newShift(guard)
			schedule.addShift(s)
		}
	}

	fmt.Println(schedule)
	fmt.Println(schedule.laziest())
}

type schedule struct {
	shifts []*shift
}

func newSchedule() *schedule {
	return &schedule{
		shifts: make([]*shift, 0),
	}
}

func (s *schedule) addShift(sh *shift) {
	s.shifts = append(s.shifts, sh)
}

func (s *schedule) laziest() string {
	guard, best := 0, 0

	// find laziest guard
	laziest := make(map[int]int)
	for _, shift := range s.shifts {
		laziest[shift.guard] += shift.total()
		total := laziest[shift.guard]
		if total > best {
			guard = shift.guard
			best = total
		}
	}

	// find laziest minute for laziest guard
	best = 0
	minute := 0
	minutes := make(map[int]int)
	for _, shift := range s.shifts {
		if shift.guard != guard {
			continue
		}

		for m := range shift.asleep {
			minutes[m]++
			if minutes[m] > best {
				minute = m
				best = minutes[m]
			}
		}
	}
	return fmt.Sprintf("Guard: %d, Minute: %d, Best: %d", guard, minute, guard*minute)
}

func (s *schedule) String() string {
	var out strings.Builder
	fmt.Fprintln(&out, "Date   ID     Minute")
	fmt.Fprintln(&out, "              000000000011111111112222222222333333333344444444445555555555")
	fmt.Fprintln(&out, "              012345678901234567890123456789012345678901234567890123456789")
	for _, shift := range s.shifts {
		fmt.Fprintln(&out, shift)
	}
	return out.String()
}

type shift struct {
	month  time.Month
	day    int
	guard  int
	asleep map[int]bool
}

func newShift(guard int) *shift {
	return &shift{
		guard:  guard,
		asleep: make(map[int]bool),
	}
}

func (s *shift) setDate(month time.Month, day int) {
	s.month = month
	s.day = day
}

func (s *shift) setAsleep(start, end int) {
	for m := start; m < end; m++ {
		s.asleep[m] = true
	}
}

func (s *shift) total() int {
	total := 0
	for range s.asleep {
		total++
	}
	return total
}

func (s *shift) String() string {
	var out strings.Builder
	fmt.Fprintf(&out, "%02d-%02d  #%4d  ", s.month, s.day, s.guard)
	for i := 0; i < 60; i++ {
		if _, ok := s.asleep[i]; ok {
			fmt.Fprintf(&out, "#")
		} else {
			fmt.Fprintf(&out, ".")
		}
	}
	return out.String()
}
