package ftoolbar

import "github.com/StevenZack/fox/fcore"

type FToolbar struct {
	fcore.FBaseView
}

func New(a fcore.IActivity) *FToolbar {
	f:=&FToolbar{}
	f.A=a
	f.Vid=fcore.NewToken()
	f.VType="Toolbar"
	fcore.ViewMap.Set(f.Vid,f)
	f.A.NewObject(f.VType,f.Vid)
	return f
}

func (f *FToolbar) Title(t string) *FToolbar {
	f.A.SetAttr(f.Vid,"Title",t,"")
	return f
}
func (f *FToolbar) SubTitle(s string) *FToolbar {
	f.A.SetAttr(f.Vid, "SubTitle", s, "")
	return f
}
func (f *FToolbar) GetTitle() string {
	return f.A.GetAttr(f.Vid,"Title")
}
func (f *FToolbar) GetSubTitle() string {
	return f.A.GetAttr(f.Vid, "SubTitle")
}
func (f *FToolbar) SubTitleColor(fcolor string) *FToolbar {
	f.A.SetAttr(f.Vid,"SubTitleColor",fcolor, "")
	return f
}

func (f *FToolbar) Menus(mis ...interface{}) *FToolbar{
	f.A.SetAttr(f.Vid,"Menus",fcore.JsonArray(mis),"")
	return f
}
func (f *FToolbar) BackNavigation(b bool) *FToolbar {
	f.A.SetAttr(f.Vid,"BackNavigation",fcore.SPrintf(b),"")
	return f
}
func (f *FToolbar) NavigationIcon(i string) *FToolbar {
	f.A.SetAttr(f.Vid,"NavigationIcon",i, "")
	return f
}
func (f *FToolbar) OnNavigationIconClick(fn func()) *FToolbar {
	fnId:=f.Vid+ ":onNavigationIconClick"
	fcore.EventMap.Set(fnId, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		fn()
		return ""
	})
	f.A.SetAttr(f.Vid,"OnNavigationIconClick",fnId, "")
	return f
}