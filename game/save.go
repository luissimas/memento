package game

import "time"

// Save represents a generic save.
type Save interface {
	GetDate() time.Time
}

// gameSave represents the safe of a game state.
type gameSave struct {
	date  time.Time
	state gameState
}

// GetDate retrieves the date when the game was saved.
func (g gameSave) GetDate() time.Time {
	return g.date
}
