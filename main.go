package main

func main() {
	// gets config
	conf, err := GetConfig("config.yml")
	if err != nil {
		panic(err)
	}

	// get applications with token
	applications, err := conf.genApplications()
	if err != nil {
		panic(err)
	}

	// starts all logic
	restarter := newRestarter(applications)
	restarter.Listen(conf.Port)
}
