package main

func main() {
    game := Game{X: 10, Y: 20, Seed: 1}
    game.CreateField()
    game.View()
}
