package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fedit"
	"github.com/StevenZack/fox/fcore/widget/ftext"
)

type IActivity interface {
	fcore.IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	return fcore.TriggerFunction(a, fnId, s, s1, s2)
}

func Main(a IActivity) {
var e *fedit.FEdit
	var t1 *ftext.FText
	fbox.New(a).DeferShow().Append(
		ftext.New(a).Assign(&t1).Text("empty"),
		fedit.New(a).Assign(&e).InputTypeEnglish().Size(-2	,-1),
		fbutton.New(a).Text("change").OnClick(func() {
			t1.Text(e.GetText())
		}),
	)
}
