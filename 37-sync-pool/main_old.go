package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	github "github.com/src-d/go-github/github"
)

func main_old() {
	http.HandleFunc("/", handle)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

func handle_old(w http.ResponseWriter, r *http.Request) {
	var data github.PullRequestEvent

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		logrus.Errorf("could not decode request: %v", err)
		http.Error(w, "could not decode request", http.StatusInternalServerError)
		return
	}
	// wingleung: data.PullRequest.ID of github.PullRequestEvent is an pointer
	fmt.Fprintf(w, "pull request id: %d", *data.PullRequest.ID)
}

// type PullRequestEvent struct {
// 	// Action is the action that was performed. Possible values are:
// 	// "assigned", "unassigned", "review_requested", "review_request_removed", "labeled", "unlabeled",
// 	// "opened", "edited", "closed", "ready_for_review", "locked", "unlocked", or "reopened".
// 	// If the action is "closed" and the "merged" key is "false", the pull request was closed with unmerged commits.
// 	// If the action is "closed" and the "merged" key is "true", the pull request was merged.
// 	// While webhooks are also triggered when a pull request is synchronized, Events API timelines
// 	// don't include pull request events with the "synchronize" action.
// 	Action      *string      `json:"action,omitempty"`
// 	Assignee    *User        `json:"assignee,omitempty"`
// 	Number      *int         `json:"number,omitempty"`
// 	PullRequest *PullRequest `json:"pull_request,omitempty"`

// 	// The following fields are only populated by Webhook events.
// 	Changes *EditChange `json:"changes,omitempty"`
// 	// RequestedReviewer is populated in "review_requested", "review_request_removed" event deliveries.
// 	// A request affecting multiple reviewers at once is split into multiple
// 	// such event deliveries, each with a single, different RequestedReviewer.
// 	RequestedReviewer *User `json:"requested_reviewer,omitempty"`
// 	// In the event that a team is requested instead of a user, "requested_team" gets sent in place of
// 	// "requested_user" with the same delivery behavior.
// 	RequestedTeam *Team         `json:"requested_team,omitempty"`
// 	Repo          *Repository   `json:"repository,omitempty"`
// 	Sender        *User         `json:"sender,omitempty"`
// 	Installation  *Installation `json:"installation,omitempty"`
// 	Label         *Label        `json:"label,omitempty"` // Populated in "labeled" event deliveries.

// 	// The following field is only present when the webhook is triggered on
// 	// a repository belonging to an organization.
// 	Organization *Organization `json:"organization,omitempty"`

// 	// The following fields are only populated when the Action is "synchronize".
// 	Before *string `json:"before,omitempty"`
// 	After  *string `json:"after,omitempty"`
// }
