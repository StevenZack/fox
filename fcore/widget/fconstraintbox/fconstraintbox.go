package fconstraintbox

import (
	"strings"

	"github.com/StevenZack/fox/fcore"
)

type FConstraintBox struct {
	fcore.FBaseView
	showAfter bool
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
func (f *FConstraintBox) Assign(i **FConstraintBox) *FConstraintBox {
	(*i) = f
	return f
}
func (f *FConstraintBox) Append(is ...fcore.IBaseView) *FConstraintBox {
	var vids []string
	for _, i := range is {
		if i == nil {
			continue
		}
		vids = append(vids, i.GetBaseView().Vid)
	}
	f.A.SetAttr(f.Vid, "Append", strings.Join(vids, ","), "")
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FConstraintBox) DeferShow() *FConstraintBox {
	f.showAfter = true
	return f
}
