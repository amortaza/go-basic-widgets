package vscroll

import (
	"github.com/amortaza/go-bellina"
)

type State struct {
	VScrollId string
	Thickness_ int
}

var gCurState *State

func Id(postfixId string) *State {
	Id := bl.Current_Node.Id + "/" + postfixId

	gCurState = ensureState(Id)

	return gCurState
}

func (s *State) Thickness(thickness int) (*State) {
	s.Thickness_ = thickness

	return s
}

func (s *State) On(cb func(interface{})) {
	logic(s, cb)
}

func Div() {
	bl.DivId(gCurState.VScrollId)
}

func End() {
	bl.End()
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
