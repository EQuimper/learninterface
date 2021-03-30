package mock

import "log"

type Client struct {

}

func New() *Client {
	return &Client{}
}

func (c *Client) Identify(userID int, email string) {
	// do segment stuff
	log.Printf("MOCK IDENTIFY: { userId: %d, email: %s }\n", userID, email)
}

func (c *Client) TrackEvent(event string, data map[string]interface{}) {
	log.Printf("MOCK TRACK_EVENT: { event: %s, data: %+v }\n", event, data)
}


