package entity

type Resource struct {
	Food           int `json:"food"`
	TradingAbility int `json:"trading_ability"`
	Production     int `json:"production"`
	Culture        int `json:"culture"`
}

var (
	Pasture Resource = Resource{
		Food:           3,
		TradingAbility: 0,
		Production:     0,
		Culture:        0,
	}
	Stone Resource = Resource{
		Food:           0,
		TradingAbility: 0,
		Production:     2,
		Culture:        1,
	}
	Metal Resource = Resource{
		Food:           0,
		TradingAbility: 1,
		Production:     2,
		Culture:        0,
	}
	Fruit Resource = Resource{
		Food:           2,
		TradingAbility: 1,
		Production:     0,
		Culture:        0,
	}
	Lake Resource = Resource{
		Food:           2,
		TradingAbility: 0,
		Production:     0,
		Culture:        1,
	}
	Gold Resource = Resource{
		Food:           0,
		TradingAbility: 3,
		Production:     0,
		Culture:        0,
	}
	Silver Resource = Resource{
		Food:           0,
		TradingAbility: 2,
		Production:     0,
		Culture:        1,
	}
	Forest Resource = Resource{
		Food:           1,
		TradingAbility: 0,
		Production:     2,
		Culture:        0,
	}
	Desert Resource = Resource{
		Food:           0,
		TradingAbility: 1,
		Production:     0,
		Culture:        1,
	}
	River Resource = Resource{
		Food:           1,
		TradingAbility: 2,
		Production:     0,
		Culture:        1,
	}
)
