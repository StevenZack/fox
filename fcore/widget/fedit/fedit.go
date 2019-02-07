package fedit

import "github.com/StevenZack/fox/fcore"

type FEdit struct {
	fcore.FBaseView
}

func New(a fcore.IActivity) *FEdit {
	f := &FEdit{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "Edit"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}
func (f *FEdit) Assign(i **FEdit) *FEdit {
	(*i)=f
	return f
}