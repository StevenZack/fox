package fcore

import (
	"encoding/json"
	"strings"
)

type FActivity struct {
	ui                                                        IActivity
	intent                                                    Intent
	fnOnCreate, fnOnResume, fnOnPause, fnOnStart, fnOnDestroy func(IActivity)
}

// Intent
type Intent struct {
	Action string
	Paths  []string
	Extras map[string]string
}

func GetIntent(a IActivity) (i Intent) {
	act := activityIdMap.Get(a.GetActivityId())
	if act != nil {
		i = act.intent
	}
	return
}

// createView method : will be invoked when new activity created ; conf is nil by default
func StartActivity(a IActivity, createView func(IActivity), conf *FAppConf) {
	act := &FActivity{}
	act.fnOnCreate = createView
	if conf == nil {
		conf = NewAppConf()
	}
	activityIdMap.Set(conf.FActivityId, act)
	a.NewObject("Activity", JsonObject(conf))
}
func startActivityTrigger(a IActivity, activityId, onEvent, intentJson string) {
	act := activityIdMap.Get(activityId)
	if act == nil {
		println("fox.startActivityTrigger.act is nil")
		return
	}
	act.ui = a
	switch onEvent {
	case "OnCreate":
		it := Intent{}
		e := json.Unmarshal([]byte(intentJson), &it)
		if e != nil {
			println("fox.startActivityTrigger.unmarshal error:", e)
		}
		act.intent = it
		act.fnOnCreate(act.ui)
	case "OnResume":
		act.fnOnResume(act.ui)
	case "OnStart":
		act.fnOnStart(act.ui)
	case "OnPause":
		act.fnOnPause(act.ui)
	case "OnDestroy":
		act.fnOnDestroy(act.ui)
	}
}
func StartUriIntent(a IActivity, uri string) {
	a.SetAttr("Activity", "StartUriIntent", uri, "")
}

func GetPackageName(a IActivity) string {
	return a.GetAttr("Activity", "PackageName")
}
func RunOnUIThreadWithFnId(a IActivity, f func(), fnId string) {
	EventMap.Set(fnId, func(IActivity, string, string, string) string {
		f()
		return ""
	})
	a.SetAttr("Activity", "RunOnUIThread", fnId, "")
}
func RunOnUIThread(a IActivity, f func()) {
	RunOnUIThreadWithFnId(a, f, NewToken())
}
func CheckSelfPermission(a IActivity, fpermission string) bool {
	return a.SetAttr("Activity","CheckSelfPermission",fpermission,"")=="true"
}
func RequestPermissions(a IActivity, perms []string, onResult func([]bool))  {
	fnId:=NewToken()
	EventMap.Set(fnId, func(activity IActivity, s string, s2 string, s3 string) string {
		if onResult == nil {
			return ""
		}
		var bs []bool
		e:=json.Unmarshal([]byte(s),&bs)
		if e != nil {
			return ""
		}
		onResult(bs)
		return ""
	})
	a.SetAttr("Activity","RequestPermissions",strings.Join(perms,":"),fnId)
}
func ShowToast(a IActivity,text string)  {
	a.SetAttr("Activity","ShowToast",text,"")
}
func GetIMEI(a IActivity) string {
	return a.GetAttr("Activity","IMEI")
}
func GetUniqueID(a IActivity)string  {
	return a.GetAttr("Activity","UniqueID")
}
func GetExternalStorageDirectory(a IActivity) string {
	return a.GetAttr("Activity","ExternalStorageDirectory")
}
func GetModel(a IActivity) string {
	return a.GetAttr("Activity","Build.MODEL")
}
func SetOnBackPressed(a IActivity, fn func() bool) {
	fnId:=NewToken()
	EventMap.Set(fnId, func(activity IActivity, s string, s2 string, s3 string) string {
		return SPrintf(fn())
	})
	a.SetAttr("Activity","OnBackPressed",fnId,"")
}
func Finish(a IActivity) {
	a.SetAttr("Activity","Finish","","")
}