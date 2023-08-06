package entity

import "fmt"

type Technology struct {
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Prerequisites map[string]bool `json:"prerequisites"`
}

const TechnologiesCount = 15

const (
	TechnologyNamePlough             = "Forest Fire"
	TechnologyNamePrimitiveWriting   = "Flood"
	TechnologyNameAdvancedWriting    = "Nessie"
	TechnologyNameBronzeWeapons      = "Blessing"
	TechnologyNameStoneWorking       = "Desertification"
	TechnologyNameIdols              = "Divine Shield"
	TechnologyNamePottery            = "Demilitarization"
	TechnologyNameAnimalHusbandry    = "Columbus"
	TechnologyNameHunting            = "Hunting"
	TechnologyNameArchery            = "Archery"
	TechnologyNameOrganizedArmy      = "OrganizedArmy"
	TechnologyNameFishing            = "Fishing"
	TechnologyNameMusicalInstruments = "MusicalInstruments"
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
		Description:   "Add 2 to harvest result. Throw 2 dice if " + TechnologyNameAnimalHusbandry + " is discovered",
		Prerequisites: map[string]bool{TechnologyNameCalendar: true},
	},
	TechnologyNamePrimitiveWriting: &Technology{
		Name:          TechnologyNamePrimitiveWriting,
		Description:   "+3 Wealth",
		Prerequisites: map[string]bool{TechnologyNamePottery: true},
	},
	TechnologyNameAdvancedWriting: &Technology{
		Name:          TechnologyNameAdvancedWriting,
		Description:   "Culture * 2",
		Prerequisites: map[string]bool{TechnologyNamePrimitiveWriting: true},
	},
	TechnologyNameBronzeWeapons: &Technology{
		Name:          TechnologyNameBronzeWeapons,
		Description:   ResourceNameMetal + " +1 wealth +1 production. Combat readiness * 2",
		Prerequisites: map[string]bool{TechnologyNameStoneWorking: true},
	},
	TechnologyNameStoneWorking: &Technology{
		Name:          TechnologyNameStoneWorking,
		Description:   ResourceNameStone + "+2 production +1culture",
		Prerequisites: map[string]bool{},
	},
	TechnologyNameIdols: &Technology{
		Name:          TechnologyNameIdols,
		Description:   ResourceNameStone + "+2 culture",
		Prerequisites: map[string]bool{TechnologyNameStoneWorking: true},
	},
	TechnologyNamePottery: &Technology{
		Name:          TechnologyNamePottery,
		Description:   "Add 2 to harvest result",
		Prerequisites: map[string]bool{},
	},
	TechnologyNameAnimalHusbandry: &Technology{
		Name:          TechnologyNameAnimalHusbandry,
		Description:   ResourceNamePasture + " +2 " + CurrencyNameFood,
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
