package substreams

import pbsubstreamsrpc "github.com/streamingfast/substreams/pb/sf/substreams/rpc/v2"

type OperationIs[T any] struct {
	operation pbsubstreamsrpc.StoreDelta_Operation
	negate    bool
	data      []Delta[T]
	index     int
}

func NewOperationIs[T any](operation pbsubstreamsrpc.StoreDelta_Operation, negate bool, data []Delta[T]) *OperationIs[T] {
	return &OperationIs[T]{
		operation: operation,
		negate:    negate,
		data:      data,
		index:     0,
	}
}

func (o *OperationIs[T]) Next() (Delta[T], bool) {
	for o.index < len(o.data) {
		x := o.data[o.index]
		o.index++

		emit := x.GetOperation() == o.operation
		if o.negate {
			emit = !emit
		}

		if emit {
			return x, true
		}
	}

	return nil, false
}
