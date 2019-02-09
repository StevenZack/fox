package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/widget/fbottomnav"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/fmenu"
	"github.com/StevenZack/fox/fcore/widget/ftext"
	"github.com/StevenZack/fox/fcore/widget/ftoolbar"
	"github.com/StevenZack/fox/fcore/widget/fviewpager"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	vp:=fviewpager.New(a)
	bn:=fbottomnav.New(a)
	fbox.NewV(a).DeferShow().Size(-2,-2).Append(
		ftoolbar.New(a).Title("title").SubTitle("sub").NavigationIcon("drawable://add").OnNavigationIconClick(func() {
			fcore.ShowToast(a,"clicked")
		}).Menus(
			fmenu.NewItem("one"),
			fmenu.NewItem("search").Icon("drawable://add").ShowAsAction().OnClick(func() {
				fcore.ShowToast(a,"search")
			}),
			fmenu.NewSub("two",fmenu.NewItem("sub").OnClick(func() {
				fcore.ShowToast(a,"sub")
			})),
			),
		vp.LayoutWeight(1).Size(-2,-2).OnGetPage(
			func(pos int) fcore.IBaseView {
				return ftext.New(a).Text(fcore.SPrintf(pos))
			},
			func() int {
				return 2
			},
			).OnPageSelected(func(i int) {
			bn.SelectedIndex(i)
		}),
		bn.Menus(
			fmenu.NewItem("one").OnClick(func() {
				vp.SelectedIndex(0)
			}),
			fmenu.NewItem("two").OnClick(func() {
				vp.SelectedIndex(1)
			}),
			),
			)
}

func secondPage(a fcore.IActivity) {
	fframebox.New(a).DeferShow().Append(
		ftext.New(a).Text("second").LayoutGravity(fgravity.Center))
}
