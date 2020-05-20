package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const sizeX = 100
const sizeY = 50
const randomSeedPercent = 0.50

func main() {
	grid := newGrid(sizeY, sizeX)
	grid.initRandom(randomSeedPercent)

	grid.print()

	t := time.NewTicker(200 * time.Millisecond)
	for range t.C {
		grid = grid.tick()
		grid.print()
	}
}

type grid [][]bool

func (g grid) initRandom(p float64) {
	rand.Seed(time.Now().UnixNano())
	for x := range g {
		for y := range g[0] {
			g[x][y] = rand.Float64() <= p
		}
	}
}

func (g grid) print() {
	fmt.Printf("\033c")

	b := strings.Builder{}

	for _, row := range g {
		b.Reset()

		for _, dot := range row {
			if dot {
				b.WriteString(" *")
			} else {
				b.WriteString("  ")
			}
		}

		fmt.Println(b.String())
	}
}

func (g grid) value(x, y int) int {
	if x < 0 {
		x = len(g) + x
	} else {
		x = x % len(g)
	}

	if y < 0 {
		y = len(g) + y
	} else {
		y = y % len(g[0])
	}

	if g[x][y] {
		return 1
	}

	return 0
}

func (g grid) tick() grid {
	ng := newGrid(len(g), len(g[0]))

	for x := range g {
		for y := range g[x] {
			sum :=
				g.value(x-1, y-1) +
					g.value(x-0, y-1) +
					g.value(x+1, y-1) +
					g.value(x-1, y-0) +
					g.value(x+1, y-0) +
					g.value(x-1, y+1) +
					g.value(x-0, y+1) +
					g.value(x+1, y+1)

			switch {
			case g[x][y] && sum == 2:
				ng[x][y] = true
			case g[x][y] && sum == 3:
				ng[x][y] = true
			case !ng[x][y] && sum == 3:
				ng[x][y] = true
			}
		}
	}

	return ng
}

func newGrid(x, y int) grid {
	s := make([][]bool, x)
	for i := range s {
		s[i] = make([]bool, y)
	}

	return s
}
