/*
   Copyright Gen Digital Inc. All Rights Reserved.

   SPDX-License-Identifier: Apache-2.0
*/

package evaluator

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage"
)

type Document = map[string]interface{}

type Module struct {
	Name string
	Code string
}

type options struct {
	modules     []*Module
	modulePaths []string
	store       storage.Store
}

type Opt func(*options)

func WithModule(name, code string) Opt {
	return func(opts *options) {
		opts.modules = append(opts.modules, &Module{
			Name: name,
			Code: code,
		})
	}
}

func WithModuleFile(path string) Opt {
	return func(opts *options) {
		opts.modulePaths = append(opts.modulePaths, path)
	}
}

func WithStore(store storage.Store) Opt {
	return func(opts *options) {
		opts.store = store
	}
}

type Evaluator struct {
	compiler    *ast.Compiler
	store       storage.Store
	cache       *PartialResultsCache
	bundlePaths []string
}

func New(opts ...Opt) (*Evaluator, error) {
	options := &options{}

	for _, opt := range opts {
		opt(options)
	}

	mods, err := resolveModules(options)
	if err != nil {
		return nil, fmt.Errorf("resolve modules: %w", err)
	}

	modules := make(map[string]string)

	for _, m := range mods {
		modules[m.Name] = m.Code
	}

	c, err := ast.CompileModules(modules)
	if err != nil {
		return nil, err
	}

	return &Evaluator{
		compiler: c,
		store:    options.store,
		cache:    NewPartialResultsCache(c, options.store),
	}, nil
}

func (e *Evaluator) Evaluate(ctx context.Context, path string, input Document) (bool, error) {
	b, err := e.EvaluateMultiple(ctx, path, input)
	if err != nil {
		return false, err
	}

	return b[0], nil
}

func (e *Evaluator) EvaluateMultiple(ctx context.Context, path string, inputs ...Document) ([]bool, error) {
	pr, err := e.cache.Get(ctx, path)
	if err != nil {
		return nil, err
	}

	results := make([]bool, len(inputs))

	for i, input := range inputs {
		rs, err := pr.Rego(rego.Input(input)).Eval(ctx)
		if err != nil {
			return nil, err
		}

		results[i] = rs.Allowed()
	}

	return results, nil
}

func resolveModules(opts *options) ([]*Module, error) {
	mods := opts.modules

	for _, path := range opts.modulePaths {
		mod, err := loadModule(path)
		if err != nil {
			return nil, fmt.Errorf("load bundle error: %w", err)
		}

		mods = append(mods, mod)
	}

	return mods, nil
}

func loadModule(path string) (*Module, error) {
	if !strings.HasSuffix(path, ".rego") {
		return nil, fmt.Errorf("module file must have extension '.rego': %s", path)
	}

	modBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read module file %q: %w", path, err)
	}

	parts := strings.Split(path, "/")

	return &Module{
		Name: parts[len(parts)-1],
		Code: string(modBytes),
	}, nil
}
