package conf

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server Server   `yaml:"server"`
	DB     DBConfig `yaml:"db"`
	//Etcd      EtcdConfig  `yaml:"etcd"`
	//Rpc       RpcConfig   `yaml:"rpc"`
	//Redis     RedisConfig `yaml:"redis"`
	//SearchDir string      `yaml:"searchdir"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

//type EtcdConfig struct {
//	Host     string `yaml:"host"`
//	Port     string `yaml:"port"`
//	BasePath string `yaml:"basepath"`
//}
//
//type RpcConfig struct {
//	Host string `yaml:"host"`
//	Port string `yaml:"port"`
//}
//
//type RedisConfig struct {
//	Host string `yaml:"host"`
//	Port string `yaml:"port"`
//}

var GConfig *Config

func init() {
	config := Config{
		DB: DBConfig{
			Host:     "localhost",
			Port:     "3306",
			UserName: "root",
			PassWord: "password_123456",
			DBName:   "chess",
		},
		//Etcd: EtcdConfig{
		//	Host:     "localhost",
		//	Port:     "2379",
		//	BasePath: "file_search_rpc",
		//},
		//Rpc: RpcConfig{
		//	Host: "localhost",
		//	Port: "9001",
		//},
		//Redis: RedisConfig{
		//	Host: "localhost",
		//	Port: "6379",
		//},
		//SearchDir: "E:/Work/searchFiles/",
	}
	cfgData, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(cfgData, &config)
	if err != nil {
		logrus.Error(err)
	}
	GConfig = &config
}
