package fradio

import "github.com/StevenZack/fox/fcore"

type FRadio struct {
	fcore.FBaseView
}


func New(a fcore.IActivity) *FRadio {
	f := &FRadio{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "Radio"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}