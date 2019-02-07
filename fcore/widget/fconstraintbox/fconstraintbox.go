package fconstraintbox

import (
	"github.com/StevenZack/fox/fcore"
)

type FConstraintBox struct {
	fcore.FBaseView
}

func New(a fcore.IActivity) *FConstraintBox {
	f := &FConstraintBox{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "ConstraintBox"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}