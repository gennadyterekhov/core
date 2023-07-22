package entity

type Territory struct {
	Food           int    `json:"food"`
	TradingAbility int    `json:"trading_ability"`
	Production     int    `json:"production"`
	Culture        int    `json:"culture"`
	Tiles          []Tile `json:"tiles"`
}

var currencyToIndexMap = map[string]int{
	"food":            0,
	"trading_ability": 1,
	"production":      2,
	"culture":         3,
}

func (territory *Territory) GetTotalFood() int {
	accumulator := 0
	for i := 0; i < len(territory.Tiles); i++ {
		accumulator += territory.Tiles[i].Resource.Food
	}
	return accumulator
}

func (territory *Territory) GetTotalTradingAbility() int {
	accumulator := 0
	for i := 0; i < len(territory.Tiles); i++ {
		accumulator += territory.Tiles[i].Resource.TradingAbility
	}
	return accumulator
}

func (territory *Territory) GetTotalProduction() int {
	accumulator := 0
	for i := 0; i < len(territory.Tiles); i++ {
		accumulator += territory.Tiles[i].Resource.Production
	}
	return accumulator
}

func (territory *Territory) GetTotalCulture() int {
	accumulator := 0
	for i := 0; i < len(territory.Tiles); i++ {
		accumulator += territory.Tiles[i].Resource.Culture
	}
	return accumulator
}

func (territory *Territory) GetTotal(name string) int {
	accumulator := 0

	return accumulator

}
