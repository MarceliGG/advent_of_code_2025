package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Dial struct {
	cur   int
	total int
}

func (d *Dial) Inc() {
	d.cur++
	if d.cur == 100 {
		d.total++
		d.cur = 0
	}
}

func (d *Dial) Dec() {
	d.cur--
	if d.cur == 0 {
		d.total++
	}
	if d.cur == -1 {
		d.cur = 99
	}
}

func (d *Dial) Rotate(r string) {
	direction := r[0]
	n, err := strconv.Atoi(r[1:])
	if err != nil {
		log.Fatal(err)
	}
	switch direction {
	case 'L':
		for range n {
			d.Dec()
		}
	case 'R':
		for range n {
			d.Inc()
		}
	default:
		log.Fatalf("Unknown direction")
	}
}

func main() {
	log.SetPrefix("aoc25: ")
	dial := Dial{50, 0}
	fmt.Printf("c: %v, t: %v\n", dial.cur, dial.total)
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		dial.Rotate(line)
	}

	fmt.Println(dial.total)
}
