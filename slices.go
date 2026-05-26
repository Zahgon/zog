package zog

import (
	"github.com/Oudwins/zog/conf"
	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

// ! INTERNALS
var _ ComplexZogSchema = &SliceSchema{}

type SliceSchema struct {
	processors  []p.ZProcessor[any]
	schema      ZogSchema
	required    *p.Test[any]
	defaultFunc func() any
	// catch          any
	coercer conf.CoercerFunc
	isNot   bool
}

type NotSliceSchema interface {
	Len(n int, options ...TestOption) *SliceSchema
	Contains(value any, options ...TestOption) *SliceSchema
}

// Returns the type of the schema
func (v *SliceSchema) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *

	// Sets the coercer for the schema
	new(zconst.ZogType)
}

func (v *SliceSchema) setCoercer(c conf.CoercerFunc) {
	_ = "STUB: not implemented"

	// ! USER FACING FUNCTIONS
	return
}

// Creates a slice schema. That is a Zog representation of a slice.
// It takes a ZogSchema which will be used to validate against all the items in the slice.
func Slice(schema ZogSchema, opts ...SchemaOption) *SliceSchema {
	_ = "STUB: not implemented"
	return nil
}

// default coercer

// Validates a slice
func (v *SliceSchema) Validate(data any, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to validate the data
func (v *SliceSchema) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// we use this to set the value to the ptr. But we still reference the ptr everywhere. This is correct even if it seems confusing.
// 2. cast data to string & handle default/required

// REQUIRED & ZERO VALUE

// 3.1 tests for slice items

// can catch here

// Only supports parsing from data=slice[any] to a dest =&slice[] (this can be typed. Doesn't have to be any)
func (v *SliceSchema) Parse(data any, dest any, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to process the data
func (v *SliceSchema) process(ctx *p.SchemaCtx) {
	_ = "STUB: not implemented"

	// 2. cast data to string & handle default/required
	return
}

// REQUIRED & ZERO VALUE

// make sure val is a slice if not try to make it one

// 3.1 tests for slice items

// Adds transform function to schema.
func (v *SliceSchema) Transform(transform Transform[any]) *SliceSchema {
	_ = "STUB: not implemented"
	return nil
}

// !MODIFIERS

// marks field as required
func (v *SliceSchema) Required(options ...TestOption) *SliceSchema {
	_ = "STUB: not implemented"
	return nil
}

// marks field as optional
func (v *SliceSchema) Optional() *SliceSchema { _ = "STUB: not implemented"; return nil }

// sets the default value
func (v *SliceSchema) Default(val any) *SliceSchema { _ = "STUB: not implemented"; return nil }

// sets the default value using a function
func (v *SliceSchema) DefaultFunc(defaultFunc func() any) *SliceSchema {
	_ = "STUB: not implemented"
	return nil
}

// NOT IMPLEMENTED YET
// sets the catch value (i.e the value to use if the validation fails)
// func (v *SliceSchema) Catch(val string) *SliceSchema {
// 	v.catch = &val
// 	return v
// }

// !TESTS

// custom test function call it -> schema.Test(t z.Test)
func (v *SliceSchema) Test(t Test[any]) *SliceSchema { _ = "STUB: not implemented"; return nil }

// Create a custom test function for the schema. This is similar to Zod's `.refine()` method.
func (v *SliceSchema) TestFunc(testFunc BoolTFunc[any], opts ...TestOption) *SliceSchema {
	_ = "STUB: not implemented"
	return nil
}

// Minimum number of items
func (v *SliceSchema) Min(n int, options ...TestOption) *SliceSchema {
	_ = "STUB: not implemented"
	return nil
}

// Maximum number of items
func (v *SliceSchema) Max(n int, options ...TestOption) *SliceSchema {
	_ = "STUB: not implemented"
	return nil
}

// Exact number of items
func (v *SliceSchema) Len(n int, options ...TestOption) *SliceSchema {
	_ = "STUB: not implemented"
	return nil
}

// Slice contains a specific value
func (v *SliceSchema) Contains(value any, options ...TestOption) *SliceSchema {
	_ = "STUB: not implemented"
	return nil
}

func sliceMin(n int) (p.Test[any], p.BoolTFunc[any]) { _ = "STUB: not implemented"; return nil, nil }

func sliceMax(n int) (p.Test[any], p.BoolTFunc[any]) { _ = "STUB: not implemented"; return nil, nil }

func sliceLength(n int) (p.Test[any], p.BoolTFunc[any]) { _ = "STUB: not implemented"; return nil, nil }

func (v *SliceSchema) Not() NotSliceSchema { _ = "STUB: not implemented"; return *new(NotSliceSchema) }

func (v *SliceSchema) addTest(t *p.Test[any], fn p.BoolTFunc[any], options ...TestOption) *SliceSchema {
	_ = "STUB: not implemented"
	return nil
}
