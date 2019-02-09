package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/fspace"
	"github.com/StevenZack/fox/fcore/widget/ftext"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	fbox.NewV(a).Size(-2,-2).DeferShow().Append(
		fspace.New(a),
		fbutton.New(a).Text("ok"))
}

func secondPage(a fcore.IActivity) {
	fframebox.New(a).DeferShow().Append(
		ftext.New(a).Text("second").LayoutGravity(fgravity.Center))
}
