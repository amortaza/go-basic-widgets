package hscroll

import (
	"github.com/amortaza/go-bellina"
)

type State struct {
	HScrollId string
	Thickness_ int
}

var gCurState *State

func Id(postfixId string) *State {
	hscrollId := bl.Current_Node.Id + "/" + postfixId

	gCurState = ensureState(hscrollId)

	return gCurState
}

func (s *State) Thickness(thickness int) *State {
	s.Thickness_ = thickness

	return s
}

func (s *State) On(cb func(interface{})) {
	logic(s, cb)
}

func Div() {
	bl.DivId(gCurState.HScrollId)
}

func End() {
	bl.End()
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
