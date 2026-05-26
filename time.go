package zog

import (
	"time"

	"github.com/Oudwins/zog/conf"
	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

// ! INTERNALS
var _ PrimitiveZogSchema[time.Time] = &TimeSchema{}

type TimeSchema struct {
	processors  []p.ZProcessor[*time.Time]
	defaultFunc func() time.Time
	required    *p.Test[*time.Time]
	catchFunc   func() time.Time
	coercer     conf.CoercerFunc
}

// Returns the type of the schema
func (v *TimeSchema) getType() zconst.ZogType {
	_ = "STUB: not implemented"
	return *

	// Sets the coercer for the schema
	new(zconst.ZogType)
}

func (v *TimeSchema) setCoercer(c conf.CoercerFunc) { _ = "STUB: not implemented"; return }

type TimeFunc func(opts ...SchemaOption) *TimeSchema

// ! USER FACING FUNCTIONS

// Returns a new Time Shape
var Time TimeFunc = func(opts ...SchemaOption) *TimeSchema {
	t := &TimeSchema{
		coercer: conf.Coercers.Time,
	}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

// WARNING ONLY SUPPOORTS Shape.Parse!
// Sets the format function for the time schema.
// Usage is:
//
//	z.Time(z.Time.FormatFunc(func(data string) (time.Time, error) {
//		return time.Parse(time.RFC3339, data)
//	}))
func (t TimeFunc) FormatFunc(format func(data string) (time.Time, error)) SchemaOption {
	_ = "STUB: not implemented"
	return *new(SchemaOption)
}

// WARNING ONLY SUPPOORTS Shape.Parse!
// Sets the string format for the  time schema
// Usage is:
// z.Time(z.Time.Format(time.RFC3339))
func (t TimeFunc) Format(format string) SchemaOption {
	_ = "STUB: not implemented"
	return *new(SchemaOption)
}

// Parses the data into the destination time.Time. Returns a list of errors
func (v *TimeSchema) Parse(data any, dest *time.Time, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// internal processes the data
func (v *TimeSchema) process(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// Validates an existing time.Time
func (v *TimeSchema) Validate(data *time.Time, options ...ExecOption) ZogIssueList {
	_ = "STUB: not implemented"
	return *new(ZogIssueList)
}

// Internal function to validate the data
func (v *TimeSchema) validate(ctx *p.SchemaCtx) { _ = "STUB: not implemented"; return }

// Adds posttransform function to schema
func (v *TimeSchema) Transform(transform Transform[*time.Time]) *TimeSchema {
	_ = "STUB: not implemented"
	return nil
}

// ! MODIFIERS

// marks field as required
func (v *TimeSchema) Required(options ...TestOption) *TimeSchema {
	_ = "STUB: not implemented"
	return nil
}

// marks field as optional
func (v *TimeSchema) Optional() *TimeSchema { _ = "STUB: not implemented"; return nil }

// sets the default value
func (v *TimeSchema) Default(val time.Time) *TimeSchema { _ = "STUB: not implemented"; return nil }

// sets the default value using a function
func (v *TimeSchema) DefaultFunc(defaultFunc func() time.Time) *TimeSchema {
	_ = "STUB: not implemented"
	return nil
}

// sets the catch value (i.e the value to use if the validation fails)
func (v *TimeSchema) Catch(val time.Time) *TimeSchema { _ = "STUB: not implemented"; return nil }

// sets the catch value (i.e the value to use if the validation fails) using a function
func (v *TimeSchema) CatchFunc(catchFunc func() time.Time) *TimeSchema {
	_ = "STUB: not implemented"
	return nil
}

// GLOBAL METHODS

// custom test function call it -> schema.Test(z.Test{Func: func (val *time.Time, ctx z.Ctx) {
// my test
// }})
func (v *TimeSchema) Test(t Test[*time.Time]) *TimeSchema { _ = "STUB: not implemented"; return nil }

// Create a custom test function for the schema. This is similar to Zod's `.refine()` method.
func (v *TimeSchema) TestFunc(testFunc BoolTFunc[*time.Time], options ...TestOption) *TimeSchema {
	_ = "STUB: not implemented"
	return nil
}

// UNIQUE METHODS

// Checks that the value is after the given time
func (v *TimeSchema) After(t time.Time, opts ...TestOption) *TimeSchema {
	_ = "STUB: not implemented"
	return nil
}

// Checks that the value is before the given time
func (v *TimeSchema) Before(t time.Time, opts ...TestOption) *TimeSchema {
	_ = "STUB: not implemented"
	return nil
}

// Checks that the value is equal to the given time
func (v *TimeSchema) EQ(t time.Time, opts ...TestOption) *TimeSchema {
	_ = "STUB: not implemented"
	return nil
}
