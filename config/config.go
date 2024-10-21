package config

import (
	"github.com/spf13/viper"
	"fmt"
)


type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	Db string `json:"db"`
}

func GetConfig() Config {
	var Cfg Config 

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".") 
	
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var e  = viper.Unmarshal(&Cfg)
	if e != nil {
		fmt.Println("unable to decode into struct, %v", err)
	}
	viper.WatchConfig()
	return Cfg
}