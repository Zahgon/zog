package zog

// Merge combines two or more schemas into a new schema.
// It performs a shallow merge, meaning:
//   - Fields with the same key from later schemas override earlier ones
//   - Modifying nested schemas may affect the original schemas
//
// Parameters:
//   - other: The first schema to merge with
//   - others: Additional schemas to merge
//
// Returns a new schema containing the merged fields and transforms
func (v *StructSchema) Merge(other *StructSchema, others ...*StructSchema) *StructSchema {
	_ = "STUB: not implemented"
	return nil
}

// processors

// cloneShallow creates a shallow copy of the schema.
// The new schema shares references to the transforms, tests and inner schema.
func (v *StructSchema) cloneShallow() *StructSchema { _ = "STUB: not implemented"; return nil }

// Omit creates a new schema with specified fields removed.
// It accepts either strings or map[string]bool as arguments:
//   - Strings directly specify fields to omit
//   - For maps, fields are omitted when their boolean value is true
//
// Returns a new schema with the specified fields removed
func (v *StructSchema) Omit(vals ...any) *StructSchema { _ = "STUB: not implemented"; return nil }

// Pick creates a new schema keeping only the specified fields.
// It accepts either strings or map[string]bool as arguments:
//   - Strings directly specify fields to keep
//   - For maps, fields are kept when their boolean value is true
//
// Returns a new schema containing only the specified fields
func (v *StructSchema) Pick(picks ...any) *StructSchema { _ = "STUB: not implemented"; return nil }

// Extend creates a new schema by adding additional fields from the provided schema.
// Fields in the provided schema override any existing fields with the same key.
//
// Parameters:
//   - schema: The schema containing fields to add
//
// Returns a new schema with the additional fields
func (v *StructSchema) Extend(schema Shape) *StructSchema { _ = "STUB: not implemented"; return nil }
