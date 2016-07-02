package edit

import (
	"math"
	"github.com/amortaza/go-bellina"
)

var cursorCycle float64
var cursorOpacity float32
var plugin *Plugin

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "edit"
}

func (c *Plugin) Init() {
	gStateByWidgetId = make(map[string] *State)
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) Reset() {
}

func (c *Plugin) Tick() {

	// There is only one cursor with blinking cursor, so this can be global "Tick"

	cursorOpacity = float32((math.Sin(cursorCycle) + 1 ) / 2 + .5)
	cursorCycle += .4
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

