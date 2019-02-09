package fcore

import (
	"fmt"
	"strings"
)

type FBaseView struct {
	FBase
	srcBackground, srcForeground string
	FnOnClick                    func()
	src                          string
	showAfter                    bool
}
type IBaseView interface {
	GetBaseView() *FBaseView
}
func (f *FBaseView) GetBaseView() *FBaseView {
	return f
}
func (f *FBaseView) GetText() string {
	return f.A.GetAttr(f.Vid, "Text")
}
func (f *FBaseView) Show() *FBaseView {
	f.A.Show(f.Vid)
	return f
}
func (f *FBaseView) Background(fbg string) *FBaseView {
	if f.srcBackground == fbg {
		return f
	}
	if StartsWith(fbg, "http") {
		go CacheNetFile(fbg,"/data/data/"+GetPackageName(f.A)+"/cacheDir/", func(s string) {
			RunOnUIThreadWithFnId(f.A, func() {
				f.A.SetAttr(f.Vid,"Background","file://"+s,"")
				f.srcBackground=fbg
			},f.Vid+":background")
		})
		return f
	}
	f.A.SetAttr(f.Vid,"Background",fbg,"")
	return f
}

func (f *FBaseView) Src(s string) *FBaseView {
	if s == f.src {
		return f
	}
	if StartsWith(s, "http") {
		go CacheNetFile(s, "/data/data/"+GetPackageName(f.A)+"/cacheDir/", func(path string) {
			RunOnUIThreadWithFnId(f.A, func() {
				f.A.SetAttr(f.Vid, "Src", "file://"+path, "")
				f.src = s
			}, f.Vid+":src")
		})
		return f
	}
	f.A.SetAttr(f.Vid, "Src", s, "")
	f.src = s
	return f
}

func (f *FBaseView) Foreground(s string) *FBaseView {
	if f.srcForeground==s {
		return f
	}
	if StartsWith(s, "http") {
		go CacheNetFile(s,"/data/data/"+GetPackageName(f.A)+"/cacheDir/", func(path string) {
			RunOnUIThreadWithFnId(f.A, func() {
				f.A.SetAttr(f.Vid,"Foreground","file://"+path,"")
				f.srcForeground=s
			},f.Vid+":foreground")
		})
		return f
	}
	f.A.SetAttr(f.Vid,"Foreground",s,"")
	return f
}
func (f *FBaseView) BackgroundColor(fcolor string) *FBaseView {
	f.A.SetAttr(f.Vid,"BackgroundColor",fcolor,"")
	return f
}
func (f *FBaseView) Size(w, h int) *FBaseView {
	f.A.SetAttr(f.Vid, "Size", fmt.Sprint(w), fmt.Sprint(h))
	return f
}

func (v *FBaseView) X(x float64) *FBaseView {
	v.A.SetAttr(v.Vid, "X", SPrintf(x), "")
	return v
}
func (v *FBaseView) Y(y float64) *FBaseView {
	v.A.SetAttr(v.Vid, "Y", SPrintf(y), "")
	return v
}
func (v *FBaseView) PivotX(x float64) *FBaseView {
	v.A.SetAttr(v.Vid, "PivotX", SPrintf(x), "")
	return v
}
func (v *FBaseView) PivotY(y float64) *FBaseView {
	v.A.SetAttr(v.Vid, "PivotY", SPrintf(y), "")
	return v
}
func (v *FBaseView) ScaleX(x float64) *FBaseView {
	v.A.SetAttr(v.Vid, "ScaleX", SPrintf(x), "")
	return v
}
func (v *FBaseView) ScaleY(y float64) *FBaseView {
	v.A.SetAttr(v.Vid, "ScaleY", SPrintf(y), "")
	return v
}
func (v *FBaseView) Rotation(r float64) *FBaseView {
	v.A.SetAttr(v.Vid, "Rotation", SPrintf(r), "")
	return v
}

func (v *FBaseView) GetX() float64 {
	x := v.A.GetAttr(v.Vid, "X")
	return a2f(x)
}
func (v *FBaseView) GetY() float64 {
	x := v.A.GetAttr(v.Vid, "Y")
	return a2f(x)
}
func (v *FBaseView) GetWidth() float64 {
	x := v.A.GetAttr(v.Vid, "Width")
	return a2f(x)
}
func (v *FBaseView) GetHeight() float64 {
	x := v.A.GetAttr(v.Vid, "Height")
	return a2f(x)
}
func (v *FBaseView) GetPivotX() float64 {
	x := v.A.GetAttr(v.Vid, "PivotX")
	return a2f(x)
}
func (v *FBaseView) GetPivotY() float64 {
	x := v.A.GetAttr(v.Vid, "PivotY")
	return a2f(x)
}
func (v *FBaseView) GetScaleX() float64 {
	x := v.A.GetAttr(v.Vid, "ScaleX")
	return a2f(x)
}
func (v *FBaseView) GetScaleY() float64 {
	x := v.A.GetAttr(v.Vid, "ScaleY")
	return a2f(x)
}
func (v *FBaseView) GetRotation() float64 {
	x := v.A.GetAttr(v.Vid, "Rotation")
	return a2f(x)
}

func (v *FBaseView) Visible() *FBaseView {
	v.A.SetAttr(v.Vid, "Visibility", "VISIBLE", "")
	return v
}
func (v *FBaseView) Invisible() *FBaseView {
	v.A.SetAttr(v.Vid, "Visibility", "INVISIBLE", "")
	return v
}
func (v *FBaseView) Gone() *FBaseView {
	v.A.SetAttr(v.Vid, "Visibility", "GONE", "")
	return v
}
func (v *FBaseView) IsVisible() bool {
	return v.A.GetAttr(v.Vid, "Visibility") == "VISIBLE"
}
func (v *FBaseView) IsInvisible() bool {
	return v.A.GetAttr(v.Vid, "Visibility") == "INVISIBLE"
}
func (v *FBaseView) IsGone() bool {
	return v.A.GetAttr(v.Vid, "Visibility") == "GONE"
}

//padding ...
func (v *FBaseView) PaddingExactly(left, top, right, bottom int) *FBaseView {
	v.A.SetAttr(v.Vid, "Padding", JsonArray([]int{left, top, right, bottom}), "")
	return v
}

func (v *FBaseView) PaddingLeft(dp int) *FBaseView {
	v.PaddingExactly(dp, 0, 0, 0)
	return v
}
func (v *FBaseView) PaddingTop(dp int) *FBaseView {
	v.PaddingExactly(0, dp, 0, 0)
	return v
}
func (v *FBaseView) PaddingRight(dp int) *FBaseView {
	v.PaddingExactly(0, 0, dp, 0)
	return v
}
func (v *FBaseView) PaddingBottom(dp int) *FBaseView {
	v.PaddingExactly(0, 0, 0, dp)
	return v
}
func (v *FBaseView) Padding(dp int) *FBaseView {
	v.PaddingExactly(dp, dp, dp, dp)
	return v
}
func (v *FBaseView) MarginExactly(left, top, right, bottom int) *FBaseView {
	v.A.SetAttr(v.Vid, "Margin", JsonArray([]int{left, top, right, bottom}), "")
	return v
}

func (v *FBaseView) MarginLeft(dp int) *FBaseView {
	v.MarginExactly(dp, 0, 0, 0)
	return v
}
func (v *FBaseView) MarginTop(dp int) *FBaseView {
	v.MarginExactly(0, dp, 0, 0)
	return v
}
func (v *FBaseView) MarginRight(dp int) *FBaseView {
	v.MarginExactly(0, 0, dp, 0)
	return v
}
func (v *FBaseView) MarginBottom(dp int) *FBaseView {
	v.MarginExactly(0, 0, 0, dp)
	return v
}
func (v *FBaseView) Margin(dp int) *FBaseView {
	v.MarginExactly(dp, dp, dp, dp)
	return v
}

func (v *FBaseView) LayoutGravity(fgravity int) *FBaseView {
	v.A.SetAttr(v.Vid, "LayoutGravity", SPrintf(fgravity), "")
	return v
}
func (v *FBaseView) Elevation(dp float64) *FBaseView {
	v.A.SetAttr(v.Vid, "Elevation", SPrintf(dp), "")
	return v
}
func (v *FBaseView) LayoutWeight(f int) *FBaseView {
	v.A.SetAttr(v.Vid, "LayoutWeight", SPrintf(f), "")
	return v
}
func (v *FBaseView) Clickable(b bool) *FBaseView {
	v.A.SetAttr(v.Vid, "Clickable", SPrintf(b), "")
	return v
}

//constraint
var parent = "_Parent_"

func (v *FBaseView) Top2TopOfParent() *FBaseView {
	v.A.SetAttr(v.Vid, "Top2TopOf", parent, "")
	return v
}

func (v *FBaseView) Top2BottomOfParent() *FBaseView {
	v.A.SetAttr(v.Vid, "Top2BottomOf", parent, "")
	return v
}

func (v *FBaseView) Bottom2TopOfParent() *FBaseView {
	v.A.SetAttr(v.Vid, "Bottom2TopOf", parent, "")
	return v
}

func (v *FBaseView) Bottom2BottomOfParent(i IBaseView) *FBaseView {
	v.A.SetAttr(v.Vid, "Bottom2BottomOf", parent, "")
	return v
}

func (v *FBaseView) Left2LeftOfParent() *FBaseView {
	v.A.SetAttr(v.Vid, "Left2LeftOf", parent, "")
	return v
}

func (v *FBaseView) Right2RightOfParent() *FBaseView {
	v.A.SetAttr(v.Vid, "Right2RightOf", parent, "")
	return v
}

func (v *FBaseView) Left2RightOfParent() *FBaseView {
	v.A.SetAttr(v.Vid, "Left2RightOf", parent, "")
	return v
}

func (v *FBaseView) Right2LeftOfParent() *FBaseView {
	v.A.SetAttr(v.Vid, "Right2LeftOf", parent, "")
	return v
}

func (v *FBaseView) Top2TopOf(i IBaseView) *FBaseView {
	v.A.SetAttr(v.Vid, "Top2TopOf", i.GetBaseView().Vid, "")
	return v
}

func (v *FBaseView) Top2BottomOf(i IBaseView) *FBaseView {
	v.A.SetAttr(v.Vid, "Top2BottomOf", i.GetBaseView().Vid, "")
	return v
}

func (v *FBaseView) Bottom2TopOf(i IBaseView) *FBaseView {
	v.A.SetAttr(v.Vid, "Bottom2TopOf", i.GetBaseView().Vid, "")
	return v
}

func (v *FBaseView) Bottom2BottomOf(i IBaseView) *FBaseView {
	v.A.SetAttr(v.Vid, "Bottom2BottomOf", i.GetBaseView().Vid, "")
	return v
}

func (v *FBaseView) Left2LeftOf(i IBaseView) *FBaseView {
	v.A.SetAttr(v.Vid, "Left2LeftOf", i.GetBaseView().Vid, "")
	return v
}

func (v *FBaseView) Right2RightOf(i IBaseView) *FBaseView {
	v.A.SetAttr(v.Vid, "Right2RightOf", i.GetBaseView().Vid, "")
	return v
}

func (v *FBaseView) Left2RightOf(i IBaseView) *FBaseView {
	v.A.SetAttr(v.Vid, "Left2RightOf", i.GetBaseView().Vid, "")
	return v
}

func (v *FBaseView) Right2LeftOf(i IBaseView) *FBaseView {
	v.A.SetAttr(v.Vid, "Right2LeftOf", i.GetBaseView().Vid, "")
	return v
}
func (v *FBaseView) CenterX() *FBaseView {
	v.A.SetAttr(v.Vid, "CenterX", "", "")
	return v
}
func (v *FBaseView) CenterY() *FBaseView {
	v.A.SetAttr(v.Vid, "CenterY", "", "")
	return v
}
func (v *FBaseView) WidthPercent(num float64) *FBaseView {
	v.A.SetAttr(v.Vid, "WidthPercent", SPrintf(num), "")
	return v
}
func (v *FBaseView) HeightPercent(num float64) *FBaseView {
	v.A.SetAttr(v.Vid, "HeightPercent", SPrintf(num), "")
	return v
}

func (f *FBaseView) OnClick(fn func()) *FBaseView {
	f.FnOnClick = fn
	return f
}
func (f *FBaseView) Assign(i **FBaseView) *FBaseView {
	(*i)=f
	return f
}
func (v *FBaseView) Text(s string) *FBaseView {
	v.A.SetAttr(v.Vid, "Text", s,"")
	return v
}
func (v *FBaseView) TextColor(s string) *FBaseView {
	v.A.SetAttr(v.Vid, "TextColor", s,"")
	return v
}
func (v *FBaseView) TextSize(dpsize int) *FBaseView {
	v.A.SetAttr(v.Vid, "TextSize", SPrintf(dpsize),"")
	return v
}
func (v *FBaseView) InputTypeSingleLineText() *FBaseView {
	v.A.SetAttr(v.Vid, "InputType", "Text","")
	return v
}
func (v *FBaseView) InputTypeNumber() *FBaseView {
	v.A.SetAttr(v.Vid, "InputType", "Number","")
	return v
}
func (v *FBaseView) InputTypePassword() *FBaseView {
	v.A.SetAttr(v.Vid, "InputType", "Password","")
	return v
}
func (v *FBaseView) InputTypeEnglish() *FBaseView {
	v.A.SetAttr(v.Vid, "InputType", "English","")
	return v
}
func (v *FBaseView) MaxLines(i int) *FBaseView {
	v.A.SetAttr(v.Vid, "MaxLines", SPrintf(i),"")
	return v
}
func (v *FBaseView) MaxEms(i int) *FBaseView {
	v.A.SetAttr(v.Vid, "MaxEms", SPrintf(i),"")
	return v
}
func (v *FBaseView) Hint(s string) *FBaseView {
	v.A.SetAttr(v.Vid, "Hint", s,"")
	return v
}
func (v *FBaseView) MaxLength(i int) *FBaseView {
	v.A.SetAttr(v.Vid, "MaxLength", SPrintf(i),"")
	return v
}
func (f *FBaseView) OnChange(fn func()) *FBaseView {
	EventMap.Set(f.Vid+":onchange", func(activity IActivity, s string, s2 string, s3 string) string {
		fn()
		return ""
	})
	f.A.SetAttr(f.Vid,"OnChange",f.Vid+":onchange", "")
	return f
}
func (f *FBaseView) Enabled(b bool)*FBaseView  {
	f.A.SetAttr(f.Vid,"Enabled",SPrintf(b), "")
	return f
}
func (f *FBaseView) IsEnabled() bool {
	return f.A.GetAttr(f.Vid,"Enabled")=="true"
}

func (f *FBaseView) Append(is ...IBaseView) *FBaseView {
	var vids []string
	for _, i := range is {
		if i == nil {
			continue
		}
		vids = append(vids, i.GetBaseView().Vid)
	}
	f.A.SetAttr(f.Vid, "Append", strings.Join(vids, ","), "")
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FBaseView) AddView(i IBaseView) *FBaseView {
	f.A.SetAttr(f.Vid,"AddView",i.GetBaseView().Vid, "")
	return f
}
func (f *FBaseView) AddViewAt(i IBaseView, pos int) *FBaseView {
	f.A.SetAttr(f.Vid,"AddViewAt",i.GetBaseView().Vid,SPrintf(pos))
	return f
}
func (f *FBaseView) DeferShow() *FBaseView {
	f.showAfter=true
	return f
}

func (f *FBaseView) Horizontal() *FBaseView {
	f.A.SetAttr(f.Vid, "Orientation", "HORIZONTAL", "")
	return f
}
func (f *FBaseView) IsVertical() bool {
	return f.A.GetAttr(f.Vid,"IsVertical")=="true"
}
func (f *FBaseView) Vertical() *FBaseView {
	f.A.SetAttr(f.Vid, "Orientation", "VERTICAL", "")
	return f
}
func (f *FBaseView) Selected(b bool) *FBaseView {
	f.A.SetAttr(f.Vid,"Selected",SPrintf(b),"")
	return f
}
func (f *FBaseView) IsSelected() bool {
	return f.A.GetAttr(f.Vid,"IsSelected")=="true"
}
func (f *FBaseView) Tabs(ts ...*FTab) *FBaseView {
	for _,t:=range ts{
		f.A.SetAttr(f.Vid,"AddTab",JsonObject(t),"")
	}
	return f
}
func (f *FBaseView) TabTextColors(normal, selected string) *FBaseView {
	f.A.SetAttr(f.Vid,"TabTextColors",normal, selected)
	return f
}
func (f *FBaseView) TabIndicatorColor(color string) *FBaseView {
	f.A.SetAttr(f.Vid,"TabIndicatorColor",color,"")
	return f
}
func (f *FBaseView) Pages(ps ...*FPage) *FBaseView {
	if ps==nil{
		return f
	}
	f.A.SetAttr(f.Vid,"Pages",JsonArray(ps), "")
	return f
}
func (f *FBaseView) OnGetPage(getView func(pos int)IBaseView,getCount func()int)*FBaseView  {
	fnId,fnId2:=NewToken(),NewToken()
	EventMap.Set(fnId, func(activity IActivity, s string, s2 string, s3 string) string {
		i,e:=a2i(s)
		if e != nil {
			return ""
		}
		return getView(i).GetBaseView().Vid
	})
	EventMap.Set(fnId2, func(activity IActivity, s string, s2 string, s3 string) string {
		return SPrintf(getCount())
	})
	f.A.SetAttr(f.Vid,"OnCreateView",fnId,"")
	f.A.SetAttr(f.Vid,"OnGetCount",fnId2,"")
	return f
}
func (f *FBaseView) BindTabLayout(t *FBaseView) *FBaseView {
	f.A.SetAttr(f.Vid,"TabLayout",t.Vid,"")
	return f
}
func (f *FBaseView) CurrentItem(pos int, soomth bool) *FBaseView {
	f.A.SetAttr(f.Vid,"CurrentItem",SPrintf(pos),SPrintf(soomth))
	return f
}
func (f *FBaseView) OnPageSelected(fn func(int))*FBaseView  {
	fnId:=NewToken()
	EventMap.Set(fnId, func(activity IActivity, s string, s2 string, s3 string) string {
		i,e:=a2i(s)
		if e != nil {
			return ""
		}
		fn(i)
		return ""
	})
	f.A.SetAttr(f.Vid,"OnPageSelected",fnId, "")
	return f
}
func (f *FBaseView) Icon(s string)*FBaseView  {
	f.A.SetAttr(f.Vid,"Icon",s,"")
	return f
}