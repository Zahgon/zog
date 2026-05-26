package internals

import (
	zconst "github.com/Oudwins/zog/zconst"
)

// Zog Context interface. This is the interface that is passed to schema tests, pre and post transforms
type Ctx interface {
	/**
	METHOD YOU ARE FREE TO USE
	*/
	// Get a value from the context
	Get(key string) any
	// Adds an issue to the schema execution.
	AddIssue(e *ZogIssue)

	// Returns a new issue with the current schema context's data prefilled
	/*
		Usage:

		func MyCustomTestFunc(val any, ctx z.Ctx) {
			if reason1 {
			   ctx.AddIssue(ctx.Issue().SetMessage("Reason 1"))
			} else if reason2 {
			   ctx.AddIssue(ctx.Issue().SetMessage("Reason 2"))
			} else {
			   ctx.AddIssue(ctx.Issue().SetMessage("Reason 3"))
			}
		}

	*/
	Issue() *ZogIssue

	/**
	METHOD YOU SHOULD NOT USE
	*/
	// Deprecated: Use Ctx.AddIssue() instead
	// Please don't depend on this interface it may change
	NewError(p *PathBuilder, e *ZogIssue)
	// Please don't depend on this interface it may change
	HasErrored() bool
}

func NewExecCtx(errs ZogIssues, fmter IssueFmtFunc) *ExecCtx { _ = "STUB: not implemented"; return nil }

type ExecCtx struct {
	Fmter  IssueFmtFunc
	Errors ZogIssues
	m      map[string]any
}

func (c *ExecCtx) HasErrored() bool { _ = "STUB: not implemented"; return false }

func (c *ExecCtx) SetIssueFormatter(fmter IssueFmtFunc) { _ = "STUB: not implemented"; return }

func (c *ExecCtx) Set(key string, val any) { _ = "STUB: not implemented"; return }

func (c *ExecCtx) Get(key string) any {
	_ = "STUB: not implemented"

	// Adds a ZogIssue to the execution context.
	return *new(any)
}

func (c *ExecCtx) AddIssue(e *ZogIssue) { _ = "STUB: not implemented"; return }

func (c *ExecCtx) Issue() *ZogIssue { _ = "STUB: not implemented"; return nil }

// Deprecated: Use Ctx.AddIssue() instead
// This is old interface. It will be removed soon
func (c *ExecCtx) NewError(path *PathBuilder, e *ZogIssue) {
	_ = "STUB: not implemented"

	// Internal. Used to format errors
	return
}

func (c *ExecCtx) FmtErr(e *ZogIssue) { _ = "STUB: not implemented"; return }

func (c *ExecCtx) NewSchemaCtx(val any, destPtr any, path *PathBuilder, dtype zconst.ZogType) *SchemaCtx {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExecCtx) NewValidateSchemaCtx(valPtr any, path *PathBuilder, dtype zconst.ZogType) *SchemaCtx {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExecCtx) Free() { _ = "STUB: not implemented"; return }

type SchemaCtx struct {
	*ExecCtx
	Data      any // input data in case of parse
	ValPtr    any // pointer to real output value in both validate & parse
	Path      *PathBuilder
	DType     zconst.ZogType
	CanCatch  bool
	Exit      bool
	HasCaught bool
	Processor any
}

func (c *SchemaCtx) AddIssue(e *ZogIssue) { _ = "STUB: not implemented"; return }

func (c *SchemaCtx) Issue() *ZogIssue {
	_ = "STUB: not implemented"
	// e := ZogIssuePool.Get().(*ZogIssue)
	// e.Code = ""
	// e.Path = c.Path.String()
	// e.Err = nil
	// e.Message = ""
	// e.Params = nil
	// e.Dtype = c.DType
	// e.Value = c.Data
	return nil
}

// Please don't depend on this method it may change
func (c *SchemaCtx) IssueFromTest(test TestInterface, val any) *ZogIssue {
	_ = "STUB: not implemented"
	return nil
}

// Please don't depend on this method it may change
func (c *SchemaCtx) IssueFromCoerce(err error) *ZogIssue { _ = "STUB: not implemented"; return nil }

// Please don't depend on this method it may change
// Wraps an error in a ZogIssue if it is not already a ZogIssue
func (c *SchemaCtx) IssueFromUnknownError(err error) *ZogIssue {
	_ = "STUB: not implemented"
	return nil
}

// Frees the context to be reused
func (c *SchemaCtx) Free() { _ = "STUB: not implemented"; return }

func (c *SchemaCtx) String() string { _ = "STUB: not implemented"; return "" }

// func (c *TestCtx) Issue() *ZogIssue {
// 	// TODO handle catch here
// 	zerr := ZogIssuePool.Get().(*ZogIssue)
// 	zerr.Code = c.Test.IssueCode
// 	zerr.Path = c.Path.String()
// 	zerr.Err = nil
// 	zerr.Message = ""
// 	zerr.Params = c.Test.Params
// 	zerr.Dtype = c.DType
// 	zerr.Value = c.Data
// 	return zerr
// }

// func (c *TestCtx) FmtErr(e *ZogIssue) {
// 	if e.Message != "" {
// 		return
// 	}

// 	if c.Test.IssueFmtFunc != nil {
// 		c.Test.IssueFmtFunc(e, c)
// 		return
// 	}

// 	c.SchemaCtx.FmtErr(e)
// }

// func (c *TestCtx) AddIssue(e *ZogIssue) {
// 	c.FmtErr(e)
// 	c.Errors.Add(c.Path.String(), e)
// }
