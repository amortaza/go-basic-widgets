package edit

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-g4"
)

func renderCursor(editId string, cursorHeight, cursorX, cursorY int ) {
	left := cursorX
	top := cursorY
	width := 1
	height := cursorHeight

	rgba := []float32{0, 0, 0, 1}

	bl.CustomRenderer1(func() {
		coords := []int{left,top,width,height}
		color := rgba

		g4.DrawColorRect(coords[0], coords[1], coords[2], coords[3], color, color, color, color)
	}, false)
}
