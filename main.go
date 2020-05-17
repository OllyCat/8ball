package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//ebiten.SetWindowSize(450, 450)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Magic ball")

	game := &Game{}
	game.Init()
	ebiten.SetMaxTPS(30)

	if err := ebiten.RunGame(game); err != nil {
		log.Println(err)
	}
}
