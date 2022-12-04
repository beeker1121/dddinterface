package mysql

import (
	"database/sql"

	"dddinterface/storage"
	"dddinterface/storage/mysql/invoice"
	"dddinterface/storage/mysql/transaction"
)

func New(db *sql.DB) *storage.Storage {
	return &storage.Storage{
		Invoice:     invoice.New(db),
		Transaction: transaction.New(db),
	}
}
