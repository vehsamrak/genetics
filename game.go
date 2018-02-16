package main

import (
    "math/rand"
    "time"
)

const CELL_TYPE_FOOD = 1
const CELL_TYPE_OBSTACLE = 2

type Game struct {
    X         int
    Y         int
    field     [][]int
    Seed      int64
    FreeSpace int
}

func init() {
    rand.Seed(time.Now().Unix())
}

func (game *Game) CreateField() {
    if game.Seed != 0 {
        rand.Seed(game.Seed)
    }

    if game.FreeSpace == 0 {
        game.FreeSpace = (game.X + game.Y) / 2
    }

    game.field = make([][]int, game.X)

    for x := 0; x < game.X; x++ {
        var gameFieldY []int
        gameFieldY = make([]int, game.Y)

        for y := 0; y < game.Y; y++ {
            gameFieldY[y] = rand.Intn(game.FreeSpace)
        }

        game.field[x] = gameFieldY
    }
}

func (game *Game) ViewAsText() (output string) {
    for _, row := range game.field {
        for _, cell := range row {
            output += game.transformCellToSymbol(cell)
        }

        output += "\n"
    }

    return
}

func (game *Game) transformCellToSymbol(cell int) (symbol string) {
    cellTypes := map[int]string{
        CELL_TYPE_FOOD:     "o",
        CELL_TYPE_OBSTACLE: "#",
    }

    symbol = cellTypes[cell]

    if symbol == "" {
        symbol = "."
    }

    return symbol
}
