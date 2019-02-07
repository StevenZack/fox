package fclipboard

import "github.com/StevenZack/fox/fcore"

type FClipboard struct {
	fcore.FBase
	a fcore.IActivity
}

func New(a fcore.IActivity) *FClipboard {
	f := &FClipboard{}
	f.a = a
	f.Vid=fcore.NewToken()
	f.VType = "Clipboard"
	f.a.NewObject(f.VType, f.Vid)
	return f
}
func (f *FClipboard) SetText(s string) *FClipboard {
	f.a.SetAttr(f.Vid,"ClipData",s,"")
	return  f
}
func (f *FClipboard) GetText() string {
	return f.a.GetAttr(f.Vid,"ClipData")
}
func (f *FClipboard) OnChange(fn func())*FClipboard  {
	fcore.EventMap.Set(f.Vid+":onchange", func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		fn()
		return ""
	})
	f.a.SetAttr(f.Vid,"OnChange",f.Vid+":onchange","")
	return f
}