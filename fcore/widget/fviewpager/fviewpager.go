package fviewpager

import "github.com/StevenZack/fox/fcore"

type FViewPager struct {
	fcore.FBaseView
}

func New(a fcore.IActivity) *FViewPager {
	f:=&FViewPager{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "ViewPager"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}
func NewPage(createView func() fcore.IBaseView) *fcore.FPage {
	p:=&fcore.FPage{}
	p.VID=fcore.NewToken()
	fcore.EventMap.Set(p.VID, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		return createView().GetBaseView().Vid
	})
	return p
}
