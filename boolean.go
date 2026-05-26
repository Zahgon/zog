package zog

import (
	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

var _ PrimitiveZogSchema[bool] = &BoolSchema[bool]{}

type BoolSchema[T ~bool] struct {
	processors  []p.ZProcessor[*T]
	defaultFunc func() T
	required    *p.Test[*T]
	catchFunc   func() T
	coercer     CoercerFunc
}

// ! INTERNALS

// Returns the type of the schema
func (v *BoolSchema[T]) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *

	// Sets the coercer for the schema
	new(zconst.ZogType)
}

func (v *BoolSchema[T]) setCoercer(c CoercerFunc) {
	_ = "STUB: not implemented"

	// ! USER FACING FUNCTIONS
	return
}

// Returns a new Bool Shape
func Bool(opts ...SchemaOption) *BoolSchema[bool] { _ = "STUB: not implemented"; return nil }

// default coercer

func BoolLike[T ~bool](opts ...SchemaOption) *BoolSchema[T] { _ = "STUB: not implemented"; return nil }

// Parse data into destination pointer
func (v *BoolSchema[T]) Parse(data any, dest *T, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to process the data
func (v *BoolSchema[T]) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// Validate data against schema
func (v *BoolSchema[T]) Validate(val *T, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to validate data
func (v *BoolSchema[T]) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// GLOBAL METHODS

func (v *BoolSchema[T]) Test(t p.Test[*T]) *BoolSchema[T] { _ = "STUB: not implemented"; return nil }

// Create a custom test function for the schema. This is similar to Zod's `.refine()` method.
func (v *BoolSchema[T]) TestFunc(testFunc p.BoolTFunc[*T], options ...p.TestOption) *BoolSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Adds a transform function to the schema. Runs in the order it is called (i.e schema.True().Transform(...) will run the transform after the True test)
func (v *BoolSchema[T]) Transform(transform p.Transform[*T]) *BoolSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// ! MODIFIERS
// marks field as required
func (v *BoolSchema[T]) Required(options ...TestOption) *BoolSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// marks field as optional
func (v *BoolSchema[T]) Optional() *BoolSchema[T] { _ = "STUB: not implemented"; return nil }

// sets the default value
func (v *BoolSchema[T]) Default(val T) *BoolSchema[T] { _ = "STUB: not implemented"; return nil }

// sets the default value using a function
func (v *BoolSchema[T]) DefaultFunc(defaultFunc func() T) *BoolSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// sets the catch value (i.e the value to use if the validation fails)
func (v *BoolSchema[T]) Catch(val T) *BoolSchema[T] { _ = "STUB: not implemented"; return nil }

// sets the catch value (i.e the value to use if the validation fails) using a function
func (v *BoolSchema[T]) CatchFunc(catchFunc func() T) *BoolSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// UNIQUE METHODS

func (v *BoolSchema[T]) True() *BoolSchema[T] { _ = "STUB: not implemented"; return nil }

func (v *BoolSchema[T]) False() *BoolSchema[T] { _ = "STUB: not implemented"; return nil }

func (v *BoolSchema[T]) EQ(val T) *BoolSchema[T] { _ = "STUB: not implemented"; return nil }

func (v *BoolSchema[T]) addTest(t *p.Test[*T], fn p.BoolTFunc[*T]) *BoolSchema[T] {
	_ = "STUB: not implemented"
	return nil
}
