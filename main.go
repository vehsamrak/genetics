package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	Bacterium "github.com/vehsamrak/genetics/src/bacterium"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	bacterium := Bacterium.New()
	fmt.Println(strconv.Itoa(bacterium.GetLifePoints()))
}
