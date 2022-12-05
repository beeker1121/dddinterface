package invoice

import (
	"dddinterface/proto"
	trancomm "dddinterface/services/transaction/comm"
	"dddinterface/storage"
	"dddinterface/storage/invoice"
)

type Service struct {
	s                  *storage.Storage
	transactionService trancomm.TransactionDescriptor
}

func New(s *storage.Storage, ts trancomm.TransactionDescriptor) *Service {
	return &Service{
		s:                  s,
		transactionService: ts,
	}
}

func (s *Service) Create(i *proto.Invoice) error {
	cs := CoreService{s.s}
	return cs.Create(i)
}

func (s *Service) Update(i *proto.Invoice) error {
	cs := CoreService{s.s}
	return cs.Update(i)
}

func (s *Service) Pay(i *proto.Invoice) error {
	// Process a transaction.
	s.transactionService.Process(&proto.Transaction{
		ID:             2,
		MerchantID:     1,
		Type:           "sale",
		AmountCaptured: 100,
	})

	// Update the invoice.
	if err := s.Update(i); err != nil {
		return err
	}

	return nil
}

type CoreService struct {
	s *storage.Storage
}

func NewCoreService(s *storage.Storage) *CoreService {
	return &CoreService{
		s: s,
	}
}

func (cs *CoreService) Create(i *proto.Invoice) error {
	cs.s.Invoice.Create(&invoice.CreateParams{})

	return nil
}

func (cs *CoreService) Update(i *proto.Invoice) error {
	cs.s.Invoice.Update(&invoice.Invoice{})

	return nil
}
