package conf

// Default error messages for all schemas. Replace the text with your own messages to customize the error messages for all zog schemas
// As a general rule of thumb, if an error message only has one parameter, the parameter name will be the same as the error code
import (
	"github.com/Oudwins/zog/i18n/en"
	p "github.com/Oudwins/zog/pkgs/internals"
	"github.com/Oudwins/zog/zconst"
)

// Deprecated: Use DefaultIssueMsgMap instead
// Default error messages for all schemas. Replace the text with your own messages to customize the error messages for all zog schemas
// As a general rule of thumb, if an error message only has one parameter, the parameter name will be the same as the error code
var DefaultErrMsgMap zconst.LangMap = en.Map

// Default error messages for all schemas. Replace the text with your own messages to customize the error messages for all zog schemas
// As a general rule of thumb, if an error message only has one parameter, the parameter name will be the same as the error code
var DefaultIssueMessageMap zconst.LangMap = en.Map

const valuePlaceholder = "{{value}}"

func NewDefaultFormatter(m zconst.LangMap) p.IssueFmtFunc {
	_ = "STUB: not implemented"
	return *new(p.IssueFmtFunc)
}

// Check if the error msg is defined do nothing if it set

// Default Issue Message formatter it uses the errors above. Please override the `IssueFormatter` variable instead of this one to customize the error messages for all zog schemas
var DefaultIssueFormatter p.IssueFmtFunc = NewDefaultFormatter(DefaultIssueMessageMap)

// Override this. This is the function use across all Zog schemas to format issue messages
/*
Usage:
```go
conf.IssueFormatter = func(e p.ZogIssue, c z.Ctx) {
     switch e.Code() {
     case zconst.IssueCodeCustom:
          e.SetMessage("Custom message")
	 default:
		conf.DefaultIssueFormatter(e, c) // fallback to default formatter
     }
}
```
*/
var IssueFormatter = DefaultIssueFormatter
