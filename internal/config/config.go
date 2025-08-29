package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct{
	Addr string
}


type Config struct{
	Env string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`

}




func MustLoad() *Config{
	var configPath string
	configPath=os.Getenv("CONFIG_PATH")

	if configPath==""{
		flags:=flag.String("config","","path to the config file")
		flag.Parse()
		configPath=*flags

		if configPath==""{
			log.Fatal("Config Path is not defined")
		}
	}

	if _,err:=os.Stat(configPath);os.IsNotExist(err){
		log.Fatal("Internal error in config")
	}


	var cfg Config
err:=	cleanenv.ReadConfig(configPath,&cfg)
	if err!=nil{
		log.Fatal("Cnnot read config file")
	}

	return &cfg
}