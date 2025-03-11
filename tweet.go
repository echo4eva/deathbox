package main

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

func sendTweet(client *gotwi.Client, text string) {
	p := &types.CreateInput{
		Text: gotwi.String(text),
	}

	_, err := managetweet.Create(context.Background(), client, p)
	if err != nil {
		fmt.Println(err)
		logger.Error("tweet failed")
	}

	logger.Info("tweet sent")
}
