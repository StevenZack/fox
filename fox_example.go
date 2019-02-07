package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fcolor"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fradio"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	rg:=fradio.NewGroup(a)
	fbox.NewV(a).Size(-2,-2).DeferShow().Append(
		rg.Size(-2,-1).Background(fcolor.Yellow).Append(
			fradio.New(a).Text("one"),
			fradio.New(a).Text("two"),
			fradio.New(a).Text("three"),
		).OnChange(func() {
			fcore.ShowToast(a,rg.GetSelected().GetText())
		}),
			)
}
