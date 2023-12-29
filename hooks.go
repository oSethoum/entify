package entify

import "entgo.io/ent/entc/gen"

func (e *Extension) generate(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		e.data.Graph = g
		// destination := rootDir()

		return next.Generate(g)
	})
}
