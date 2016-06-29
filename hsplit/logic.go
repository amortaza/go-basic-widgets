package hsplit

import "github.com/amortaza/go-bellina"

func lineup(parent, kid1, kid2 *bl.Node, dividerWidth int32) (dividerLeft int32) {
	
	kid1Shadow, ok := bl.GetShadowById(kid1.Id)

	if !ok {
		kid1Shadow = bl.EnsureShadowByID(kid1.Id)

		kid1Shadow.Left = 0
		kid1Shadow.Top = 0
		kid1Shadow.Width = 160
		kid1Shadow.Height = parent.Height
	}

	kid2Shadow, ok2 := bl.GetShadowById(kid2.Id)

	if !ok2 {
		kid2Shadow = bl.EnsureShadowByID(kid2.Id)

		var kid2Width int32 = parent.Width - kid1Shadow.Width - dividerWidth

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
