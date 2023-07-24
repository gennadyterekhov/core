package entity

import (
	"time"
)

type Game struct {
	Players          []Player          `json:"players"`
	StartDate        time.Time         `json:"start_date"`
	EndDate          time.Time         `json:"end_date"`
	Name             string            `json:"name"`
	IsFinished       bool              `json:"is_finished"`
	Winner           *Player           `json:"winner"`
	WinningCondition *WinningCondition `json:"winning_condition"`
}

func StartNewGame(name string, playerNames []string) *Game {
	game := Game{
		Players:          createPlayers(playerNames),
		StartDate:        time.Now().Local(),
		Name:             name,
		IsFinished:       false,
		Winner:           nil,
		WinningCondition: nil,
	}
	return &game
}

func (game *Game) Finish() {
	game.IsFinished = true
	game.EndDate = time.Now().Local()
	game.WinningCondition = game.calculateWinningCondition()
	game.Winner = game.calculateWinner()
}

func (game *Game) calculateWinningCondition() *WinningCondition {
	// TODO: implement
	return WinningConditionAliasToWinningConditionMap[WinningConditionMilitaryAlias]
}

func (game *Game) calculateWinner() *Player {
	if game.WinningCondition == nil {
		return nil
	}
	// TODO: implement
	return &game.Players[0]
}

func createPlayers(playerNames []string) []Player {
	players := make([]Player, len(playerNames))
	tribeNames := TribeNames[:len(playerNames)]

	for i := 0; i < len(playerNames); i++ {
		players[i] = Player{
			Tribe: Tribe{
				Name:       tribeNames[i],
				Wealth:     0,
				Points:     0,
				Population: createStarterPopulation(),
				Territory:  createStarterTerritory(),
			},
			Name: playerNames[i],
		}
	}
	return players
}
