module github.com/UnderAnder/D2T

go 1.19

require (
	github.com/diamondburned/arikawa/v3 v3.1.0 // direct
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1 // direct
)

replace github.com/diamondburned/arikawa/v3 => github.com/UnderAnder/arikawa/v3 v3.1.0

require github.com/ilyakaznacheev/cleanenv v1.3.0

require (
	github.com/BurntSushi/toml v1.2.0 // indirect
	github.com/gorilla/schema v1.2.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
