package zog

import (
	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

type CreateBoxFunc[T any, B any] func(data T, ctx Ctx) (B, error)
type UnboxFunc[B any, T any] func(data B, ctx Ctx) (T, error)

var _ ComplexZogSchema = &BoxedSchema[any, any]{}

type BoxedSchema[B any, T any] struct {
	schema ZogSchema
	unbox  UnboxFunc[B, T]
	box    CreateBoxFunc[T, B]
}

func Boxed[B any, T any](schema ZogSchema, unboxFunc UnboxFunc[B, T], boxFunc CreateBoxFunc[T, B]) *BoxedSchema[B, T] {
	_ = "STUB: not implemented"
	return nil
}

func (s *BoxedSchema[B, T]) Parse(data any, dest any, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

func (s *BoxedSchema[B, T]) Validate(dest *B, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

func (s *BoxedSchema[B, T]) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// Re-box and propagate back

func (s *BoxedSchema[B, T]) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// 1. Handle ctx.Data - could be B, *B, or raw data

// Unbox B to get T

// Dereference and unbox

// Raw data - pass directly to inner schema

// 2. Create temporary T for inner schema and pass data

// 3. Process through inner schema (keeps pointer to inner)

// 4. Re-box and set to original destination

// TODO maybe some kind of flag that you executed process with boxFunc is nil

func (s *BoxedSchema[B, T]) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *new(zconst.ZogType)
}

func (s *BoxedSchema[B, T]) setCoercer(c CoercerFunc) { _ = "STUB: not implemented"; return }
