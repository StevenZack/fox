package fbottomnav

import "github.com/StevenZack/fox/fcore"

type FBottomNav struct {
	fcore.FBaseView
}


func New(a fcore.IActivity) *FBottomNav {
	f := &FBottomNav{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "BottomNav"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}