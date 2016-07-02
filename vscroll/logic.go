package vscroll

import (
	"github.com/amortaza/go-bellina"
	"math"
	"github.com/amortaza/go-bellina-plugins/drag"
)

var gStateByNode map[string] *State

func logic(state *State, cb func(interface{})) {
	scrollId := state.VScrollId
	handleId := scrollId + ":handle"
	thickness := state.Thickness_

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

			handleHeight := 80

			bl.Dim(thickness, handleHeight)

			bl.Color(.5,.5,.1)

			drag.On(func(v interface{}) {
				e := v.(drag.Event)
				handle := e.Target

				totalHeight := handle.Parent.Height

				handleTop := int(math.Max(0, float64(handle.Top)))

				maxTop := totalHeight - handle.Height
				handleTop = int(math.Min(float64(maxTop), float64(handleTop)))

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

func ensureState(Id string) *State {
	state, ok := gStateByNode[Id]

	if !ok {
		state = &State{VScrollId: Id}
		gStateByNode[Id] = state
	}

	return state
}
