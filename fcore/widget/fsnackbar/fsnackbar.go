package fsnackbar

import "github.com/StevenZack/fox/fcore"

type FSnackbar struct {
	fcore.FBaseView
}

func New(a fcore.IActivity)*FSnackbar  {
	f:=&FSnackbar{}
	f.Vid=fcore.NewToken()
	f.VType="Snackbar"
	f.A=a
	fcore.ViewMap.Set(f.Vid,f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}
