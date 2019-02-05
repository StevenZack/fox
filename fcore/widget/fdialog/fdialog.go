package fdialog

import "github.com/StevenZack/fox/fcore"

type FDialog struct {
	fcore.FBaseView
}

func New(a fcore.IActivity) *FDialog {
	v:=&FDialog{}
	v.A=a
	v.Vid=fcore.NewToken()
	v.VType= "Dialog"
	fcore.ViewMap.Set(v.Vid,v)
	v.A.NewObject(v.VType,v.Vid)
	return v
}


func (v *FDialog) Title(s string) *FDialog {
	v.A.SetAttr(v.Vid, "Title", s,"")
	return v
}
func (v *FDialog) Append(iv fcore.IBaseView) *FDialog {
	v.A.SetAttr(v.Vid, "View", iv.GetBaseView().Vid,"")
	return v
}
func (v *FDialog) PositiveButton(text string, onClick func(*FDialog)) *FDialog {
	fnId := v.Vid+"PositiveButtonClicked"
	fcore.EventMap.Set(fnId, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		onClick(v)
		return ""
	})
	v.A.SetAttr(v.Vid, "PositiveButton", text, fnId)
	return v
}

func (v *FDialog) NegativeButton(text string, onClick func(*FDialog)) *FDialog {
	fnId := v.Vid+"NegativeButtonClicked"
	fcore.EventMap.Set(fnId, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		onClick(v)
		return ""
	})
	v.A.SetAttr(v.Vid, "NegativeButton", text, fnId)
	return v
}
func (v *FDialog) Show() *FDialog {
	v.A.SetAttr(v.Vid, "Show", "","")
	return v
}
func (v *FDialog) Dismiss() *FDialog {
	v.A.SetAttr(v.Vid, "Dismiss", "","")
	return v
}
