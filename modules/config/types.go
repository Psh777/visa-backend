package config

type MyConfig struct {
	Env        Env
}

type Env struct {
	BackendPort                     string `json:"backend_port"`
	TelegramBot                     string `json:"telegram_bot"`
}
