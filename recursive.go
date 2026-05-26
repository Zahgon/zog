package zog

import (
	"sync"

	p "github.com/Oudwins/zog/pkgs/internals"
	zss "github.com/Oudwins/zog/pkgs/zss/core"
	"github.com/Oudwins/zog/zconst"
)

type lazySchema struct {
	innerSchema ZogSchema
	fn          func() ZogSchema
	once        sync.Once
}

var _ ZogSchema = &lazySchema{}

func (l *lazySchema) get() ZogSchema { _ = "STUB: not implemented"; return *new(ZogSchema) }

func (l *lazySchema) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

func (l *lazySchema) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

func (l *lazySchema) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *new(zconst.ZogType)
}
func (l *lazySchema) setCoercer(c CoercerFunc) { _ = "STUB: not implemented"; return }
func (l *lazySchema) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func lazy(fn func() ZogSchema) *lazySchema { _ = "STUB: not implemented"; return nil }

type RecursiveSchemaUpdater[T ZogSchema] func(self T) T
type RecursiveSchema[T ZogSchema] func(updaters ...RecursiveSchemaUpdater[T]) ZogSchema
type RecursiveSchemaBuilder[T ZogSchema] func(self RecursiveSchema[T]) T

// Experimental API.
// Do not use unless you know what you are doing.
func EXPERIMENTAL_RECURSIVE[T ZogSchema](build RecursiveSchemaBuilder[T]) T {
	_ = "STUB: not implemented"
	return *new(T)
}
