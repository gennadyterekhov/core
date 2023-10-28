package entity

type Action struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Radius      int    `json:"radius"`
}

const ActionsCount = 19 // circa
const (
	ActionNameForestFire            = "Forest Fire"
	ActionNameFlood                 = "Flood"
	ActionNameNessie                = "Nessie"
	ActionNameBlessing              = "Blessing"
	ActionNameDesertification       = "Desertification"
	ActionNameDivineShield          = "Divine Shield"
	ActionNamePacifism              = "Pacifism"
	ActionNameColumbus              = "Columbus"
	ActionNamePoliticalInstablility = "Vulnerability"
	ActionNameNothing               = "Nothing"
)

var ActionNames = [ActionsCount]string{
	ActionNameForestFire,
	ActionNameFlood,
	ActionNameNessie,
	ActionNameBlessing,
	ActionNameDesertification,
	ActionNameDivineShield,
	ActionNamePacifism,
	ActionNameColumbus,
	ActionNamePoliticalInstablility,
	ActionNameNothing,
}

var ActionNameToActionMap map[string](*Action) = map[string](*Action){
	ActionNameForestFire: &Action{
		Name:        ActionNameForestFire,
		Description: "Remove " + ResourceNameForest + " tile. If you don't have one - remove one from the next player",
		Radius:      4,
	},
	ActionNameFlood: &Action{
		Name:        ActionNameFlood,
		Description: "Remove " + ResourceNameRiver + " tile. If you don't have one - remove one from the next player",
		Radius:      5,
	},
	ActionNameNessie: &Action{
		Name:        ActionNameNessie,
		Description: "Remove " + ResourceNameLake + " tile. If you don't have one - remove one from the next player",
		Radius:      1,
	},
	ActionNameDesertification: &Action{
		Name: ActionNameDesertification,
		Description: "Turn one " +
			ResourceNamePasture +
			" tile into " +
			ResourceNameDesert +
			". If you don't have one - do so with one from the next player",
		Radius: 2,
	},
	ActionNameBlessing: &Action{
		Name:        ActionNameBlessing,
		Description: "+1 Action this round",
		Radius:      2,
	},
	ActionNameDivineShield: &Action{
		Name:        ActionNameDivineShield,
		Description: "You win every fight if defending",
		Radius:      2,
	},
	ActionNamePacifism: &Action{
		Name:        ActionNamePacifism,
		Description: "Nobody can arm next round",
		Radius:      5,
	},
	ActionNameColumbus: &Action{
		Name:        ActionNameColumbus,
		Description: "Only you can discover new tiles next round",
		Radius:      2,
	},
	ActionNamePoliticalInstablility: &Action{
		Name:        ActionNamePoliticalInstablility,
		Description: "Rome is weaker 1000pts next round",
		Radius:      2,
	},
	ActionNameNothing: &Action{
		Name:        ActionNameNothing,
		Description: "Nothing happens",
		Radius:      10,
	},
}
