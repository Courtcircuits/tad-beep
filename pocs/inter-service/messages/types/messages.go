package types

type Message struct {
	MessageID string `json:"messageID" db:"id"`
	Content   string `json:"content" db:"content"`
	Owner     string `json:"ownerID" db:"owner_id"`
	Channel   string `json:"channelID" db:"channel_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
