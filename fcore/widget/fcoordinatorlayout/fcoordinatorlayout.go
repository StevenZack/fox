package fcoordinatorlayout

import "github.com/StevenZack/fox/fcore"

type FCoordinatorLayout struct {
	fcore.FBaseView
}

func New(a fcore.IActivity) *FCoordinatorLayout {
	f := &FCoordinatorLayout{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "CoordinatorLayout"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}
