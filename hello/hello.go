package hello

import context "context"

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	return &Message{Body: "Hello World"}, nil
}
