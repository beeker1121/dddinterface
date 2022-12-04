package comm

import "dddinterface/proto"

type TransactionDescriptor interface {
	Process(t *proto.Transaction)
}
