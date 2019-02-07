package fclipboard

import "github.com/StevenZack/fox/fcore"

type FClipboard struct {
	fcore.FBase
}

func New(a fcore.IActivity) *FClipboard {
	f := &FClipboard{}
	f.A = a
	f.Vid=fcore.NewToken()
	f.VType = "Clipboard"
	f.A.NewObject(f.VType, f.Vid)
	return f
}
func (f *FClipboard) SetText(s string) *FClipboard {
	f.A.SetAttr(f.Vid,"ClipData",s,"")
	return  f
}
func (f *FClipboard) GetText() string {
	return f.A.GetAttr(f.Vid,"ClipData")
}
func (f *FClipboard) OnChange(fn func())*FClipboard  {
	fcore.EventMap.Set(f.Vid+":onchange", func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		fn()
		return ""
	})
	f.A.SetAttr(f.Vid,"OnChange",f.Vid+":onchange","")
	return f
}