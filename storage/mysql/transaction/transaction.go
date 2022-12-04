package transaction

import (
	"fmt"

	"database/sql"
	"dddinterface/storage/transaction"
)

var transactionMap map[uint]*transaction.Transaction = make(map[uint]*transaction.Transaction)

type Database struct {
	db *sql.DB
}

func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

func (db *Database) Create(params *transaction.CreateParams) error {
	fmt.Println("Created transaction")
	return nil
}
