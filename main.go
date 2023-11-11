package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lucaspiller/orbital-go/game"
)

func main() {
	game := game.NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Orbital")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
