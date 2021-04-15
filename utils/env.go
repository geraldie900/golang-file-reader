/*
=====================================
;	Author : Geraldie Tanu Saputra
;	Email  : geraldie.saputra@soluix.ai
;	Date   : 15-04-2021
=====================================
*/

package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	FOLDER_PATH = getEnv("FOLDER_PATH")
	PORT        = getEnv("PORT")
)

func getEnv(nameEnv string) string {
	if exists, value := viperGetEnv(nameEnv); exists {
		log.Println(fmt.Sprintf("[SETTING_ENVIRONMENT_VARIABLES] FOUND %s : %v", nameEnv, value))
		return value
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, nameEnv))
}

func viperGetEnv(env string) (bool, string) {
	viperSettings()
	value := ""
	empty := false
	if empty = viper.IsSet(env); empty {
		value = fmt.Sprintf("%v", viper.Get(env))
		return empty, value
	}
	return empty, ""
}

func viperSettings() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if err, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("error: ", err)
		} else {
			fmt.Println("another error: ", err)
		}
	}
}
