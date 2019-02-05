package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fscaletype"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fconstraintbox"
	"github.com/StevenZack/fox/fcore/widget/fimage"
	"github.com/StevenZack/fox/fcore/widget/flist"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	s := []string{
		`https://ss3.baidu.com/9fo3dSag_xI4khGko9WTAnF6hhy/image/h%3D300/sign=08d982d4b3096b639e1958503c328733/3801213fb80e7bec5678461a222eb9389a506bae.jpg`,
		`https://ss3.baidu.com/9fo3dSag_xI4khGko9WTAnF6hhy/image/h%3D300/sign=01fb5a25fa36afc3110c39658318eb85/908fa0ec08fa513d24237b3b306d55fbb3fbd94d.jpg`,
		`https://ss3.baidu.com/-fo3dSag_xI4khGko9WTAnF6hhy/image/h%3D300/sign=ea5d1597e6f81a4c3932eac9e72b6029/2e2eb9389b504fc2330c17e6e8dde71191ef6d86.jpg`,
		`https://ss1.baidu.com/-4o3dSag_xI4khGko9WTAnF6hhy/image/h%3D300/sign=6215326447ed2e73e3e9802cb700a16d/6a63f6246b600c33d32b0328174c510fd8f9a184.jpg`,
		`https://ss1.baidu.com/9vo3dSag_xI4khGko9WTAnF6hhy/image/h%3D300/sign=a8d02e1ccecec3fd943ea175e68ad4b6/1f178a82b9014a90e7c1956da4773912b21bee67.jpg`,
		`https://ss0.baidu.com/7Po3dSag_xI4khGko9WTAnF6hhy/image/h%3D300/sign=7392271076cb0a469a228d395b62f63e/7dd98d1001e939015e5e48d476ec54e737d196f8.jpg`,
		`https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1549171978958&di=8cbf62e48c374d8d4f6edcb72a72ec22&imgtype=0&src=http%3A%2F%2Fcimage.tianjimedia.com%2Fimagelist%2F2009%2F287%2F77a67c44t76e.jpg`,
	}
	fconstraintbox.New(a).DeferShow().Append(
		flist.New(a).Funcs(
			func(l *flist.FList) fcore.IBaseView {
				return fbox.NewH(a).Append(
					flist.SetItemId(fimage.New(a).ScaleType(fscaletype.CenterInside).Size(30, 30), l, "icon"),
					flist.SetItemId(fbutton.New(a), l, "text"),
				)
			},
			func(vh *flist.ViewHolder, pos int) {
				vh.GetImageByItemId("icon").Src(s[pos])
				vh.GetButtonByItemId("text").Text(s[pos])
			},
			func() int {
				return len(s)
			},
		).HeightPercent(0.3).WidthPercent(1),
	).Size(-2, -2)
}
