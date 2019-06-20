package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	"go.zenithar.org/protoc-gen-blueprint/modules/db"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		db.Entity(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
