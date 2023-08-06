package entity

import "fmt"

type Technology struct {
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Prerequisites map[string]bool `json:"prerequisites"`
}

const TechnologiesCount = 15

const (
	TechnologyNamePlough             = "Plough"
	TechnologyNamePrimitiveWriting   = "Primitive Writing"
	TechnologyNameAdvancedWriting    = "Advanced Writing"
	TechnologyNameBronzeWeapons      = "Bronze Weapons"
	TechnologyNameStoneWorking       = "Stone Working"
	TechnologyNameIdols              = "Idols"
	TechnologyNamePottery            = "Pottery"
	TechnologyNameAnimalHusbandry    = "Animal Husbandry"
	TechnologyNameHunting            = "Hunting"
	TechnologyNameArchery            = "Archery"
	TechnologyNameOrganizedArmy      = "Organized Army"
	TechnologyNameFishing            = "Fishing"
	TechnologyNameMusicalInstruments = "Musical Instruments"
	TechnologyNamePoetry             = "Poetry"
	TechnologyNameCalendar           = "Calendar"
)

var TechnologyNames = [TechnologiesCount]string{
	TechnologyNamePlough,
	TechnologyNamePrimitiveWriting,
	TechnologyNameAdvancedWriting,
	TechnologyNameBronzeWeapons,
	TechnologyNameStoneWorking,
	TechnologyNameIdols,
	TechnologyNamePottery,
	TechnologyNameAnimalHusbandry,
	TechnologyNameHunting,
	TechnologyNameArchery,
	TechnologyNameOrganizedArmy,
	TechnologyNameFishing,
	TechnologyNameMusicalInstruments,
	TechnologyNamePoetry,
	TechnologyNameCalendar,
}

var TechnologyNameToTechnologyMap map[string](*Technology) = map[string](*Technology){
	TechnologyNamePlough: &Technology{
		Name:          TechnologyNamePlough,
		Description:   "Harvest result +2. Throw 2 dice if " + TechnologyNameAnimalHusbandry + " is discovered",
		Prerequisites: map[string]bool{TechnologyNameCalendar: true},
	},
	TechnologyNamePrimitiveWriting: &Technology{
		Name:          TechnologyNamePrimitiveWriting,
		Description:   fmt.Sprintf("%v +3", CurrencyNameWealth),
		Prerequisites: map[string]bool{TechnologyNamePottery: true},
	},
	TechnologyNameAdvancedWriting: &Technology{
		Name:          TechnologyNameAdvancedWriting,
		Description:   fmt.Sprintf("%v *2", CurrencyNameCulture),
		Prerequisites: map[string]bool{TechnologyNamePrimitiveWriting: true},
	},
	TechnologyNameBronzeWeapons: &Technology{
		Name: TechnologyNameBronzeWeapons,
		Description: fmt.Sprintf("%v +1 %v +1 %v. %v *2",
			ResourceNameMetal,
			CurrencyNameWealth,
			CurrencyNameProduction,
			CurrencyNameCombatReadiness,
		),
		Prerequisites: map[string]bool{TechnologyNameStoneWorking: true},
	},
	TechnologyNameStoneWorking: &Technology{
		Name:        TechnologyNameStoneWorking,
		Description: fmt.Sprintf("%v +2 %v +1 %v", ResourceNameStone, CurrencyNameProduction, CurrencyNameCulture),

		Prerequisites: map[string]bool{},
	},
	TechnologyNameIdols: &Technology{
		Name:          TechnologyNameIdols,
		Description:   fmt.Sprintf("%v +2 %v", ResourceNameStone, CurrencyNameCulture),
		Prerequisites: map[string]bool{TechnologyNameStoneWorking: true},
	},
	TechnologyNamePottery: &Technology{
		Name:          TechnologyNamePottery,
		Description:   "Harvest result +2",
		Prerequisites: map[string]bool{},
	},
	TechnologyNameAnimalHusbandry: &Technology{
		Name:          TechnologyNameAnimalHusbandry,
		Description:   fmt.Sprintf("%v +2 %v", ResourceNamePasture, CurrencyNameFood),
		Prerequisites: map[string]bool{},
	},

	TechnologyNameHunting: &Technology{
		Name:          TechnologyNameHunting,
		Description:   fmt.Sprintf("%v +1 %v", ResourceNameForest, CurrencyNameFood),
		Prerequisites: map[string]bool{},
	},
	TechnologyNameArchery: &Technology{
		Name:          TechnologyNameArchery,
		Description:   fmt.Sprintf("%v *2", CurrencyNameCombatReadiness),
		Prerequisites: map[string]bool{},
	},
	TechnologyNameOrganizedArmy: &Technology{
		Name:          TechnologyNameOrganizedArmy,
		Description:   fmt.Sprintf("%v *3", CurrencyNameCombatReadiness),
		Prerequisites: map[string]bool{TechnologyNameBronzeWeapons: true},
	},
	TechnologyNameFishing: &Technology{
		Name:          TechnologyNameFishing,
		Description:   fmt.Sprintf("%v and %v +2 %v", ResourceNameRiver, ResourceNameLake, CurrencyNameFood),
		Prerequisites: map[string]bool{},
	},
	TechnologyNameMusicalInstruments: &Technology{
		Name:          TechnologyNameMusicalInstruments,
		Description:   fmt.Sprintf("%v *2", CurrencyNameCulture),
		Prerequisites: map[string]bool{},
	},
	TechnologyNamePoetry: &Technology{
		Name:          TechnologyNamePoetry,
		Description:   fmt.Sprintf("%v *2", CurrencyNameCulture),
		Prerequisites: map[string]bool{},
	},
	TechnologyNameCalendar: &Technology{
		Name:          TechnologyNameCalendar,
		Description:   fmt.Sprintf("%v *2", CurrencyNameFood),
		Prerequisites: map[string]bool{TechnologyNamePrimitiveWriting: true},
	},
}
