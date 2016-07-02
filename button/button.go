package button

import (
	"github.com/amortaza/go-bellina-plugins/hover"
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina"
)

type State struct {
	ButtonId string
	IsHover  bool
	IsDown   bool

	Label_   string

	onClick  func()
}

// Shared variable across Div()/End()
var gCurState *State

func Id(postfixButtonId string) *State {
	buttonId := bl.Current_Node.Id + "/" + postfixButtonId

	gCurState = ensureState(buttonId)

	return gCurState
}

func (s *State) On(cb func(interface{})) {

	gCurState = s

	Div()
	End()
}

func Div() {

	buttonId := gCurState.ButtonId
	state := gCurState

	bl.Div()
	{
		bl.Id(buttonId)
		bl.Pos(10, 10)
		bl.Dim(96,48)

		bl.Color(.21, .21, 0)

		if state.IsHover {
			bl.Color(.3, .3, .0)
		}

		if state.IsDown {
			bl.Color(.6, .6, .1)
		}

		bl.BorderThickness([]int{1,1,1,1})
		bl.BorderColor(.6,.6,.1)
		bl.BorderTopsCanvas()

		bl.Font("arial", 7)
		bl.FontColor(1,1,1)
		
		bl.Flag(bl.Z_COLOR_SOLID | bl.Z_BORDER_ALL | bl.Z_LABEL_ALIGN_HCENTER | bl.Z_LABEL_ALIGN_VCENTER)
	}
}

func (s *State) Label(label string) (*State){
	s.Label_ = label

	return s
}

func OnHover(cb func()) {
	hover.On( func(i interface{}){
		cb()
	})
}

func OnClick(cb func()) {
	gCurState.onClick = cb
}

func End() {
	state := gCurState

	bl.Label(state.Label_)

	hover.On(func(i interface{}){
		e := i.(*hover.Event)

		if e.IsInEvent {
			state.IsHover = true
		} else {
			state.IsHover = false
		}
	})

	click.On2(

		// click
		func(i interface{}) {
			state.IsDown = false

			if state.onClick != nil {
				state.onClick()
			}
		},

		// on down
		func(i interface{}) {
			state.IsDown = true
			//fmt.Println("down")
		},

		// on miss
		func(i interface{}) {
			state.IsDown = false
		} )

	bl.End()
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}

