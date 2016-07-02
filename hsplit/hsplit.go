package hsplit

import "github.com/amortaza/go-bellina"

var gCurState *State

func Id(postfixId string) *State {
	Id := bl.Current_Node.Id + "/" + postfixId

	gCurState = ensureState(Id)

	return gCurState
}

func Use(id string) {
	Id(id).On(nil)
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
