package internals

import (
	zconst "github.com/Oudwins/zog/zconst"
)

// this is the function that formats the error message given a zog error
type IssueFmtFunc = func(e *ZogIssue, p Ctx)

// ZogIssue represents an issue that occurred during parsing or validation.
// When printed it looks like:
// ZogIssue{Code: coercion_issue, Params: map[], Type: number, Value: not_empty, Message: number is invalid, Error: failed to coerce string int: strconv.Atoi: parsing "not_empty": invalid syntax}
type ZogIssue struct {
	// Code is the unique identifier for the issue. Generally also the ID for the Test that caused the issue.
	Code zconst.ZogIssueCode
	// Path is the path to the field that caused the issue
	Path []string
	// Value is the data value that caused the issue.
	// If using Schema.Parse(data, dest) then this will be the value of data.
	Value any
	// Dtype is the destination type. i.e The zconst.ZogType of the value that was validated.
	// If using Schema.Parse(data, dest) then this will be the type of dest.
	Dtype zconst.ZogType
	// Params is the params map for the issue. Taken from the Test that caused the issue.
	// This may be nil if Test has no params.
	Params map[string]any
	// Message is the human readable, user-friendly message for the issue.
	// This is safe to expose to the user.
	Message string
	// Err is the wrapped error or nil if none
	Err error
}

func NewZogIssue() *ZogIssue { _ = "STUB: not implemented"; return nil }

// SetCode sets the issue code for the issue and returns the issue for chaining
func (i *ZogIssue) SetCode(c zconst.ZogIssueCode) *ZogIssue { _ = "STUB: not implemented"; return nil }

// SetPath sets the path for the issue and returns the issue for chaining
func (i *ZogIssue) SetPath(p []string) *ZogIssue { _ = "STUB: not implemented"; return nil }

// SetValue sets the data value that caused the issue and returns the issue for chaining
func (i *ZogIssue) SetValue(v any) *ZogIssue { _ = "STUB: not implemented"; return nil }

// SetDType sets the destination type for the issue and returns the issue for chaining
func (i *ZogIssue) SetDType(t zconst.ZogType) *ZogIssue { _ = "STUB: not implemented"; return nil }

// SetParams sets the params map for the issue and returns the issue for chaining
func (i *ZogIssue) SetParams(p map[string]any) *ZogIssue { _ = "STUB: not implemented"; return nil }

// SetMessage sets the human readable, user-friendly message for the issue and returns the issue for chaining
func (i *ZogIssue) SetMessage(m string) *ZogIssue { _ = "STUB: not implemented"; return nil }

// SetError sets the wrapped error for the issue and returns the issue for chaining
func (i *ZogIssue) SetError(e error) *ZogIssue { _ = "STUB: not implemented"; return nil }

// Unwrap returns the wrapped error or nil if none
func (i *ZogIssue) Unwrap() error {
	_ = "STUB: not implemented"

	// Error returns the string representation of the ZogIssue (same as String())
	return nil
}

func (i *ZogIssue) Error() string {
	_ = "STUB: not implemented"

	// String returns the string representation of the ZogIssue (same as Error())
	return ""
}

func (i *ZogIssue) String() string { _ = "STUB: not implemented"; return "" }

// Returns a stringified version of the path based on the flatten path string logic
func (i *ZogIssue) PathString() string { _ = "STUB: not implemented"; return "" }

func FreeIssue(i *ZogIssue) { _ = "STUB: not implemented"; return }

// ZogIssueList is the unified return type for all schema Parse/Validate operations
type ZogIssueList = []*ZogIssue

// Deprecated: ZogIssueMap is deprecated. All schemas now return ZogIssueList.
// Users should migrate to using ZogIssueList and access paths via issue.Path
type ZogIssueMap = map[string]ZogIssueList

// INTERNAL ONLY: Interface used to add errors during parsing & validation. It represents a group of errors
type ZogIssues interface {
	Add(err *ZogIssue)
	IsEmpty() bool
	Free()
}

// ErrsList - internal structure for collecting issues during schema execution
type ErrsList struct {
	List ZogIssueList
}

// internal only
func NewErrsList() *ErrsList { _ = "STUB: not implemented"; return nil }

func (e *ErrsList) Add(err *ZogIssue) { _ = "STUB: not implemented"; return }

// Slightly larger initial capacity

// Path is already set on the issue by SchemaCtx.Issue() or IssueFrom* methods

func (e *ErrsList) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (e *ErrsList) Free() { _ = "STUB: not implemented"; return }
