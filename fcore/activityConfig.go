package fcore

type FAppConf struct {
	FActivityId string
	FLaunchMode string
	FIntent     Intent
}

func NewAppConf() *FAppConf {
	ac := &FAppConf{}
	ac.FActivityId = NewToken()
	ac.FIntent.Extras = make(map[string]string)
	return ac
}
func (c *FAppConf) LaunchMode_Standard() *FAppConf {
	c.FLaunchMode = "Standard"
	return c
}
func (c *FAppConf) LaunchMode_SingleTask() *FAppConf {
	c.FLaunchMode = "SingleTask"
	return c
}
func (c *FAppConf) LaunchMode_SingleInstance() *FAppConf {
	c.FLaunchMode = "SingleInstance"
	return c
}
func (c *FAppConf) LaunchMode_SingleTop() *FAppConf {
	c.FLaunchMode = "SingleTop"
	return c
}
func (c *FAppConf) IntentAction(s string) *FAppConf {
	c.FIntent.Action = s
	return c
}
func (f *FAppConf) ActionSend() *FAppConf {
	f.FIntent.Action = "android.intent.action.SEND"
	return f
}
func (f *FAppConf) ActionView() *FAppConf {
	f.FIntent.Action = "android.intent.action.VIEW"
	return f
}

func (f *FAppConf) ActionMultiSend() *FAppConf {
	f.FIntent.Action = "android.intent.action.SEND_MULTIPLE"
	return f
}
func (c *FAppConf) Paths(ps ...string) *FAppConf {
	c.FIntent.Paths = ps
	return c
}
func (c *FAppConf) PutExtra(key, value string) *FAppConf {
	c.FIntent.Extras[key] = value
	return c
}
