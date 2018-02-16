package main

func main() {
    game := Game{X: 10, Y: 10, FoodPercent: 10}
    game.CreateField()
    game.View()
}
