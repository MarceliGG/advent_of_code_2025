package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.SetPrefix("aoc25: ")

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	text := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}

	l := len(text[0])

	total := 0
	i := 0
	for i < l {
		var eq func(x, y int) int
		switch text[len(text)-1][i] {
		case '*':
			eq = func(x, y int) int {
				return x * y
			}
		case '+':
			eq = func(x, y int) int {
				return x + y
			}
		default:
			log.Fatal("incorrect equation")
		}

		ns := []int{}
		for i < l {
			if i+1 < l-1 && text[len(text)-1][i+1] != ' ' {
				break
			}
			num := []byte{}
			for _, n := range text[:len(text)-1] {
				num = append(num, n[i])
			}
			n, err := strconv.Atoi(strings.TrimSpace(string(num)))
			if err != nil {
				log.Fatal(err)
			}

			ns = append(ns, n)
			i++
		}

		t := ns[0]
		for _, n := range ns[1:] {
			t = eq(t, n)
		}
		total += t
		i++
	}

	fmt.Printf("Result: %v\n", total)
}
