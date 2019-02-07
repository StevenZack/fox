package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fpermission"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fedit"
	"github.com/StevenZack/fox/fcore/widget/fmenu"
	"github.com/StevenZack/fox/fcore/widget/fpopupMenu"
	"github.com/StevenZack/fox/fcore/widget/fprogress"
	"github.com/StevenZack/fox/fcore/widget/ftext"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	var t1 *ftext.FText
	var e *fedit.FEdit
	fbox.New(a).DeferShow().Append(
		ftext.New(a).Assign(&t1).Text("asd").OnClick(func() {
			fpopupMenu.NewFrom(t1).DeferShow().Menus(
				fmenu.NewItem("item").ShowAsAction().OnClick(func() {
					fcore.ShowToast(a,"item clicked")
				}),
				fmenu.NewSub("sub",
					fmenu.NewItem("one"),
					fmenu.NewItem("two"),
					),
				)
		}),
		fedit.New(a).Assign(&e).Size(-2,-1).Text("a").Hint("input").OnChange(func() {
			t1.Text(e.GetText())
		}),
		fprogress.New(a),
		fbutton.New(a).Text("text").OnClick(func() {
			fcore.ShowToast(a,"clicked")
			if !fcore.CheckSelfPermission(a,fpermission.WRITE_EXTERNAL_STORAGE){
				fcore.RequestPermissions(a,[]string{fpermission.CAMERA,fpermission.WRITE_EXTERNAL_STORAGE}, func(bools []bool) {

				})
			}
		}),
	)
}
