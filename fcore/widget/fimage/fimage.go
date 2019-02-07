package fimage

import (
	"github.com/StevenZack/fox/fcore"
)

type FImage struct {
	fcore.FBaseView
	fnOnClick func()
	src       string
}

func New(a fcore.IActivity) *FImage {
	f := &FImage{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "Image"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	fnId := fcore.NewToken()
	fcore.EventMap.Set(fnId, func(fcore.IActivity, string, string, string) string {
		if f.fnOnClick != nil {
			f.fnOnClick()
		}
		return ""
	})
	f.A.SetAttr(f.Vid, "OnClick", fnId, "")
	return f
}
func (f *FImage) Assign(i **FImage) *FImage {
	(*i) = f
	return f
}
func (f *FImage) OnClick(fn func()) *FImage {
	f.fnOnClick = fn
	return f
}
func (f *FImage) ScaleType(fscaletype string) *FImage {
	f.A.SetAttr(f.Vid, "ScaleType", fscaletype, "")
	return f
}
func (f *FImage) Src(s string) *FImage {
	if s == f.src {
		return f
	}
	if fcore.StartsWith(s, "http") {
		go fcore.CacheNetFile(s, "/data/data/"+fcore.GetPackageName(f.A)+"/cacheDir/", func(path string) {
			fcore.RunOnUIThreadWithFnId(f.A, func() {
				f.A.SetAttr(f.Vid, "Src", "file://"+path, "")
				f.src = s
			}, f.Vid+":src")
		})
		return f
	}
	f.A.SetAttr(f.Vid, "Src", s, "")
	f.src = s
	return f
}
