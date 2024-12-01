package main

import (
	"fmt"
	"time"

	"github.com/luissimas/memento/game"
	"github.com/luissimas/memento/manager"
)

func main() {
	manager := manager.NewGameManager(game.NewGame(1, 0))

	// Save the initial game state
	manager.Save()
	time.Sleep(3 * time.Second)

	// Set a new score
	manager.Game.SetScore(3)
	manager.Save()
	time.Sleep(3 * time.Second)

	// Set a new score and level up
	manager.Game.SetScore(9)
	manager.Game.SetLevel(2)

	// Rollback the state
	fmt.Println("---------")
	fmt.Println(manager.Game)
	fmt.Println("---------")
	manager.PreviousSave()
	fmt.Println(manager.Game)
	fmt.Println("---------")
	manager.PreviousSave()
	fmt.Println(manager.Game)
	fmt.Println("---------")
}
