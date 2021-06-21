package main

import (
	"time"

	"github.com/lrstanley/girc"
)

var (
	// updateFrequency is the number of times in a timer
	// to update the user. An updateFrequency of 10 means every 10%.
	// An updateFrequency of 4 means every 25%
	updateFrequency = 4
)

func RunTimer(c *girc.Client, user, channel, msg string, d time.Duration) {
	interval := d / time.Duration(updateFrequency)

	for step := 1; step <= updateFrequency; step++ {
		time.Sleep(interval)

		// No need to tell us it's at 100%; that's what the next thing is
		if step < updateFrequency {
			c.Cmd.Messagef(user, "timebox is %d%% complete", step*(100/updateFrequency))
		}
	}

	c.Cmd.Messagef(channel, "%s: %s", user, msg)

	if channel != user {
		c.Cmd.Message(user, msg)
	}
}
