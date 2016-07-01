package hsplit

import "github.com/amortaza/go-bellina"

func Use() {
	plugin.On(nil)
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
