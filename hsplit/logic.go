package hsplit

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/drag"
	"math"
)

func logic(state *State, cb func(interface{})) {
	cur := bl.Current_Node
	parentId := cur.Id

	kid1, kid2 := get2Kids(cur)
	kid1Id := kid1.Id
	kid2Id := kid2.Id

	dividerWidth := 16
	dividerId := parentId + "/vdivider"

	dividerLeft := lineup(cur, kid1, kid2, dividerWidth)

	// divider
	bl.Div()
	{
		bl.Id(dividerId)
		bl.Pos(dividerLeft,0)
		bl.Dim(dividerWidth,cur.Width)
		bl.Color(.61,.61,0)

		drag.On(func(v interface{}) {
			parent := bl.GetNodeById(parentId)

			e := v.(drag.Event)

			divider := e.Target
			dividerShadow := bl.EnsureShadowById(dividerId)

			maxAllowed := parent.Width - dividerWidth - 16

			dividerShadow.Left = int(math.Max(float64(dividerShadow.Left), 16))
			dividerShadow.Left = int(math.Min(float64(dividerShadow.Left), float64(maxAllowed)))

			divider.Left = dividerShadow.Left
			dividerShadow.Top = 0
			divider.Top = 0

			kid1 := bl.GetNodeById(kid1Id)
			kid2 := bl.GetNodeById(kid2Id)

			kid1Shadow := bl.EnsureShadowById(kid1Id)
			kid2Shadow := bl.EnsureShadowById(kid2Id)

			X := divider.Left

			kid1Shadow.Width = X
			kid1.Width = kid1Shadow.Width

			kid2Width := parent.Width - kid1Shadow.Width - dividerWidth

			kid2Shadow.Left = dividerShadow.Left + dividerWidth
			kid2.Left = kid2Shadow.Left

			kid2Shadow.Width = kid2Width
			kid2.Width = kid2Width
		})

		bl.EnsureShadow()
	}
	bl.End()
}

func lineup(parent, kid1, kid2 *bl.Node, dividerWidth int) (dividerLeft int) {
	
	kid1Shadow, ok := bl.TestShadowById(kid1.Id)

	if !ok {
		kid1Shadow = bl.EnsureShadowById(kid1.Id)

		kid1Shadow.Left = 0
		kid1Shadow.Top = 0
		kid1Shadow.Width = 160
		kid1Shadow.Height = parent.Height
	}

	kid2Shadow, ok2 := bl.TestShadowById(kid2.Id)

	if !ok2 {
		kid2Shadow = bl.EnsureShadowById(kid2.Id)

		kid2Width := parent.Width - kid1Shadow.Width - dividerWidth

		kid2Shadow.Left = parent.Width - kid2Width
		kid2Shadow.Top = 0
		kid2Shadow.Width = kid2Width
		kid2Shadow.Height = parent.Height
	}

	kid1.Left = kid1Shadow.Left
	kid1.Top = kid1Shadow.Top
	kid1.Width = kid1Shadow.Width
	kid1.Height = kid1Shadow.Height
	
	kid2.Left = kid2Shadow.Left
	kid2.Top = kid2Shadow.Top
	kid2.Width = kid2Shadow.Width
	kid2.Height = kid2Shadow.Height

	dividerLeft = kid1.Width

	return dividerLeft
}

func ensureState(hsplidId string) *State {
	state, ok := gStateByWidgetId[hsplidId]

	if !ok {
		state = &State{HSplitId: hsplidId}
		gStateByWidgetId[hsplidId] = state
	}

	return state
}

func get2Kids(node *bl.Node)(kid1, kid2 *bl.Node) {
	e := node.Kids.Front()

	kid1 = e.Value.(*bl.Node)
	kid2 = e.Next().Value.(*bl.Node)

	return kid1, kid2
}



