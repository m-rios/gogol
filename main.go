package main

import (
	r "github.com/m-rios/gogol/renderer"
	w "github.com/m-rios/gogol/world"
)

func main() {
	world := w.NewWorld(16)
	// w.SpawnGlider(world)
	newWorld := w.NewWorld(16)
	w.UpdateWorld(&world, &newWorld)
	r.PrintWorld(&world)
}
