package blanket

import (
	"github.com/amortaza/go-bellina"
)

var gId string

func logic() {
	w := bl.EnsureShadowByNode(bl.Current_Node).Width
	h := bl.EnsureShadowByNode(bl.Current_Node).Height

	bl.Div()
	{
		bl.Id(gId)

		bl.Pos(10,10)
		bl.Dim(w-20, h-20)
		bl.Color(.1,.1,.1)
		bl.NodeOpacity1f(.5)

		bl.PreventBubbling()
	}
	bl.End()
}
