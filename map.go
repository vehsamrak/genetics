package main

type Game struct {
    X int
    Y int
    Command int
    field map[int] map[int] int
}

func (game *Game) init() {
    game.createField()
}

func (game *Game) createField() {
    var gameFieldY map[int] int
    gameFieldY = make(map[int]int, game.Y)

    game.field = make(map[int] map[int] int, game.X)

    for x := 0; x < game.X ;x++ {
        for y := 0; y < game.Y ; y++ {
            gameFieldY[y] = game.Command
        }

        game.field[x] = gameFieldY
    }
}

func (game *Game) view() {

}
