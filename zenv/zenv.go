package zenv

import (
	"reflect"

	p "github.com/Oudwins/zog/pkgs/internals"
)

var _ p.DataProvider = &envDataProvider{}

var (
	envTag string = "env"
)

type envDataProvider struct {
}

func (e *envDataProvider) Get(key string) any { _ = "STUB: not implemented"; return *new(any) }

func (e *envDataProvider) GetByField(field reflect.StructField, fallback string) (any, string) {
	_ = "STUB: not implemented"
	return *new(any), ""
}

func (e *envDataProvider) GetNestedProvider(key string) p.DataProvider {
	_ = "STUB: not implemented"
	return *new(p.DataProvider)
}

func NewDataProvider() *envDataProvider { _ = "STUB: not implemented"; return nil }

func (e *envDataProvider) GetUnderlying() any { _ = "STUB: not implemented"; return *new(any) }
