package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start, end int
}

func (r Range) contains(n int) bool {
	return n >= r.start && n <= r.end
}

func main() {
	log.SetPrefix("aoc25: ")

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fresh []Range

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rn := strings.Split(line, "-")
		start, err := strconv.Atoi(rn[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(rn[1])
		if err != nil {
			log.Fatal(err)
		}

		fresh = append(fresh, Range{start, end})
	}

	total := 0

	sort.Slice(fresh, func(i, j int) bool {
		return fresh[i].start < fresh[j].start
	})

	merged := []Range{fresh[0]}

	for _, r := range fresh[1:] {
		if merged[len(merged)-1].end >= r.start {
			if r.end > merged[len(merged)-1].end {
				merged[len(merged)-1].end = r.end
			}
		} else {
			merged = append(merged, r)
		}
	}

	for _, r := range merged {
		total += r.end - r.start + 1
	}

	fmt.Println(merged)

	fmt.Printf("Result: %v\n", total)
}
