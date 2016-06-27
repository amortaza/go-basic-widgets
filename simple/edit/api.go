package edit

import (
	"github.com/amortaza/go-bellina"
)

func Pos(x,y int32) {
	bl.Pos(x,y)
}

func Size(size int32) {
	bl.Font("tahoma", size)
}

func Width(width int32) {
	editWidth = width
}

func Padding(left, top, bottom int32) {
	paddingLeft = left
	paddingTop = top
	paddingBottom = bottom
}

func GetText(editId string) string {
	shadow, _ := bl.GetShadowById(editId)

	return shadow.Label
}
