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
	Dbname string `json:"db_name"`
}

func GetConfig() Config {
	var C Config 

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".") 
	
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.WatchConfig()
	var e  = viper.Unmarshal(&C)
	if e != nil {
		fmt.Println("unable to decode into struct, %v", err)
	}

	return C
}