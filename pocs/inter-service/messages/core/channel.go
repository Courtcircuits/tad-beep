package core

import (
	"errors"
	"fmt"

	"github.com/Courtcircuits/tad-beep/pocs/inter-service/messages/types"
	"github.com/optique-dev/optique"
)

type ChannelService interface {
	ConnectChannel(channelID string, ownerID string)
	Broadcast(message types.Message, channelID string)
	GetChannel(channelID string, ownerID string) (chan types.Message, error)
}

type channel struct {
	channels map[string]map[string]chan types.Message
}

func NewChannelService() ChannelService {
	return &channel{
		channels: make(map[string]map[string]chan types.Message),
	}
}

func (c *channel) ConnectChannel(channelID string, ownerID string) {
	if _, ok := c.channels[ownerID]; !ok {
		c.channels[ownerID] = make(map[string]chan types.Message)
	}
	chans := make([](chan types.Message), 0)
	for _, channel := range c.channels[ownerID] {
		chans = append(chans, channel)
	}
	if _, ok := c.channels[ownerID][channelID]; !ok {
		c.channels[ownerID][channelID] = make(chan types.Message, 1000)
		optique.Info(fmt.Sprintf("Core: channel %s created", channelID))
	}
	optique.Info(fmt.Sprintf("Core: Connected channel %s to owner %s", channelID, ownerID))
}

func (c *channel) Broadcast(message types.Message, channelID string) {
	for owner_id, channs := range c.channels {
		if owner_id == message.Owner {
			continue
		}
		channs[channelID] <- message
	}
}

func (c *channel) GetChannel(channelID string, ownerID string) (chan types.Message, error) {
	if _, ok := c.channels[ownerID]; !ok {
		return nil, ErrChannelNotFound
	}
	if _, ok := c.channels[ownerID][channelID]; !ok {
		return nil, ErrChannelNotFound
	}
	return c.channels[ownerID][channelID], nil
}

var ErrChannelNotFound = errors.New("channel not found")
