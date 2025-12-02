package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func is_valid(n int) bool {
	s := strconv.Itoa(n)
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			seg := s[:i]
			all := true
			for j := i; j <= len(s)-i; j += i {
				if seg != s[j:j+i] {
					all = false
					break
				}
			}
			if all {
				return false
			}
		}
	}

	return true
}

func main() {
	log.SetPrefix("aoc25: ")
	total := 0

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.SplitSeq(scanner.Text(), ",")
	for r := range line {
		rn := strings.Split(r, "-")
		start, err := strconv.Atoi(rn[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(rn[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := start; i <= end; i++ {
			if !is_valid(i) {
				total += i
			}
		}
	}

	fmt.Printf("Result: %v\n", total)
}
