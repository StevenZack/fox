package fframebox

import (
	"github.com/StevenZack/fox/fcore"
)

type FFrameBox struct {
	fcore.FBaseView
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