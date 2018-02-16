package main

import (
    term "github.com/buger/goterm"
)

func main() {
    game := Game{X: 10, Y: 30}
    game.CreateField()

    term.Clear()
    term.MoveCursor(1, 1)
    term.Println(game.ViewAsText())
    term.Flush()
}
