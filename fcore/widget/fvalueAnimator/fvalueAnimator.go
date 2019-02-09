package fvalueAnimator

import "github.com/StevenZack/fox/fcore"

type FValueAnimator struct {
	fcore.FBase
}

func New(a fcore.IActivity) *FValueAnimator {
	v := &FValueAnimator{}
	v.Vid = fcore.NewToken()
	v.VType= "ValueAnimator"
	v.A = a
	v.A.NewObject(v.VType,v.Vid)
	return v
}
func (v *FValueAnimator) OfFloat(fs ...float64) *FValueAnimator {
	v.A.SetAttr(v.Vid, "OfFloat", fcore.JsonArray(fs),"")
	return v
}
func (v *FValueAnimator) OfInt(fs ...int) *FValueAnimator {
	v.A.SetAttr(v.Vid, "OfInt", fcore.JsonArray(fs),"")
	return v
}
func (v *FValueAnimator) Duration(ms int64) *FValueAnimator {
	v.A.SetAttr(v.Vid, "Duration",fcore. SPrintf(ms),"")
	return v
}
func (v *FValueAnimator) OnValueChanged(f func(string)) *FValueAnimator {
	fnId := fcore.NewToken()
	fcore.EventMap.Set(fnId, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		f(s)
		return ""
	})
	v.A.SetAttr(v.Vid, "ValueChangedListener", fnId,"")
	return v
}
func (v *FValueAnimator) Start() *FValueAnimator {
	v.A.SetAttr(v.Vid, "Start", "","")
	return v
}
