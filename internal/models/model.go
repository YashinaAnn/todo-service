package models

type Todo struct {
	ID        int64  `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

type Config struct {
	Server struct {
		Url string `yaml:"url", envconfig:"SERVER_URL"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"username", envconfig:"DB_USERNAME"`
		Password string `yaml:"password", envconfig:"DB_PASSWORD"`
		Name     string `yaml:"name", envconfig:"DB_NAME"`
		Url      string `yaml:"url", envconfig:"DB_URL"`
	} `yaml:"database"`
}
