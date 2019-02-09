package flist

import (
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fcheck"
	"github.com/StevenZack/fox/fcore/widget/fconstraintbox"
	"github.com/StevenZack/fox/fcore/widget/fframebox"
	"github.com/StevenZack/fox/fcore/widget/ftext"

	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/widget/fimage"
)

// ViewHolder
type ViewHolder struct {
	VID   string
	Vlist map[string]string
}

type TypeOnBindDataArgsBundle struct {
	Str      string
	Position int
}

func SetItemId(i fcore.IBaseView, l *FList, iid string) fcore.IBaseView {
	if l.vh.Vlist == nil {
		l.vh.Vlist = make(map[string]string)
	}
	l.vh.Vlist[iid] = i.GetBaseView().Vid
	return i
}
func (vh *ViewHolder) GetListByItemId(iid string) *FList {
	if v, ok := vh.Vlist[iid]; ok {
		if bt, ok := fcore.ViewMap.Get(v).(*FList); ok {
			return bt
		}
	}
	return nil
}
func (vh *ViewHolder) GetButtonByItemId(iid string) *fbutton.FButton {
	if v, ok := vh.Vlist[iid]; ok {
		if bt, ok := fcore.ViewMap.Get(v).(*fbutton.FButton); ok {
			return bt
		}
	}
	return nil
}

func (vh *ViewHolder) GetImageByItemId(iid string) *fimage.FImage {
	if v, ok := vh.Vlist[iid]; ok {
		if bt, ok := fcore.ViewMap.Get(v).(*fimage.FImage); ok {
			return bt
		}
	}
	return nil
}

func (vh *ViewHolder) GetBoxByItemId(iid string) *fbox.FBox {
	if v, ok := vh.Vlist[iid]; ok {
		if bt, ok := fcore.ViewMap.Get(v).(*fbox.FBox); ok {
			return bt
		}
	}
	return nil
}

func (vh *ViewHolder) GetConstraintBoxByItemId(iid string) *fconstraintbox.FConstraintBox {
	if v, ok := vh.Vlist[iid]; ok {
		if bt, ok := fcore.ViewMap.Get(v).(*fconstraintbox.FConstraintBox); ok {
			return bt
		}
	}
	return nil
}

func (vh *ViewHolder) GetFrameBoxByItemId(iid string) *fframebox.FFrameBox {
	if v, ok := vh.Vlist[iid]; ok {
		if bt, ok := fcore.ViewMap.Get(v).(*fframebox.FFrameBox); ok {
			return bt
		}
	}
	return nil
}

func (vh *ViewHolder) GetTextByItemId(iid string) *ftext.FText {
	if v, ok := vh.Vlist[iid]; ok {
		if bt,ok:=fcore.ViewMap.Get(v).(*ftext.FText);ok {
			return bt
		}
	}
	return nil
}

func (vh *ViewHolder) GetCheckBoxByItemId(s string) *fcheck.FCheck {
	if v, ok := vh.Vlist[s]; ok {
		if bt,ok:=fcore.ViewMap.Get(v).(*fcheck.FCheck);ok {
			return bt
		}
	}
	return nil
}


