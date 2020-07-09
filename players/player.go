package player

import (
	"gttt/game"
)

//Player interface must implement action
type Player interface {
	Action(State) (int, int)
}
