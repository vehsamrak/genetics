package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Game struct {
    X           int
    Y           int
    field       [][]int
    Seed        int64
    FoodPercent int
}

func init() {
    rand.Seed(time.Now().Unix())
}

func (game *Game) CreateField() {
    if game.Seed != 0 {
        rand.Seed(game.Seed)
    }

    if game.FoodPercent == 0 {
        game.FoodPercent = 10
    }

    game.field = make([][]int, game.X)

    for x := 0; x < game.X; x++ {
        var gameFieldY []int
        gameFieldY = make([]int, game.Y)

        for y := 0; y < game.Y; y++ {
            gameFieldY[y] = rand.Intn(10)
        }

        game.field[x] = gameFieldY
    }
}

func (game *Game) View() {
    for _, row := range game.field {
        fmt.Println(row)
    }
}
