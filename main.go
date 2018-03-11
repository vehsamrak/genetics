package main

import (
    "time"
    "math/rand"

    term "github.com/buger/goterm"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
    game := Game{X: 10, Y: 30}
    game.CreateField()

    term.Clear()
    term.MoveCursor(1, 1)
    term.Println(game.ViewAsText())
    term.Flush()
}
