package fsnackbar

import (
	"github.com/StevenZack/fox/fcore"
	"strings"
)

type FSnackbar struct {
	fcore.FBaseView
}

func New(a fcore.IActivity,anchor fcore.IBaseView)*FSnackbar  {
	f:=&FSnackbar{}
	f.Vid=fcore.NewToken()
	f.VType="Snackbar"
	f.A=a
	fcore.ViewMap.Set(f.Vid,f)
	f.A.NewObject(f.VType, strings.Join([]string{f.Vid,anchor.GetBaseView().Vid},","))
	return f
}
