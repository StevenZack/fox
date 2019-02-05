package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fdialog"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/fimage"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	fframebox.New(a).DeferShow().Append(
		fbutton.New(a).Text("showDialog").OnClick(func() {
			fdialog.New(a).PositiveButton("ok", func(dialog *fdialog.FDialog) {
				dialog.Dismiss()
			}).NegativeButton("quit", func(dialog *fdialog.FDialog) {
				dialog.Dismiss()
			}).Append(
				fimage.New(a).Src("https://jywjl.github.io/icons/transparent.png"),
				).Title("title").Show()
		}).LayoutGravity(fgravity.Center))
}
