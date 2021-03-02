package calab

// RenderFunction ...
type RenderFunction func(n uint64, s Space)

// Renderer represents a new image renderer.
type Renderer interface {
	Render(n uint64, s Space)
}
