package entify

import (
	"path"

	"entgo.io/ent/entc/gen"
)

func (e *Extension) generate(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		e.data.Graph = g
		destination := rootDir()

		if in(DB, e.data.Config.Files) {
			writeFile(path.Join(destination, "db/db.go"), parseTemplate("db", e.data))
		}

		if in(Env, e.data.Config.Files) {
			writeFile(path.Join(destination, ".env"), parseTemplate("env", e.data))
		}

		if in(Input, e.data.Config.Files) {
			writeFile(path.Join(destination, "ent/input.go"), parseTemplate("input", e.data))
		}

		if in(Query, e.data.Config.Files) {
			writeFile(path.Join(destination, "ent/query.go"), parseTemplate("query", e.data))
		}

		if in(Token, e.data.Config.Files) {
			writeFile(path.Join(destination, e.data.Config.ClientPath, "middleware/token.ts"), parseTemplate("token", e.data))
		}

		if in(Api, e.data.Config.Files) {
			writeFile(path.Join(destination, e.data.Config.ClientPath, "api.ts"), parseTemplate("api", e.data))
		}

		if in(Types, e.data.Config.Files) {
			writeFile(path.Join(destination, e.data.Config.ClientPath, "types.ts"), parseTemplate("types", e.data))
		}

		if in(Auth, e.data.Config.Files) {
			writeFile(path.Join(destination, "handlers/auth.go"), parseTemplate("input", e.data))
		}

		if in(Privacy, e.data.Config.Files) {
			writeFile(path.Join(destination, "privacy/privacy.go"), parseTemplate("privacy", e.data))
		}

		if in(Utils, e.data.Config.Files) {
			writeFile(path.Join(destination, "utils/utils.go"), parseTemplate("utils", e.data))
		}

		return next.Generate(g)
	})
}
