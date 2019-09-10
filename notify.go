package p

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
	"google.golang.org/api/cloudbuild/v1"
	"google.golang.org/api/pubsub/v1"
)

var (
	slackURL   string
	channel    string
	branchName string
	repoName   string

	// いらないステータスをfalseにしあとでスキップする
	useStatus = map[string]bool{
		"STATUS_UNKNOWN": false,
		"QUEUED":         false,
		"WORKING":        true,
		"SUCCESS":        true,
		"FAILURE":        true,
		"INTERNAL_ERROR": true,
		"TIMEOUT":        true,
		"CANCELLED":      true,
	}
	colors = map[string]string{
		"GOOD":    "good",
		"WARNING": "warning",
		"DANGER":  "danger",
		"BLUE":    "#439FE0",
	}
)

func init() {
	slackURL = os.Getenv("SLACK_URL")
	channel = os.Getenv("CHANNEL")
	repoName = os.Getenv("REPO_NAME")
	branchName = os.Getenv("BRANCH_NAME")
}

// NotifyGCB2Slack consumes a Pub/Sub message.
func NotifyGCB2Slack(ctx context.Context, m *pubsub.PubsubMessage) error {
	build, err := newCloudBuild(m)
	if err != nil {
		log.Println(err)
		return err
	}

	// どうでもいいrepositoryとbranchにfileterをかける
	if build.Source.RepoSource.RepoName != repoName || build.Source.RepoSource.BranchName != branchName {
		return nil
	}

	// どうでも良いステータスは無視する
	if onStatus, ok := useStatus[build.Status]; !ok || !onStatus {
		log.Printf("%s status is skipped\n", build.Status)
		return nil
	}

	msg := newWebhookMessage(build)

	err = slack.PostWebhook(slackURL, msg)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func newWebhookMessage(build *cloudbuild.Build) *slack.WebhookMessage {
	var msg slack.WebhookMessage
	attachment1 := slack.Attachment{}
	attachment1.Title = "Build Status: " + build.Status
	attachment1.TitleLink = build.LogUrl
	attachment1.Text = fmt.Sprintf(
		"project_id: *%s*\nrepository: *%s*\nbranch: *%s*",
		build.ProjectId,
		build.Source.RepoSource.RepoName,
		build.Source.RepoSource.BranchName,
	)

	switch build.Status {
	case "WORKING":
		attachment1.Color = colors["BLUE"]

	case "SUCCESS":
		attachment1.Color = colors["GOOD"]

	case "FAILURE", "TIMEOUT", "CANCELLED":
		attachment1.Color = colors["DANGER"]

	default:
		attachment1.Color = colors["WARNING"]

	}
	msg = slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment1},
		Channel:     "#" + channel,
	}
	return &msg
}

func newCloudBuild(m *pubsub.PubsubMessage) (*cloudbuild.Build, error) {
	d, err := base64.StdEncoding.DecodeString(m.Data)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	build := cloudbuild.Build{}
	err = json.Unmarshal(d, &build)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &build, nil
}
