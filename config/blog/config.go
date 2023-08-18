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
	RssSys int `yaml:"rss_sys"`
}

type SystemConfig struct {
	JwtKey    string `yaml:"jwt_key"`
	Db        string `yaml:"db"`
	Port      int    `yaml:"port"`
	ApiPrefix string `yaml:"api_prefix"`
	CosId     string `yaml:"cos_id"`
	CosKey    string `yaml:"cos_key"`
	CosAppid  string `yaml:"cos_appid"`
	CosBucket string `yaml:"cos_bucket"`
	CosRegion string `yaml:"cos_region"`
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
