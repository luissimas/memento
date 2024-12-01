package manager

import (
	"fmt"
	"time"

	"github.com/luissimas/memento/game"
)

// GameManager manages a game and its saves.
type GameManager struct {
	Game  game.Game
	saves []game.Save
}

// NewGameManager creates a new `GameManager`.
func NewGameManager(g game.Game) GameManager {
	return GameManager{Game: g, saves: make([]game.Save, 0)}
}

// Save saves the current game state and stores it on the manager's saves.
func (s *GameManager) Save() {
	s.saves = append(s.saves, s.Game.Save())
}

// PreviousSave restores the game to the last previous known save.
func (s *GameManager) PreviousSave() {
	if len(s.saves) > 0 {
		save := s.saves[len(s.saves)-1]
		fmt.Printf("Restoring game to save made at: %s\n", save.GetDate().Format(time.RFC3339))
		s.Game.Restore(save)
		s.saves = s.saves[:len(s.saves)-1]
	}
}
