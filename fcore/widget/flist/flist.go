package flist

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/StevenZack/fox/fcore"
)

type FList struct {
	fcore.FBaseView
	vh ViewHolder
}

func (f *FList) Assign(i **FList) *FList {
	(*i) = f
	return f
}
func NewV(a fcore.IActivity) *FList {
	f := &FList{}
	f.A = a
	f.vh.Vlist = make(map[string]string)
	f.Vid = fcore.NewToken()
	f.VType = "VList"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}
func NewH(a fcore.IActivity) *FList {
	f := &FList{}
	f.A = a
	f.vh.Vlist = make(map[string]string)
	f.Vid = fcore.NewToken()
	f.VType = "HList"
	fcore.ViewMap.Set(f.Vid, f)
	f.A.NewObject(f.VType, f.Vid)
	return f
}
func New(a fcore.IActivity) *FList {
	return NewV(a)
}
func (f *FList) Funcs(createView func(l *FList) fcore.IBaseView, bindData func(vh *ViewHolder, pos int), getCount func() int) *FList {
	fnId1 := fcore.NewToken()
	fnId2 := fcore.NewToken()
	fnId3 := fcore.NewToken()
	fcore.EventMap.Set(fnId1, func(a fcore.IActivity, v, v1, v2 string) string {
		f.vh.VID = createView(f).GetBaseView().Vid
		return fcore.JsonObject(f.vh)
	})
	fcore.EventMap.Set(fnId2, func(a fcore.IActivity, v, v1, v2 string) string {
		obd := TypeOnBindDataArgsBundle{}
		e := json.Unmarshal([]byte(v), &obd)
		if e != nil {
			fmt.Println(`FList.Funcs unmarshal1 error:`, e)
			return ""
		}
		vh := ViewHolder{}
		e = json.Unmarshal([]byte(obd.Str), &vh)
		if e != nil {
			fmt.Println(`FList.Funcs unmarshal2 error:`, e)
			return ""
		}
		bindData(&vh, obd.Position)
		return ""
	})
	fcore.EventMap.Set(fnId3, func(a fcore.IActivity, v, v1, v2 string) string {
		return fcore.SPrintf(getCount())
	})

	f.A.SetAttr(f.Vid, "OnGetView", fnId1, "")
	f.A.SetAttr(f.Vid, "OnBindData", fnId2, "")
	f.A.SetAttr(f.Vid, "OnGetCount", fnId3, "")
	return f
}
func (f *FList) NotifyDataSetChanged() *FList {
	f.A.SetAttr(f.Vid, "NotifyDataSetChanged", "", "")
	return f
}
func (f *FList) OnItemClick(fn func(int)) *FList {
	fnId := fcore.NewToken()
	fcore.EventMap.Set(fnId, func(a fcore.IActivity, v, v1, v2 string) string {
		pos, e := strconv.ParseInt(v, 10, 64)
		if e != nil {
			fmt.Println(`List.OnItemClick.parseInt error:`, e)
			return ""
		}
		fn(int(pos))
		return ""
	})
	f.A.SetAttr(f.Vid, "OnItemClick", fnId, "")
	return f
}

func (f *FList) OnItemLongClick(fn func(int)) *FList {
	fnId := fcore.NewToken()
	fcore.EventMap.Set(fnId, func(a fcore.IActivity, v, v1, v2 string) string {
		pos, e := strconv.ParseInt(v, 10, 64)
		if e != nil {
			fmt.Println(`List.OnItemLongClick.parseInt error:`, e)
			return ""
		}
		fn(int(pos))
		return ""
	})
	f.A.SetAttr(f.Vid, "OnItemLongClick", fnId, "")
	return f
}
