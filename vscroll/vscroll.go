package vscroll

import (
	"github.com/amortaza/go-bellina"
)

func SetThickness(thickness int) {
	bl.SetI(NAME, "thickness", thickness)
}

func Use() {
	plugin.On(nil)
}

func On(cb func(interface{})) {
	plugin.On(cb)
}


func Div(scrollId_postfix string) {
	parentId := bl.Current_Node.Id
	scrollId := parentId + scrollId_postfix
	bl.DivId(scrollId)
}

func End() {
	bl.End()
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
