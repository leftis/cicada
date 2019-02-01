package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var App app

type app struct {
	CurrentDirectory string
	Config
}

type Config struct {
	Environment string
	Secret string
	Database
}

type Database struct {
	SSL    string `yaml:"ssl_mode"'`
	Host   string
	Port   string
	User   string
	Pass   string
	Name   string
	Driver string
}

func Init() {
	App.setCurrentDirectory()
	App.setConfiguration()
}

func (a *app) setConfiguration() {
	unmarshalErr := yaml.Unmarshal(openConfigurationFile(a.CurrentDirectory), &a.Config)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
}

func (a *app) setCurrentDirectory() {
	var err error
	a.CurrentDirectory, err = os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
}

func openConfigurationFile(currentDirectory string) (bytes []byte) {
	b, err := ioutil.ReadFile(currentDirectory + "/config/global.yml")

	if err != nil {
		log.Fatal(err)
	}

	return b
}