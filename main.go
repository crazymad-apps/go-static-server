package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type ConfigServer struct {
	Location string `json: "location"`
	Root     string `json: "root"`
}

type Config struct {
	Port    string         `json:"port"`
	Servers []ConfigServer `json:"servers"`
}

func readConfig() (Config, error) {
	content, err := os.ReadFile("./conf.json")
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(content, &config); err != nil {
		println(err.Error())
		return Config{}, err
	}

	return config, nil
}

func main() {
	config, err := readConfig()
	if err != nil {
		println("读取配置文件失败")
		println(err.Error())
		os.Exit(0)
	}

	r := gin.Default()

	for _, server := range config.Servers {
		r.Static(server.Location, server.Root)
	}

	r.Run(":" + config.Port)
}
