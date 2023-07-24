package entity

type Population struct {
	Men             int `json:"men"`
	Women           int `json:"women"`
	Total           int `json:"total"`
	CombatReadiness int `json:"combat_readiness"`
	Civilizedness   int `json:"civilizedness"`
}

const DefaultMen = 25
const DefaultWomen = 25
const DefaultCombatReadiness = 25
const DefaultCivilizedness = 25

func createStarterPopulation() Population {
	return Population{
		Men:             DefaultMen,
		Women:           DefaultWomen,
		Total:           DefaultWomen + DefaultMen,
		CombatReadiness: DefaultCombatReadiness,
		Civilizedness:   DefaultCivilizedness,
	}
}
