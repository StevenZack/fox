package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fcolor"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fimage"
	"github.com/StevenZack/fox/fcore/widget/ftext"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
	fbox.New(a).DeferShow().Append(
		ftext.New(a).Text("one").TextColor(fcolor.Teal).TextSize(30).Foreground(fcolor.RippleEffect).LayoutGravity(fgravity.CenterHorizontal),
		fimage.New(a).Src("https://ss0.baidu.com/7Po3dSag_xI4khGko9WTAnF6hhy/image/h%3D300/sign=a1ee7108f01f4134ff37037e151d95c1/c995d143ad4bd11374005bd957afa40f4afb050f.jpg").Foreground(fcolor.RippleEffect).BackgroundColor(fcolor.White).Elevation(10))
}
