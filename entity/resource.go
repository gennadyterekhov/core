package entity

type Resource struct {
	Name           string `json:"name"`
	Quantity       int    `json:"quantity"` // how many cards of this type are there
	Food           int    `json:"food"`
	TradingAbility int    `json:"trading_ability"`
	Production     int    `json:"production"`
	Culture        int    `json:"culture"`
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
	ResourceNamePasture,
	ResourceNameStone,
	ResourceNameMetal,
	ResourceNameFruit,
	ResourceNameLake,
	ResourceNameGold,
	ResourceNameSilver,
	ResourceNameForest,
	ResourceNameDesert,
	ResourceNameRiver,
}

var ResourceNameToResourceMap map[string](*Resource) = map[string](*Resource){
	ResourceNamePasture: &Resource{
		Name:           ResourceNamePasture,
		Quantity:       10,
		Food:           3,
		TradingAbility: 0,
		Production:     0,
		Culture:        0,
	},
	ResourceNameStone: &Resource{
		Name:           ResourceNameStone,
		Quantity:       10,
		Food:           0,
		TradingAbility: 0,
		Production:     2,
		Culture:        1,
	},
	ResourceNameMetal: &Resource{
		Name:           ResourceNameMetal,
		Quantity:       10,
		Food:           0,
		TradingAbility: 1,
		Production:     2,
		Culture:        0,
	},
	ResourceNameFruit: &Resource{
		Name:           ResourceNameFruit,
		Quantity:       8,
		Food:           2,
		TradingAbility: 1,
		Production:     0,
		Culture:        0,
	},
	ResourceNameLake: &Resource{
		Name:           ResourceNameLake,
		Quantity:       4,
		Food:           2,
		TradingAbility: 0,
		Production:     0,
		Culture:        1,
	},
	ResourceNameGold: &Resource{
		Name:           ResourceNameGold,
		Quantity:       2,
		Food:           0,
		TradingAbility: 4,
		Production:     0,
		Culture:        0,
	},
	ResourceNameSilver: &Resource{
		Name:           ResourceNameSilver,
		Quantity:       4,
		Food:           0,
		TradingAbility: 2,
		Production:     0,
		Culture:        1,
	},
	ResourceNameForest: &Resource{
		Name:           ResourceNameForest,
		Quantity:       20,
		Food:           1,
		TradingAbility: 0,
		Production:     2,
		Culture:        0,
	},
	ResourceNameDesert: &Resource{
		Name:           ResourceNameDesert,
		Quantity:       6,
		Food:           0,
		TradingAbility: 1,
		Production:     0,
		Culture:        1,
	},
	ResourceNameRiver: &Resource{
		Name:           ResourceNameRiver,
		Quantity:       6,
		Food:           1,
		TradingAbility: 2,
		Production:     0,
		Culture:        1,
	},
}
