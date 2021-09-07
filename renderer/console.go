package renderer

import (
    "fmt"
    w "github.com/m-rios/gogol/world"
)

func PrintWorld(world w.World) {
    for r := 1; r <= world.Width; r ++ {
        for c := 1; c <= world.Width; c ++ {
            cell := world.Cells[w.Sub2Ind(world, r, c)]
            if (cell) {
                fmt.Print("X")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}
