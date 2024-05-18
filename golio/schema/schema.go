package schema

import "embed"

//go:embed postgres/migrations/*.sql
var PostgresMigrations embed.FS
