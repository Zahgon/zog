//go:build !zogmeta
// +build !zogmeta

package zog

const (
	EXHAUSTIVE_METADATA = false
)

var exMetaRegistry ExMetaRegistry = nil

func RegistryAdd(_ ExMetaRegistry, _ any, _ string, _ any) {
	_ = "STUB: not implemented"
	// no op
	return
}

func RegistryGet(_ ExMetaRegistry, _ any, _ string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}
