package entity

import (
	"math/rand"
	"time"
)

type Tribe struct {
	Name       string     `json:"name"`
	Wealth     int        `json:"wealth"`
	Points     int        `json:"points"`
	Population Population `json:"population"`
	Territory  Territory  `json:"territory"`
}

const TribesCount = 8

const TribeNameNorth = "North"
const TribeNameNortheast = "Northeast"
const TribeNameEast = "East"
const TribeNameSoutheast = "Southeast"
const TribeNameSouth = "South"
const TribeNameSouthwest = "Southwest"
const TribeNameWest = "West"
const TribeNameNorthwest = "Northwest"

var TribeNames = [TribesCount]string{
	"North",
	"Northeast",
	"East",
	"Southeast",
	"South",
	"Southwest",
	"West",
	"Northwest",
}

var TribeNameToAliasMap = map[string]string{
	"North":     "Saami",
	"Northeast": "Chukchi",
	"East":      "Chinese",
	"Southeast": "Javanese",
	"South":     "Zulu",
	"Southwest": "Lusitanians",
	"West":      "Lakota",
	"Northwest": "Aleut",
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
