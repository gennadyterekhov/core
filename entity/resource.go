package entity

type Resource struct {
	Quantity       int `json:"quantity"` // how many cards of this type are there
	Food           int `json:"food"`
	TradingAbility int `json:"trading_ability"`
	Production     int `json:"production"`
	Culture        int `json:"culture"`
}

const ResourcesCount = 10

const ResourceNamePasture = "Pasture"
const ResourceNameStone = "Stone"
const ResourceNameMetal = "Metal"
const ResourceNameFruit = "Fruit"
const ResourceNameLake = "Lake"
const ResourceNameGold = "Gold"
const ResourceNameSilver = "Silver"
const ResourceNameForest = "Forest"
const ResourceNameDesert = "Desert"
const ResourceNameRiver = "River"

var ResourceNames = [ResourcesCount]string{
	"Pasture",
	"Stone",
	"Metal",
	"Fruit",
	"Lake",
	"Gold",
	"Silver",
	"Forest",
	"Desert",
	"River",
}

var (
	Pasture Resource = Resource{
		Quantity:       10,
		Food:           3,
		TradingAbility: 0,
		Production:     0,
		Culture:        0,
	}
	Stone Resource = Resource{
		Quantity:       10,
		Food:           0,
		TradingAbility: 0,
		Production:     2,
		Culture:        1,
	}
	Metal Resource = Resource{
		Quantity:       10,
		Food:           0,
		TradingAbility: 1,
		Production:     2,
		Culture:        0,
	}
	Fruit Resource = Resource{
		Quantity:       8,
		Food:           2,
		TradingAbility: 1,
		Production:     0,
		Culture:        0,
	}
	Lake Resource = Resource{
		Quantity:       4,
		Food:           2,
		TradingAbility: 0,
		Production:     0,
		Culture:        1,
	}
	Gold Resource = Resource{
		Quantity:       2,
		Food:           0,
		TradingAbility: 4,
		Production:     0,
		Culture:        0,
	}
	Silver Resource = Resource{
		Quantity:       4,
		Food:           0,
		TradingAbility: 2,
		Production:     0,
		Culture:        1,
	}
	Forest Resource = Resource{
		Quantity:       20,
		Food:           1,
		TradingAbility: 0,
		Production:     2,
		Culture:        0,
	}
	Desert Resource = Resource{
		Quantity:       6,
		Food:           0,
		TradingAbility: 1,
		Production:     0,
		Culture:        1,
	}
	River Resource = Resource{
		Quantity:       6,
		Food:           1,
		TradingAbility: 2,
		Production:     0,
		Culture:        1,
	}
)

var ResourceNameToResourceMap map[string](*Resource) = map[string](*Resource){
	"Pasture": &Pasture,
	"Stone":   &Stone,
	"Metal":   &Metal,
	"Fruit":   &Fruit,
	"Lake":    &Lake,
	"Gold":    &Gold,
	"Silver":  &Silver,
	"Forest":  &Forest,
	"Desert":  &Desert,
	"River":   &River,
}
