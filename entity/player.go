package entity

type Player struct {
	Tribe Tribe  `json:"tribe"`
	Name  string `json:"name"`
}
