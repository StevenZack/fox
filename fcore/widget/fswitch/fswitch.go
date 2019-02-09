package fswitch

import "github.com/StevenZack/fox/fcore"

type FSwitch struct {
	fcore.FBaseView
}


func New(a fcore.IActivity) *FSwitch {
	f := &FSwitch{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "Switch"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	fnId := f.Vid+":onclick"
	fcore.EventMap.Set(fnId, func(fcore.IActivity, string, string, string) string {
		if f.FnOnClick != nil {
			f.FnOnClick()
		}
		return ""
	})
	f.A.SetAttr(f.Vid, "OnClick", fnId, "")
	return f
}