package enums

type GameType int

const (
	Survival GameType = iota
	Creative
	Adventure
	Spectator
)

func (g GameType) String() string {
	return [...]string{"Survival", "Creative", "Adventure", "Spectator"}[g]
}

func (g GameType) ID() int {
	return int(g)
}
