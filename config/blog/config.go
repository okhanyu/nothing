package blog

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Global *Config

type Config struct {
	System   SystemConfig   `yaml:"system"`
	Business BusinessConfig `yaml:"business"`
}

type BusinessConfig struct {
	RowNum int `yaml:"row_num"`
}

type SystemConfig struct {
	JwtKey    string `yaml:"jwt_key"`
	Db        string `yaml:"db"`
	Port      int    `yaml:"port"`
	ApiPrefix string `yaml:"api_prefix"`
}

func ParseConfig(configFile string) *Config {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Printf("无法读取文件: %v", err)
	}
	err = yaml.Unmarshal(data, &Global)
	if err != nil {
		log.Printf("无法解析YAML数据: %v", err)
	}
	return Global
}
