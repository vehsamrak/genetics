package main

import "fmt"

func main() {
    game := Game{X: 10, Y: 10}
    game.init()

    fmt.Println(game)
}
