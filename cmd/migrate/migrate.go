package main

import "job-application/cmd"

func main() {
	repo, err := cmd.InitRepository()
	if err != nil {
		return
	}

	repo.Automigrate()
}
