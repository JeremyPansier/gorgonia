package stdops

import (
	"context"
	"runtime/trace"

	"gorgonia.org/gorgonia/values"
	"gorgonia.org/tensor"
)

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

// addOp is the base op for elementwise addition.
type addOp struct{ binop }

// String implements fmt.Stringer.
func (op addOp) String() string { return "+" }

// Do performs elementwise addition.
func (op addOp) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Add(a, b, tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise addition but with a preallocated return value.
// PreallocDo allows add to implement ops.PreallocOp.
func (op addOp) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Add(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// addVV is a tensor-tensor elementwise addition.
type addVV struct {
	addOp
	binopVV
}

// addVS is a tensor-scalar elementwise addition.
type addVS struct {
	addOp
	binopVS
}

// String implements fmt.Stringer.
func (op addVS) String() string { return "+·" }

// addSV is a scalar-tensor elementwise addition.
type addSV struct {
	addOp
	binopSV
}

// String implements fmt.Stringer.
func (op addSV) String() string { return "·+" }
