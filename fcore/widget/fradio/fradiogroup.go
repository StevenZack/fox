package fradio

import "github.com/StevenZack/fox/fcore"

type FRadioGroup struct {
	fcore.FBaseView
	Children []*FRadio
}

func NewGroup(a fcore.IActivity) *FRadioGroup {
	f := &FRadioGroup{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "RadioGroup"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}
func (f *FRadioGroup) Assign(i **FRadioGroup) *FRadioGroup {
	(*i) = f
	return f
}