package fmenu

import "github.com/StevenZack/fox/fcore"

type FMenuItem struct {
	MyTitle        string
	MyIcon         string
	MyOnClick      string
	MyShowAsAction string
}
type FSubMenu struct {
	MyTitle   string
	MySubMenu []interface{}
}

func NewItem(title string) *FMenuItem {
	mi:=&FMenuItem{}
	mi.MyTitle= title
	return mi
}
func (m *FMenuItem) OnClick(f func()) *FMenuItem {
	fnId := fcore.NewToken()
	m.MyOnClick = fnId
	fcore.EventMap.Set(fnId, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		f()
		return ""
	})
	return m
}
func (m *FMenuItem) Icon(s string) *FMenuItem {
	m.MyIcon = s
	return m
}
func (m *FMenuItem) ShowAsAction() *FMenuItem {
	m.MyShowAsAction = "IF_ROOM"
	return m
}
func NewSub(title string, menuItems ...interface{}) *FSubMenu {
	m:=&FSubMenu{}
	m.MyTitle=title
	m.MySubMenu=menuItems
	return m
}