package hscroll

import (
	"github.com/amortaza/go-bellina"
	"math"
	"github.com/amortaza/go-bellina-plugins/drag"
)

func logic(scrollId_postfix string, cb func(interface{})) {
	parent := bl.Current_Node
	scrollId := parent.Id + scrollId_postfix
	handleId := scrollId + ":handle"
	thickness := bl.GetI(NAME, "thickness")

	bl.Div()
	{
		bl.Id(scrollId)
		bl.Pos(0,0)
		bl.Dim(400,thickness)
		bl.Color(.1,.3,.1)

		bl.Div()
		{
			bl.Id(handleId)
			bl.Pos(0,0)

			var handleWidth int = 80

			bl.Dim(handleWidth,thickness)

			bl.Color(.5,.5,.1)

			drag.On(func(v interface{}) {
				e := v.(drag.Event)
				handle := e.Target

				totalWidth := handle.Parent.Width

				handleLeft := int(math.Max(0, float64(handle.Left)))

				maxLeft := totalWidth - handle.Width
				handleLeft = int(math.Min(float64(maxLeft), float64(handleLeft)))

				bl.EnsureShadowByNode(handle).PosLeft(handleLeft)
				bl.EnsureShadowByNode(handle).PosTop(0)

				pctStart := float32(handle.Left) / float32(totalWidth)
				pctEnd := float32(handle.Left + handle.Width) / float32(totalWidth)

				if cb != nil {
					cb(newEvent(pctStart, pctEnd))
				}
			})

			bl.EnsureShadow().PosTop_to_Node()
		}
		bl.End()
	}
	bl.End()
}

func newEvent(pctStart, pctEnd float32) *Event {
	return &Event{pctStart, pctEnd}
}