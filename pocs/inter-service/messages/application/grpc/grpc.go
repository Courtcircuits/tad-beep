package grpc

import (
	context "context"
	"fmt"
	"net"
	"time"

	"github.com/Courtcircuits/tad-beep/pocs/inter-service/messages/core"
	"github.com/optique-dev/optique"
	"google.golang.org/grpc"
)

type GRPCServer interface {
	Ignite() error
	Stop() error
	SwapServer(server MessagesServer)
}

type grpcMessageServer struct {
	listen_addr string
	server      MessagesServer
	conn        *grpc.Server
}

func NewGRPCServer(config Config, messageService core.MessageService, channelService core.ChannelService) GRPCServer {
	return &grpcMessageServer{
		listen_addr: config.ListenAddr,
		server: &messageServer{
			messageService: messageService,
			channelService: channelService,
		},
	}
}

// the one that really implements the grpc server
type messageServer struct {
	messageService core.MessageService
	channelService core.ChannelService
	UnimplementedMessagesServer
}

func (s *messageServer) SendMessage(ctx context.Context, in *MessageRequest) (*Message, error) {
	message, err := s.messageService.SendMessage(in.ChannelID, in.OwnerID, in.Content)
	if err != nil {
		return nil, err
	}
	return &Message{Content: message.Content, ChannelID: message.Channel, OwnerID: message.Owner, CreatedAt: message.CreatedAt, MessageID: message.MessageID}, nil
}

func (s *messageServer) GetMessages(in *GetMessagesQuery, stream grpc.ServerStreamingServer[Message]) error {
	messages, err := s.messageService.GetMessages(in.ChannelID)
	if err != nil {
		return err
	}
	for _, message := range messages {
		err := stream.Send(&Message{
			Content:   message.Content,
			ChannelID: message.Channel,
			OwnerID:   message.Owner,
			MessageID: message.MessageID,
			CreatedAt: message.CreatedAt,
		})
		if err != nil {
			optique.Error(err.Error())
			return err
		}
	}
	s.channelService.ConnectChannel(in.ChannelID, in.OwnerID)
	channel, err := s.channelService.GetChannel(in.ChannelID, in.OwnerID)
	if err != nil {
		return nil
	}
	for {
		select {
		case message := <-channel:
			err := stream.Send(&Message{
				Content:   message.Content,
				ChannelID: message.Channel,
				OwnerID:   message.Owner,
				MessageID: message.MessageID,
				CreatedAt: message.CreatedAt,
			})
			if err != nil {
				return err
			}
		default:
			time.Sleep(time.Microsecond)
		}
	}
}

func (s *messageServer) SearchMessage(in *SearchQuery, stream grpc.ServerStreamingServer[Message]) error {
	search_results, err := s.messageService.SearchMessage(in.ChannelID, in.Query)
	if err != nil {
		return err
	}
	for _, message := range search_results {
		err := stream.Send(&Message{
			MessageID: message.MessageID,
			Content:   message.Content,
			ChannelID: message.Channel,
			OwnerID:   message.Owner,
			CreatedAt: message.CreatedAt,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *grpcMessageServer) Ignite() error {
	var opts []grpc.ServerOption
	// auth_option := grpc.WithTransportCredentials(insecure.NewCredentials())
	lis, err := net.Listen("tcp", s.listen_addr)
	// opts = append(opts, op
	grpcServer := grpc.NewServer(opts...)
	RegisterMessagesServer(grpcServer, s.server)
	if err != nil {
		return err
	}
	s.conn = grpcServer
	optique.Info(fmt.Sprintf("GRPC server listening on %s", s.listen_addr))
	grpcServer.Serve(lis)
	return nil
}

func (s *grpcMessageServer) Stop() error {
	s.conn.GracefulStop()
	return nil
}

func (s *grpcMessageServer) SwapServer(server MessagesServer) {
	s.server = server
}
