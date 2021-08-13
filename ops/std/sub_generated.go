package stdops

import (
	"context"
	"runtime/trace"

	"gorgonia.org/gorgonia/values"
	"gorgonia.org/tensor"
)

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

// subOp is the base op for elementwise subtraction.
type subOp struct{ binop }

// String implements fmt.Stringer.
func (op subOp) String() string { return "-" }

// Do performs elementwise subtraction.
func (op subOp) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Sub(a, b, tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise subtraction but with a preallocated return value.
// PreallocDo allows sub to implement ops.PreallocOp.
func (op subOp) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	retVal, err = tensor.Sub(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	task.End()
	return retVal, err
}

// subVV is a tensor-tensor elementwise subtraction.
type subVV struct {
	subOp
	binopVV
}

// subVS is a tensor-scalar elementwise subtraction.
type subVS struct {
	subOp
	binopVS
}

// String implements fmt.Stringer.
func (op subVS) String() string { return "-·" }

// subSV is a scalar-tensor elementwise subtraction.
type subSV struct {
	subOp
	binopSV
}

// String implements fmt.Stringer.
func (op subSV) String() string { return "·-" }
