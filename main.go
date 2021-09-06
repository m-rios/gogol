package main

import (
    "fmt"
)

func main() {
    world := NewWorld(2)
    UpdateWorld(world)
    PrintWorld(world)
    fmt.Println(world.cells)
}
