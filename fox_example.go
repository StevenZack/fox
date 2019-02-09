package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/values/fpermission"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/ffilechooser"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/ftext"
	strings2 "strings"
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
		bt.Text("hold").OnClick(func() {
			if fcore.CheckSelfPermission(a, fpermission.WRITE_EXTERNAL_STORAGE) {
			ffilechooser.New(a).DeferShow().TypeSelectFiles(func(strings []string) {
				fcore.ShowToast(a,strings2.Join(strings,"\n"))
			})
			}else{
				fcore.RequestPermissions(a,[]string{fpermission.WRITE_EXTERNAL_STORAGE},nil)
			}
		}))
}

func secondPage(a fcore.IActivity) {
	fframebox.New(a).DeferShow().Append(
		ftext.New(a).Text("second").LayoutGravity(fgravity.Center))
}
