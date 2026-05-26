package internals

// type PathBuilder string

// func (p PathBuilder) Push(path string) PathBuilder {
// 	if p == "" {
// 		return PathBuilder(path)
// 	}
// 	if path[0] == '[' {
// 		return p + PathBuilder(path)
// 	}
// 	return p + PathBuilder("."+path)
// }

// func (p PathBuilder) String() string {
// 	return string(p)
// }

func NewPathBuilder() *PathBuilder { _ = "STUB: not implemented"; return nil }

type PathBuilder []string

func (p *PathBuilder) Push(path *string) *PathBuilder { _ = "STUB: not implemented"; return nil }

func (p *PathBuilder) Pop() { _ = "STUB: not implemented"; return }

func (p *PathBuilder) ToListClone() []string { _ = "STUB: not implemented"; return nil }

func (p *PathBuilder) String() string { _ = "STUB: not implemented"; return "" }

func FlattenPath(s []string) string { _ = "STUB: not implemented"; return "" }

func (p *PathBuilder) Free() { _ = "STUB: not implemented"; return }
