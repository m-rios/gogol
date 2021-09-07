package main

import (
	r "github.com/m-rios/gogol/renderer"
	w "github.com/m-rios/gogol/world"
)

func main() {
	world := w.NewWorld(16)
	w.SpawnGlider(world)
	// w.UpdateWorld(world)
	r.PrintWorld(world)
}
