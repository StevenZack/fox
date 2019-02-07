package fbox

import (
	"github.com/StevenZack/fox/fcore"
)

type FBox struct {
	fcore.FBaseView
}


func New(a fcore.IActivity) *FBox {
	return NewV(a)
}
func NewV(a fcore.IActivity) *FBox {
	f := &FBox{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "Box"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}
func NewH(a fcore.IActivity) *FBox {
	f := NewV(a)
	f.Horizontal()
	return f
}