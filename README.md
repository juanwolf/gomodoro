# Gomodoro

Gomodoro is an integrated Timer. It will help you to keep your productivity to the top by disabling most of the distracting thing ever.

## How to install it

```
go get -u github.com/juanwolf/gomodoro
```

## Config

Your configuration is read in ~/.gomodoro.toml /etc/gomodoro/config.toml

```
pomodoro_duration = "25m"
break_duration = "5m"
refresh_rate = "1s"

[outputs]

  [outputs.stdout]
  show_percent = false
  size = 80 # Size of the progress bar in characters
  activated = true

  [outputs.file]
  activated = true
  path = "/home/juanwolf/.gomodoro"

  [outputs.slack]
  activated = false
  token = ""
  do_not_disturb = false
  emoji = ":tomato:"
```

## Outputs available

### Stdout

Output the state of the pomodoro in the current stdout session with a progress bar. Classic.

### File

Output the state of the pomodoro inside a file. Handy when you use tmux or any kind of hackable status bar.

### Slack

Output the state of the pomodoro inside your slack status. An option is available to activate the do not disturb mode.
You'll need a token with those permissions:

```
users.profile:write
dnd:write
```

You can create an app to get a token [here](https://api.slack.com/apps?new_app=1).

## License

MIT
