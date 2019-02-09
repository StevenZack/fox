package ffilechooser

import (
	"encoding/json"
	"fmt"
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fcolor"
	"github.com/StevenZack/fox/fcore/values/fgravity"
	"github.com/StevenZack/fox/fcore/values/fpermission"
	"github.com/StevenZack/fox/fcore/values/fscaletype"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fcheck"
	"github.com/StevenZack/fox/fcore/widget/fdialog"
	"github.com/StevenZack/fox/fcore/widget/fedit"
	"github.com/StevenZack/fox/fcore/widget/fimage"
	"github.com/StevenZack/fox/fcore/widget/flist"
	"github.com/StevenZack/fox/fcore/widget/fmenu"
	"github.com/StevenZack/fox/fcore/widget/ftext"
	"github.com/StevenZack/fox/fcore/widget/ftoolbar"
	"io/ioutil"
	"mime"
	"os"
	"strings"
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
func New(a fcore.IActivity) *FFileChooser {
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
func (f *FFileChooser) Title(s string) *FFileChooser {
	f.title = s
	return f
}

func (f *FFileChooser) StartsFromLocation(dir string) *FFileChooser {
	f.currentDir = dir
	return f
}

// filter:"image/*.png"
func (f *FFileChooser) Filter(s func(fname string) bool) *FFileChooser {
	f.filter = s
	return f
}
func (f *FFileChooser) TypeSelectFile(rp func(string)) *FFileChooser {
	f.onResult = rp
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSelectFiles(rp func([]string)) *FFileChooser {
	f.onResults = rp
	f.multiple = true
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSelectFolder(rp func(string)) *FFileChooser {
	f.folderMode = true
	f.onResult = rp
	if f.showAfter {
		f.Show()
	}
	return f
}

func (f *FFileChooser) Show() {
	fcore.StartActivity(f.a, func(a fcore.IActivity) {
		if !fcore.CheckSelfPermission(a, fpermission.WRITE_EXTERNAL_STORAGE) {
			fcore.RequestPermissions(a, []string{fpermission.WRITE_EXTERNAL_STORAGE}, nil)
		}
		toolbar:=ftoolbar.New(a)
		listview:=flist.NewV(a)
		loadCurrentDir := func() {
			go func() {
				if f.currentDir == "" {
					f.currentDir = f.rootPath
				}
				simplifiedPath := strings.Replace(f.currentDir, fcore.Getrpath(fcore.GetExternalStorageDirectory(a)), "内部存储/", -1)
				infos, e := ioutil.ReadDir(f.currentDir)
				if e != nil {
					fmt.Println(`fileChooser loadCurrentDir error :`, e)
					fcore.RunOnUIThread(a, func() {
						fcore.ShowToast(a, e.Error())
					})
					return
				}
				f.flist = nil
				var fileList, dirList []*FFileItem
				for _, v := range infos {
					if f.folderMode && !v.IsDir() {
						continue
					}
					if f.filter != nil && !f.filter(v.Name()) {
						continue
					}
					if f.hiddenFilter != nil && !f.hiddenFilter(v.Name()) {
						continue
					}
					ffi := &FFileItem{
						info: v,
						path: f.currentDir + v.Name(),
						icon: getIconURLByFileType(f.currentDir + v.Name()),
					}
					if v.IsDir() {
						dirList = append(dirList, ffi)
					} else {
						fileList = append(fileList, ffi)
					}
				}
				f.flist = append(dirList, fileList...)
				fcore.RunOnUIThread(a, func() {
					toolbar.SubTitle(simplifiedPath)
					listview.NotifyDataSetChanged()
				})
			}()
		}

		fcore.SetOnBackPressed(a, func() bool {
			if fcore.Getrpath(f.currentDir) == f.rootPath {
				return true
			}
			f.currentDir = fcore.Getrpath(getDirOfFile(f.currentDir))
			loadCurrentDir()
			return false
		})

		getTwoButtons := func() fcore.IBaseView {
			if f.multiple || f.folderMode {
				bt_submit:=fbutton.New(a)
				return fbox.NewH(a).Size(-2, -1).Padding(10).Append(
					bt_submit.Text(f.positiveButtonTitle).Background(fcolor.RadiusCorner).TextColor(fcolor.White).Foreground(fcolor.RippleEffect).LayoutWeight(1).OnClick(func() {
						if f.multiple {
							var flist []string
							for k, v := range f.checkedMap {
								if v {
									flist = append(flist, k)
								}
							}
							f.onResults(flist)
							fcore.Finish(a)
						} else {
							if f.folderMode {
								f.onResult(f.currentDir)
								fcore.Finish(a)
							} else {
								fcore.ShowToast(a, "You're not support to click me")
								fcore.Finish(a)
							}
						}
					}),
					fbutton.New(a).Text(f.negativeButtonTitle).LayoutWeight(1).BackgroundColor(fcolor.Grey).TextColor(fcolor.White).Foreground(fcolor.RippleEffect).OnClick(func() {
						fcore.Finish(a)
					}),
				)
			}
			return nil
		}
		et_search:=fedit.New(a)
		fbox.NewV(a).DeferShow().Size(-2, -2).Append(
			toolbar.Title(f.title).Menus(
				fmenu.NewItem("反选").OnClick(func() {
					go func() {
						for _, v := range f.flist {
							if v.info.IsDir() {
								continue
							}
							if checked, ok := f.checkedMap[v.path]; ok {
								f.checkedMap[v.path] = !checked
							} else {
								f.checkedMap[v.path] = true
							}
						}
						fcore.RunOnUIThread(a, func() {
							listview.NotifyDataSetChanged()
						})
					}()
				}),
				fmenu.NewItem("隐藏文件").OnClick(func() {
					if f.hiddenFilter == nil {
						f.hiddenFilter = func(f string) bool {
							if f[:1] == "." {
								return false
							}
							return true
						}
					} else {
						f.hiddenFilter = nil
					}
					loadCurrentDir()
				}),
				fmenu.NewItem("其他存储位置").OnClick(func() {
					go func() {
						var storageNameList []string
						dis, e := ioutil.ReadDir("/storage")
						if e != nil {
							fcore.RunOnUIThread(a, func() {
								fcore.ShowToast(a, "读取存储设备失败:"+e.Error())
							})
						}
						for _, v := range dis {
							str := v.Name()
							if str == "emulated" {
								str = "内部存储"
							} else if str == "self" {
								continue
							}
							storageNameList = append(storageNameList, str)
						}
						fcore.RunOnUIThread(a, func() {
							d:=fdialog.New(a)
							d.Title("选择存储位置").Append(
								flist.NewV(a).Funcs(
									func(lv *flist.FList) fcore.IBaseView{
										return fbox.NewH(a).Foreground(fcolor.RippleEffect).Padding(10).Append(
											flist.SetItemId(ftext.New(a).LayoutGravity(fgravity.CenterVertical),lv,"title"),
										)
									},
									func(vh *flist.ViewHolder, pos int) {
										txt := vh.GetTextByItemId("title")
										txt.Text(storageNameList[pos])
									},
									func() int {
										return len(storageNameList)
									},
								).OnItemClick(func(pos int) {
									defer d.Dismiss()
									name := storageNameList[pos]
									if name == "内部存储" {
										f.rootPath = fcore.Getrpath(fcore.GetExternalStorageDirectory(a))
										f.currentDir = f.rootPath
										loadCurrentDir()
										return
									}
									f.rootPath = "/storage/" + name + "/"
									f.currentDir = f.rootPath
									loadCurrentDir()
								}).Padding(10),
							)
							d.Show()
						})
					}()
				}),
			),
			et_search.InputTypeSingleLineText().Size(-2, -1).Hint("在本页快速搜索").OnChange(func() {
				text:=et_search.GetText()
				go func() {
					if text == "" {
						loadCurrentDir()
						return
					}
					text = strings.ToLower(text)
					var fl []*FFileItem
					for _, v := range f.flist {
						if strings.Contains(strings.ToLower(v.info.Name()), text) {
							fl = append(fl, v)
						}
					}
					f.flist = fl
					fcore.RunOnUIThread(a, func() {
						listview.NotifyDataSetChanged()
					})
				}()
			}),
			listview.Funcs(
				func(lv *flist.FList) fcore.IBaseView{
					return flist.SetItemId(fbox.NewH(a).Size(-2, -1).Foreground(fcolor.RippleEffect).Padding(10).Append(
						flist.SetItemId(fcheck.New(a).LayoutGravity(fgravity.CenterVertical).Clickable(false),lv,"check"),
						flist.SetItemId(fimage.New(a).LayoutGravity(fgravity.CenterVertical).Size(60, 60).MarginRight(20).ScaleType(fscaletype.FitXY),lv,"icon"),
						flist.SetItemId(ftext.New(a).LayoutGravity(fgravity.CenterVertical),lv,"text"),
					),lv,"ctn")
				},
				func(vh *flist.ViewHolder, pos int) {
					check := vh.GetCheckBoxByItemId("check")
					_, checked := f.checkedMap[f.flist[pos].path]
					check.Checked(checked)
					if !f.multiple {
						check.Gone()
					} else {
						if f.folderMode {
							if f.flist[pos].info.IsDir() {
								check.Visible()
								check.Enabled(true)
							}
						} else {
							if f.flist[pos].info.IsDir() {
								check.Invisible()
								check.Enabled(false)
							} else {
								check.Visible()
								check.Enabled(true)
							}
						}
					}
					vh.GetImageByItemId("icon").Src(f.flist[pos].icon)
					vh.GetTextByItemId("text").Text(f.flist[pos].info.Name())
				},
				func() int {
					return len(f.flist)
				},
			).OnItemClick(func(pos int) {
				if f.flist[pos].info.IsDir() {
					if f.folderMode {
						if f.multiple {
							v := f.flist[pos]
							if checked, ok := f.checkedMap[v.path]; ok {
								f.checkedMap[v.path] = !checked
							} else {
								f.checkedMap[v.path] = true
							}
							listview.NotifyDataSetChanged()
						} else {
							f.currentDir = fcore.Getrpath(f.currentDir) + f.flist[pos].info.Name() + "/"
							loadCurrentDir()
						}
					} else {
						f.currentDir =fcore. Getrpath(f.currentDir) + f.flist[pos].info.Name() + "/"
						loadCurrentDir()
					}
				} else {
					if !f.folderMode {
						if f.multiple {
							go func() {
								v := f.flist[pos]
								if checked, ok := f.checkedMap[v.path]; ok {
									f.checkedMap[v.path] = !checked
								} else {
									f.checkedMap[v.path] = true
								}
								fcore.RunOnUIThread(a, func() {
									listview.NotifyDataSetChanged()
								})
							}()
						} else {
							f.onResult(f.flist[pos].path)
							fcore.Finish(a)
						}
					}
				}
			}).Size(-2, -2).LayoutWeight(1),
			getTwoButtons(),
		)
		loadCurrentDir()
	}, nil)
}

func getDirOfFile(path string) string {
	if path[len(path)-1:] == "/" {
		path = path[:len(path)-1]
	}
	for i := len(path) - 1; i > -1; i-- {
		if path[i:i+1] == "/" {
			return path[:i]
		}
	}
	return path
}
func (f *FFileChooser) DeferShow() *FFileChooser {
	f.showAfter = true
	return f
}
func getIconURLByFileType(fpath string) string {
	f, e := os.Open(fpath)
	if e != nil {
		fmt.Println(`os.open error :`, e)
		return ""
	}
	server := "https://jywjl.github.io/images/icons/"
	info, e := f.Stat()
	if e != nil {
		return server + "file.png"
	}
	if info.IsDir() {
		return server + "folder.png"
	}
	nameS := strings.Split(f.Name(), ".")
	ext := nameS[len(nameS)-1]
	mimeTypes := strings.Split(mime.TypeByExtension("."+ext), "/")
	switch mimeTypes[0] {
	case "audio":
		return server + "audio.png"
	case "image":
		return "file://" + fpath
	case "video":
		return server + "video.png"
	default:
		return server + "file.png"
	}
}
