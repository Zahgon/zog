package internals

const defaultString = "<nil>"

func SafeString(x any) string { _ = "STUB: not implemented"; return "" }

func SafeError(x error) string { _ = "STUB: not implemented"; return "" }

func UnwrapPtr(x any) any { _ = "STUB: not implemented"; return *new(any) }
