package edit

import "github.com/amortaza/go-bellina"

func renderCursor(editId string, cursorHeight, cursorX, cursorY int32 ) {
	bl.Rect1(cursorX, cursorY, 1, cursorHeight, 0, 0, 0, 1, true)
}
