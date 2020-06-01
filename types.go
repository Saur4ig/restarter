package main

type Application struct {
	token          string
	folder         string
	endpoint       string
	secretParamKey string
	repo           ReposotoryData
	commands       Commands
}

type App struct {
	Folder         string         `yaml:"folder"`
	Endpoint       string         `yaml:"endpoint"`
	SecretParamKey string         `yaml:"secretParamKey"`
	Repo           ReposotoryData `yaml:"repository"`
	Commands       Commands       `yaml:"commands"`
}

type Commands struct {
	Start   string `yaml:"start"`
	Stop    string `yaml:"stop"`
	Restart string `yaml:"restart"`
}
type Credentials struct {
	Login string `yaml:"login"`
	Pass  string `yaml:"pass"`
}

type Spec struct {
	Apps []App `yaml:"apps"`
}

type ReposotoryData struct {
	Creds  Credentials `yaml:"creds"`
	ID     int         `yaml:"id"`
	Name   string      `yaml:"name"`
	Branch string      `yaml:"branch"`
	Owner  int         `yaml:"owner_id"`
	Sender string      `yaml:"sender"`
}

type Conf struct {
	ApiVersion string `yaml:"apiVersion"`
	Spec       Spec   `yaml:"spec"`
	Port       int    `yaml:"port"`
}
