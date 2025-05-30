package quickwit

type Quickwit interface {
	Setup() error
	Shutdown() error
	Ingest(indexID string, data []any) error
	CreateIndex(payload CreateIndexPayload) error
	Search(indexID string, query *SearchQuery) (*SearchResult, error)
	// Add more methods here
	NewMessage(messageID string, channelID string, ownerID string, content string, created_at string) error
}

type quickwit struct {
	client Client
}

func NewQuickwit(config Config) (Quickwit, error) {
	client, err := NewClient(config.Endpoint)
	if err != nil {
		return nil, err
	}
	quickwit := &quickwit{
		client: client,
	}
	return quickwit, nil
}

func (m *quickwit) Setup() error {
	return nil
}

func (m *quickwit) Shutdown() error {
	return nil
}

func (m *quickwit) Ingest(indexID string, data []any) error {
	return m.client.Ingest(indexID, data)
}

func (m *quickwit) CreateIndex(payload CreateIndexPayload) error {
	return m.client.CreateIndexIfNotExists(payload)
}

func (m *quickwit) Search(indexID string, query *SearchQuery) (*SearchResult, error) {
	return m.client.Search(indexID, query)
}

func (m *quickwit) NewMessage(messageID string, channelID string, ownerID string, content string, created_at string) error {
	// create new index if new channel
	err := m.CreateIndex(CreateIndexPayload{
		Version: "0.7",
		IndexID: channelID,
		SearchSettings: &SearchSettings{
			DefaultSearchFields: []string{"content", "ownerID"},
		},
		IndexingSettings: &IndexingSettings{
			CommitTimeoutSecs: 30,
		},
		DocMapping: &DocMapping{
			TimestampField: "created_at",
			FieldMappings: []map[string]any{
				{
					"name":      "content",
					"type":      "text",
					"tokenizer": "default",
					"record":    "position",
					"stored":    true,
				},
				{
					"name":      "ownerID",
					"type":      "text",
					"tokenizer": "default",
					"record":    "position",
					"stored":    true,
				},
				{
					"name":   "messageID",
					"type":   "text",
					"stored": true,
					"tokenizer": "default",
					"record":    "position",
				},
				{
					"name":   "created_at",
					"type":   "datetime",
					"stored": true,
					"input_formats": []string{
						"rfc3339",
					},
					"fast_precision": "seconds",
					"fast":           true,
				},
			},
		},
	})
	if err != nil {
		return err
	}

	return m.Ingest(channelID, []any{
		map[string]any{
			"content":    content,
			"ownerID":    ownerID,
			"created_at": created_at,
			"messageID":  messageID,
		},
	})
}
