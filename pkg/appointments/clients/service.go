package clients

import "context"

type service struct {
	repo ClientRepository
}

func (s *service) FetchByUsername(ctx context.Context, username string) (Client, error) {
	return Client{}, nil
}
