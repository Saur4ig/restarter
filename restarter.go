package main

type Restarter struct {
	apps []Application
}

func newRestarter(apps []Application) *Restarter {
	return &Restarter{
		apps: apps,
	}
}
