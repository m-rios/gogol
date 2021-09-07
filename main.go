package main

import (
	"fmt"
	r "github.com/m-rios/gogol/renderer"
	w "github.com/m-rios/gogol/world"
)

func main() {
	world := w.NewWorld(2)
	w.UpdateWorld(world)
	r.PrintWorld(world)
	fmt.Println(world.Cells)
}
