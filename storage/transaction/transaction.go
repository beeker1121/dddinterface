package transaction

type Database interface {
	Create(params *CreateParams) error
}

type Transaction struct {
	ID             uint
	MerchantID     uint
	Type           string
	AmountCaptured uint
}

type CreateParams struct {
	ID             uint
	MerchantID     uint
	Type           string
	AmountCaptured uint
}
