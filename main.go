package main

import (
    "fmt"
    w "github.com/m-rios/gogol/world"
    r "github.com/m-rios/gogol/renderer"
)

func main() {
    world := w.NewWorld(2)
    w.UpdateWorld(world)
    r.PrintWorld(world)
    fmt.Println(world.Cells)
}
