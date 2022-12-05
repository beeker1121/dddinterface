package services

import (
	"dddinterface/services/invoice"
	"dddinterface/services/transaction"
	"dddinterface/storage"
)

type Services struct {
	Invoice     *invoice.Service
	Transaction *transaction.Service
}

func New(s *storage.Storage) *Services {
	return &Services{
		Invoice:     invoice.New(s, transaction.NewCoreService(s)),
		Transaction: transaction.New(s, invoice.NewCoreService(s)),
	}
}
