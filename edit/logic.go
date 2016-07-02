package edit

import (
	"github.com/amortaza/go-xel"
	"math"
	"github.com/amortaza/go-bellina/core"
	"github.com/amortaza/go-bellina"
)

var fontNudgeLeft = 4
var fontNudgeTop = 5

var paddingLeft int
var paddingTop int
var paddingBottom int

var gStateByWidgetId map[string] *State

func insertChar(src string, pos int, char string) string {

	if pos >= len(src) {
		return src + char
	}

	p1 := src[0:pos]
	p2 := src[pos:]

	return p1 + char + p2
}

func backspace(src string, pos int) string {

	if pos < 1 {
		return src
	}

	p1 := src[0:pos-1]
	p2 := src[pos:]

	return p1 + p2
}

func doDelete(src string, pos int) string {

	if pos >= len(src) {
		return src
	}

	p1 := src[0:pos]
	p2 := src[pos+1:]

	return p1 + p2
}

func processKeyDown(key xel.KeyboardKey, alt, ctrl, shift bool, shadow *bl.ShadowNode, state *State) {
	bt := xel.KeyToBehavior(key, false, true)

	if bt == xel.Key_Behavior_CHAR {
		shadow.Label = insertChar(shadow.Label, state.CursorPos, xel.KeyToChar(key, shift, true))
		state.CursorPos++

	} else {
		if key == xel.Key_HOME {
			state.CursorPos = 0
		}
		if key == xel.Key_END {
			state.CursorPos = len(shadow.Label)
		}
		if key == xel.Key_DELETE {
			shadow.Label = doDelete(shadow.Label, state.CursorPos)
		}
		if key == xel.Key_BACKSPACE {
			shadow.Label = backspace(shadow.Label, state.CursorPos)
			state.CursorPos = int(math.Max(0, float64(state.CursorPos-1)))
		}
		if key == xel.Key_LEFT {
			state.CursorPos = int(math.Max(0, float64(state.CursorPos-1)))
		}
		if key == xel.Key_RIGHT {
			state.CursorPos = int(math.Min(float64(state.CursorPos+1), float64(len(shadow.Label))))
		}
	}
}

func getCursorX(cursorPos int, text string, fontname string, fontsize int) int {
	if cursorPos > len(text) {
		cursorPos = len(text)
	}

	substr := text[:cursorPos]

	g4font := core.GetG4Font(fontname, fontsize)

	return g4font.Width(substr)
}

func ensureState(editId string) *State {
	state, ok := gStateByWidgetId[editId]

	if !ok {
		state = &State{EditId:editId, Width: 200}

		gStateByWidgetId[editId] = state
	}

	return state
}

