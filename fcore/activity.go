package fcore

import (
	"encoding/json"
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
