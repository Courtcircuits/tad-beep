package types

type Message struct {
	MessageID     string `json:"messageID"`
	ChannelID     string `json:"channelID"`
	OwnerID       string `json:"ownerID"`
	Content       string `json:"content"`
	CreatedAt     string `json:"createdAt"`
}
