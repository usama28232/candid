package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var appConfigMap map[string]interface{}

func initConfigMap() {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&appConfigMap); err != nil {
		fmt.Println("Error decoding config file:", err)
		return
	}
}

func GetConfigByKey(key string) (interface{}, error) {
	if len(appConfigMap) == 0 {
		initConfigMap()
	}
	if len(key) == 0 {
		return "", errors.New("config key is required")
	}
	val, ok := appConfigMap[key]
	// If the key exists
	if ok {
		return val, nil
	}
	return "", fmt.Errorf("key %v is not defined in app config", key)
}
