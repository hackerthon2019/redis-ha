package setting

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	App struct {
		Port string
	}
	Redis struct {
		CliPath string
		Network string
		Addr    string
		Port    string
		Expire  int
	}
}

var Config config

func init() {
	Config.App.Port = "8888"
	Config.Redis.CliPath = "redis-cli"
	Config.Redis.Network = "tcp"
	Config.Redis.Addr = "127.0.0.1"
	Config.Redis.Port = "6379"
	Config.Redis.Expire = 3600

	f, err := os.Open("config.yaml")
	if err != nil {
		log.Println("failed to read config file, will use default setting:", err.Error())
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("failed to read config file, will use default setting:", err.Error())
		return
	}

	err = yaml.Unmarshal(bs, &Config)
	if err != nil {
		log.Println("failed to read config file, will use default setting:", err.Error())
		return
	}
}
