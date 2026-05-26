package zog

import (
	"github.com/Oudwins/zog/conf"
	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

var _ ComplexZogSchema = &StructSchema{}

type StructSchema struct {
	schema     Shape
	processors []p.ZProcessor[any]
	// defaultVal     any
	required *p.Test[any]
	// catch          any
}

// Returns the type of the schema
func (v *StructSchema) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *

	// Sets the coercer for the schema
	new(zconst.ZogType)
}

func (v *StructSchema) setCoercer(c conf.CoercerFunc) {
	_ = "STUB: not implemented"
	// noop

	// ! USER FACING FUNCTIONS
	return
}

// A map of field names to zog schemas
type Shape map[string]ZogSchema

// Deprecated: use z.Struct(z.Shape{}) instead
// A map of field names to zog schemas
type Schema = Shape

// Returns a new StructSchema which can be used to parse input data into a struct
func Struct(schema Shape) *StructSchema { _ = "STUB: not implemented"; return nil }

// Parses val into destPtr and validates each field based on the schema. Only supports val = map[string]any & dest = &struct
func (v *StructSchema) Parse(data any, destPtr any, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

func (v *StructSchema) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// 2. cast data as DataProvider

// This is a little bit hacky. But we want to exit here because the error came from zhttp. Meaning we had an error trying to parse the request.
// I'm not sure if this is the best behaviour? Do we want to exit here or do we want to continue processing (ofc we add the error always)

// 3. Process / validate struct fields

// Use a size that fits your max key length

// Catch here

// Validate a struct pointer given the struct schema. Usage:
// userSchema.Validate(&User, ...options)
func (v *StructSchema) Validate(dataPtr any, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to validate the data
func (v *StructSchema) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// 2. cast data to string & handle default/required

// 3.1 tests for struct fields

// Use a size that fits your max key length

// Catch here

// Adds posttransform function to schema
func (v *StructSchema) Transform(transform p.Transform[any]) *StructSchema {
	_ = "STUB: not implemented"
	return nil
}

// ! MODIFIERS

// Deprecated: structs are not required or optional. They pass through to the fields. If you want to say that an entire struct may not exist you should use z.Ptr(z.Struct(...))
// This now is a noop. But I believe most people expect it to work how it does now.
// marks field as required
func (v *StructSchema) Required(options ...TestOption) *StructSchema {
	_ = "STUB: not implemented"

	// Deprecated: structs are not required or optional. They pass through to the fields. If you want to say that an entire struct may not exist you should use z.Ptr(z.Struct(...))
	// marks field as optional
	return nil
}

func (v *StructSchema) Optional() *StructSchema {
	_ = "STUB: not implemented"

	// // sets the default value
	//
	//	func (v *StructSchema) Default(val any) *StructSchema {
	//		v.defaultVal = val
	//		return v
	//	}
	return nil
}

// // sets the catch value (i.e the value to use if the validation fails)
// func (v *StructSchema) Catch(val any) *StructSchema {
// 	v.catch = val
// 	return v
// }

// ! VALIDATORS
// custom test function call it -> schema.Test(t z.Test)
func (v *StructSchema) Test(t Test[any]) *StructSchema { _ = "STUB: not implemented"; return nil }

// Create a custom test function for the schema. This is similar to Zod's `.refine()` method.
func (v *StructSchema) TestFunc(testFunc BoolTFunc[any], options ...TestOption) *StructSchema {
	_ = "STUB: not implemented"
	return nil
}
