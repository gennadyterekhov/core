package core

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
