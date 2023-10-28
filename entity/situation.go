package entity

type Situation struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"` // how many cards are there
}

const SituationsCount = 10

const SituationNameForestFire = "Forest Fire"
const SituationNameFlood = "Flood"
const SituationNameNessie = "Nessie"
const SituationNameBlessing = "Blessing"
const SituationNameDesertification = "Desertification"
const SituationNameDivineShield = "Divine Shield"
const SituationNamePacifism = "Pacifism"
const SituationNameColumbus = "Columbus"
const SituationNamePoliticalInstablility = "Vulnerability"
const SituationNameNothing = "Nothing"

var SituationNames = [SituationsCount]string{
	SituationNameForestFire,
	SituationNameFlood,
	SituationNameNessie,
	SituationNameBlessing,
	SituationNameDesertification,
	SituationNameDivineShield,
	SituationNamePacifism,
	SituationNameColumbus,
	SituationNamePoliticalInstablility,
	SituationNameNothing,
}

var SituationNameToSituationMap map[string](*Situation) = map[string](*Situation){
	SituationNameForestFire: &Situation{
		Name:        SituationNameForestFire,
		Description: "Remove " + ResourceNameForest + " tile. If you don't have one - remove one from the next player",
		Quantity:    4,
	},
	SituationNameFlood: &Situation{
		Name:        SituationNameFlood,
		Description: "Remove " + ResourceNameRiver + " tile. If you don't have one - remove one from the next player",
		Quantity:    5,
	},
	SituationNameNessie: &Situation{
		Name:        SituationNameNessie,
		Description: "Remove " + ResourceNameLake + " tile. If you don't have one - remove one from the next player",
		Quantity:    1,
	},
	SituationNameDesertification: &Situation{
		Name: SituationNameDesertification,
		Description: "Turn one " +
			ResourceNamePasture +
			" tile into " +
			ResourceNameDesert +
			". If you don't have one - do so with one from the next player",
		Quantity: 2,
	},
	SituationNameBlessing: &Situation{
		Name:        SituationNameBlessing,
		Description: "+1 Action this round",
		Quantity:    2,
	},
	SituationNameDivineShield: &Situation{
		Name:        SituationNameDivineShield,
		Description: "You win every fight if defending",
		Quantity:    2,
	},
	SituationNamePacifism: &Situation{
		Name:        SituationNamePacifism,
		Description: "Nobody can arm next round",
		Quantity:    5,
	},
	SituationNameColumbus: &Situation{
		Name:        SituationNameColumbus,
		Description: "Only you can discover new tiles next round",
		Quantity:    2,
	},
	SituationNamePoliticalInstablility: &Situation{
		Name:        SituationNamePoliticalInstablility,
		Description: "Rome is weaker 1000pts next round",
		Quantity:    2,
	},
	SituationNameNothing: &Situation{
		Name:        SituationNameNothing,
		Description: "Nothing happens",
		Quantity:    10,
	},
}
