package fcore

import (
	"fmt"
)

type FBaseView struct {
	FBase
	A IActivity
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
