package ftablayout

import "github.com/StevenZack/fox/fcore"

type FTabLayout struct {
	fcore.FBaseView
}

func NewLayout(a fcore.IActivity) *FTabLayout {
	f:=&FTabLayout{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "TabLayout"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}
