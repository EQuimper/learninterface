package main

import (
	"os"

	"learninterface/analytics"
	"learninterface/analytics/mock"
	"learninterface/analytics/segment"
)

func main() {
	var analyticsClient *analytics.Analytics

	if os.Getenv("APP_ENV") == "production" {
		analyticsClient = analytics.New(segment.New())
	} else {
		analyticsClient = analytics.New(mock.New())
	}

	analyticsClient.Register(1, "bob@gmail.com")

	analyticsClient.Login(2, "jon@gmail.com")
}
