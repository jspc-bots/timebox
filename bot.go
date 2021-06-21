package main

import (
	"strings"
	"time"

	"github.com/ergochat/irc-go/ircfmt"
	"github.com/jspc-bots/bottom"
	"github.com/lrstanley/girc"
)

var (
	DefaultDuration = "25m"
	DefaultMessage  = "Time to relax!"
)

const (
	HelpText = `$b$c[light blue]Timeboxer Help$r
  $brun timer                                    - Run a timer which sends "Time to relax!" after 25 minutes
  $brun timer for 10m                            - Run a timer which sends "Time to relax!" after 10 minutes
                                                 The 10m can be pretty much anything, like 5m, or 30s or 12h30m14s
                                                 Note the lack of spaces. See: https://golang.org/pkg/time/#ParseDuration
  $brun timer for 10m and say "Wake up!"         - Run a timer which sends "Wake up!" after 10 minutes.
                                                 The message can be pretty much anything, so long as it's between two
                                                 quotation marks.
`
)

type Bot struct {
	bottom bottom.Bottom
}

func New(user, password, server string, verify bool) (b Bot, err error) {
	b.bottom, err = bottom.New(user, password, server, verify)
	if err != nil {
		return
	}

	b.bottom.Client.Handlers.Add(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
		c.Cmd.Join(Chan)
	})

	b.bottom.ErrorFunc = func(ctx bottom.Context, err error) {
		b.bottom.Client.Cmd.Message(ctx["sender"].(string), err.Error())

		b.help(ctx["sender"].(string), "", []string{})
	}

	router := bottom.NewRouter()
	router.AddRoute(`(?i)^help$`, b.help)
	router.AddRoute(`(?i)^run\s+timer$`, b.timer)
	router.AddRoute(`(?i)^run\s+timer\s+for\s+(\S*)$`, b.timer)
	router.AddRoute(`(?i)^run\s+timer\s+for\s+(\S*)\s+and\s+say\s+\"(.+)\"$`, b.timer)

	b.bottom.Middlewares.Push(router)

	return
}

func (b Bot) help(sender, _ string, _ []string) (err error) {
	for _, line := range strings.Split(HelpText, "\n") {
		b.bottom.Client.Cmd.Message(sender, ircfmt.Unescape(line))
	}

	return
}

func (b Bot) timer(sender, channel string, groups []string) (err error) {
	b.bottom.Client.Cmd.Messagef(channel, "Creating timer for %s", sender)

	msg := DefaultMessage
	duration := DefaultDuration

	if len(groups) >= 2 {
		duration = groups[1]

		if len(groups) == 3 {
			msg = groups[2]
		}
	}

	d, err := time.ParseDuration(duration)
	if err != nil {
		return
	}

	go RunTimer(b.bottom.Client, sender, channel, msg, d)

	return
}
