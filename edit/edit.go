package edit

import (
	"github.com/amortaza/go-bellina-plugins/focus"
	"github.com/amortaza/go-xel"
	"math"
	"github.com/amortaza/go-bellina"
)

var gCurState *State

func Id(postfixId string) *State {
	editId := bl.Current_Node.Id + "/" + postfixId

	gCurState = ensureState(editId)

	return gCurState
}

func Div() {
	paddingLeft = 0  // default
	paddingTop = 0  // default
	paddingBottom = 0  // default

	state := gCurState

	bl.Div()
	{
		bl.Id(state.EditId)

		shadow := bl.EnsureShadow()

		bl.Color(1,1,1)
		bl.Flag(bl.Z_COLOR_SOLID | bl.Z_BORDER_ALL)
		bl.BorderColor(0,1,1)
		bl.BorderThickness(bl.FourTwosInt)
		bl.Font("tahoma", 5)
		bl.FontColor(0,0,0)
		
		focus.On2( func(focusEvent interface{}) {
			e := focusEvent.(focus.Event)

			if e.KeyEvent.Action == xel.Button_Action_Down {
				key := e.KeyEvent.Key
				processKeyDown(key, e.KeyEvent.Alt, e.KeyEvent.Ctrl, e.KeyEvent.Shift, shadow, state)
			}

		}, func(focusEvent interface{}) {
			state.HasFocus = true
			state.CursorPos = len(shadow.Label)

		}, func(focusEvent interface{}) {
			state.HasFocus = false
		})
	}
}

func End() {
	fontheight :=  bl.GetFontHeight()

	cursorHeight := fontheight + 6 // <<<
	editHeight :=  cursorHeight + paddingTop + paddingBottom

	shadow := bl.EnsureShadow()
	editId := shadow.Id
	state := ensureState(editId)

	bl.Dim(state.Width, editHeight)

	bl.FontNudge(fontNudgeLeft + paddingLeft, fontNudgeTop + paddingTop)
	bl.Label(shadow.Label)
	parent := bl.Current_Node

	if state.HasFocus {
		if math.Abs(float64(cursorOpacity)) > .9 {

			cursorY := (editHeight - cursorHeight)/2 - 3 // <<<

			renderCursor(
				editId,
				cursorHeight,
				paddingLeft + fontNudgeLeft + getCursorX(state.CursorPos, shadow.Label, parent.FontName, parent.FontSize),
				cursorY,
			)
		}
	}

	bl.End()
}

/*
func Pos(x,y int) *State {
	bl.Pos(x,y)
}

func Size(size int) {
	bl.Font("tahoma", size)
}

func Width(width int) {
	editWidth = width
}

func Padding(left, top, bottom int) {
	paddingLeft = left
	paddingTop = top
	paddingBottom = bottom
}

func GetText(editId string) string {
	shadow := bl.EnsureShadowById(editId)

	return shadow.Label
}
*/

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
