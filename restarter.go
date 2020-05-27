package main

type Creds struct {
	login string
	pass  string
}

type Restarter struct {
	secretToken    string
	secretTokenKey string
	dir            string
	commands       Commands
	repo           Repo
	creds          Creds
}

func newRestarter(
	token, dir, secName string,
	commands Commands,
	repo Repo,
	login, pass string,
) *Restarter {
	return &Restarter{
		secretToken:    token,
		dir:            dir,
		commands:       commands,
		repo:           repo,
		secretTokenKey: secName,
		creds: Creds{
			login: login,
			pass:  pass,
		},
	}
}
