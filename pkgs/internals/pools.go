package internals

import (
	"strings"
	"sync"
)

var ExecCtxPool = sync.Pool{
	New: func() any {
		return &ExecCtx{}
	},
}

var SchemaCtxPool = sync.Pool{
	New: func() any {
		return &SchemaCtx{}
	},
}

var InternalIssueListPool = sync.Pool{
	New: func() any {
		return &ErrsList{}
	},
}

var ZogIssuePool = sync.Pool{
	New: func() any {
		return &ZogIssue{}
	},
}

var PathBuilderPool = sync.Pool{
	New: func() any {
		pb := make(PathBuilder, 0, 5)
		return &pb
	},
}

var StringBuilderPool = sync.Pool{
	New: func() any {
		sb := strings.Builder{}
		return &sb
	},
}

func NewStringBuilder() *strings.Builder { _ = "STUB: not implemented"; return nil }

func FreeStringBuilder(sb *strings.Builder) { _ = "STUB: not implemented"; return }

func ClearPools() { _ = "STUB: not implemented"; return }

func Clear() { _ = "STUB: not implemented"; return }
