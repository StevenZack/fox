package fspinner

import "github.com/StevenZack/fox/fcore"

type FSpinner struct {
	fcore.FBaseView
}

func New(a fcore.IActivity) *FSpinner {
	f:=&FSpinner{}
	f.A=a
	f.Vid=fcore.NewToken()
	f.VType= "Spinner"
	fcore.ViewMap.Set(f.Vid,f)
	f.A.NewObject(f.VType,f.Vid)
	return f
}

func (f *FSpinner) NotifyDataSetChanged() *FSpinner {
	f.A.SetAttr(f.Vid,"NotifyDataSetChanged","","")
	return f
}

func (v *FSpinner) List(ls []string) *FSpinner {
	v.A.SetAttr(v.Vid, "List", fcore.JsonArray(ls),"")
	return v
}

func (v *FSpinner) ListAdd(s string) *FSpinner {
	v.A.SetAttr(v.Vid, "ListAdd", s,"")
	return v
}

func (v *FSpinner) ListRemove(i int) *FSpinner {
	v.A.SetAttr(v.Vid, "ListRemove", fcore.SPrintf(i),"")
	return v
}
