package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Game struct {
    X       int
    Y       int
    Command int
    field   [][]int
    Seed    int64
}

func init() {
    rand.Seed(time.Now().Unix())
}

func (game *Game) CreateField() {
    if game.Seed != 0 {
        rand.Seed(game.Seed)
    }

    game.field = make([][]int, game.X)

    var gameFieldY []int
    gameFieldY = make([]int, game.Y)

    for x := 0; x < game.X; x++ {
        for y := 0; y < game.Y; y++ {
            gameFieldY[y] = game.Command
        }

        game.field[x] = gameFieldY
    }
}

func (game *Game) View() {
    for _, row := range game.field {
        for range row {
            fmt.Print(".")
        }
        fmt.Println("")
    }
}
