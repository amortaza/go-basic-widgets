package edit

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-g4"
)

func renderCursor(editId string, cursorHeight, cursorX, cursorY int32 ) {
	var left int32 = cursorX
	var top int32 = cursorY
	var width int32 = 1
	var height int32 = cursorHeight

	rgba := []float32{0, 0, 0, 1}

	bl.CustomRenderer1(func() {
		coords := []int32{left,top,width,height}
		color := rgba

		g4.DrawColorRect(coords[0], coords[1], coords[2], coords[3], color, color, color, color)
	}, false)
}
