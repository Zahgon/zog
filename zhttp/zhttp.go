package zhttp

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"

	"github.com/Oudwins/zog/parsers/zjson"
	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

type ParserFunc = func(r *http.Request) p.DpFactory

var (
	formTag    string = "form"
	queryParam string = "query"
)

var Config = struct {
	Parsers struct {
		JSON          ParserFunc
		Form          ParserFunc
		Query         ParserFunc
		MultipartForm ParserFunc
	}
}{
	Parsers: struct {
		JSON          ParserFunc
		Form          ParserFunc
		Query         ParserFunc
		MultipartForm ParserFunc
	}{
		JSON: func(r *http.Request) p.DpFactory {
			return zjson.Decode(r.Body)
		},
		Form: func(r *http.Request) p.DpFactory {
			return func() (p.DataProvider, *p.ZogIssue) {
				if r.Form == nil { // Check in case user already parsed the form
					err := r.ParseForm()
					if err != nil {
						return nil, &p.ZogIssue{Code: zconst.IssueCodeZHTTPInvalidForm, Err: err}
					}
				}
				return form(r.Form, &formTag), nil
			}
		},
		MultipartForm: func(r *http.Request) p.DpFactory {
			return func() (p.DataProvider, *p.ZogIssue) {
				if r.MultipartForm == nil {
					// See this article on why/how to correctly parse multipart form data: https://medium.com/@owlwalks/dont-parse-everything-from-client-multipart-post-golang-9280d23cd4ad
					return nil, &p.ZogIssue{Code: zconst.IssueCodeZHTTPInvalidMultipartForm, Err: errors.New("You must parse multipart form data before using it with zhttp")}
				}
				return form(r.MultipartForm.Value, &formTag), nil
			}
		},
		Query: func(r *http.Request) p.DpFactory {
			return func() (p.DataProvider, *p.ZogIssue) {
				// This handles generic GET request from browser. We treat it as url.Values
				return form(r.URL.Query(), &queryParam), nil
			}
		},
	},
}

type urlDataProvider struct {
	Data url.Values
	tag  *string
}

var _ p.DataProvider = urlDataProvider{}

func (u urlDataProvider) Get(key string) any { _ = "STUB: not implemented"; return *new(any) }

// if query param ends with [] its always a slice

func (u urlDataProvider) GetByField(field reflect.StructField, fallback string) (any, string) {
	_ = "STUB: not implemented"
	return *new(any), ""
}

func (u urlDataProvider) GetNestedProvider(key string) p.DataProvider {
	_ = "STUB: not implemented"
	return *new(p.DataProvider)
}

func (u urlDataProvider) GetUnderlying() any {
	_ = "STUB: not implemented"

	// Parses JSON, Form & Query data from request based on Content-Type header
	// Usage:
	// schema.Parse(zhttp.Request(r), &dest)
	// WARNING: FOR JSON PARSING DOES NOT SUPPORT JSON ARRAYS OR PRIMITIVES
	return *new(any)
}

func Request(r *http.Request) p.DpFactory { _ = "STUB: not implemented"; return *new(p.DpFactory) }

// Content-Type follows this format: Content-Type: <media-type> [; parameter=value]

func form(data url.Values, tag *string) p.DataProvider {
	_ = "STUB: not implemented"
	return *new(p.DataProvider)
}

// func params(data url.Values) p.DataProvider {
// 	return form(data)
// }
