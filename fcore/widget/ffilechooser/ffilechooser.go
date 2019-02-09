package ffilechooser

import (
	"encoding/json"
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fpermission"
	"os"
)

type FFileChooser struct {
	a                                               fcore.IActivity
	showAfter                                       bool
	title, positiveButtonTitle, negativeButtonTitle string
	onResult                                        func(string)
	onResults                                       func([]string)
	filter, hiddenFilter                            func(string) bool
	multiple, folderMode                            bool

	rootPath   string
	currentDir string
	flist      []*FFileItem
	checkedMap map[string]bool
}

type FFileItem struct {
	info os.FileInfo
	path string
	icon string
}
func FileChooser(a fcore.IActivity) *FFileChooser {
	f := &FFileChooser{a: a, positiveButtonTitle: "确定", negativeButtonTitle: "取消", rootPath: fcore.Getrpath(fcore.GetExternalStorageDirectory(a)), checkedMap: make(map[string]bool)}
	f.hiddenFilter = func(f string) bool {
		if f[:1] == "." {
			return false
		}
		return true
	}
	return f
}
func (f *FFileChooser) ButtonTitle(positive, negative string) *FFileChooser {
	f.positiveButtonTitle = positive
	f.negativeButtonTitle = negative
	return f
}
func OpenSystemFileChooser(a fcore.IActivity, selectType string, allowMultiple bool, callback func([]string)) {
	if !fcore.CheckSelfPermission(a, fpermission.READ_EXTERNAL_STORAGE) {
		fcore.RequestPermissions(a, []string{fpermission.READ_EXTERNAL_STORAGE}, func(bs []bool) {
			OpenSystemFileChooser(a, selectType, allowMultiple, callback)
		})
		return
	}
	fnId := fcore.NewToken()
	fcore.EventMap.Set(fnId, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		var strs []string
		e:=json.Unmarshal([]byte(s),&strs)
		if e==nil{
			callback(strs)
		}
		return ""
	})
	a.SetAttr("Activity", "OpenFileChooser", fcore.JsonArray([]string{
		selectType,
		fcore.SPrintf(allowMultiple),
		fnId,
	}),"")
}