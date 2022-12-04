package comm

import "dddinterface/proto"

type InvoiceDescriptor interface {
	Create(i *proto.Invoice) error
	Update(i *proto.Invoice) error
}
