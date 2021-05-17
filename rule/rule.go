package rule

import "time"

type Rule struct{}

// TimedTrade --
func (r *Rule) TimedTrade(fn func(), at time.Time) *time.Timer {
	current := time.Now()
	duration := at.Sub(current)
	execute := time.NewTimer(duration)
	time.AfterFunc(duration, fn)
	return execute
}
