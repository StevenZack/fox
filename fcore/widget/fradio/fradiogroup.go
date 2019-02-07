package fradio

import "github.com/StevenZack/fox/fcore"

type FRadioGroup struct {
	fcore.FBaseView
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
func (f *FRadioGroup) GetSelected() *FRadio {
	vid:=f.A.GetAttr(f.Vid, "Selected")
	if fcore.ViewMap.Exists(vid) {
		if v, ok := fcore.ViewMap.Get(vid).(*FRadio); ok {
			return v
		}
	}
	return nil
}