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
```

## Outputs available

### Stdout

Output the state of the pomodoro in the current stdout session with a progress bar. Classic.

2. File

Output the state of the pomodoro inside a file. Handy when you use tmux or any kind of status hackable status bar.

## Outputs incoming

1. Slack

Output the state of the pomodoro inside your slack status. An option will be available to activate the disturb mode.


## License

MIT
