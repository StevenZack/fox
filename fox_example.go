package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/fspace"
	"github.com/StevenZack/fox/fcore/widget/fspinner"
	"github.com/StevenZack/fox/fcore/widget/ftext"
	"time"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	orgData:=[]string{"one","two"}
	sp:=fspinner.New(a).List(orgData)
	fbox.NewV(a).Size(-2,-2).DeferShow().Append(
		sp.Size(-2,-1).OnItemClick(func(i int) {
			fcore.ShowToast(a,fcore.SPrintf(orgData[i]))
		}),
		fspace.New(a),
		fbutton.New(a).Text("ok").OnClick(func() {
			orgData=append(orgData,fcore.SPrintf(time.Now().Second()))
			sp.List(orgData)
			sp.NotifyDataSetChanged()
		}))
}

func secondPage(a fcore.IActivity) {
	fframebox.New(a).DeferShow().Append(
		ftext.New(a).Text("second").LayoutGravity(fgravity.Center))
}
