package timeholder

import (
	"time"
)

var DefaultHolder = NewHolder()

type Holder struct {
	Pool     map[string]time.Time
	Duration time.Duration
}

func (this *Holder) SetDuration(d time.Duration) {
	this.Duration = d
}

func (this *Holder) Hold(key string) bool {
	oldTime, ok := this.Pool[key]
	if ok {
		if time.Since(oldTime) < this.Duration {
			return true
		}
	}
	this.Pool[key] = time.Now()
	return false
}

func NewHolder() (r *Holder) {
	r = &Holder{
		Pool:     make(map[string]time.Time),
		Duration: 1 * time.Minute,
	}
	return
}

func SetDuration(d time.Duration) {
	DefaultHolder.SetDuration(d)
}

func Hold(key string) bool {
	return DefaultHolder.Hold(key)
}
