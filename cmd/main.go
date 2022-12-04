package main

import (
	"database/sql"
	"fmt"

	"dddinterface/proto"
	"dddinterface/services"
	"dddinterface/storage/mysql"
)

func main() {
	fmt.Println("Starting...")

	// Create new storage.
	store := mysql.New(&sql.DB{})
	fmt.Println(store)

	// Create new services.
	serv := services.New(store)
	fmt.Println(serv)

	// Create an invoice.
	inv := &proto.Invoice{
		ID:         1,
		BillTo:     "John Doe",
		PayTo:      "John Smith",
		AmountDue:  100,
		AmountPaid: 0,
	}
	if err := serv.Invoice.Create(inv); err != nil {
		panic(err)
	}

	// Update an invoice.
	if err := serv.Invoice.Update(inv); err != nil {
		panic(err)
	}

	// Process a transaction.
	if err := serv.Transaction.Process(&proto.Transaction{
		ID:             1,
		MerchantID:     1,
		Type:           "sale",
		AmountCaptured: 100,
	}); err != nil {
		panic(err)
	}
}
