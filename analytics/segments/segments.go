package segments

import "log"

type Client struct {

}

func New() *Client {
	return &Client{}
}

func (c *Client) Identify(userID int, email string) {
	// do segment stuff
	log.Printf("SEGMENT IDENTIFY: { userId: %d, email: %s }\n", userID, email)
}

func (c *Client) TrackEvent(event string, data map[string]interface{}) {
	log.Printf("SEGMENT TRACK_EVENT: { event: %s, data: %+v }\n", event, data)
}

