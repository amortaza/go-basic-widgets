package vscroll

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
		bl.Dim(thickness,400)
		bl.Color(.1,.3,.1)

		bl.Div()
		{
			bl.Id(handleId)
			bl.Pos(0,0)

			var handleHeight int32 = 80

			bl.Dim(thickness, handleHeight)

			bl.Color(.5,.5,.1)

			drag.On(func(v interface{}) {
				e := v.(drag.Event)
				handle := e.Target

				totalHeight := handle.Parent.Height

				handleTop := int32(math.Max(0, float64(handle.Top)))

				maxTop := totalHeight - handle.Height
				handleTop = int32(math.Min(float64(maxTop), float64(handleTop)))

				bl.EnsureShadowByNode(handle).PosTop(handleTop)
				bl.EnsureShadowByNode(handle).PosLeft(0)

				pctStart := float32(handle.Top) / float32(totalHeight)
				pctEnd := float32(handle.Top + handle.Height) / float32(totalHeight)

				if cb != nil {
					cb(newEvent(pctStart, pctEnd))
				}
			})

			bl.EnsureShadow().PosLeft_to_Node()
		}
		bl.End()
	}
	bl.End()
}

func newEvent(pctStart, pctEnd float32) *Event {
	return &Event{pctStart, pctEnd}
}