/*
   Copyright SecureKey Technologies Inc.

   This file contains software code that is the intellectual property of SecureKey.
   SecureKey reserves all rights in the code and you may not use it without
	 written permission from SecureKey.
*/

package evaluator

import (
	"bytes"
	"context"
	"testing"

	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/stretchr/testify/require"
)

const (
	allowPath    = "data.module1.allow"
	notAllowPath = "data.module1.not_allow"
)

const data = `{
	"roles": [
		{
			"resources": ["documentA", "documentB"],
			"operations": ["read"],
			"name": "analyst"
		},
		{
			"resources": ["documentC"],
			"operations": ["read"],
			"name": "poet"
		},
		{
			"resources": ["*"],
			"operations": ["*"],
			"name": "admin"
		}
	],
	"bindings": [
		{
			"user": "bob",
			"role": "admin"
		},
		{
			"user": "alice",
			"role": "analyst"
		},
		{
			"user": "jim",
			"role": "poet"
		}
	]
}`

func TestEvaluate(t *testing.T) {
	store := inmem.NewFromReader(bytes.NewBufferString(data))

	e, err := New(
		WithModuleFile("./testdata/module1.rego"),
		WithModuleFile("./testdata/module2.rego"),
		WithStore(store),
	)
	require.NoError(t, err)

	input1 := Document{
		"resource":  "documentA",
		"operation": "write",
		"subject": Document{
			"user": "bob",
		},
	}

	input2 := Document{
		"resource":  "documentB",
		"operation": "read",
		"subject": Document{
			"user": "alice",
		},
	}

	t.Run("allow -> true", func(t *testing.T) {
		allowed, err := e.Evaluate(context.TODO(), allowPath, input1)
		require.NoError(t, err)
		require.True(t, allowed)
	})

	t.Run("not_allow -> false", func(t *testing.T) {
		allowed, err := e.Evaluate(context.TODO(), notAllowPath, input1)
		require.NoError(t, err)
		require.False(t, allowed)
	})

	t.Run("allow -> false", func(t *testing.T) {
		allowed, err := e.Evaluate(context.TODO(), allowPath, input2)
		require.NoError(t, err)
		require.False(t, allowed)
	})
}

func TestEvaluateMultiple(t *testing.T) {
	store := inmem.NewFromReader(bytes.NewBufferString(data))

	e, err := New(
		WithModuleFile("./testdata/module1.rego"),
		WithModuleFile("./testdata/module2.rego"),
		WithStore(store),
	)
	require.NoError(t, err)

	t.Run("Evaluates to [true,false]", func(t *testing.T) {
		allowed, err := e.EvaluateMultiple(context.TODO(),
			allowPath,
			Document{
				"resource":  "documentA",
				"operation": "write",
				"subject": map[string]interface{}{
					"user": "bob",
				},
			},
			Document{
				"resource":  "documentB",
				"operation": "write",
				"subject": map[string]interface{}{
					"user": "alice",
				},
			},
			Document{
				"resource":  "documentC",
				"operation": "read",
				"subject": map[string]interface{}{
					"user": "jim",
				},
			},
		)
		require.NoError(t, err)
		require.Len(t, allowed, 3)
		require.True(t, allowed[0])
		require.False(t, allowed[1])
		require.True(t, allowed[2])
	})
}
