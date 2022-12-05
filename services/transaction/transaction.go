package transaction

import (
	"dddinterface/proto"
	invcomm "dddinterface/services/invoice/comm"
	"dddinterface/storage"
	"dddinterface/storage/transaction"
)

type Service struct {
	s              *storage.Storage
	invoiceService invcomm.InvoiceDescriptor
}

func New(s *storage.Storage, is invcomm.InvoiceDescriptor) *Service {
	return &Service{
		s:              s,
		invoiceService: is,
	}
}

func (s *Service) Create(t *proto.Transaction) error {
	cs := CoreService{s: s.s}
	return cs.Create(t)
}

func (s *Service) Process(t *proto.Transaction) error {
	cs := CoreService{s: s.s}
	if err := cs.Process(t); err != nil {
		return err
	}

	// if voided, check if transaction was for invoice, if so change status.
	// Update invoice.
	s.invoiceService.Update(&proto.Invoice{
		ID:         1,
		BillTo:     "John Doe",
		PayTo:      "John Smith",
		AmountDue:  0,
		AmountPaid: 100,
	})

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

func (cs *CoreService) Create(t *proto.Transaction) error {
	cs.s.Transaction.Create(&transaction.CreateParams{})

	return nil
}

func (cs *CoreService) Process(t *proto.Transaction) error {
	// Create a transaction.
	if err := cs.Create(t); err != nil {
		return err
	}

	return nil
}
