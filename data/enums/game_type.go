package enums

type GameType uint8

const (
	Survival GameType = iota
	Creative
	Adventure
	Spectator
)

func (g GameType) String() string {
	return [...]string{"Survival", "Creative", "Adventure", "Spectator"}[g]
}
