package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/ftablayout"
	"github.com/StevenZack/fox/fcore/widget/ftext"
	"github.com/StevenZack/fox/fcore/widget/fviewpager"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	tl := ftablayout.NewLayout(a)
	vp := fviewpager.New(a).BindTabLayout(&tl.FBaseView)
	fbox.NewV(a).DeferShow().Size(-2, -2).Append(
		tl.Tabs(
			fcore.NewTab("one"),
			fcore.NewTab("two"),
			fcore.NewTab("three"),
		),
		vp.Size(-2, -2).OnGetPage(func(pos int) fcore.IBaseView {
			return fframebox.New(a).Size(-2, -2).Append(
				ftext.New(a).Text(fcore.SPrintf(pos)).LayoutGravity(fgravity.Center))
		}, func() int {
			return 3
		}),
	)
}
