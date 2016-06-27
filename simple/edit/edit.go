package edit

import (
	"github.com/amortaza/go-bellina-plugins/focus"
	"github.com/amortaza/go-xel"
	"math"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina/core"
)

var cursorCycle float64
var cursorOpacity float32

func (c *Plugin) Tick() {

	// There is only one cursor with blinking cursor, so this can be global "Tick"

	cursorOpacity = float32((math.Sin(cursorCycle) + 1 ) / 2 + .5)
	cursorCycle += .4
}

var fontNudgeLeft int32 = 4
var fontNudgeTop int32 = 5

var editWidth int32
var paddingLeft int32
var paddingTop int32
var paddingBottom int32

func Div(editId string) {
	editWidth = 200  // default
	paddingLeft = 0  // default
	paddingTop = 0  // default
	paddingBottom = 0  // default

	editInfo := ensureEditInfo(editId)

	bl.Div()
	{
		bl.ID(editId)

		shadow := bl.EnsureShadow()

		bl.Color(1,1,1)
		bl.Flag(bl.Z_COLOR_SOLID | bl.Z_BORDER_ALL)
		bl.BorderColor(0,1,1)
		bl.BorderThickness(bl.FourTwosInt)
		bl.Font("tahoma", 5)
		bl.FontColor(0,0,0)
		
		bl.On2("focus", func(focusEvent interface{}) {
			e := focusEvent.(focus.Event)

			if e.KeyEvent.Action == xel.Button_Action_Down {
				key := e.KeyEvent.Key
				processKeyDown(key, e.KeyEvent.Alt, e.KeyEvent.Ctrl, e.KeyEvent.Shift, shadow, editInfo)
			}

		}, func(focusEvent interface{}) {
			editInfo.hasFocus = true
			editInfo.cursorPos = len(shadow.Label)

		}, func(focusEvent interface{}) {
			editInfo.hasFocus = false
		})
	}
}

func Extend(maxLimit int32) {

	shadow, _ := bl.GetShadow()
	editId := shadow.Id

	substr := GetText(editId)

	fontname, fontsize := bl.GetFont()

	g4font := core.GetG4Font(fontname, fontsize)

	test := int32(math.Max( float64(editWidth), float64(g4font.Width(substr + "  ") )))

	editWidth = int32(math.Min(float64(test), float64(maxLimit)))
}

func End() {
	fontheight :=  bl.GetFontHeight()

	cursorHeight := fontheight + 6 // <<<
	editHeight :=  cursorHeight + paddingTop + paddingBottom

	shadow, _ := bl.GetShadow()
	editId := shadow.Id
	editInfo := ensureEditInfo(editId)

	bl.Dim(editWidth, editHeight)

	bl.FontNudge(fontNudgeLeft + paddingLeft, fontNudgeTop + paddingTop)
	bl.Label(shadow.Label)
	parent := bl.Current_Node

	if editInfo.hasFocus {
		if math.Abs(float64(cursorOpacity)) > .9 {

			cursorY := (editHeight - cursorHeight)/2 - 3 // <<<

			renderCursor(
				editId,
				cursorHeight,
				paddingLeft + fontNudgeLeft + getCursorX(editInfo.cursorPos, shadow.Label, parent.FontName, parent.FontSize),
				cursorY,
			)
		}
	}

	bl.End()
}

