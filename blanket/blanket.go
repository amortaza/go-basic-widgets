package blanket

import (
	"github.com/amortaza/go-bellina"
)

type State struct {
	BlanketId string
}

// this is used for the "local" Div() function below
var gBlanketId string

func Id(postfixBlanketId string) *State {
	blanketId := bl.Current_Node.Id + "/" + postfixBlanketId

	gBlanketId = blanketId

	return EnsureState(blanketId)
}

// make sure blanket.Id(id) was called before this

func Div() {
	bl.DivId(gBlanketId)
}

func End() {
	bl.End()
}

func EnsureState(blanketId string) *State {
	state, ok := gStateByNode[blanketId]

	if !ok {
		state = &State{BlanketId: blanketId}
		gStateByNode[blanketId] = state
	}

	return state
}

func (s *State) Use() (*State) {
	logic(s)

	return s
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
