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

