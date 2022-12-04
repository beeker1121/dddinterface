package invoice

import (
	"dddinterface/proto"
	"dddinterface/storage"
	"dddinterface/storage/invoice"
)

type Service struct {
	s *storage.Storage
}

func New(s *storage.Storage) *Service {
	return &Service{
		s: s,
	}
}

func (is *Service) Create(i *proto.Invoice) error {
	is.s.Invoice.Create(&invoice.CreateParams{})

	return nil
}

func (is *Service) Update(i *proto.Invoice) error {
	is.s.Invoice.Update(&invoice.Invoice{})

	return nil
}
