package entity

import (
	"math/rand"
	"time"
)

type Tribe struct {
	Population Population `json:"population"`
	Name       string     `json:"name"`
	Wealth     int        `json:"wealth"`
	Points     int        `json:"points"`
	Territory  Territory  `json:"territory"`
}

func (tribe *Tribe) GetNewPopulationCount(fertility int) int {
	food := tribe.Territory.GetTotalFood()
	cropsYield := food * fertility
	upperBound := tribe.Population.Total * 10

	if cropsYield < upperBound {
		return cropsYield
	}
	return upperBound
}

func (tribe *Tribe) MakeTerritorialDiscovery() {
	newTile := discoverNewTile()
	tribe.Territory.addTile(newTile)
	tribe.Territory.updateResources()
}

func discoverNewTile() Tile {
	newTile := Tile{
		Resource: getRandomResource(),
	}
	return newTile
}

func getRandomResource() *Resource {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := randomGenerator.Intn(ResourcesCount)
	randomName := ResourceNames[randomIndex]
	randomResource := ResourceNameToResourceMap[randomName]

	return randomResource
}
