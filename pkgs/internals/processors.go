package internals

// Internal Processor interface
type ZProcessor[T any] interface {
	ZProcess(valPtr T, ctx Ctx)
}

type TransformProcessor[T any] struct {
	Transform Transform[T]
}

func (p *TransformProcessor[T]) ZProcess(valPtr T, ctx Ctx) { _ = "STUB: not implemented"; return }

func (p *TransformProcessor[T]) GetName() string {
	_ = "STUB: not implemented"
	// Temporary. TODO actually implement this
	return ""
}

type TransformerInterface interface {
	GetName() string
}
