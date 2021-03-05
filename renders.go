package calab

// ObserverFunction ...
type ObserverFunction func(n uint64, s Space)

// Renderer represents a new image renderer.
type Renderer func(n uint64, s Space)
