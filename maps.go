package zog

import (
	"github.com/Oudwins/zog/conf"
	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

// ! INTERNALS
var _ ComplexZogSchema = &MapSchema[string, any]{}

type MapSchema[K p.ZogPrimitive, V any] struct {
	processors  []p.ZProcessor[any]
	keySchema   PrimitiveZogSchema[K]
	valueSchema ZogSchema
	required    *p.Test[any]
	defaultFunc func() map[K]V
}

// Returns the type of the schema
func (v *MapSchema[K, V]) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *

	// Sets the coercer for the schema (no-op for maps, kept for interface compliance)
	new(zconst.ZogType)
}

func (v *MapSchema[K, V]) setCoercer(c conf.CoercerFunc) {
	_ = "STUB: not implemented"
	// Maps don't need coercers as we use generics
	return
}

// ! USER FACING FUNCTIONS

// Creates a map schema. That is a Zog representation of a map.
// It takes a PrimitiveZogSchema for keys and a ZogSchema for values.
func EXPERIMENTAL_MAP[K p.ZogPrimitive, V any](keySchema PrimitiveZogSchema[K], valueSchema ZogSchema, opts ...SchemaOption) *MapSchema[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Validates a map
func (v *MapSchema[K, V]) Validate(data *map[K]V, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to validate the data
func (v *MapSchema[K, V]) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// REQUIRED & ZERO VALUE

// Create contexts once for reuse

// Validate map entries

// Validate key

// Validate value - use map's value type, not runtime type of the value
// This ensures proper handling of interface types like `any`

// Parse the data into the destination map
func (v *MapSchema[K, V]) Parse(data any, dest any, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to process the data
func (v *MapSchema[K, V]) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// REQUIRED & ZERO VALUE

// Verify input is a map

// Create destination map

// Create contexts once for reuse

// Process map entries

// Parse key - create a zero value and get pointer to it

// Parse value - create a zero value and get pointer to it

// Only add to map if no errors occurred
// Use reflect.Value directly to avoid type assertion issues with nil interfaces

// Adds transform function to schema.
func (v *MapSchema[K, V]) Transform(transform Transform[any]) *MapSchema[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// !MODIFIERS

// marks field as required
func (v *MapSchema[K, V]) Required(options ...TestOption) *MapSchema[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// marks field as optional
func (v *MapSchema[K, V]) Optional() *MapSchema[K, V] { _ = "STUB: not implemented"; return nil }

// sets the default value
func (v *MapSchema[K, V]) Default(val map[K]V) *MapSchema[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// sets the default value using a function
func (v *MapSchema[K, V]) DefaultFunc(defaultFunc func() map[K]V) *MapSchema[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// !TESTS

// custom test function call it -> schema.Test(t z.Test)
func (v *MapSchema[K, V]) Test(t Test[any]) *MapSchema[K, V] { _ = "STUB: not implemented"; return nil }

// Create a custom test function for the schema. This is similar to Zod's `.refine()` method.
func (v *MapSchema[K, V]) TestFunc(testFunc BoolTFunc[any], opts ...TestOption) *MapSchema[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Minimum number of entries
func (v *MapSchema[K, V]) Min(n int, options ...TestOption) *MapSchema[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Maximum number of entries
func (v *MapSchema[K, V]) Max(n int, options ...TestOption) *MapSchema[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Exact number of entries
func (v *MapSchema[K, V]) Len(n int, options ...TestOption) *MapSchema[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func mapMin(n int) (p.Test[any], p.BoolTFunc[any]) { _ = "STUB: not implemented"; return nil, nil }

func mapMax(n int) (p.Test[any], p.BoolTFunc[any]) { _ = "STUB: not implemented"; return nil, nil }

func mapLength(n int) (p.Test[any], p.BoolTFunc[any]) { _ = "STUB: not implemented"; return nil, nil }

func (v *MapSchema[K, V]) addTest(t *p.Test[any], fn p.BoolTFunc[any], options ...TestOption) *MapSchema[K, V] {
	_ = "STUB: not implemented"
	return nil
}
