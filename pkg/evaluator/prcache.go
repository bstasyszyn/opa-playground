/*
   Copyright Gen Digital Inc.

   This file contains software code that is the intellectual property of SecureKey.
   SecureKey reserves all rights in the code and you may not use it without
	 written permission from SecureKey.
*/

package evaluator

import (
	"context"
	"fmt"
	"sync"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage"
)

type PartialResultsCache struct {
	cache map[string]*rego.PartialResult
	mutex sync.RWMutex

	compiler *ast.Compiler
	store    storage.Store
}

func NewPartialResultsCache(compiler *ast.Compiler, store storage.Store) *PartialResultsCache {
	return &PartialResultsCache{
		cache:    make(map[string]*rego.PartialResult),
		compiler: compiler,
		store:    store,
	}
}

func (c *PartialResultsCache) Get(ctx context.Context, path string) (*rego.PartialResult, error) {
	c.mutex.RLock()
	r, ok := c.cache[path]
	c.mutex.RUnlock()

	if ok {
		return r, nil
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	pr, err := c.create(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("create rego partial result for path [%s]: %w", path, err)
	}

	c.cache[path] = pr

	return pr, nil
}

func (c *PartialResultsCache) create(ctx context.Context, path string) (*rego.PartialResult, error) {
	pr, err := rego.New(
		rego.Query(path),
		rego.Compiler(c.compiler),
		rego.Store(c.store),
	).PartialResult(ctx)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}
