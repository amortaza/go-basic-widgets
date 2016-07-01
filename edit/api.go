package edit

import (
	"github.com/amortaza/go-bellina"
)

func Pos(x,y int) {
	bl.Pos(x,y)
}

func Size(size int) {
	bl.Font("tahoma", size)
}

func Width(width int) {
	editWidth = width
}

func Padding(left, top, bottom int) {
	paddingLeft = left
	paddingTop = top
	paddingBottom = bottom
}

func GetText(editId string) string {
	shadow, _ := bl.GetShadowById(editId)

	return shadow.Label
}
