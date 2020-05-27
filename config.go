package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const configFilePath = ""

type Commands struct {
	Start   string `yaml:"start"`
	Stop    string `yaml:"stop"`
	Restart string `yaml:"restart"`
}
type Cr struct {
	Login string `yaml:"login"`
	Pass  string `yaml:"pass"`
}

type Repo struct {
	Creds  Cr     `yaml:"creds"`
	ID     int    `yaml:"id"`
	Name   string `yaml:"name"`
	Branch string `yaml:"branch"`
	Owner  int    `yaml:"owner_id"`
	Sender string `yaml:"sender"`
}

type Conf struct {
	ApiVersion     string   `yaml:"apiVersion"`
	SecretParamKey string   `json:"secretParamKey"`
	Folder         string   `yaml:"folder"`
	Repo           Repo     `yaml:"repository"`
	Endpoint       string   `yaml:"endpoint"`
	Port           int      `yaml:"port"`
	Commands       Commands `yaml:"commands"`
}

func GetConfig(fileName string) (*Conf, error) {
	dat, err := ioutil.ReadFile(configFilePath + fileName)
	if err != nil {
		return nil, err
	}

	conf := Conf{}
	err = yaml.Unmarshal(dat, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
