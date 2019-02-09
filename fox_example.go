package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/ftext"
	"github.com/StevenZack/fox/fcore/widget/fvalueAnimator"
	"strconv"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	bt:=fbutton.New(a)
	fframebox.New(a).DeferShow().Size(-2,-2).Append(
		bt.Text("text").OnClick(func() {
			fvalueAnimator.New(a).OfInt(0,100).Duration(3000).OnValueChanged(func(s string) {
				i,e:=strconv.ParseFloat(s,64)
				if e != nil {
					return
				}
				bt.X(i)
			}).Start()
		}))
}

func secondPage(a fcore.IActivity) {
	fframebox.New(a).DeferShow().Append(
		ftext.New(a).Text("second").LayoutGravity(fgravity.Center))
}
