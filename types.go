package entify

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
)

type Extension struct {
	entc.DefaultExtension
	hooks []gen.Hook
	data  data
}

type Driver = string

type File = uint

const (
	Db File = iota * 2
	Api
	Auth
	Handlers
	Input
	Query
	Privacy
	Routes
	Token
	Types
	Upload
	Utils
)

const (
	SQLite     Driver = "sqlite3"
	MySQL      Driver = "mysql"
	PostgreSQL Driver = "postgres"
)

type option = func(*Extension)

type data struct {
	*gen.Graph
	Config *Config
	Schema *load.Schema
}

type Config struct {
	Driver     *Driver
	Dsn        *string
	ClientPath string

	WithFiles      []File
	WithSwagger    bool
	WithValidation bool
}

type comparable interface {
	~string | ~int | ~float32 | ~uint
}

var go_ts = map[string]string{
	"time.Time": "string",
	"bool":      "boolean",
	"int":       "number",
	"uint":      "number",
	"float":     "number",
	"enum":      "string",
	"string":    "string",
	"any":       "any",
	"other":     "any",
	"json":      "any",
}
