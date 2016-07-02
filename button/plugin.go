package button

import "github.com/amortaza/go-bellina"

var plugin *Plugin

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "button"
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


