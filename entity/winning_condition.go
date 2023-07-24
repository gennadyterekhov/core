package entity

type WinningCondition struct {
	Condition   string `json:"condition"`
	WinningText string `json:"winning_text"`
}

const WinningConditionMilitaryAlias = "military"
const WinningConditionMilitaryText = "Your tribe has successfully conquered a state. It will be remembered not as barbaric, but as heroic."

const WinningConditionCulturalAlias = "cultural"
const WinningConditionCulturalText = "Your tribe has successfully been idolized by a state. It will be remembered not as barbaric, but as heroic."

const WinningConditionPoliticalAlias = "political"
const WinningConditionPoliticalText = "Your tribe has successfully infiltrated and subjugated a state. It will be remembered not as barbaric, but as heroic."

var WinningConditionAliasToWinningConditionMap = map[string]*WinningCondition{
	WinningConditionMilitaryAlias:  &WinningCondition{Condition: WinningConditionMilitaryAlias, WinningText: WinningConditionMilitaryText},
	WinningConditionCulturalAlias:  &WinningCondition{Condition: WinningConditionCulturalAlias, WinningText: WinningConditionCulturalText},
	WinningConditionPoliticalAlias: &WinningCondition{Condition: WinningConditionPoliticalAlias, WinningText: WinningConditionPoliticalText},
}
