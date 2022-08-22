package main

type Config struct {
	DiscordLogin     string  `yaml:"discord_login" env:"D2T_DISCORD_LOGIN"`
	DiscordPassword  string  `yaml:"discord_password" env:"D2T_DISCORD_PASSWORD"`
	DiscordWhiteList []int64 `yaml:"discord_channels_white_list"`
	TelegramApiToken string  `yaml:"telegram_apitoken" env:"D2T_TELEGRAM_APITOKEN"`
	TelegramChannel  int64   `yaml:"telegram_channel" env:"D2T_TELEGRAM_CHANNEL"`
}
