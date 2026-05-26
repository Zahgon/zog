package internals

import (
	"reflect"
)

func GetKeyFromField(field reflect.StructField, fallback string, tag *string) string {
	_ = "STUB: not implemented"
	return ""
}

type DpFactory = func() (DataProvider, *ZogIssue)

// This is used for parsing structs & maps
type DataProvider interface {
	Get(key string) any
	GetByField(field reflect.StructField, fallback string) (any, string)
	GetNestedProvider(key string) DataProvider
	GetUnderlying() any // returns the underlying value the dp is wrapping
}

// checks that we implement the interface
var _ DataProvider = &MapDataProvider[string]{}
var _ DataProvider = &StructDataProvider{}

type StructDataProvider struct {
	value reflect.Value
	tag   *string
}

func (s *StructDataProvider) Get(key string) any { _ = "STUB: not implemented"; return *new(any) }

func (s *StructDataProvider) GetByField(field reflect.StructField, fallback string) (any, string) {
	_ = "STUB: not implemented"
	return *new(any), ""
}

func (s *StructDataProvider) GetNestedProvider(key string) DataProvider {
	_ = "STUB: not implemented"
	return *new(DataProvider)
}

func (s *StructDataProvider) GetUnderlying() any { _ = "STUB: not implemented"; return *new(any) }

type MapDataProvider[T any] struct {
	M   map[string]T
	tag *string
}

func (m *MapDataProvider[T]) Get(key string) any { _ = "STUB: not implemented"; return *new(any) }

// returns value + key used
func (m *MapDataProvider[T]) GetByField(field reflect.StructField, fallback string) (any, string) {
	_ = "STUB: not implemented"
	return *new(any), ""
}

func (m *MapDataProvider[T]) GetNestedProvider(key string) DataProvider {
	_ = "STUB: not implemented"
	return *new(DataProvider)
}

func (m *MapDataProvider[T]) GetUnderlying() any { _ = "STUB: not implemented"; return *new(any) }

func NewMapDataProvider[T any](m map[string]T, tag *string) DataProvider {
	_ = "STUB: not implemented"
	return *new(DataProvider)
}

func NewSafeMapDataProvider[T any](m map[string]T) DataProvider {
	_ = "STUB: not implemented"
	return *new(DataProvider)
}

type EmptyDataProvider struct {
	Underlying any
}

func (e *EmptyDataProvider) Get(key string) any { _ = "STUB: not implemented"; return *new(any) }

func (e *EmptyDataProvider) GetByField(field reflect.StructField, fallback string) (any, string) {
	_ = "STUB: not implemented"
	return *new(any), ""
}

func (e *EmptyDataProvider) GetNestedProvider(key string) DataProvider {
	_ = "STUB: not implemented"
	return *new(DataProvider)
}

func (e *EmptyDataProvider) GetUnderlying() any { _ = "STUB: not implemented"; return *new(any) }

func TryNewAnyDataProvider(val any) (DataProvider, error) {
	_ = "STUB: not implemented"
	return *new(DataProvider), nil
}

// TODO: add more types
