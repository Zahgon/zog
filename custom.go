package zog

import (
	p "github.com/Oudwins/zog/pkgs/internals"
	zss "github.com/Oudwins/zog/pkgs/zss/core"
	"github.com/Oudwins/zog/zconst"
)

type Custom[T any] struct {
	test p.Test[*T]
}

func CustomFunc[T any](fn func(ptr *T, ctx Ctx) bool, opts ...TestOption) *Custom[T] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Custom[T]) Parse(data any, destPtr *T, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

func (c *Custom[T]) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// set the value

// run the test

func (c *Custom[T]) Validate(dataPtr *T, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

func (c *Custom[T]) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

func (c *Custom[T]) setCoercer(coercer CoercerFunc) {
	_ = "STUB: not implemented"
	// no op
	return
}

func (c *Custom[T]) getType() zconst.ZogType {
	_ = "STUB: not implemented"

	// Experimental API. Expect breaking changes and no documentation unfortunately for now
	return *new(zconst.ZogType)
}

type EXPERIMENTAL_PUBLIC_ZOG_SCHEMA interface {
	Process(ctx *p.SchemaCtx)
	Validate(ctx *p.SchemaCtx)
	GetType() zconst.ZogType
	SetCoercer(c CoercerFunc)
	ToZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema
}

// Experimental API
func Use(schema EXPERIMENTAL_PUBLIC_ZOG_SCHEMA) *CustomSchema {
	_ = "STUB: not implemented"
	return nil
}

// Experimental API
type CustomSchema struct {
	schema EXPERIMENTAL_PUBLIC_ZOG_SCHEMA
}

var _ ZogSchema = &CustomSchema{}

func (c *CustomSchema) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

func (c *CustomSchema) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

func (c *CustomSchema) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *new(zconst.ZogType)
}

func (c *CustomSchema) setCoercer(coercer CoercerFunc) { _ = "STUB: not implemented"; return }

func (c *CustomSchema) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}
