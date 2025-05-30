package core

import (
	"fmt"
	"time"

	"github.com/Courtcircuits/tad-beep/pocs/inter-service/messages/infrastructure/quickwit"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/messages/infrastructure/sql"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/messages/types"
	"github.com/optique-dev/optique"
)

type MessageService interface {
	SendMessage(channelID string, ownerID string, content string) (*types.Message, error)
	GetMessages(channelID string) ([]types.Message, error)
	SearchMessage(channelID string, query string) ([]types.Message, error)
}

type messages struct {
	channelService  ChannelService
	sqlService      sql.Sql
	quickwitService quickwit.Quickwit
}

func NewMessageService(channelService ChannelService, sqlService sql.Sql, quickwitService quickwit.Quickwit) MessageService {
	return &messages{
		channelService:  channelService,
		sqlService:      sqlService,
		quickwitService: quickwitService,
	}
}

func (m *messages) SendMessage(channelID string, ownerID string, content string) (*types.Message, error) {
	now := time.Now().Format(time.RFC3339)
	message, err := m.sqlService.CreateMessage(&types.Message{
		Content:   content,
		Channel:   channelID,
		Owner:     ownerID,
		CreatedAt: now,
	})
	optique.Info(fmt.Sprintf("Core: New message %v at %s", message, now))
	if err != nil {
		optique.Error(err.Error())
		return nil, err
	}
	optique.Info(fmt.Sprintf("SQL: created message %s", content))
	err = m.quickwitService.NewMessage(message.MessageID, message.Channel, message.Owner, message.Content, message.CreatedAt)
	optique.Info(fmt.Sprintf("Quickwit: created index %s", channelID))
	message_copy := types.Message{
		Content:   message.Content,
		Owner:     message.Owner,
		Channel:   message.Channel,
		CreatedAt: message.CreatedAt,
		MessageID: message.MessageID,
	}
	go func() {
		m.channelService.Broadcast(message_copy, channelID)
	}()
	optique.Info(fmt.Sprintf("Core: message sent to channel %s", channelID))

	return message, nil
}

func (m *messages) GetMessages(channelID string) ([]types.Message, error) {
	messages, err := m.sqlService.GetMessages(channelID)
	if err != nil {
		optique.Error(err.Error())
		return nil, err
	}
	optique.Info(fmt.Sprintf("Core: found %v messages", messages))
	return messages, nil
}

func (m *messages) SearchMessage(channelID string, query string) ([]types.Message, error) {
	optique.Info(fmt.Sprintf("Core: searching query %s", query))
	search_query := quickwit.NewSearchQuery(query)
	search_results, err := m.quickwitService.Search(channelID, search_query)
	if err != nil {
		optique.Error(err.Error())
		return nil, err
	}
	optique.Info(fmt.Sprintf("Quickwit: found %d messages", search_results.NumHits))
	var messages []types.Message
	for _, hit := range search_results.Hits {
		optique.Info(fmt.Sprintf("Quickwit: found message %s", hit))
		messages = append(messages, types.Message{
			MessageID: hit["messageID"].(string),
			Content:   hit["content"].(string),
			Owner:     hit["ownerID"].(string),
			Channel:   channelID,
			CreatedAt: hit["created_at"].(string),
		})
	}
	return messages, nil
}
