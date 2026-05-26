package zog

import (
	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

type Numeric = p.Numeric

var _ PrimitiveZogSchema[int] = &NumberSchema[int]{}

type NumberSchema[T Numeric] struct {
	processors  []p.ZProcessor[*T]
	defaultFunc func() T
	required    *p.Test[*T]
	catchFunc   func() T
	coercer     CoercerFunc
	isNot       bool
}

type NotNumberSchema[T Numeric] interface {
	OneOf(enum []T, options ...TestOption) *NumberSchema[T]
	EQ(n T, options ...TestOption) *NumberSchema[T]
}

// ! INTERNALS

// Returns the type of the schema
func (v *NumberSchema[T]) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *

	// Sets the coercer for the schema
	new(zconst.ZogType)
}

func (v *NumberSchema[T]) setCoercer(c CoercerFunc) {
	_ = "STUB: not implemented"

	// ! USER FACING FUNCTIONS
	return
}

// Deprecated: Use Float64 instead
// creates a new float64 schema. No plans to remove it but recommended to use Float64 instead.
func Float(opts ...SchemaOption) *NumberSchema[float64] { _ = "STUB: not implemented"; return nil }

func FloatLike[T Numeric](opts ...SchemaOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

func Float64(opts ...SchemaOption) *NumberSchema[float64] { _ = "STUB: not implemented"; return nil }

func Float32(opts ...SchemaOption) *NumberSchema[float32] { _ = "STUB: not implemented"; return nil }

// creates a new int schema
func Int(opts ...SchemaOption) *NumberSchema[int] { _ = "STUB: not implemented"; return nil }

func IntLike[T Numeric](opts ...SchemaOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

func Int64(opts ...SchemaOption) *NumberSchema[int64] { _ = "STUB: not implemented"; return nil }

func Int32(opts ...SchemaOption) *NumberSchema[int32] { _ = "STUB: not implemented"; return nil }

// creates a uint schema
func Uint(opts ...SchemaOption) *NumberSchema[uint] { _ = "STUB: not implemented"; return nil }

func UintLike[T Numeric](opts ...SchemaOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// parses the value and stores it in the destination
func (v *NumberSchema[T]) Parse(data any, dest *T, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to process the data
func (v *NumberSchema[T]) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// Validates a number pointer
func (v *NumberSchema[T]) Validate(data *T, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

func (v *NumberSchema[T]) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// GLOBAL METHODS

// Adds a transform function to the schema. Runs in the order it is called
func (v *NumberSchema[T]) Transform(transform p.Transform[*T]) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// ! MODIFIERS

// marks field as required
func (v *NumberSchema[T]) Required(options ...TestOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// marks field as optional
func (v *NumberSchema[T]) Optional() *NumberSchema[T] { _ = "STUB: not implemented"; return nil }

// sets the default value
func (v *NumberSchema[T]) Default(val T) *NumberSchema[T] { _ = "STUB: not implemented"; return nil }

// sets the default value using a function
func (v *NumberSchema[T]) DefaultFunc(defaultFunc func() T) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// sets the catch value (i.e the value to use if the validation fails)
func (v *NumberSchema[T]) Catch(val T) *NumberSchema[T] { _ = "STUB: not implemented"; return nil }

// sets the catch value (i.e the value to use if the validation fails) using a function
func (v *NumberSchema[T]) CatchFunc(catchFunc func() T) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// custom test function call it -> schema.Test(test, options)
func (v *NumberSchema[T]) Test(t Test[*T]) *NumberSchema[T] { _ = "STUB: not implemented"; return nil }

// Create a custom test function for the schema. This is similar to Zod's `.refine()` method.
func (v *NumberSchema[T]) TestFunc(testFunc BoolTFunc[*T], options ...TestOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// UNIQUE METHODS

// Check that the value is one of the enum values
func (v *NumberSchema[T]) OneOf(enum []T, options ...TestOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// checks for equality
func (v *NumberSchema[T]) EQ(n T, options ...TestOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// checks for lesser or equal
func (v *NumberSchema[T]) LTE(n T, options ...TestOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// checks for greater or equal
func (v *NumberSchema[T]) GTE(n T, options ...TestOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// checks for lesser
func (v *NumberSchema[T]) LT(n T, options ...TestOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// checks for greater
func (v *NumberSchema[T]) GT(n T, options ...TestOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

func (v *NumberSchema[T]) Not() NotNumberSchema[T] { _ = "STUB: not implemented"; return nil }

func (v *NumberSchema[T]) addTest(t *p.Test[*T], fn p.BoolTFunc[*T], options ...TestOption) *NumberSchema[T] {
	_ = "STUB: not implemented"
	return nil
}
