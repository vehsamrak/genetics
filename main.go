package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	Cell "github.com/vehsamrak/genetics/src/cell"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cell := Cell.New()
	fmt.Println(strconv.Itoa(cell.GetLifePoints()))
}
