package fimage

import (
	"github.com/StevenZack/fox/fcore"
)

type FImage struct {
	fcore.FBaseView
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
		if f.FnOnClick != nil {
			f.FnOnClick()
		}
		return ""
	})
	f.A.SetAttr(f.Vid, "OnClick", fnId, "")
	return f
}