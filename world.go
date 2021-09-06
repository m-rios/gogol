package main

import (
    "math"
    "fmt"
)


type World struct {
    cells []bool
    width int
    paddedWidth int
    size int
}

func NewWorld(width int) World {
    // We use a padding of 1 to avoid boundary conditions when updating the world
    paddedWidth := 2 + width
    size := int(math.Pow(float64(paddedWidth), 2))

    return World {
        width: width,
        cells: make([]bool, size),
        size: size,
        paddedWidth: paddedWidth }
}

func UpdateWorld(world World) {
    world.cells[Sub2Ind(world, 1, 1)] = true
}

// Ind2Sub returns row and column that corresponds to an index. The index is a
// position in the World cells. All values are 0 indexed (the first possible position is 0)
func Ind2Sub(world World, ind int) (int, int) {
    return ind / world.paddedWidth, ind % world.paddedWidth
}

func Sub2Ind(world World, row int, column int) int {
    return row % world.paddedWidth + world.paddedWidth * column
}

func PrintWorld(world World) {
    for r := 1; r <= world.width; r ++ {
        for c := 1; c <= world.width; c ++ {
            cell := world.cells[Sub2Ind(world, r, c)]
            if (cell) {
                fmt.Print("X")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}
