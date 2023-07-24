package entity

type Territory struct {
	Food           int    `json:"food"`
	TradingAbility int    `json:"trading_ability"`
	Production     int    `json:"production"`
	Culture        int    `json:"culture"`
	Tiles          []Tile `json:"tiles"`
}

func createStarterTerritory() Territory {
	territory := Territory{
		Food:           0,
		TradingAbility: 0,
		Production:     0,
		Culture:        0,
		Tiles:          createStarterTiles(),
	}
	territory.updateResources()

	return territory
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

func (territory *Territory) addTile(newTile Tile) {
	territory.Tiles = append(territory.Tiles, newTile)
}

func (territory *Territory) updateResources() {
	territory.Food = territory.GetTotalFood()
	territory.TradingAbility = territory.GetTotalTradingAbility()
	territory.Production = territory.GetTotalProduction()
	territory.Culture = territory.GetTotalCulture()
}
