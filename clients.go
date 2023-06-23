package nixmq

type Clients struct {
	internal map[string]*Client
}

func NewClients() *Clients {
	return &Clients{
		make(map[string]*Client),
	}
}

func (c *Clients) Add(client *Client) {
	if client.Propreties.ClientID == "" {
		client.GenerateClientID()
	}

	c.internal[client.Propreties.ClientID] = client
}

func (c *Clients) Get(clientID string) (*Client, bool) {
	client, ok := c.internal[clientID]
	return client, ok
}

func (c *Clients) Remove(clientID string) {
	delete(c.internal, clientID)
}