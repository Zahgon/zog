package zog

import (
	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

type PreprocessSchema[F any, T any] struct {
	schema ZogSchema
	fn     func(data F, ctx Ctx) (out T, err error)
}

// out should never be a pointer type
func Preprocess[F any, T any](fn func(data F, ctx Ctx) (out T, err error), schema ZogSchema) *PreprocessSchema[F, T] {
	_ = "STUB: not implemented"
	return nil
}

func (s *PreprocessSchema[F, T]) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

func (s *PreprocessSchema[F, T]) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

func (s *PreprocessSchema[F, T]) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *new(zconst.ZogType)
}

func (s *PreprocessSchema[F, T]) setCoercer(coercer CoercerFunc) { _ = "STUB: not implemented"; return }

func (s *PreprocessSchema[F, T]) Parse(data F, destPtr *T, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

func (s *PreprocessSchema[F, T]) Validate(data *T, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}
