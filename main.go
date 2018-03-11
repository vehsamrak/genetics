package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/vehsamrak/genetics/src/cell"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cell := cell.New()
	fmt.Println(strconv.Itoa(cell.GetLifePoints()))
}
