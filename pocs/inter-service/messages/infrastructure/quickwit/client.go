package quickwit

import "fmt"

type Client interface {
	CreateIndex(payload CreateIndexPayload) error
	Ingest(indexID string, data []any) error
	Search(indexID string, query *SearchQuery) (*SearchResult, error)
	CheckIndexExists(indexID string) error
	CreateIndexIfNotExists(payload CreateIndexPayload) error
}

type client struct {
	endpoint string
}

func NewClient(endpoint string) (*client, error) {
	return &client{endpoint}, nil
}

func (c *client) CreateIndex(payload CreateIndexPayload) error {
	return CreateIndex(fmt.Sprintf("%s/api/v1/indexes", c.endpoint), payload)
}

func (c *client) CreateIndexIfNotExists(payload CreateIndexPayload) error {
	err := c.CheckIndexExists(payload.IndexID)
	if err != nil {
		return c.CreateIndex(payload)
	}
	return nil
}

func (c *client) Ingest(indexID string, data []any) error {
	return Ingest(fmt.Sprintf("%s/api/v1/%s/ingest", c.endpoint, indexID), data)
}

func (c *client) Search(indexID string, query *SearchQuery) (*SearchResult, error) {
	return Search(indexID, c.endpoint, query)
}

func (c *client) CheckIndexExists(indexID string) error {
	return CheckIndexExists(indexID, c.endpoint)
}
