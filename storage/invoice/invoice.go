package invoice

type Database interface {
	Create(params *CreateParams) error
	Update(i *Invoice) error
}

type Invoice struct {
	ID         uint
	BillTo     string
	PayTo      string
	AmountDue  uint
	AmountPaid uint
}

type CreateParams struct {
	ID         uint
	BillTo     string
	PayTo      string
	AmountDue  uint
	AmountPaid uint
}
