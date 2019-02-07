package fprogress

import "github.com/StevenZack/fox/fcore"

type FProgress struct {
	fcore.FBaseView
}

func New(a fcore.IActivity) *FProgress {
	f := &FProgress{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "Progress"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	fnId := fcore.NewToken()
	fcore.EventMap.Set(fnId+":onclick", func(fcore.IActivity, string, string, string) string {
		if f.FnOnClick != nil {
			f.FnOnClick()
		}
		return ""
	})
	f.A.SetAttr(f.Vid, "OnClick", fnId+":onclick", "")
	return f
}