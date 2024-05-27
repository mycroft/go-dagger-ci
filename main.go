package main

import (
	"context"
	"fmt"
)

type App struct{}

const (
	GO_VERSION = "1.20"
)

// Returns a container for testing & build
func (m *App) getContainer(source *Directory) *Container {
	return dag.Container().
		From(fmt.Sprintf("golang:%s-alpine", GO_VERSION)).
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"apk", "add", "curl"})
}

// Build static-server
func (m *App) build(source *Directory) *Container {
	return m.getContainer(source).
		WithExec([]string{"go", "build"})
}

// Simply build the project with go build
func (m *App) Build(ctx context.Context, source *Directory) (string, error) {
	return m.build(source).
		Stdout(ctx)
}

// Run tests
func (m *App) Test(ctx context.Context, source *Directory) (string, error) {
	return m.getContainer(source).
		WithExec([]string{"go", "test"}).
		Stdout(ctx)
}

// Build & test the project
func (m *App) BuildTest(ctx context.Context, source *Directory) (string, error) {
	_, err := m.Build(ctx, source)
	if err != nil {
		return "", err
	}

	return m.Test(ctx, source)
}

// Run the whole pipeline: test, build & publish
func (m *App) Publish(ctx context.Context, source *Directory, regUsername string, regPassword *Secret, regAddress string, imageName string, tag string) (string, error) {
	_, err := m.Test(ctx, source)
	if err != nil {
		return "", err
	}

	address, err := m.build(source).
		WithRegistryAuth(regAddress, regUsername, regPassword).
		Publish(ctx, fmt.Sprintf("%s/%s:%s", regAddress, imageName, tag))
	if err != nil {
		return "", err
	}

	return address, err
}
