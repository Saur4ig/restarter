package main

func main() {
	// gets config
	conf, err := GetConfig("config.yml")
	if err != nil {
		panic(err)
	}

	// generate security token
	token, err := tokenGenerator()
	if err != nil {
		panic(err)
	}

	restarter := newRestarter(
		token,
		conf.Folder,
		conf.SecretParamKey,
		conf.Commands,
		conf.Repo,
		conf.Repo.Creds.Login,
		conf.Repo.Creds.Pass,
	)
	restarter.Listen(conf.Port, conf.Endpoint)
}
