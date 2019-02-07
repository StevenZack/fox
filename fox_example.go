package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/widget/fbox"
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
	var t1 *ftext.FText
	var e *fedit.FEdit
	fbox.New(a).DeferShow().Append(
		ftext.New(a).Assign(&t1).Text("asd"),
		fedit.New(a).Assign(&e).Size(-2,-1).Text("a").Hint("input").OnChange(func() {
			t1.Text(e.GetText())
		}),
	)
}
