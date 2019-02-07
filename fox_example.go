package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fconstraintbox"
	"github.com/StevenZack/fox/fcore/widget/fedit"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	var e1 *fcore.FBaseView
	var b1,b2 *fcore.FBaseView
	var c *fcore.FBaseView
	fconstraintbox.New(a).DeferShow().Assign(&c).Size(-2,-2).Append(
		fbutton.New(a).Assign(&b1).Size(-1,-1).Text("one"),
		fbutton.New(a).Assign(&b2).Left2RightOf(b1).Right2RightOfParent().Top2BottomOf(b1).OnClick(func() {
			c.AddViewAt(fbutton.New(a).Text("asd").Top2BottomOf(e1),0)
		}),
		fedit.New(a).Size(-2,-1).Top2BottomOf(b2).Assign(&e1))
}
