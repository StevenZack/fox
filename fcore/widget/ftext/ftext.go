package ftext

import "github.com/StevenZack/fox/fcore"

type FText struct {
	fcore.FBaseView
	fnOnClick func()
}

func New(a fcore.IActivity) *FText {
	f := &FText{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "Text"
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
func (f *FText) Assign(i **FText) *FText {
	(*i) = f
	return f
}
func (f *FText) OnClick(fn func()) *FText {
	f.fnOnClick = fn
	return f
}

func (f *FText) Text(s string) *FText {
	f.A.SetAttr(f.Vid, "Text", s, "")
	return f
}
func (f *FText) TextColor(fcolor string) *FText {
	f.A.SetAttr(f.Vid,"TextColor",fcolor,"")
	return f
}
func (f *FText) TextSize(dp int) *FText {
	f.A.SetAttr(f.Vid,"TextSize",fcore.SPrintf(dp),"")
	return  f
}