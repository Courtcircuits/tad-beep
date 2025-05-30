package grpc

import (
	"context"
	"io"

	"github.com/Courtcircuits/tad-beep/pocs/inter-service/gateway/types"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient interface {
	Setup() error
	Shutdown() error
	SendMessage(context.Context, *types.Message) (*types.Message, error)
	GetMessages(context.Context, string, string) (chan types.Message, error)
	SearchMessage(context.Context, string, string) ([]types.Message, error)
}

type grpcClient struct {
	conn   *grpc.ClientConn
	client MessagesClient
}

func NewGRPCClient(config Config) (GRPCClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(config.Endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &grpcClient{
		conn:   conn,
		client: NewMessagesClient(conn),
	}, nil
}

func (g *grpcClient) Setup() error {
	return nil
}

func (g *grpcClient) Shutdown() error {
	g.conn.Close()
	return nil
}

func (g *grpcClient) SendMessage(ctx context.Context, message *types.Message) (*types.Message, error) {
	req := &MessageRequest{
		ChannelID: message.ChannelID,
		OwnerID:   message.OwnerID,
		Content:   message.Content,
	}
	res, err := g.client.SendMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	return &types.Message{
		MessageID: res.MessageID,
		ChannelID: res.ChannelID,
		OwnerID:   res.OwnerID,
		Content:   res.Content,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (g *grpcClient) GetMessages(ctx context.Context, channelID string, ownerID string) (chan types.Message, error) {
	query := &GetMessagesQuery{
		ChannelID: channelID,
		OwnerID: ownerID,
	}
	stream, err := g.client.GetMessages(ctx, query)
	if err != nil {
		return nil, err
	}
	messages := make(chan types.Message)
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				close(messages)
				return
			}
			messages <- types.Message{
				MessageID: res.MessageID,
				ChannelID: res.ChannelID,
				OwnerID:   res.OwnerID,
				Content:   res.Content,
				CreatedAt: res.CreatedAt,
			}
		}
	}()
	return messages, nil
}

func (g *grpcClient) SearchMessage(ctx context.Context, channelID, query string) ([]types.Message, error) {
	search_query := &SearchQuery{
		ChannelID: channelID,
		Query:     query,
	}
	stream, err := g.client.SearchMessage(ctx, search_query)
	if err != nil {
		return nil, err
	}
	messages := make([]types.Message, 0)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return messages, err
		}
		messages = append(messages, types.Message{
			MessageID: res.MessageID,
			ChannelID: res.ChannelID,
			OwnerID:   res.OwnerID,
			Content:   res.Content,
			CreatedAt: res.CreatedAt,
		})
	}
	return messages, nil
}
