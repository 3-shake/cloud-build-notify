package notify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/nlopes/slack"
	"google.golang.org/api/cloudbuild/v1"
	"google.golang.org/api/pubsub/v1"
)

var (
	cloudBuildStatus = map[string]bool{
		"STATUS_UNKNOWN": false,
		"QUEUED":         false,
		"WORKING":        false,
		"SUCCESS":        true,
		"FAILURE":        true,
		"INTERNAL_ERROR": true,
		"TIMEOUT":        true,
		"CANCELLED":      false,
	}
)

func init() {
	_, err := NewSlack()
	if err != nil {
		log.Fatal(err)
	}
}

type Slack struct {
	Token     string
	ChannelID string
}

func NotifyCloudBuild(ctx context.Context, m *pubsub.PubsubMessage) error {
	build, err := NewCloudBuild(m)
	if err != nil {
		log.Println(err)
		return err
	}
	if !cloudBuildStatus[build.Status] {
		return nil
	}

	cli, err := NewSlack()
	if err != nil {
		log.Println(err)
		return err
	}

	err = cli.Notify(build.Status)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func NewCloudBuild(m *pubsub.PubsubMessage) (*cloudbuild.Build, error) {
	decoded, err := base64.StdEncoding.DecodeString(m.Data)
	if err != nil {
		return nil, err
	}
	build := &cloudbuild.Build{}
	err = json.Unmarshal(decoded, build)
	if err != nil {
		return nil, err
	}

	return build, nil
}

func NewSlack() (*Slack, error) {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		return nil, errors.New("Required SLACK_TOKEN")
	}

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		return nil, errors.New("Required SLACK_CHANNEL_ID")
	}

	return &Slack{
		Token:     token,
		ChannelID: channelID,
	}, nil
}

func (sl *Slack) Notify(msg string) error {
	cli := slack.New(sl.Token)
	attachment := slack.Attachment{
		Pretext: msg,
	}

	_, _, err := cli.PostMessage(sl.ChannelID, slack.MsgOptionText("[Deploy Status]", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		return err
	}
	return nil
}
