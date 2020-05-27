package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// starts listening requests
func (r *Restarter) Listen(port int, endpoint string) {
	router := mux.NewRouter()
	router.HandleFunc(endpoint, r.PullAndRestart)

	err := logInfo(port, endpoint, r.secretToken, r.secretTokenKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

// first of all - pull project in folder, then restart all inside
func (r *Restarter) PullAndRestart(w http.ResponseWriter, req *http.Request) {
	log.Println("request -> " + req.RequestURI)
	// is it github webhook
	err := r.validate(req)
	if err != nil {
		log.Println("failed validate " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// git pull in folder
	err = r.pull()
	if err != nil {
		log.Println("failed pull " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// restart pulled service
	go r.restart()

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("done!"))
	if err != nil {
		log.Println("ERR - " + err.Error())
	}
}

// exec restart script inside project
func (r *Restarter) restart() {
	cmd := exec.Command(r.commands.Restart)
	cmd.Dir = r.dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println("executing command...")
	log.Println("command - " + r.commands.Restart)
	log.Println("dir - " + r.dir)
	err := cmd.Run()
	if err != nil {
		log.Println("restart err - " + err.Error())
	}
}

// run git pull
func (r *Restarter) pull() error {
	repo, err := git.PlainOpen(r.dir)
	if err != nil {
		log.Printf("open ERR - %s", err.Error())
		return fmt.Errorf("failed open dir")
	}

	// Get the working directory for the repository
	w, err := repo.Worktree()
	if err != nil {
		log.Printf("worktree ERR - %s", err.Error())
		return fmt.Errorf("failed get worktree")
	}

	err = w.Pull(
		&git.PullOptions{
			RemoteName: "origin",
			Auth: &gitHttp.BasicAuth{
				Username: r.creds.login,
				Password: r.creds.pass,
			},
		},
	)

	if err != nil {
		log.Printf("err - %s, dir - %s", err.Error(), r.dir)
	}

	return err
}

// it should be github hook with generated token inside
func (r *Restarter) validate(req *http.Request) error {
	// if not push - ignore
	event := req.Header.Get("x-github-event")
	if event != "push" {
		if event == "" || event == " " {
			return fmt.Errorf("not alowed")
		}
		return fmt.Errorf("wrong github event - %s", event)
	}

	// check is valid github delivery
	delivery := req.Header.Get("x-github-delivery")
	_, err := uuid.Parse(delivery)
	if err != nil {
		if delivery == "" || delivery == " " {
			return fmt.Errorf("not alowed")
		}
		return fmt.Errorf("wrong github delivery - %s", delivery)
	}

	// get generated token
	tokens, ok := req.URL.Query()[r.secretTokenKey]
	if !ok {
		return fmt.Errorf("secret token missing")
	}
	if tokens[0] != r.secretToken {
		return fmt.Errorf("%s secret token is invalid", tokens[0])
	}

	// get and check github hook body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v\n", err)
		return err
	}
	defer req.Body.Close()

	// Unmarshal
	var h Hook
	err = json.Unmarshal(body, &h)
	if err != nil {
		return err
	}
	// check base repository data
	if h.Repository.ID != r.repo.ID {
		return fmt.Errorf("wrong repository id - %d", h.Repository.ID)
	}
	if h.Repository.Name != r.repo.Name {
		return fmt.Errorf("wrong repository name - %s", h.Repository.Name)
	}
	if h.Repository.Owner.ID != r.repo.Owner {
		return fmt.Errorf("wrong owner id - %d", h.Repository.Owner.ID)
	}
	if h.Sender.Login != r.repo.Sender {
		return fmt.Errorf("wrong sender - %s", h.Sender.Login)
	}
	return nil
}
