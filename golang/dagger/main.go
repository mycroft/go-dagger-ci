// A generated module for Golang functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import "fmt"

type Golang struct{}

// Returns a container for testing & build
func (m *Golang) getContainer(golangVersion string, source *Directory) *Container {
	return dag.Container().
		From(fmt.Sprintf("golang:%s-alpine", golangVersion)).
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"apk", "add", "curl"})
}

// Build static-server
func (m *Golang) build(source *Directory) *Container {
	return m.getContainer(source).
		WithExec([]string{"go", "build"})
}
