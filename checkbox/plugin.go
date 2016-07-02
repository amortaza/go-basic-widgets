package checkbox

import "github.com/amortaza/go-bellina"

var plugin *Plugin

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "checkbox"
}

func (c *Plugin) Init() {
	gStateByNode = make(map[string] *State)
}

func (c *Plugin) Tick() {
}

func (c *Plugin) Reset() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On(cb func(interface{})) {
	logic(cb)
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supoorted in checkbox")
}

