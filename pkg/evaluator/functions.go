/*
   Copyright Gen Digital Inc. All Rights Reserved.

   SPDX-License-Identifier: Apache-2.0
*/

package evaluator

import (
	"fmt"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

func init() {
	fmt.Println("***** Registering builtin functions *****")

	rego.RegisterBuiltin1(&rego.Function{
		Name: "custom_print",
		Decl: types.NewFunction(types.Args(types.S), types.NewNull()),
	}, func(_ rego.BuiltinContext, op *ast.Term) (*ast.Term, error) {
		fmt.Println(op.Value.(ast.String))

		return ast.NullTerm(), nil
	})

	rego.RegisterBuiltin1(&rego.Function{
		Name: "custom_allow",
		Decl: types.NewFunction(types.Args(types.B), types.NewBoolean()),
	}, func(bctx rego.BuiltinContext, op *ast.Term) (*ast.Term, error) {
		b := bool(op.Value.(ast.Boolean))
		fmt.Printf("***** Returning %t from custom_allow\n", b)
		return ast.BooleanTerm(b), nil
	})
}
