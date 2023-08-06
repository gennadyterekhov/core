package entity

type Tile struct {
	Resource *Resource `json:"resource"`
}

const DefaultTilesCount = 2

func createStarterTiles() []Tile {
	return []Tile{
		Tile{ResourceNameToResourceMap[ResourceNamePasture]},
		Tile{ResourceNameToResourceMap[ResourceNameForest]},
	}
}
