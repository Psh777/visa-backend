package config

var projectConfig MyConfig

func GetConfig(envState string) MyConfig {
	env := Env{}
	if envState == "prod" {
		env = Env{
			TelegramBot: "1613187411:AAGyl5sbbRZQoUHjt7mtm_-BUBQsnlwGjTk",
			BackendPort: "3010",
		}
	} else {
		env = Env{
			TelegramBot: "1613187411:AAGyl5sbbRZQoUHjt7mtm_-BUBQsnlwGjTk",
			BackendPort: "3010",
		}
	}
	return MyConfig{
		Env: env,
	}
}


func GetMyConfig() MyConfig {
	return projectConfig
}

