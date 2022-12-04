package invoice

import (
	"fmt"

	"database/sql"
	"dddinterface/storage/invoice"
)

var invoiceMap map[uint]*invoice.Invoice = make(map[uint]*invoice.Invoice)

type Database struct {
	db *sql.DB
}

func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

func (db *Database) Create(params *invoice.CreateParams) error {
	fmt.Println("Created invoice")
	return nil
}

func (db *Database) Update(i *invoice.Invoice) error {
	fmt.Println("Updated invoice")
	invoiceMap[i.ID] = i
	return nil
}
