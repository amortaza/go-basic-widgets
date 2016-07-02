package edit

type State struct {
	EditId string
	CursorPos int
	HasFocus bool
	Width int
}

func (s *State) On() {
	Div()
	End()
}

