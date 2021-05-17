package rule_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stevemurr/cbot/rule"
)

func TestTimedTrade(t *testing.T) {
	r := rule.Rule{}
	at := time.Now().Add(time.Second * 2)
	r.TimedTrade(func() {
		fmt.Println(time.Now())
	}, at)
	time.Sleep(time.Second * 3)
}
