package tutils

import (
	"testing"

	"github.com/Oudwins/zog/pkgs/internals"
)

// VerifyDefaultIssueMessages verifies that all issues have valid default messages
func VerifyDefaultIssueMessages(t *testing.T, errs internals.ZogIssueList) {
	_ = "STUB: not implemented"
	return
}

// Deprecated: Use VerifyDefaultIssueMessages instead.
// All schemas now return ZogIssueList.
func VerifyDefaultIssueMessagesMap(t *testing.T, errs internals.ZogIssueMap) {
	_ = "STUB: not implemented"
	return
}

// FindByPath returns all issues with the given path
func FindByPath(errs internals.ZogIssueList, path string) internals.ZogIssueList {
	_ = "STUB: not implemented"
	return *new(internals.ZogIssueList)
}

// HasPath checks if any issue has the given path
func HasPath(errs internals.ZogIssueList, path string) bool {
	_ = "STUB: not implemented"
	return false
}
