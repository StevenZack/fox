package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/widget/fcoordinatorlayout"
	"github.com/StevenZack/fox/fcore/widget/ffab"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/fsnackbar"
	"github.com/StevenZack/fox/fcore/widget/ftext"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	var fc *fcore.FBaseView
	fc=fcoordinatorlayout.New(a).DeferShow().Size(-2,-2).Append(
		ffab.New(a).Icon("drawable://add").LayoutGravity(fgravity.Bottom|fgravity.Right).OnClick(func() {
			fsnackbar.New(a,fc).Text("text").Show()
		}),
		)
}

func secondPage(a fcore.IActivity) {
	fframebox.New(a).DeferShow().Append(
		ftext.New(a).Text("second").LayoutGravity(fgravity.Center))
}
