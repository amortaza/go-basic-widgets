package edit

import "github.com/amortaza/go-bellina"

func renderCursor(editId string, cursorHeight, cursorX, cursorY int32 ) {
	bl.Div()
	{
		bl.ID(editId + ":cursor")
		bl.Dim(1, cursorHeight)
		bl.Pos(cursorX,cursorY)
		bl.Color(0,0,0)
		bl.NodeOpacity1f(1)
	}
	bl.End()
}
