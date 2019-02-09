package fservice

import "github.com/StevenZack/fox/fcore"

type FService struct {
	a fcore.IActivity
}

func StartService(a fcore.IActivity, onCreate func()) *FService {
	fnId:=fcore.NewToken()
	fcore.EventMap.Set(fnId, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		onCreate()
		return ""
	})
	f:=&FService{a:a}
	f.a.SetAttr("Service","OnCreate",fnId,"")
	return f
}
func New(a fcore.IActivity) *FService {
	return &FService{a:a}
}
func (f *FService) QuitButton(title string) *FService {
	f.a.SetAttr("Service","QuitButton",title,"")
	return f
}
func (f *FService) Title(title string) *FService {
	f.a.SetAttr("Service","Title",title,"")
	return f
}
func (f *FService) SubTitle(sub string) *FService {
	f.a.SetAttr("Service","SubTitle",sub,"")
	return f
}
func (f *FService) OnCreate(fn func()) *FService {
	fnId:=fcore.NewToken()
	fcore.EventMap.Set(fnId, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		fn()
		return ""
	})
	f.a.SetAttr("Service","OnCreate",fnId,"")
	return f
}
func (f *FService) OnQuit(fn func()) *FService {
	fnId:=fcore.NewToken()
	fcore.EventMap.Set(fnId, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		fn()
		return ""
	})
	f.a.SetAttr("Service","OnQuit",fnId,"")
	return f
}
func (f *FService) Show() *FService {
	f.a.SetAttr("Service","Show","","")
	return f
}
func (f *FService) FinishAllActivity() *FService {
	f.a.SetAttr("Service","FinishAllActivity","","")
	return f
}
func (f *FService) KillAll() *FService {
	f.a.SetAttr("Service","KillAll","","")
	return f
}
func (f *FService) Stop() *FService {
	f.a.SetAttr("Service","Stop","","")
	return f
}