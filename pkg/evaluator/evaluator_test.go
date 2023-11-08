/*
   Copyright Gen Digital Inc. All Rights Reserved.

   SPDX-License-Identifier: Apache-2.0
*/

package evaluator

import (
	"bytes"
	"context"
	"sync"
	"testing"

	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/stretchr/testify/require"
)

const (
	allowPath      = "data.module1.allow"
	notAllowPath   = "data.module1.not_allow"
	customFuncPath = "data.module1.custom_func"

	module1 = "./testdata/module1.rego"
	module2 = "./testdata/module2.rego"
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
		WithModuleFile(module1),
		WithModuleFile(module2),
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

	t.Run("allow -> true", func(t *testing.T) {
		allowed, err := e.Evaluate(context.TODO(), allowPath, input2)
		require.NoError(t, err)
		require.True(t, allowed)
	})

	t.Run("custom_func -> true", func(t *testing.T) {
		allowed, err := e.Evaluate(context.TODO(), customFuncPath, nil)
		require.NoError(t, err)
		require.True(t, allowed)
	})
}

func TestEvaluateMultiple(t *testing.T) {
	store := inmem.NewFromReader(bytes.NewBufferString(data))

	e, err := New(
		WithModuleFile(module1),
		WithModuleFile(module2),
		WithStore(store),
	)
	require.NoError(t, err)

	t.Run("Evaluates to [true,false,true]", func(t *testing.T) {
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

type testData struct {
	input         Document
	expectedAllow bool
}

type testResult struct {
	path     string
	input    Document
	expected bool
	result   bool
	err      error
}

func TestConcurrency(t *testing.T) {
	inputData := []*testData{
		{
			input: Document{
				"resource":  "documentA",
				"operation": "write",
				"subject": Document{
					"user": "bob",
				},
			},
			expectedAllow: true,
		},
		{
			input: Document{
				"resource":  "documentB",
				"operation": "write",
				"subject": map[string]interface{}{
					"user": "alice",
				},
			},
			expectedAllow: false,
		},
		{
			input: Document{
				"resource":  "documentC",
				"operation": "read",
				"subject": map[string]interface{}{
					"user": "jim",
				},
			},
			expectedAllow: true,
		},
	}

	var inputs []*testResult

	for i := 0; i < 1000; i++ {
		for _, d := range inputData {
			inputs = append(inputs, &testResult{
				input:    d.input,
				path:     allowPath,
				expected: d.expectedAllow,
			})

			inputs = append(inputs, &testResult{
				input:    d.input,
				path:     notAllowPath,
				expected: !d.expectedAllow,
			})
		}
	}

	eval, err := New(
		WithModuleFile(module1),
		WithModuleFile(module2),
		WithStore(inmem.NewFromReader(bytes.NewBufferString(data))),
	)
	require.NoError(t, err)

	var wg sync.WaitGroup

	wg.Add(len(inputs))

	for _, input := range inputs {
		go func(input *testResult) {
			input.result, input.err = eval.Evaluate(context.TODO(), input.path, input.input)

			wg.Done()
		}(input)
	}

	wg.Wait()

	for i, input := range inputs {
		require.NoError(t, input.err)
		require.Equalf(t, input.expected, input.result, "Input[%d] - Path [%s]", i, input.path)
	}
}
