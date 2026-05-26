package internals

func Flatten(issues ZogIssueList) map[string][]string { _ = "STUB: not implemented"; return nil }

func GroupByFlattenedPath(issues ZogIssueList) map[string]ZogIssueList {
	_ = "STUB: not implemented"
	return nil
}

func Treeify(issues ZogIssueList) map[string]any { _ = "STUB: not implemented"; return nil }

// Root level errors (empty or nil path)

// Navigate to the target location in the tree

// Process all path segments except the last one to build the structure

// This is an array index
// Ensure current is a map with "items" key

// This shouldn't happen in normal flow, but handle it

// Ensure the slice is large enough

// If this is the last segment, create error structure

// Not the last segment, continue navigating

// This is a property name

// This shouldn't happen in normal flow, but handle it

// If this is the last segment, add error to this property

// Not the last segment, continue navigating

// parseArrayIndex attempts to parse a path segment as an array index.
// Returns (index, true) if it's an array index, (0, false) otherwise.
// Handles both numeric strings (like "1") and bracket notation (like "[1]").
func parseArrayIndex(segment string) (int, bool) {
	_ = "STUB: not implemented"
	// Check if it's bracket notation like "[1]"
	return 0, false
}

// Check if it's a plain numeric string

func Prettify(issues ZogIssueList) string { _ = "STUB: not implemented"; return "" }
