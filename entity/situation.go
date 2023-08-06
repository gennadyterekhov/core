package entity

type Situation struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

const SituationsCount = 8

const SituationNameForestFire = "Forest Fire"
const SituationNameFlood = "Flood"
const SituationNameNessie = "Nessie"
const SituationNameBlessing = "Blessing"
const SituationNameDesertification = "Desertification"
const SituationNameDivineShield = "Divine Shield"
const SituationNameDemilitarization = "Demilitarization"
const SituationNameColumbus = "Columbus"

var SituationNames = [SituationsCount]string{
	SituationNameForestFire,
	SituationNameFlood,
	SituationNameNessie,
	SituationNameBlessing,
	SituationNameDesertification,
	SituationNameDivineShield,
	SituationNameDemilitarization,
	SituationNameColumbus,
}

var SituationNameToSituationMap map[string](*Situation) = map[string](*Situation){
	SituationNameForestFire: &Situation{
		Name:        SituationNameForestFire,
		Description: "desc",
	},
	SituationNameFlood: &Situation{
		Name:        SituationNameFlood,
		Description: "desc",
	},
	SituationNameNessie: &Situation{
		Name:        SituationNameNessie,
		Description: "desc",
	},
	SituationNameBlessing: &Situation{
		Name:        SituationNameBlessing,
		Description: "desc",
	},
	SituationNameDesertification: &Situation{
		Name:        SituationNameDesertification,
		Description: "desc",
	},
	SituationNameDivineShield: &Situation{
		Name:        SituationNameDivineShield,
		Description: "desc",
	},
	SituationNameDemilitarization: &Situation{
		Name:        SituationNameDemilitarization,
		Description: "desc",
	},
	SituationNameColumbus: &Situation{
		Name:        SituationNameColumbus,
		Description: "desc",
	},
}
