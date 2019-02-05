package fcore

import "sync"

var (
	activityIdMap = newFactivityIdMap()
	EventMap      = NewFEventMap()
	ViewMap       = NewFViewMap()
)

// eventMap type
type FEventMap struct {
	Data map[string]func(IActivity, string, string, string) string
	Lock sync.Mutex
}

func NewFEventMap() *FEventMap {
	return &FEventMap{Data: make(map[string]func(IActivity, string, string, string) string)}
}
func (w *FEventMap) Set(k string, v func(IActivity, string, string, string) string) {
	w.Lock.Lock()
	w.Data[k] = v
	w.Lock.Unlock()
}
func (w *FEventMap) Get(k string) func(IActivity, string, string, string) string {
	return w.Data[k]
}
func (w *FEventMap) Exists(k string) bool {
	_, ok := w.Data[k]
	return ok
}
func (w *FEventMap) Remove(k string) {
	w.Lock.Lock()
	delete(w.Data, k)
	w.Lock.Unlock()
}

// activityIdMap type
type factivityIdMap struct {
	Data map[string]*FActivity
	Lock sync.Mutex
}

func newFactivityIdMap() *factivityIdMap {
	return &factivityIdMap{Data: make(map[string]*FActivity)}
}
func (w *factivityIdMap) Set(k string, v *FActivity) {
	w.Lock.Lock()
	w.Data[k] = v
	w.Lock.Unlock()
}
func (w *factivityIdMap) Get(k string) *FActivity {
	return w.Data[k]
}
func (w *factivityIdMap) Exists(k string) bool {
	_, ok := w.Data[k]
	return ok
}
func (w *factivityIdMap) Remove(k string) {
	w.Lock.Lock()
	delete(w.Data, k)
	w.Lock.Unlock()
}

// viewMap
type FViewMap struct {
	Data map[string]IBaseView
	Lock sync.Mutex
}

func NewFViewMap() *FViewMap {
	return &FViewMap{Data: make(map[string]IBaseView)}
}
func (w *FViewMap) Set(k string, v IBaseView) {
	w.Lock.Lock()
	w.Data[k] = v
	w.Lock.Unlock()
}
func (w *FViewMap) Get(k string) IBaseView {
	return w.Data[k]
}
func (w *FViewMap) Exists(k string) bool {
	_, ok := w.Data[k]
	return ok
}
func (w *FViewMap) Remove(k string) {
	w.Lock.Lock()
	delete(w.Data, k)
	w.Lock.Unlock()
}
