package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/fservice"
	"github.com/StevenZack/fox/fcore/widget/ftext"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	bt1,bt2,bt3:=fbutton.New(a),fbutton.New(a),fbutton.New(a)
	s:=fservice.New(a).OnCreate(func() {
		fcore.ShowToast(a,"service oncreate")
	})
	s.OnQuit(func() {
		fcore.ShowToast(a,"service quit")
		s.FinishAllActivity()
	}).QuitButton("quit").Title("title").SubTitle("sub")
	fbox.NewV(a).Size(-2,-2).DeferShow().Append(
		bt1.Size(-2,-1).Text("start").OnClick(func() {
			s.Show()
		}),
		bt2.Size(-2,-1).Text("stop").OnClick(func() {
			s.Stop()
		}),
		bt3.Size(-2,-1).Text("kill").OnClick(func() {
			s.KillAll()
		}),
		fbutton.New(a).Text("new activity").Size(-2,-1).OnClick(func() {
			fcore.StartActivity(a, func(activity fcore.IActivity) {
				fframebox.New(a).Size(-2,-2).DeferShow().Append(
					ftext.New(a).Text("second").LayoutGravity(fgravity.Center))
			},nil)
		}))
}
