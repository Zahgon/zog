package zog

import (
	"github.com/Oudwins/zog/conf"
	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

var _ ComplexZogSchema = &PointerSchema{}

type PointerSchema struct {
	schema   ZogSchema
	required *p.Test[any]
	// postTransforms []PostTransform
	// defaultVal     *any
	// catch          *any
}

func (v *PointerSchema) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	// return zconst.TypePtr
	return *new(zconst.ZogType)
}

func (v *PointerSchema) setCoercer(c conf.CoercerFunc) { _ = "STUB: not implemented"; return }

// Ptr creates a pointer ZogSchema
func Ptr(schema ZogSchema) *PointerSchema { _ = "STUB: not implemented"; return nil }

// Parse the data into the destination pointer
func (v *PointerSchema) Parse(data any, dest any, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

func (v *PointerSchema) process(ctx *p.SchemaCtx) {
	_ = "STUB: not implemented"

	// TODO this is a mess. But couldn't figure out a simple way to support top level optional structs without doing this.
	// Companion code to this codde is in struct.go > process
	return
}

// End of messy code

// We set the destination type to the schema type because pointer doesn't have any issue messages. They pass through to the schema type

// this sets the primitive also

// this generates a new nil pointer
//newVal := reflect.Zero(destPtr.Type())

// Validates a pointer pointer
func (v *PointerSchema) Validate(data any, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

func (v *PointerSchema) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// We set the destination type to the schema type because pointer doesn't have any issue messages. They pass through to the schema type

// Validate Existing Pointer

func (v *PointerSchema) NotNil(options ...TestOption) *PointerSchema {
	_ = "STUB: not implemented"
	return nil
}
