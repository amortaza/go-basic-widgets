package hsplit

type State struct {
	HSplitId string
}

var gStateByWidgetId map[string] *State

func (s *State) On(cb func(interface{})) {
	logic(s, cb)
}