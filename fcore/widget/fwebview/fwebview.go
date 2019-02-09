package fwebview

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/values/fcolor"
)

type FWebView struct {
	fcore.FBaseView
}

func New(a fcore.IActivity) *FWebView {
	f:=&FWebView{}
	f.A=a
	f.Vid=fcore.NewToken()
	f.VType="WebView"
	fcore.ViewMap.Set(f.Vid,f)
	f.A.NewObject(f.VType,f.Vid)
	fnId := f.Vid+":onclick"
	fcore.EventMap.Set(fnId, func(fcore.IActivity, string, string, string) string {
		if f.FnOnClick != nil {
			f.FnOnClick()
		}
		return ""
	})
	f.A.SetAttr(f.Vid, "OnClick", fnId, "")
	return f
}

func NewItem(a fcore.IActivity, uri string) *fcore.FBaseView {
	w := New(a)
	if len(uri) > 0 && uri[:1] == "/" {
		uri = "file://" + uri
	}
	return w.Focusable(false).LoadUri(uri).BackgroundColor(fcolor.Transparent).Size(-1, -1)
}
// --------------------------------------------------------

func (v *FWebView) LoadUri(s string) *FWebView {
	v.A.SetAttr(v.Vid, "Uri", s,"")
	return v
}
func (v *FWebView) Focusable(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "Focusable", "true","")
	} else {
		v.A.SetAttr(v.Vid, "Focusable", "false","")
	}
	return v
}
func (v *FWebView) SupportZoom(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "SupportZoom", "true","")
	} else {
		v.A.SetAttr(v.Vid, "SupportZoom", "false","")
	}
	return v
}
func (v *FWebView) BuiltInZoomControls(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "BuiltInZoomControls", "true","")
	} else {
		v.A.SetAttr(v.Vid, "BuiltInZoomControls", "false","")
	}
	return v
}
func (v *FWebView) UseWideViewPort(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "UseWideViewPort", "true","")
	} else {
		v.A.SetAttr(v.Vid, "UseWideViewPort", "false","")
	}
	return v
}
func (v *FWebView) AllowContentAccess(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "AllowContentAccess", "true","")
	} else {
		v.A.SetAttr(v.Vid, "AllowContentAccess", "false","")
	}
	return v
}
func (v *FWebView) AllowFileAccess(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "AllowFileAccess", "true","")
	} else {
		v.A.SetAttr(v.Vid, "AllowFileAccess", "false","")
	}
	return v
}
func (v *FWebView) AllowFileAccessFromFileURLs(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "AllowFileAccessFromFileURLs", "true","")
	} else {
		v.A.SetAttr(v.Vid, "AllowFileAccessFromFileURLs", "false","")
	}
	return v
}
func (v *FWebView) AllowUniversalAccessFromFileURLs(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "AllowUniversalAccessFromFileURLs", "true","")
	} else {
		v.A.SetAttr(v.Vid, "AllowUniversalAccessFromFileURLs", "false","")
	}
	return v
}
func (v *FWebView) AppCacheEnabled(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "AppCacheEnabled", "true","")
	} else {
		v.A.SetAttr(v.Vid, "AppCacheEnabled", "false","")
	}
	return v
}
func (v *FWebView) BlockNetworkImage(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "BlockNetworkImage", "true","")
	} else {
		v.A.SetAttr(v.Vid, "BlockNetworkImage", "false","")
	}
	return v
}
func (v *FWebView) BlockNetworkLoads(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "BlockNetworkLoads", "true","")
	} else {
		v.A.SetAttr(v.Vid, "BlockNetworkLoads", "false","")
	}
	return v
}
func (v *FWebView) JavaScriptEnabled(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "JavaScriptEnabled", "true","")
	} else {
		v.A.SetAttr(v.Vid, "JavaScriptEnabled", "false","")
	}
	return v
}
func (v *FWebView) OffscreenPreRaster(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "OffscreenPreRaster", "true","")
	} else {
		v.A.SetAttr(v.Vid, "OffscreenPreRaster", "false","")
	}
	return v
}
func (v *FWebView) SaveFormData(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "SaveFormData", "true","")
	} else {
		v.A.SetAttr(v.Vid, "SaveFormData", "false","")
	}
	return v
}
func (v *FWebView) SupportMultipleWindows(b bool) *FWebView {
	if b {
		v.A.SetAttr(v.Vid, "SupportMultipleWindows", "true","")
	} else {
		v.A.SetAttr(v.Vid, "SupportMultipleWindows", "false","")
	}
	return v
}
func (v *FWebView) TextZoom(i int) *FWebView {
	v.A.SetAttr(v.Vid, "TextZoom",fcore. SPrintf(i),"")
	return v
}
func (v *FWebView) UserAgentString(s string) *FWebView {
	v.A.SetAttr(v.Vid, "UserAgentString", s,"")
	return v
}
func (v *FWebView) HandleUrl(m map[string]func(string) bool) *FWebView {
	if m == nil {
		return v
	}
	ms := make(map[string]string)
	for k, h := range m {
		fnId :=fcore. NewToken()
		fcore.EventMap.Set(fnId, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
			return fcore.SPrintf(h(s))
		})
		ms[k] = fnId
	}
	v.A.SetAttr(v.Vid, "HandleUrl", fcore.JsonObject(ms),"")
	return v
}
func (v *FWebView) OnDownload(f func(string)) *FWebView {
	fnId := fcore.NewToken()
	fcore.EventMap.Set(fnId, func(activity fcore.IActivity, s string, s2 string, s3 string) string {
		f(s)
		return ""
	})
	v.A.SetAttr(v.Vid, "DownloadListener", fnId,"")
	return v
}
func (v *FWebView) LoadHtmlData(s string) *FWebView {
	v.A.SetAttr(v.Vid, "LoadHtmlData", s,"")
	return v
}
