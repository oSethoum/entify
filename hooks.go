package entify

import (
	"fmt"
	"path"

	"entgo.io/ent/entc/gen"
)

func (e *Extension) generate(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		e.data.Graph = g
		destination := rootDir()

		// backend
		if in(DB, e.data.Config.Files) {
			writeFile(path.Join(destination, "db/db.go"), parseTemplate("backend/db", e.data))
		}

		if in(Env, e.data.Config.Files) {
			writeFile(path.Join(destination, ".env"), parseTemplate("backend/env", e.data))
		}

		if in(Input, e.data.Config.Files) {
			writeFile(path.Join(destination, "ent/input.go"), parseTemplate("backend/input", e.data))
		}

		if in(Query, e.data.Config.Files) {
			writeFile(path.Join(destination, "ent/query.go"), parseTemplate("backend/query", e.data))
		}

		if in(Query, e.data.Config.Files) {
			writeFile(path.Join(destination, "routes/routes.go"), parseTemplate("backend/routes", e.data))
		}

		if in(Model, e.data.Config.Files) {
			writeFile(path.Join(destination, "models/models.go"), parseTemplate("backend/models", e.data))
		}

		if in(Swagger, e.data.Config.Files) {
			writeFile(path.Join(destination, "models/swagger.go"), parseTemplate("backend/swagger", e.data))
		}

		if in(Token, e.data.Config.Files) {
			writeFile(path.Join(destination, "middleware/token.go"), parseTemplate("backend/token", e.data))
		}

		for _, s := range e.data.Schemas {
			e.data.CurrentSchema = s
			writeFile(path.Join(destination, fmt.Sprintf("handlers/%s.go", plural(kebab(s.Name)))), parseTemplate("backend/handler", e.data))
		}

		if in(Files, e.data.Config.Files) {
			writeFile(path.Join(destination, "handlers/u-files.go"), parseTemplate("backend/u-files", e.data))
		}

		if in(Auth, e.data.Config.Files) {
			writeFile(path.Join(destination, "handlers/auth.go"), parseTemplate("backend/auth", e.data))
		}

		if in(Privacy, e.data.Config.Files) {
			writeFile(path.Join(destination, "privacy/privacy.go"), parseTemplate("backend/privacy", e.data))
		}

		if in(Utils, e.data.Config.Files) {
			writeFile(path.Join(destination, "utils/utils.go"), parseTemplate("backend/utils", e.data))
		}

		// frontend
		if in(Api, e.data.Config.Files) {
			writeFile(path.Join(destination, e.data.Config.ClientPath, "api.ts"), parseTemplate("frontend/api", e.data))
		}

		if in(Types, e.data.Config.Files) {
			writeFile(path.Join(destination, e.data.Config.ClientPath, "types.ts"), parseTemplate("frontend/types", e.data))
		}

		return next.Generate(g)
	})
}
