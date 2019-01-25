package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var config map[string]map[string]string

func Config() map[string]map[string]string {
	err :=  yaml.Unmarshal(openConfigurationFile(), &config)

	if err != nil {
		log.Fatal(err)
	}

	return config
}

func openConfigurationFile() (bytes []byte) {
	pwd, _ := os.Getwd()
	b, err := ioutil.ReadFile(pwd+"/configuration.yml")

	if err != nil {
		log.Fatal(err)
	}

	return b
}