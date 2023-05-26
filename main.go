package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/amjadjibon/2048/game"
)

func main() {
	newGame, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("2048")
	if err := ebiten.RunGame(newGame); err != nil {
		log.Fatal(err)
	}
}
