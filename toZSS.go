package zog

import (
	"reflect"

	"github.com/Oudwins/zog/pkgs/internals"
	zss "github.com/Oudwins/zog/pkgs/zss/core"
	"github.com/Oudwins/zog/zconst"
)

// EXPERIMENTAL. PLEASE DO NOT USE UNLESS YOU KNOW WHAT YOU ARE DOING!
type ExMetaRegistry map[any]map[string]any

const (
	EX_META_KEY_FORMAT  = "format"
	EX_META_KEY_MESSAGE = "message"
	EX_META_KEY_ID      = "id"
)

func getZSSGoType[T any]() zss.ZSSGoType { _ = "STUB: not implemented"; return *new(zss.ZSSGoType) }

// If T is a pointer or interface, handle nil case

type ZSSSerializable interface {
	toZSS(*ZSSSerializeCtx) *zss.ZSSSchema
}

type ZSSSerializeCtx struct {
	defs     map[string]*zss.ZSSSchema
	defNames map[ZSSSerializable]int
	nextDef  int
}

func newZSSSerializeCtx() *ZSSSerializeCtx { _ = "STUB: not implemented"; return nil }

func (ctx *ZSSSerializeCtx) refFor(schema ZSSSerializable) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func EXPERIMENTAL_TO_ZSS(s ZSSSerializable) zss.ZSSDocument {
	_ = "STUB: not implemented"
	return *new(zss.ZSSDocument)
}

func (s *StringSchema[T]) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func (s *NumberSchema[T]) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func (s *BoolSchema[T]) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func (s *TimeSchema) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func (s *PointerSchema) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func (s *SliceSchema) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func (s *MapSchema[K, V]) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

// Helper function to shallow copy a map for ZSS
func shallowCopyMap[K comparable, V any](m map[K]V) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func shallowCopyMapFromFunc[K comparable, V any](defaultFunc func() map[K]V) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (s *StructSchema) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func (s *Custom[T]) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

// TODO not sure this is the right place for this info
// TODO THIS create this

func toZSSFields(m Shape, ctx *ZSSSerializeCtx) map[string]*zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func (s *PreprocessSchema[F, T]) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func (s *BoxedSchema[B, T]) toZSS(ctx *ZSSSerializeCtx) *zss.ZSSSchema {
	_ = "STUB: not implemented"
	return nil
}

func processRVtoZSS(rv reflect.Value, dtype zconst.ZogType) *zss.ZSSProcessor {
	_ = "STUB: not implemented"
	return nil

	// TODO better error messages + maybe not panic
}

// TODO better error messages + maybe not panic

func toZSSRequired(test any, dtype zconst.ZogType) *zss.ZSSTest {
	_ = "STUB: not implemented"
	return nil
}

// Check if the underlying value is actually nil using reflection
// This handles the case where a nil pointer is passed as an interface

func processorsToZSS(l reflect.Value, dtype zconst.ZogType) []zss.ZSSProcessor {
	_ = "STUB: not implemented"
	return nil
}

func toZSSProcessorList(test any, dtype zconst.ZogType) []zss.ZSSProcessor {
	_ = "STUB: not implemented"
	return nil
}

// Check if the underlying value is actually nil using reflection
// This handles the case where a nil pointer is passed as an interface

// fakeCtx is a minimal implementation of Ctx interface for extracting default messages
type fakeCtx struct{}

func (f *fakeCtx) Get(key string) any             { _ = "STUB: not implemented"; return *new(any) }
func (f *fakeCtx) AddIssue(e *internals.ZogIssue) { _ = "STUB: not implemented"; return }
func (f *fakeCtx) Issue() *internals.ZogIssue     { _ = "STUB: not implemented"; return nil }
func (f *fakeCtx) NewError(p *internals.PathBuilder, e *internals.ZogIssue) {
	_ = "STUB: not implemented"
	return
}
func (f *fakeCtx) HasErrored() bool { _ = "STUB: not implemented"; return false }

var fakeCtxInstance = &fakeCtx{}

func toZSSTest(test internals.TestInterface, dtype zconst.ZogType) *zss.ZSSTest {
	_ = "STUB: not implemented"
	return nil
}

// Check for custom message in registry first

// If no message is set, extract message using the test's formatter or default

// Use the test's custom formatter if available, otherwise use default

func toZSSTransformer(transformer internals.TransformerInterface) *zss.ZSSTransformer {
	_ = "STUB: not implemented"
	// TODO issue here is that I can't get the code for the transformer and we currently do not have IDs so no way to actually know what this will be
	return nil
}

// extra

func deepCopyPrimitivePtr(v any) any { _ = "STUB: not implemented"; return *new(any) }

func defaultValueFromFunc[T any](defaultFunc func() T) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func defaultValueFromAnyFunc(defaultFunc func() any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
