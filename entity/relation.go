package entity

type Relation struct {
	Name           string `json:"name"`
	AgentBonus     int    `json:"agent_bonus"`
	RecipientBonus int    `json:"recipient_bonus"`
}

const RelationsCount = 10

const RelationNameCannibals = "Cannibals"
const RelationNameBarbarians = "Barbarians"
const RelationNameRogues = "Rogues"
const RelationNameBourgeois = "Bourgeois"
const RelationNameKnowItAlls = "KnowItAlls"
const RelationNameEquals = "Equals"
const RelationNameRespectable = "Respectable"
const RelationNameMates = "Mates"
const RelationNameProteges = "Proteges"
const RelationNameIdols = "Idols"

var RelationNames = [RelationsCount]string{
	"Cannibals",
	"Barbarians",
	"Rogues",
	"Bourgeois",
	"KnowItAlls",
	"Equals",
	"Respectable",
	"Mates",
	"Proteges",
	"Idols",
}

var (
	Cannibals Relation = Relation{
		Name:           RelationNameCannibals,
		AgentBonus:     0,
		RecipientBonus: -3,
	}
	Barbarians Relation = Relation{
		Name:           RelationNameBarbarians,
		AgentBonus:     1,
		RecipientBonus: -2,
	}
	Rogues Relation = Relation{
		Name:           RelationNameRogues,
		AgentBonus:     2,
		RecipientBonus: -1,
	}
	Bourgeois Relation = Relation{
		Name:           RelationNameBourgeois,
		AgentBonus:     2,
		RecipientBonus: 0,
	}
	KnowItAlls Relation = Relation{
		Name:           RelationNameKnowItAlls,
		AgentBonus:     1,
		RecipientBonus: -1,
	}
	Equals Relation = Relation{
		Name:           RelationNameEquals,
		AgentBonus:     0,
		RecipientBonus: 0,
	}
	Respectable Relation = Relation{
		Name:           RelationNameRespectable,
		AgentBonus:     1,
		RecipientBonus: 2,
	}
	Mates Relation = Relation{
		Name:           RelationNameMates,
		AgentBonus:     2,
		RecipientBonus: 2,
	}
	Proteges Relation = Relation{
		Name:           RelationNameProteges,
		AgentBonus:     1,
		RecipientBonus: 4,
	}
	Idols Relation = Relation{
		Name:           RelationNameIdols,
		AgentBonus:     3,
		RecipientBonus: 2,
	}
)

var RelationNameToRelationMap map[string](*Relation) = map[string](*Relation){
	"Cannibals":   &Cannibals,
	"Barbarians":  &Barbarians,
	"Rogues":      &Rogues,
	"Bourgeois":   &Bourgeois,
	"KnowItAlls":  &KnowItAlls,
	"Equals":      &Equals,
	"Respectable": &Respectable,
	"Mates":       &Mates,
	"Proteges":    &Proteges,
	"Idols":       &Idols,
}
