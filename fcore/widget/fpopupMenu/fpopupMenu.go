package fpopupMenu

import (
	"github.com/StevenZack/fox/fcore"
	"strings"
)

type FPopupMenu struct {
	fcore.FBase
	showAfter bool
}

func NewFrom(i fcore.IBaseView)*FPopupMenu  {
	f:=&FPopupMenu{}
	f.Vid=fcore.NewToken()
	f.VType="PopupMenu"
	f.A=i.GetBaseView().A
	f.A.NewObject(f.VType,strings.Join([]string{f.Vid,i.GetBaseView().Vid},":"))
	return f
}
func (f *FPopupMenu) Menus(menuItems ...interface{}) *FPopupMenu {
	f.A.SetAttr(f.Vid,"Menus",fcore.JsonArray(menuItems),"")
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FPopupMenu) Show() *FPopupMenu {
	f.A.SetAttr(f.Vid,"Show","","")
	return f
}
func (f *FPopupMenu) Dismiss() *FPopupMenu {
	f.A.SetAttr(f.Vid,"Dismiss","","")
	return f
}
func (f *FPopupMenu) DeferShow() *FPopupMenu {
	f.showAfter=true
	return f
}