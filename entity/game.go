package entity

import "time"

type Game struct {
	Players          []Tribe           `json:"players"`
	Date             time.Time         `json:"date"`
	Name             string            `json:"name"`
	IsFinished       bool              `json:"is_finished"`
	Winner           *Tribe            `json:"winner"`
	WinningCondition *WinningCondition `json:"winning_condition"`
}
