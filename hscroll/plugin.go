package hscroll

import (
	"github.com/amortaza/go-bellina"
)

var NAME = "hscroll"

var plugin *Plugin

type Plugin struct {
}

type Event struct {
	PercentStart float32
	PercentEnd float32
}

func (c *Plugin) Name() string {
	return NAME
}

func (c *Plugin) Init() {
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
	logic("one", cb)
}

func (c *Plugin) On2(cb func(interface{}), onDown func(interface{}), onUpAndMiss func(interface{})) {
	panic("On2 not supported for hscroll")
}

