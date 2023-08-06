package entity

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
		Description:   "desc",
		Prerequisites: map[string]bool{TechnologyNameCalendar: true},
	},
	TechnologyNamePrimitiveWriting: &Technology{
		Name:          TechnologyNamePrimitiveWriting,
		Description:   "desc",
		Prerequisites: map[string]bool{TechnologyNamePottery: true},
	},
	TechnologyNameAdvancedWriting: &Technology{
		Name:          TechnologyNameAdvancedWriting,
		Description:   "desc",
		Prerequisites: map[string]bool{TechnologyNamePrimitiveWriting: true},
	},
	TechnologyNameBronzeWeapons: &Technology{
		Name:          TechnologyNameBronzeWeapons,
		Description:   "desc",
		Prerequisites: map[string]bool{TechnologyNameStoneWorking: true},
	},
	TechnologyNameStoneWorking: &Technology{
		Name:          TechnologyNameStoneWorking,
		Description:   "desc",
		Prerequisites: map[string]bool{},
	},
	TechnologyNameIdols: &Technology{
		Name:          TechnologyNameIdols,
		Description:   "desc",
		Prerequisites: map[string]bool{TechnologyNameStoneWorking: true},
	},
	TechnologyNamePottery: &Technology{
		Name:          TechnologyNamePottery,
		Description:   "desc",
		Prerequisites: map[string]bool{},
	},
	TechnologyNameAnimalHusbandry: &Technology{
		Name:          TechnologyNameAnimalHusbandry,
		Description:   "desc",
		Prerequisites: map[string]bool{},
	},

	TechnologyNameHunting: &Technology{
		Name:          TechnologyNameHunting,
		Description:   "desc",
		Prerequisites: map[string]bool{},
	},
	TechnologyNameArchery: &Technology{
		Name:          TechnologyNameArchery,
		Description:   "desc",
		Prerequisites: map[string]bool{},
	},
	TechnologyNameOrganizedArmy: &Technology{
		Name:          TechnologyNameOrganizedArmy,
		Description:   "desc",
		Prerequisites: map[string]bool{TechnologyNameBronzeWeapons: true},
	},
	TechnologyNameFishing: &Technology{
		Name:          TechnologyNameFishing,
		Description:   "desc",
		Prerequisites: map[string]bool{},
	},
	TechnologyNameMusicalInstruments: &Technology{
		Name:          TechnologyNameMusicalInstruments,
		Description:   "desc",
		Prerequisites: map[string]bool{},
	},
	TechnologyNamePoetry: &Technology{
		Name:          TechnologyNamePoetry,
		Description:   "desc",
		Prerequisites: map[string]bool{},
	},
	TechnologyNameCalendar: &Technology{
		Name:          TechnologyNameCalendar,
		Description:   "desc",
		Prerequisites: map[string]bool{TechnologyNamePrimitiveWriting: true},
	},
}
