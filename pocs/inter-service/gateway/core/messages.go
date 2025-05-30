package core

import (
	"context"

	"github.com/Courtcircuits/tad-beep/pocs/inter-service/gateway/infrastructure/grpc"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/gateway/types"
)

type MessageService interface {
	SendMessage(context context.Context, message *types.Message) (*types.Message, error)
	GetMessages(context context.Context, channelID string, ownerID string) (chan types.Message, error)
	SearchMessage(context context.Context, channelID string, query string) ([]types.Message, error)
}

type messageService struct {
	grpcClient grpc.GRPCClient
}

func NewMessageService(grpcClient grpc.GRPCClient) MessageService {
	return &messageService{
		grpcClient: grpcClient,
	}
}

func (m *messageService) SendMessage(ctx context.Context, message *types.Message) (*types.Message, error) {
	return m.grpcClient.SendMessage(ctx, message)
}

func (m *messageService) GetMessages(ctx context.Context, channelID string, ownerID string) (chan types.Message, error) {
	return m.grpcClient.GetMessages(ctx, channelID, ownerID)
}

func (m *messageService) SearchMessage(ctx context.Context, channelID string, query string) ([]types.Message, error) {
	return m.grpcClient.SearchMessage(ctx, channelID, query)
}
