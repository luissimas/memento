package game

import (
	"fmt"
	"time"
)

// gameState represents the state of a game.
type gameState struct {
	level uint
	score uint
}

// Game represents a generic game.
type Game struct {
	state gameState
}

// NewGame creates a new `Game`.
func NewGame(level, score uint) Game {
	return Game{state: gameState{level: level, score: score}}
}

// Save generates a new save based on the current game state.
func (g *Game) Save() gameSave {
	return gameSave{state: g.state, date: time.Now()}
}

// SetScore mutates the game score.
func (g *Game) SetScore(score uint) {
	g.state.score = score
}

// SetScore mutates the game level.
func (g *Game) SetLevel(level uint) {
	g.state.level = level
}

// Restore restores the game state from a `save`.
func (g *Game) Restore(save Save) {
	v, ok := save.(gameSave)
	if !ok {
		panic("Save is not a GameSave")
	}
	g.state = v.state
}

// String returns the string representation of a game.
func (g Game) String() string {
	return fmt.Sprintf("Level: %d\nScore: %d", g.state.level, g.state.score)
}
