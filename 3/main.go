package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Map[T any, U any](in []T, f func(T) U) []U {
	out := make([]U, 0, len(in))
	for _, v := range in {
		out = append(out, f(v))
	}
	return out
}

func MaxIdx[T cmp.Ordered](in []T) (int, T) {
	biggest := in[0]
	biggest_idx := 0
	for i, n := range in {
		if n > biggest {
			biggest = n
			biggest_idx = i
		}
	}
	return biggest_idx, biggest
}

func find_bigggest_joltage(bank string) int {
	iBank := Map(strings.Split(bank, ""), func(s string) int {
		j, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return j
	})

	out := 0

	_ = iBank

	idx := 0

	skipped := 0

	for i := range 12 {
		bi, b := MaxIdx(iBank[idx : idx+(len(bank)-11)-skipped])

		skipped += bi

		idx = idx + bi + 1

		out += int(math.Pow10(11-i)) * b
	}

	return out
}

func main() {
	log.SetPrefix("aoc25: ")

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		a := find_bigggest_joltage(line)
		total += a
	}

	fmt.Println(total)
}
