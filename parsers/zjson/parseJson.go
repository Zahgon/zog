package zjson

import (
	"io"

	p "github.com/Oudwins/zog/pkgs/internals"
)

// func Unmarshal(data []byte) p.DpFactory {
// 	return func() (p.DataProvider, p.ZogIssue) {
// 		var m map[string]any
// 		err := json.Unmarshal(data, &m)
// 		if err != nil {
// 			return nil, &p.ZogErr{C: zconst.IssueCodeInvalidJSON, Err: err}
// 		}
// 		if m == nil {
// 			return nil, &p.ZogErr{C: zconst.IssueCodeInvalidJSON, Err: errors.New("nill json body")}
// 		}
// 		return p.NewMapDataProvider(m), nil
// 	}
// }

var (
	jsonTag string = "json"
)

// Decodes JSON data. Does not support json arrays or primitives
/*
- "null" -> nil -> Not accepted by zhttp -> errs["zconst.ISSUE_KEY_ROOT"]-> required issue
- "{}" -> okay -> map[]{}
- "" -> parsing error -> errs["zconst.ISSUE_KEY_ROOT"]-> parsing error
- "1213" -> zhttp -> plain value
  - struct schema -> hey this valid input
  - "string is not an object"
*/
func Decode(r io.Reader) p.DpFactory { _ = "STUB: not implemented"; return *new(p.DpFactory) }
