package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fedit"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/ftext"
	"github.com/StevenZack/fox/fcore/widget/fwebview"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	bt:=fbutton.New(a)
	et:=fedit.New(a)
	wv:=fwebview.New(a)
	fbox.New(a).DeferShow().Size(-2,-2).Append(
		fwebview.NewItem(a,"http://stevenzack.coding.me/asd/out.webp"),
		et.Size(-2,-1),
		bt.Text("load").OnClick(func() {
			wv.LoadUri(et.GetText())
		}),
		wv.Size(-2,-2),
		)
}

func secondPage(a fcore.IActivity) {
	fframebox.New(a).DeferShow().Append(
		ftext.New(a).Text("second").LayoutGravity(fgravity.Center))
}
