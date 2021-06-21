# Timebox bot

Given a duration ("1h", "2m", "8h10m3s" etc. - see https://golang.org/pkg/time/#ParseDuration) and an optional message ("Stop! Times up!", "And.... relax", etc.) create a timebox.

That is to say, create a timer which counts down the minutes (sending a message to the requestor at intervals) prior to sending the final message.

This is useful for doing things like setting a so-called 'productivity timer' (like in [pomodoro](https://en.wikipedia.org/wiki/Pomodoro_Technique)), or by ensuring you don't spend too long working on a specific tasks (like a [timebox](https://en.wikipedia.org/wiki/Timeboxing)).

This bot makes a couple of assumptions:

1. You've a SASL account for this bot to use
2. You've enabled actions notifications in github for failed/successful runs

This bot requires the following env vars:

* `$SASL_USER` - the user to connect with
* `$SASL_PASSWORD` - the password to connect with
* `$SERVER` - IRC connection details, as `irc://server:6667` or `ircs://server:6697` (`ircs` implies irc-over-tls)
* `$VERIFY_TLS` - Verify TLS, or sack it off. This is of interest to people, like me, running an ircd on localhost with a self-signed cert. Matches "true" as true, and anything else as false

The SASL mechanism is hardcoded to PLAIN.

## Building

This bot can be built using pretty standard go tools:

```bash
$ go build
```

Or via docker:

```bash
$ docker build -t foo .
```

## Running

If you've built the app yourself, then happy day- there's your binary!

Otherwise I suggest via docker:

```bash
$ docker build -t foo .
$ docker run foo
```

(Setting the above environment variables accordingly)
