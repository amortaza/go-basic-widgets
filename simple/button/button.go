package button

import (
	"bellina"
	"plugin/mouse-hover"
	"plugin/click"
	"fmt"
)

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "button"
}

func (c *Plugin) Init() {
	bl.Plugin( click.NewPlugin() )
	bl.Plugin( mouse_hover.NewPlugin() )

	g_stateByButtonId = make(map[string] *ButtonState)
}

func (c *Plugin) Tick() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On(cb func(interface{})) {
}

var State *ButtonState

func Div(id string) {

	State = getOrCreateState(id)

	bl.Div()
	{
		bl.ID(id)
		bl.Pos(10, 10)
		bl.Dim(96,48)

		bl.Color(.21, .21, 0)

		if State.IsHover {
			bl.Color(.3, .3, .0)

		}

		if State.IsDown {
			bl.Color(.6, .6, .1)
		}

		bl.BorderThickness([]int32{1,1,1,1})
		bl.BorderColor(.6,.6,.1)
		bl.BorderTopsCanvas()

		bl.Font("arial", 7)
		bl.FontColor(1,1,1)
		
		bl.Label("OK")

		bl.Flag(bl.Z_COLOR_SOLID | bl.Z_BORDER_ALL | bl.Z_LABEL_ALIGN_HCENTER | bl.Z_LABEL_ALIGN_VCENTER)
	}
}

func Label(label string) {
	bl.Label(label)
}

func OnHover(cb func(*mouse_hover.Event)) {
	bl.On("hover", func(i interface{}){
		e := i.(*mouse_hover.Event)

		cb(e)
	})
}

func End() {
	state := State

	bl.On("hover", func(i interface{}){
		e := i.(*mouse_hover.Event)

		if e.IsInEvent {
			state.IsHover = true
		} else {
			state.IsHover = false
		}
	})

	bl.On2("click",

		// click
		func(i interface{}) {
			//fmt.Println("click")
			state.IsDown = false
			fmt.Println("click")
		},

		// on down
		func(i interface{}) {
			state.IsDown = true
			fmt.Println("down")
		},

		// on miss
		func(i interface{}) {
			state.IsDown = false
			fmt.Println("miss")
		} )

	bl.End()
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supoorted in click.Plugin")
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}
