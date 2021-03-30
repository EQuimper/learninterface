package analytics

type Client interface {
	Identify(userID int, email string)
	TrackEvent(event string, data map[string]interface{})
}

type Analytics struct {
	Client
}

func New(client Client) *Analytics  {
	return &Analytics{
		Client: client,
	}
}

func (a *Analytics) Register(userID int, email string) {
	a.Client.Identify(userID, email)
}

func (a *Analytics) Login(userID int, email string) {
	a.Client.TrackEvent("LOGIN", map[string]interface{}{
		"userId": userID,
		"email": email,
	})
}