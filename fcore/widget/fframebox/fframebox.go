package fframebox

import (
	"github.com/StevenZack/fox/fcore"
)

type FFrameBox struct {
	fcore.FBaseView
	showAfter bool
}

func New(a fcore.IActivity) *FFrameBox {
	f := &FFrameBox{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "FrameBox"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}
func (f *FFrameBox) Assign(i **FFrameBox) *FFrameBox {
	(*i) = f
	return f
}
func (f *FFrameBox) Append(is ...fcore.IBaseView) *FFrameBox {
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
func (f *FFrameBox) DeferShow() *FFrameBox {
	f.showAfter = true
	return f
}
