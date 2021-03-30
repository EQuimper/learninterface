package main

import (
	"os"

	"learninterface/analytics"
	"learninterface/analytics/mock"
	"learninterface/analytics/segments"
)

func main() {
	var analyticsClient *analytics.Analytics

	if os.Getenv("APP_ENV") == "production" {
		analyticsClient = analytics.New(segments.New())
	} else {
		analyticsClient = analytics.New(mock.New())
	}

	analyticsClient.Register(1, "bob@gmail.com")

	analyticsClient.Login(2, "jon@gmail.com")
}
