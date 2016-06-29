package button

type ButtonState struct {
	IsHover bool
	IsDown bool

	onClick func()
}

var g_stateByButtonId map[string] *ButtonState

func newState() *ButtonState {
	return &ButtonState{}
}

func getOrCreateState(buttonId string) *ButtonState {
	state, ok := g_stateByButtonId[buttonId]

	if !ok {
		state = newState()
		g_stateByButtonId[buttonId] = state
	}

	return state
}