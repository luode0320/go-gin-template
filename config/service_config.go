package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const (
	configFile = "config.yaml"
)

// ServiceConfig 服务端全局配置
var ServiceConfig ServiceCfg

// ServiceCfg 总体结构配置
type ServiceCfg struct {
	Version string    `yaml:"version"`
	Web     WebCfg    `yaml:"web"`
	Db      DbCfg     `yaml:"db"`
	Logger  LoggerCfg `yaml:"logger"`
	Redis   RedisCfg  `yaml:"redis"`
	Kafka   KafkaCfg  `yaml:"kafka"`
	Oss     MinioCfg  `yaml:"oss"`
}

type WebCfg struct {
	Port  string `yaml:"port"`
	Mode  string `yaml:"mode"`
	Token bool   `yaml:"token"`
}

type DbCfg struct {
	Driver       string `yaml:"driver"`
	Url          string `yaml:"url"`
	UserName     string `yaml:"userName"`
	Password     string `yaml:"password"`
	DbName       string `yaml:"dbname"`
	Port         string `yaml:"port"`
	LogLevel     int    `yaml:"logLevel"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

type LoggerCfg struct {
	LogLevel       int    `yaml:"logLevel"`
	FileName       string `yaml:"fileName"`
	MaxSize        int    `yaml:"maxSize"`
	MaxAge         int    `yaml:"maxAge"`
	MaxBackups     int    `yaml:"maxBackups"`
	Compress       bool   `yaml:"compress"`
	DisableConsole bool   `yaml:"disableConsole"`
}

// RedisCfg redis配置
type RedisCfg struct {
	Ip       string `yaml:"ip"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type KafkaCfg struct {
	Topic   string `yaml:"topic"`
	Ip      string `yaml:"ip"`
	Port    string `yaml:"port"`
	GroupID string `yaml:"groupID"`
}

type MinioCfg struct {
	Endpoint   string `yaml:"endpoint"`
	AccessKey  string `yaml:"accessKey"`
	SecretKey  string `yaml:"secretKey"`
	BucketName string `yaml:"bucketName"`
	FileType   string `yaml:"fileType"`
}

// InitConfig 解析通用配置信息
func InitConfig() {
	var cfg ServiceCfg
	yamlFile, err := ioutil.ReadFile(configFile)
	//若出现错误，打印错误提示
	if err != nil {
		log.Panicf("解析yaml配置异常 -> [%s]", err.Error())
	}

	//将读取的字符串转换成结构体conf
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Panicf("读取yaml的字符串转换成结构体异常 -> [%s]", err.Error())
	}

	ServiceConfig = cfg
}
