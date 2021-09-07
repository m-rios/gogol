package world

import (
    "math"
)


type World struct {
    Cells []bool
    Width int
    PaddedWidth int
    Size int
}

func NewWorld(width int) World {
    // We use a padding of 1 to avoid boundary conditions when updating the world
    paddedWidth := 2 + width
    size := int(math.Pow(float64(paddedWidth), 2))

    return World {
        Width: width,
        Cells: make([]bool, size),
        Size: size,
        PaddedWidth: paddedWidth }
}

func UpdateWorld(world World) {
    world.Cells[Sub2Ind(world, 1, 1)] = true
}

// Ind2Sub returns row and column that corresponds to an index. The index is a
// position in the World cells. All values are 0 indexed (the first possible position is 0)
func Ind2Sub(world World, ind int) (int, int) {
    return ind / world.PaddedWidth, ind % world.PaddedWidth
}

func Sub2Ind(world World, row int, column int) int {
    return row % world.PaddedWidth + world.PaddedWidth * column
}

