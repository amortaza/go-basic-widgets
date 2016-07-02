package checkbox

import "github.com/amortaza/go-bellina"

type CheckboxState int

const (
	Unchecked CheckboxState = iota
	Semi
	Checked
)

type State struct {
	Label string
	CheckState CheckboxState
	checkStateAtStart CheckboxState
}

func Id(id string) *Plugin {
	gId = bl.Current_Node.Id + "/" + id

	return plugin
}

func SetLabel(label string) *Plugin {
	ensureState(gId).Label = label

	return plugin
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
