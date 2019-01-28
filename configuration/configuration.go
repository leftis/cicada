package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type App struct {
	CurrentDirectory string
	Config
}

type Config struct {
	Environment string
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

func Init() (a App) {
	a.setCurrentDirectory()
	a.setConfiguration()
	return a
}

func (a *App) setConfiguration() {
	unmarshalErr := yaml.Unmarshal(openConfigurationFile(a.CurrentDirectory), &a.Config)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
}

func (a *App) setCurrentDirectory() {
	var err error
	a.CurrentDirectory, err = os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
}

func openConfigurationFile(currentDirectory string) (bytes []byte) {
	b, err := ioutil.ReadFile(currentDirectory + "/configuration/global.yml")

	if err != nil {
		log.Fatal(err)
	}

	return b
}
