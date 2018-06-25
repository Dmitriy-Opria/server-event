package config

type Settings struct {
	Bind string `json:"bind"`
}

var (
	settings = &Settings{
		Bind: ":3000",
	}
)

func GetConfig() *Settings {
	return settings
}
