package fedit

import "github.com/StevenZack/fox/fcore"

type FEdit struct {
	fcore.FBaseView
}

func New(a fcore.IActivity) *FEdit {
	f := &FEdit{}
	f.A = a
	f.Vid = fcore.NewToken()
	f.VType = "Edit"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}
func (f *FEdit) Assign(i **FEdit) *FEdit {
	(*i)=f
	return f
}
func (v *FEdit) Text(s string) *FEdit {
	v.A.SetAttr(v.Vid, "Text", s,"")
	return v
}
func (v *FEdit) TextColor(s string) *FEdit {
	v.A.SetAttr(v.Vid, "TextColor", s,"")
	return v
}
func (v *FEdit) TextSize(dpsize int) *FEdit {
	v.A.SetAttr(v.Vid, "TextSize", fcore.SPrintf(dpsize),"")
	return v
}
func (v *FEdit) InputTypeSingleLineText() *FEdit {
	v.A.SetAttr(v.Vid, "InputType", "Text","")
	return v
}
func (v *FEdit) InputTypeNumber() *FEdit {
	v.A.SetAttr(v.Vid, "InputType", "Number","")
	return v
}
func (v *FEdit) InputTypePassword() *FEdit {
	v.A.SetAttr(v.Vid, "InputType", "Password","")
	return v
}
func (v *FEdit) InputTypeEnglish() *FEdit {
	v.A.SetAttr(v.Vid, "InputType", "English","")
	return v
}
func (v *FEdit) MaxLines(i int) *FEdit {
	v.A.SetAttr(v.Vid, "MaxLines", fcore.SPrintf(i),"")
	return v
}
func (v *FEdit) MaxEms(i int) *FEdit {
	v.A.SetAttr(v.Vid, "MaxEms", fcore.SPrintf(i),"")
	return v
}
func (v *FEdit) Hint(s string) *FEdit {
	v.A.SetAttr(v.Vid, "Hint", s,"")
	return v
}
func (v *FEdit) MaxLength(i int) *FEdit {
	v.A.SetAttr(v.Vid, "MaxLength", fcore.SPrintf(i),"")
	return v
}
