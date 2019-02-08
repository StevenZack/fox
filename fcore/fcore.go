package fcore

type IActivity interface {
	GetActivityId() string
	NewObject(string, string) string
	SetAttr(string, string, string, string) string
	GetAttr(string, string) string
	Show(string) string
}
type FBase struct {
	Vid, VType string
	A                            IActivity
}

func TriggerFunction(a IActivity, fnId, s, s1, s2 string) string {
	switch fnId {
	case "StartActivity":
		startActivityTrigger(a, s, s1, s2)
		return ""
	}
	return EventMap.Get(fnId)(a, s, s1, s2)
}

type FPage struct {
	VID string
}
type FTab struct {
	Text, Icon string
}

func NewTab(text string) *FTab {
	t := &FTab{}
	t.Text = text
	return t
}
func (t *FTab) IconSrc(s string) *FTab {
	t.Icon = s
	return t
}