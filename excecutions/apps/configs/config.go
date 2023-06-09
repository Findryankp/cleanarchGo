package configs

import (
	"fmt"

	"github.com/Findryankp/cleanarchGo/excecutions/generates"
)

func configCreate() {
	file, err := generates.FilesCreate("./apps/configs/config.go")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		generates.FilesAddContent(file, configContent())
		fmt.Println("Config File Created")
	}
}

func configContent() string {
	var text = `package configs
	
type AppConfig struct {
	DBUSERNAME string
	DBPASS     string
	DBHOST     string
	DBPORT     string
	DBNAME     string
	JWTKEY     string
}

func InitConfig() *AppConfig {
	return InitEnv()
}
	`

	return text
}
