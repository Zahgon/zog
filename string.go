package zog

import (
	"regexp"

	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

var (
	_ PrimitiveZogSchema[string] = (*StringSchema[string])(nil)
	_ NotStringSchema[string]    = (*StringSchema[string])(nil)
)

type likeString interface {
	~string
}

type NotStringSchema[T likeString] interface {
	OneOf(enum []T, options ...TestOption) *StringSchema[T]
	Len(n int, options ...TestOption) *StringSchema[T]
	Email(options ...TestOption) *StringSchema[T]
	URL(options ...TestOption) *StringSchema[T]
	HasPrefix(s T, options ...TestOption) *StringSchema[T]
	HasSuffix(s T, options ...TestOption) *StringSchema[T]
	Contains(sub T, options ...TestOption) *StringSchema[T]
	ContainsUpper(options ...TestOption) *StringSchema[T]
	ContainsDigit(options ...TestOption) *StringSchema[T]
	ContainsSpecial(options ...TestOption) *StringSchema[T]
	UUID(options ...TestOption) *StringSchema[T]
	Match(regex *regexp.Regexp, options ...TestOption) *StringSchema[T]

	// `Test` method is missing here as we require the user to define their own test for their use case.
	// `Not` method is missing here as we do not want the user to do `Not` chaining.
	// `NotNil`, `Min`, `Max` methods are not included as they are opposites of each other.
}

type StringSchema[T likeString] struct {
	processors  []p.ZProcessor[*T]
	defaultFunc func() T
	required    *p.Test[*T]
	catchFunc   func() T
	coercer     CoercerFunc
	isNot       bool
}

// ! INTERNALS

// Returns the type of the schema
func (v *StringSchema[T]) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *

	// Sets the coercer for the schema
	new(zconst.ZogType)
}

func (v *StringSchema[T]) setCoercer(c CoercerFunc) {
	_ = "STUB: not implemented"

	// ! USER FACING FUNCTIONS
	return
}

func StringLike[T likeString](opts ...SchemaOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Returns a new String Shape
func String(opts ...SchemaOption) *StringSchema[string] { _ = "STUB: not implemented"; return nil }

// default coercer

// Parses the data into the destination string. Returns a list of ZogIssues
func (v *StringSchema[T]) Parse(data any, dest *T, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to process the data
func (v *StringSchema[T]) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// Validate Given string
func (v *StringSchema[T]) Validate(data *T, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to validate the data
func (v *StringSchema[T]) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// Transform: trims the input data of whitespace if it is a string
func (v *StringSchema[T]) Trim() *StringSchema[T] { _ = "STUB: not implemented"; return nil }

// Adds a transform function to the schema. Runs in the order it is called
func (v *StringSchema[T]) Transform(transform p.Transform[*T]) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// ! MODIFIERS

// marks field as required
func (v *StringSchema[T]) Required(options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// marks field as optional
func (v *StringSchema[T]) Optional() *StringSchema[T] { _ = "STUB: not implemented"; return nil }

// sets the default value
func (v *StringSchema[T]) Default(val T) *StringSchema[T] { _ = "STUB: not implemented"; return nil }

// sets the default value using a function
func (v *StringSchema[T]) DefaultFunc(defaultFunc func() T) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// sets the catch value (i.e the value to use if the validation fails)
func (v *StringSchema[T]) Catch(val T) *StringSchema[T] { _ = "STUB: not implemented"; return nil }

// sets the catch value (i.e the value to use if the validation fails) using a function
func (v *StringSchema[T]) CatchFunc(catchFunc func() T) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// ! Tests
// custom test function call it -> schema.Test(t z.Test, opts ...TestOption)
func (v *StringSchema[T]) Test(t Test[*T]) *StringSchema[T] { _ = "STUB: not implemented"; return nil }

// Create a custom test function for the schema. This is similar to Zod's `.refine()` method.
func (v *StringSchema[T]) TestFunc(testFunc BoolTFunc[*T], options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value is one of the enum values
func (v *StringSchema[T]) OneOf(enum []T, options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value is at least n characters long
func (v *StringSchema[T]) Min(n int, options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value is at most n characters long
func (v *StringSchema[T]) Max(n int, options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value is exactly n characters long
func (v *StringSchema[T]) Len(n int, options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value is a valid email address
func (v *StringSchema[T]) Email(options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value is a valid URL
func (v *StringSchema[T]) URL(options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value is a valid IP address (IPv4 or IPv6)
func (v *StringSchema[T]) IP(options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value is a valid IPv4 address
func (v *StringSchema[T]) IPv4(options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value is a valid IPv6 address
func (v *StringSchema[T]) IPv6(options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value has the prefix
func (v *StringSchema[T]) HasPrefix(s T, options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value has the suffix
func (v *StringSchema[T]) HasSuffix(s T, options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value contains the substring
func (v *StringSchema[T]) Contains(sub T, options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value contains an uppercase letter
func (v *StringSchema[T]) ContainsUpper(options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value contains a digit
func (v *StringSchema[T]) ContainsDigit(options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value contains a special character
func (v *StringSchema[T]) ContainsSpecial(options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that the value is a valid uuid
func (v *StringSchema[T]) UUID(options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Test: checks that value matches to regex
func (v *StringSchema[T]) Match(regex *regexp.Regexp, options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}

// Not returns a schema that negates the next validation test.
// For example, `z.String().Not().Email()` validates that the string is NOT a valid email.
// Note: The negation only applies to the next validation test and is reset afterward.
func (v *StringSchema[T]) Not() NotStringSchema[T] { _ = "STUB: not implemented"; return nil }

func (v *StringSchema[T]) addTest(t p.Test[*T], fn p.BoolTFunc[*T], options ...TestOption) *StringSchema[T] {
	_ = "STUB: not implemented"
	return nil
}
