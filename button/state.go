package button


var gStateByNode  map[string] *State

func ensureState(buttonId string) *State {
	state, ok := gStateByNode[buttonId]

	if !ok {
		state = &State{ButtonId: buttonId, Label_: "Ace"}

		gStateByNode[buttonId] = state
	}

	return state
}
