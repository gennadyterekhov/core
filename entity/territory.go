package entity

type Territory struct {
	Food       int    `json:"food"`
	Wealth     int    `json:"wealth"`
	Production int    `json:"production"`
	Culture    int    `json:"culture"`
	Tiles      []Tile `json:"tiles"`
}

func (territory *Territory) GetTotalFood() int {
	accumulator := 0
	for i := 0; i < len(territory.Tiles); i++ {
		accumulator += territory.Tiles[i].Resource.Food
	}
	return accumulator
}
