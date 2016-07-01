package blanket

import (
	"github.com/amortaza/go-bellina"
)

func Id(id string) *Plugin {
	gId = id

	return plugin
}

func (c *Plugin) Use() {
	c.On(nil)
}

// make sure blanket.Id(id) was called before this
func Div() {
	bl.DivId(gId)
}

func End() {
	bl.End()
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
