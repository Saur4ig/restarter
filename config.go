package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const configFilePath = ""

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

func (c *Conf) genApplications() ([]Application, error) {
	apps := make([]Application, len(c.Spec.Apps))

	for i, val := range c.Spec.Apps {
		// generate security token
		token, err := tokenGenerator()
		if err != nil {
			return nil, err
		}

		app := Application{
			token:          token,
			folder:         val.Folder,
			endpoint:       val.Endpoint,
			secretParamKey: val.SecretParamKey,
			repo:           val.Repo,
			commands:       val.Commands,
		}
		apps[i] = app
	}
	return apps, nil
}
