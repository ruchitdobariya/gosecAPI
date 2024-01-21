package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ruchitdobariya/gosecAPI/database"
)

func GetConfigsObject() []Config {
	var configs []Config

	
	path := database.GetPath() + "config.json"

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("The database could not be found. If you have executed the 'gosec' program and the database did not get created, please provide feedback on github.com/ruchitdobariya/GoSec.")
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &configs)

	return configs
}

func FindConfigByUserId(id string, configs []Config) (*Config, *database.MyError) {
	userId, err := database.ConverToFloat64(id)

	if err != nil {
		return nil, err
	}

	for _, u := range configs {
		if u.UserId == userId {
			return &u, nil
		}
	}

	
	return &Config{ConfigId: -1}, nil
}
