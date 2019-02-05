package fbutton

import (
	"github.com/StevenZack/fox/fcore"
)

type FButton struct {
	fcore.FBaseView
	fnOnClick func()
}

func New(a fcore.IActivity) *FButton {
	f := &FButton{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "Button"
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
func (f *FButton) Assign(i **FButton) *FButton {
	(*i) = f
	return f
}
func (f *FButton) OnClick(fn func()) *FButton {
	f.fnOnClick = fn
	return f
}

func (f *FButton) Text(s string) *FButton {
	f.A.SetAttr(f.Vid, "Text", s, "")
	return f
}
