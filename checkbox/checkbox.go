package checkbox

import "github.com/amortaza/go-bellina"

type CheckboxState int

const (
	Unchecked CheckboxState = iota
	Semi
	Checked
)

func Id(postfixCheckboxId string) *State {
	checkboxId := bl.Current_Node.Id + "/" + postfixCheckboxId

	return EnsureState(checkboxId)
}

type State struct {
	CheckboxId string
	Label string
	CheckState CheckboxState
	checkStateAtStart CheckboxState
}

func (s *State) SetLabel(label string) *State {
	s.Label = label

	return s
}

func (s *State) On(cb func(interface{})) {
	logic(s, cb)
}

func EnsureState(checkboxId string) *State {
	state, ok := gStateByNode[checkboxId]

	if !ok {
		state = &State{CheckboxId: checkboxId}
		gStateByNode[checkboxId] = state
	}

	return state
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
