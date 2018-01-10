package buqi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config 配置文件
type Config struct {
	Local    string `json:"Local"`
	Server   string `json:"Server"`
	Current  string `json:"Current"`
	Password string `json:"Password"`
}

// SaveConfig 保存配置文件
func (config *Config) SaveConfig() {
	cfg, _ := json.MarshalIndent(config, "", "	")
	err := ioutil.WriteFile("Config.json", cfg, 0644)
	if err != nil {
		log.Println(err)
	}
}

// ReadConfig 读取配置文件
func (config *Config) ReadConfig() (exists bool) {
	if _, err := os.Stat("Config.json"); !os.IsNotExist(err) {
		file, err := os.Open("Config.json")
		defer file.Close()
		if err != nil {
			exists = false
			log.Println(err)
		} else {
			err = json.NewDecoder(file).Decode(config)
			if err != nil {
				exists = false
				log.Println(err)
			} else {
				exists = true
			}
		}
	} else {
		exists = false
	}
	return
}
