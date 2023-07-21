package core

type Population struct {
	Men             int `json:"men"`
	Women           int `json:"women"`
	Total           int `json:"total"`
	CombatReadiness int `json:"combat_readiness"`
	Civilizedness   int `json:"civilizedness"`
}
