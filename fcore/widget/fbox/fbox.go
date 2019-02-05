package fbox

import (
	"github.com/StevenZack/fox/fcore"
)

type FBox struct {
	fcore.FBaseView
	showAfter bool
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
func (f *FBox) Horizontal() *FBox {
	f.A.SetAttr(f.Vid, "Orientation", "HORIZONTAL", "")
	return f
}
func (f *FBox) Assign(i **FBox) *FBox {
	(*i) = f
	return f
}
func (f *FBox) Append(is ...fcore.IBaseView) *FBox {
	for _, v := range is {
		if v == nil {
			continue
		}
		f.A.SetAttr(f.Vid, "AddView", v.GetBaseView().Vid, "")
	}
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FBox) DeferShow() *FBox {
	f.showAfter = true
	return f
}
