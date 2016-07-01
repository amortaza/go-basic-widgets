package hsplit

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/drag"
	"math"
)

var NAME = "hsplit"

var plugin *Plugin

type Plugin struct {
}

func (c *Plugin) Name() string {
	return NAME
}

func (c *Plugin) GetState() interface{} {
	return nil
}

func (c *Plugin) Init() {
}

func (c *Plugin) Tick() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Uninit() {
}

func get2Kids(node *bl.Node)(kid1, kid2 *bl.Node) {
	e := node.Kids.Front()

	kid1 = e.Value.(*bl.Node)
	kid2 = e.Next().Value.(*bl.Node)

	return kid1, kid2
}

func (c *Plugin) On(cb func(interface{})) {
	cur := bl.Current_Node
	parentId := cur.Id

	kid1, kid2 := get2Kids(cur)
	kid1Id := kid1.Id
	kid2Id := kid2.Id

	var dividerWidth int32 = 16
	dividerId := "wtf" + parentId

	dividerLeft := lineup(cur, kid1, kid2, dividerWidth)

	// divider
	bl.Div()
	{
		bl.Id(dividerId)
		bl.Pos(dividerLeft,0)
		bl.Dim(dividerWidth,cur.Width)
		bl.Color(.61,.61,0)

		drag.On(func(v interface{}) {
			parent := bl.GetNodeByID(parentId)

			e := v.(drag.Event)

			divider := e.Target
			dividerShadow := bl.EnsureShadowByID(dividerId)

			var maxAllowed int32 = parent.Width - dividerWidth - 16

			dividerShadow.Left = int32(math.Max(float64(dividerShadow.Left), 16))
			dividerShadow.Left = int32(math.Min(float64(dividerShadow.Left), float64(maxAllowed)))

			divider.Left = dividerShadow.Left
			dividerShadow.Top = 0
			divider.Top = 0

			kid1 := bl.GetNodeByID(kid1Id)
			kid2 := bl.GetNodeByID(kid2Id)

			kid1Shadow := bl.EnsureShadowByID(kid1Id)
			kid2Shadow := bl.EnsureShadowByID(kid2Id)



			X := divider.Left

			kid1Shadow.Width = X
			kid1.Width = kid1Shadow.Width

			var kid2Width int32 = parent.Width - kid1Shadow.Width - dividerWidth

			kid2Shadow.Left = dividerShadow.Left + dividerWidth
			kid2.Left = kid2Shadow.Left

			kid2Shadow.Width = kid2Width
			kid2.Width = kid2Width
		})

		bl.EnsureShadow()
	}
	bl.End()
}

func (c *Plugin) On2(cb func(interface{}), onDown func(interface{}), onUpAndMiss func(interface{})) {
	panic("On2 not supported for hsplit")
}

func NewPlugin() *Plugin {
	plugin = &Plugin{}

	return plugin
}
