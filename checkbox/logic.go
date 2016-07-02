package checkbox

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/click"
)

var gStateByNode map[string] *State

func logic(state *State, cb func(interface{})) {
	cbid := state.CheckboxId
	ctrlid := cbid + "/ctrl"
	tickid := ctrlid + "/tick"

	bl.Div()
	{
		bl.Id(cbid)

		bl.Pos(10, 50)
		bl.Dim(220,39)
		bl.Color(0.2,0,0)
		bl.BorderThickness(bl.FourOnesInt)
		bl.BorderColor(1,1,1)
		bl.BorderTopsCanvas()

		bl.Font("arial", 8)
		bl.FontColor(1,1,1)
		bl.FontNudge(43,9)
		bl.Label( state.Label)

		bl.Div()
		{
			bl.Id(ctrlid)

			bl.Pos(7, 7)
			bl.Dim(26,26)

			bl.BorderThickness(bl.FourOnesInt)
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			bl.Color(0, 0.2, 0)

			drawTick(tickid, state.CheckState)

			click.On2(
				func(v interface{}) {
					state := EnsureState(cbid)

					if state.checkStateAtStart == Checked{
						state.CheckState = Unchecked

					} else if state.checkStateAtStart == Unchecked{
						state.CheckState = Checked
					}

					state.checkStateAtStart = state.CheckState
				},

				func(v interface{}) {
					state := EnsureState(cbid)

					state.checkStateAtStart = state.CheckState

					state.CheckState = Semi
				},

				func(v interface{}) {
					state := EnsureState(cbid)

					if state.CheckState == Semi {

						state.CheckState = state.checkStateAtStart
					}
				},
			)
		}
		bl.End()
	}
	bl.End()
}

func drawTick(id string, s CheckboxState) {

	if s == Unchecked {
		return
	}

	bl.Div()
	{
		bl.Id(id)

		bl.Pos(5, 5)
		bl.Dim(16, 16)

		bl.Color(1, 1, 0)

		if s == Semi {
			bl.NodeOpacity1f(.2)
		}

		bl.BorderThickness(bl.FourOnesInt)
		bl.BorderColor(1, 1, 1)
		bl.BorderTopsCanvas()
		bl.InvisibleToEvents()
	}
	bl.End()

}

