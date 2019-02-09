package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/values/fpermission"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/ftext"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	fbox.NewV(a).DeferShow().Size(-2,-2).Append(
		fbutton.New(a).Text("imei").OnClick(func() {
			if fcore.CheckSelfPermission(a, fpermission.READ_PHONE_STATE) {
				fcore.ShowToast(a,fcore.GetIMEI(a))
			}else{
				fcore.RequestPermissions(a,[]string{fpermission.READ_PHONE_STATE},nil)
			}
		}),
		fbutton.New(a).Text("id").OnClick(func() {
			fcore.ShowToast(a,fcore.GetUniqueID(a))
		}))
}

func secondPage(a fcore.IActivity) {
	fframebox.New(a).DeferShow().Append(
		ftext.New(a).Text("second").LayoutGravity(fgravity.Center))
}
