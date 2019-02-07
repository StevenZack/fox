package fox

import (
	"github.com/StevenZack/fox/fcore"
	"github.com/StevenZack/fox/fcore/widget/fbox"
	"github.com/StevenZack/fox/fcore/widget/fbutton"
	"github.com/StevenZack/fox/fcore/widget/fclipboard"
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
	c:=fclipboard.New(a)
	c.OnChange(func() {
		if t1 != nil {
			t1.Text(c.GetText())
		}
	})
	fbox.New(a).DeferShow().Append(
		ftext.New(a).Assign(&t1).Text("empty"),
		fbutton.New(a).Text("change").OnClick(func() {
			c.SetText("asd")
		}),
	)
}
