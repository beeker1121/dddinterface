package transaction

import (
	"dddinterface/proto"
	"dddinterface/storage"
	"dddinterface/storage/transaction"
)

type Service struct {
	s *storage.Storage
}

func New(s *storage.Storage) *Service {
	return &Service{
		s: s,
	}
}

func (is *Service) Create(t *proto.Transaction) error {
	is.s.Transaction.Create(&transaction.CreateParams{})

	return nil
}

func (is *Service) Process(t *proto.Transaction) error {
	// Create a transaction.
	if err := is.Create(t); err != nil {
		return err
	}

	return nil
}
