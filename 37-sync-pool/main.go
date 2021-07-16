package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	// github "github.com/src-d/go-github/github"
)

func main() {
	http.HandleFunc("/", handle)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

var prPool = sync.Pool{
	New: func() interface{} {
		// ref: https://stackoverflow.com/questions/34543430/golang-basics-struct-and-new-keyword/34543716
		// new allocates zeroed storage for a new item or type whatever
		// and then returns a pointer to it.
		// new is the same as short variable declaration := type{}
		// return new(github.PullRequestEvent)
		return new(pullRequestEvent)
	},
}

type pullRequest struct {
	ID *int `json:"id"`
}

type pullRequestEvent struct {
	PullRequest *pullRequest `json:"pull_request"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	// var data github.PullRequestEvent
	data := prPool.Get().(*pullRequestEvent)
	defer prPool.Put(data)
	if data.PullRequest != nil {
		*data.PullRequest.ID = 0
	}

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		logrus.Errorf("could not decode request: %v", err)
		http.Error(w, "could not decode request", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "pull request id: %d", *data.PullRequest.ID)
}
