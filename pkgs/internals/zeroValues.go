package internals

type IsZeroValueFunc = func(val any, ctx Ctx) bool

// checks that the value is the zero value for its type
func IsZeroValue(x any) bool { _ = "STUB: not implemented"; return false }

// checks if the value is the zero value but only for parsing purposes (i.e the parse function)
func IsParseZeroValue(val any, ctx Ctx) bool { _ = "STUB: not implemented"; return false }
