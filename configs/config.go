package configs

import (
	"fmt"
	"os"

	"github.com/YashinaAnn/todo-service/internal/models"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

func ReadConfig() models.Config {
	var cfg models.Config
	readFile(&cfg)
	readEnv(&cfg)
	return cfg
}

func readFile(cfg *models.Config) {
	fmt.Println(os.Getwd())
	f, err := os.Open("../configs/config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *models.Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
