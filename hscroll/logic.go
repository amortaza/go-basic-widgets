package hscroll

import (
	"github.com/amortaza/go-bellina"
	"math"
	"github.com/amortaza/go-bellina-plugins/drag"
)

var gStateByNode map[string] *State

func logic(state *State, cb func(interface{})) {
	scrollId := state.HScrollId
	handleId := scrollId + ":handle"
	thickness := state.Thickness_

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

func ensureState(Id string) *State {
	state, ok := gStateByNode[Id]

	if !ok {
		state = &State{HScrollId: Id}
		gStateByNode[Id] = state
	}

	return state
}
