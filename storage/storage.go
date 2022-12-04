package storage

import (
	"dddinterface/storage/invoice"
	"dddinterface/storage/transaction"
)

type Storage struct {
	Invoice     invoice.Database
	Transaction transaction.Database
}

func New(s *Storage) *Storage {
	return s
}
