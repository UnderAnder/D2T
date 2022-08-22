# D2T
## forward messages from selected Discrord channels to Telegram channel

### Installation
```shell script
go build -o ./D2T .
```

### Usage of D2T:
  -cfg string
        path to config file (default "config.yml")

### Config
```
discord_login:
discord_password:
telegram_apitoken:
telegram_channel:
discord_channels_white_list: # [ID, ID]
```

### Environment variables:
* D2T_DISCORD_LOGIN
* D2T_DISCORD_PASSWORD
* D2T_TELEGRAM_APITOKEN
* D2T_TELEGRAM_CHANNEL
