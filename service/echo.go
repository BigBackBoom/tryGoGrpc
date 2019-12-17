package echoService

import (
	"context"
)

type EchoService struct {
}

func (s *EchoService) GetEcho(ctx context.Context, message *GetEchoMessage) (*EchoResponse, error) {
	return &EchoResponse{
		Input: message.TargetEcho,
	}, nil
}
