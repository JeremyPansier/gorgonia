package stdops

import (
	"context"
	"runtime/trace"

	"github.com/chewxy/hm"
	"gorgonia.org/gorgonia/types"
	"gorgonia.org/gorgonia/values"
	"gorgonia.org/tensor"
)

// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT.

// elEqOp is the base op for elementwise equal-to.
type elEqOp struct {
	binop
	retSame bool
}

// String implements fmt.Stringer.
func (op elEqOp) String() string { return "=" }

// Do performs elementwise equal-to.
func (op elEqOp) Do(ctx context.Context, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	// Do the actual operation
	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.ElEq(a, b, tensor.WithContext(ctx2), tensor.AsSameType())
	} else {
		retVal, err = tensor.ElEq(a, b, tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}

// PreallocDo performs elementwise equal-to but with a preallocated return value.
// PreallocDo allows elEq to implement ops.PreallocOp.
func (op elEqOp) PreallocDo(ctx context.Context, prealloc values.Value, vs ...values.Value) (retVal values.Value, err error) {
	if err := handleCtx(ctx); err != nil {
		return nil, err
	}

	a := vs[0].(tensor.Tensor)
	b := vs[1].(tensor.Tensor)

	ctx2, task := trace.NewTask(ctx, op.String())
	if op.retSame {
		retVal, err = tensor.ElEq(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2), tensor.AsSameType())
	} else {
		retVal, err = tensor.ElEq(a, b, tensor.WithReuse(prealloc), tensor.WithContext(ctx2))
	}
	task.End()
	return retVal, err
}                                           // DiffWRT returns {false, false} for elEq
func (op elEqOp) DiffWRT(inputs int) []bool { return twofalses }

// elEqVV is a tensor-tensor elementwise equal-to.
type elEqVV struct {
	elEqOp
	binopVV
}

// Type returns the type: (·) : a → a → a or (·) :  a → a → b
func (op elEqVV) Type() hm.Type {
	a := hm.TypeVariable('a') // (T U) or U
	if op.retSame {
		return hm.NewFnType(a, a, a)
	}
	b := types.MakeDependent(a, tensor.Bool) // (T Bool) or Bool
	return hm.NewFnType(a, a, b)
}

// elEqVS is a tensor-scalar elementwise equal-to.
type elEqVS struct {
	elEqOp
	binopVS
}

// String implements fmt.Stringer.
func (op elEqVS) String() string { return "=·" }

// Type returns the type: (·) : a → b → a or (·) :  a → b → c
func (op elEqVS) Type() hm.Type {
	a := hm.TypeVariable('a') // (T U) or U
	b := hm.TypeVariable('b') // U
	if op.retSame {
		return hm.NewFnType(a, b, a)
	}
	c := types.MakeDependent(a, tensor.Bool) // (T Bool) or Bool
	return hm.NewFnType(a, b, c)
}

// elEqSV is a scalar-tensor elementwise equal-to.
type elEqSV struct {
	elEqOp
	binopSV
}

// String implements fmt.Stringer.
func (op elEqSV) String() string { return "·=" }

// Type returns the type: (·) : a → b → b or (·) :  a → b → c
func (op elEqSV) Type() hm.Type {
	a := hm.TypeVariable('a') // U
	b := hm.TypeVariable('b') // (T U) or U
	if op.retSame {
		return hm.NewFnType(a, b, b)
	}
	c := types.MakeDependent(b, tensor.Bool) // (T Bool) or Bool
	return hm.NewFnType(a, b, c)
}
