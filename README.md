# D2T
## forward messages from selected Discrord channels to Telegram channel

### Installation

Build the bot binary using the make target `build`:
```shell script
go build -o ./D2T .
```

### Usage of D2T:
  -cfg string
        path to config file (default "config.yml")

### Environment variables:
  D2T_DISCORD_LOGIN
  D2T_DISCORD_PASSWORD
  D2T_TELEGRAM_APITOKEN
  D2T_TELEGRAM_CHANNEL
