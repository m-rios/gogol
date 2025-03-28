package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// x,y coordinates (so x is rows and y columns)
type world [80][22]int

const generationMode string = "random"

// x,y coordinates of the directions where to find neighbors
var directions = [][]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func initializeWorld(w *world) {
	// Glider
	// w[3][1] = 1
	// w[3][2] = 1
	// w[3][3] = 1
	// w[1][2] = 1
	// w[2][3] = 1

	//Blinker
	// w[30][3] = 1
	// w[30][4] = 1
	// w[30][5] = 1

	// Block
	// w[30][10] = 1
	// w[31][10] = 1
	// w[30][11] = 1
	// w[31][11] = 1

	// Gosper glider gun
	w[1][5] = 1
	w[1][6] = 1

	w[2][5] = 1
	w[2][6] = 1

	w[11][5] = 1
	w[11][6] = 1
	w[11][7] = 1

	w[12][4] = 1
	w[12][8] = 1

	w[13][3] = 1
	w[13][9] = 1

	w[14][3] = 1
	w[14][9] = 1

	w[15][6] = 1

	w[16][4] = 1
	w[16][8] = 1

	w[17][5] = 1
	w[17][6] = 1
	w[17][7] = 1

	w[18][6] = 1

	w[21][3] = 1
	w[21][4] = 1
	w[21][5] = 1

	w[22][3] = 1
	w[22][4] = 1
	w[22][5] = 1

	w[23][2] = 1
	w[23][6] = 1

	w[25][1] = 1
	w[25][2] = 1
	w[25][6] = 1
	w[25][7] = 1

	w[35][3] = 1
	w[35][4] = 1

	w[36][3] = 1
	w[36][4] = 1
}

func initializeWorldRandomly(w *world) {
	for x := 1; x < 80; x++ {
		for y := 1; y < 22; y++ {
			rnd := rand.Int()
			w[x][y] = rnd % 2
		}
	}
}

func printWorld(w world) {
	for y := 0; y < 22; y++ {
		for x := 0; x < 80; x++ {
			if x == 0 || x == 80-1 || y == 0 || y == 22-1 {
				fmt.Print("\033[47m", " ", "\033[0m")
				continue
			}
			cell := w[x][y]
			if cell > 0 {
				fmt.Print("\033[42m", "O", "\033[0m")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func countNeightbours(w *world, x, y int) int {
	count := 0
	for _, direction := range directions {
		xInc := direction[0]
		yInc := direction[1]
		if w[x+xInc][y+yInc] > 0 {
			count += 1
		}
	}

	return count
}

func updateWorld(oldWorld, newWorld *world) {
	// Iterate from 1 to range - 1 to avoid boundary conditions
	for y := 1; y < 22-1; y++ {
		for x := 1; x < 80-1; x++ {
			count := countNeightbours(oldWorld, x, y)
			cell := oldWorld[x][y]

			// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
			if cell == 1 && count < 2 {
				newWorld[x][y] = 0
				// Any live cell with two or three live neighbours lives on to the next generation.
			} else if cell == 1 && (count == 2 || count == 3) {
				newWorld[x][y] = 1
				// Any live cell with more than three live neighbours dies, as if by overpopulation.
			} else if cell == 1 && count > 3 {
				newWorld[x][y] = 0
				// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
			} else if cell == 0 && count == 3 {
				newWorld[x][y] = 1
			}
		}
	}
}

func main() {
	var oldWorld, world world

	if generationMode == "random" {
		initializeWorldRandomly(&world)
	} else {
		initializeWorld(&world)
	}

	// Create a channel to receive signals
	sigs := make(chan os.Signal, 1)

	// Notify the channel of incoming SIGINT signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	iter := 0

	// Disable cursor, it's distracting
	fmt.Print("\033[?25l")

	fmt.Println("Iteration: ", iter)
	printWorld(world)

	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	for {
		select {
		// signal caught, cleanup
		case <-sigs:
			// re-enable cursor
			fmt.Print("\033[?25h")
			os.Exit(0)
		// signal not caught, main loop
		default:
			time.Sleep(10 * time.Millisecond)
			updateWorld(&world, &oldWorld)
			fmt.Print("\033[23F") // move cursor position 22 + 1 rows up and to the start
			iter += 1
			fmt.Println("Iteration: ", iter)
			printWorld(oldWorld)
			world = oldWorld
		}
	}
}
