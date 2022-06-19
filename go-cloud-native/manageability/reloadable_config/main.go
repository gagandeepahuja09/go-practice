package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host string
	Port string
	Tags map[string]string
}

var config Config

func init() {
	updates, errors, err := watchConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	go startListening(updates, errors)
}

func loadConfiguration(filepath string) (Config, error) {
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}

	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func startListening(updates <-chan string, errors <-chan error) {
	for {
		select {
		case filepath := <-updates:
			c, err := loadConfiguration(filepath)
			if err != nil {
				log.Println("error loading config:", err)
				continue
			}
			config = c
		case err := <-errors:
			log.Println("error watching config:", err)
		}
	}
}

func main() {

}
