package zog

import (
	p "github.com/Oudwins/zog/pkgs/internals"
	zss "github.com/Oudwins/zog/pkgs/zss/core"
	"github.com/Oudwins/zog/zconst"
)

var _ ZogSchema = &AnySchema{}

type AnySchema struct {
	processors  []p.ZProcessor[*any]
	defaultFunc func() any
	required    *p.Test[*any]
	catchFunc   func() any
}

// ! INTERNALS

// Returns the type of the schema
func (v *AnySchema) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *

	// Sets the coercer for the schema (no-op for Any schema)
	new(zconst.ZogType)
}

func (v *AnySchema) setCoercer(c CoercerFunc) {
	_ = "STUB: not implemented"
	// no-op - Any schema doesn't need coercion
	return
}

// ! USER FACING FUNCTIONS

// Returns a new EXPERIMENTAL_ANY Schema
// Do not use unless you know what you are doing & are okay with possible breaking changes & bugs.
func EXPERIMENTAL_ANY(opts ...SchemaOption) *AnySchema { _ = "STUB: not implemented"; return nil }

// Parse data into destination pointer
func (v *AnySchema) Parse(data any, dest *any, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to process the data
func (v *AnySchema) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// Handle default/required for nil values

// This handles optional case

// is required & zero value

// For Any schema, we just assign the value as-is

// Process all processors (tests and transforms)

// Validate data against schema
func (v *AnySchema) Validate(val *any, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to validate data
func (v *AnySchema) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// Handle default/required for zero values

// This handles optional case

// is required & zero value

// Process all processors (tests and transforms)

// GLOBAL METHODS

func (v *AnySchema) Test(t p.Test[*any]) *AnySchema { _ = "STUB: not implemented"; return nil }

// Create a custom test function for the schema. This is similar to Zod's `.refine()` method.
func (v *AnySchema) TestFunc(testFunc p.BoolTFunc[*any], options ...p.TestOption) *AnySchema {
	_ = "STUB: not implemented"
	return nil
}

// Adds a transform function to the schema. Runs in the order it is called
func (v *AnySchema) Transform(transform p.Transform[*any]) *AnySchema {
	_ = "STUB: not implemented"
	return nil
}

// ! MODIFIERS

// marks field as required
func (v *AnySchema) Required(options ...TestOption) *AnySchema {
	_ = "STUB: not implemented"
	return nil
}

// marks field as optional
func (v *AnySchema) Optional() *AnySchema { _ = "STUB: not implemented"; return nil }

// sets the default value
func (v *AnySchema) Default(val any) *AnySchema { _ = "STUB: not implemented"; return nil }

// sets the default value using a function
func (v *AnySchema) DefaultFunc(defaultFunc func() any) *AnySchema {
	_ = "STUB: not implemented"
	return nil
}

// sets the catch value (i.e the value to use if the validation fails)
func (v *AnySchema) Catch(val any) *AnySchema { _ = "STUB: not implemented"; return nil }

// sets the catch value (i.e the value to use if the validation fails) using a function
func (v *AnySchema) CatchFunc(catchFunc func() any) *AnySchema {
	_ = "STUB: not implemented"
	return nil
}

// toZSS converts the schema to ZSS format
func (v *AnySchema) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}
