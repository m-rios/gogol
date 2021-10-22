package world

import (
	"errors"
	"math"
)

type World struct {
	Cells       []bool
	Width       int
	PaddedWidth int
	Size        int
}

func NewWorld(width int) World {
	// We use a padding of 1 to avoid boundary conditions when updating the world
	paddedWidth := 2 + width
	size := int(math.Pow(float64(paddedWidth), 2))

	return World{
		Width:       width,
		Cells:       make([]bool, size),
		Size:        size,
		PaddedWidth: paddedWidth}
}

// Spawns a glider patters in the world's top left corner
func SpawnGlider(world *World) {
	world.Cells[Sub2Ind(world, 1, 3)] = true
	world.Cells[Sub2Ind(world, 2, 3)] = true
	world.Cells[Sub2Ind(world, 3, 3)] = true

	world.Cells[Sub2Ind(world, 2, 1)] = true

	world.Cells[Sub2Ind(world, 3, 2)] = true
}

func UpdateWorld(world *World, newWorld *World) error {
	if (world.Size > newWorld.Size) {
		return errors.New("world must be equal or bigger in size than newWorld")
	}

	// We have a padding of 1 row and 1 column on the top-left column, so we must
	// start counting from the (1,1) cell
	for ind := Sub2Ind(world, 1, 1); ind < world.Size; ind++ {
		neighbourCount := countNeighbours(world, ind)
		if (neighbourCount < 2) {
			// Underpopulation, cell dies
			newWorld.Cells[ind] = false
		} else if (neighbourCount == 2) {
			// Equilibrium, cell remains as is
			newWorld.Cells[ind] = world.Cells[ind]
		} else if (neighbourCount == 3) {
			// Reproduction, a new cell is born
			newWorld.Cells[ind] = true
		} else {
			// Overpopulation, cell dies
			newWorld.Cells[ind] = false
		}
	}

	return nil
}

func countNeighbours(world *World, ind int) int {
	// top row
	// row of ind
	// bottom row
	return 0
}

// Ind2Sub returns row and column that corresponds to an index. The index is a
// position in the World cells. All values are 0 indexed (the first possible position is 0)
func Ind2Sub(world *World, ind int) (int, int) {
	return ind / world.PaddedWidth, ind % world.PaddedWidth
}

func Sub2Ind(world *World, row int, column int) int {
	return row%world.PaddedWidth + world.PaddedWidth*column
}

func Test1() {

}
